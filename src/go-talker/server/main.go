package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"

	"github.com/astaxie/beego/logs"
	cli "github.com/urfave/cli"
)

var (
	logger *logs.BeeLogger
	stopIt chan os.Signal = make(chan os.Signal, 1)
	port   string         = ":34567"
	pool   *ConnectionPool
)

const VERSION = "0.0.1"

type Message struct {
	Content string
	Time    uint64
}

type Client struct {
	Name       string
	Connection net.Conn
	Address    net.Addr
}

type ConnectionPool struct {
	Locker  sync.Mutex
	Clients []Client
}

func (self *ConnectionPool) Insert(client Client) {
	self.Locker.Lock()
	self.Clients = append(pool.Clients, client)
	self.Locker.Unlock()
}

func (self *ConnectionPool) Remove(client Client) {
	self.Locker.Lock()
	for i := 0; i < len(self.Clients); i++ {
		if self.Clients[i].Name == client.Name {
			self.Clients = append(self.Clients[:i], self.Clients[i+1:]...)
			break
		}
	}
	self.Locker.Unlock()
}

func main() {
	Init()
	app := cli.NewApp()
	app.Version = VERSION
	app.UsageText = "go-talker"
	start := cli.Command{
		Name:      "start",
		ShortName: "s",
		Usage:     "start go-talker",
		Action:    startAction,
		HelpName:  "help",
	}
	app.Commands = []cli.Command{
		start,
	}
	app.Run(os.Args)
}

func startAction(ctx *cli.Context) {
	listener, err := net.Listen("tcp", ":34567")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	go func(_listener net.Listener) {
		for {
			conn, err := _listener.Accept()
			fmt.Println("open connection:", conn.RemoteAddr())
			go func(_conn net.Conn) {
				if err != nil {
					logger.Info("connection failure , ", _conn.RemoteAddr, err)
					return
				}
				err = _conn.SetReadDeadline(time.Now().Add(time.Duration(5) * time.Second))
				if err != nil {
					logger.Info("connection timeout", _conn.RemoteAddr, err)
					_conn.Close()
					return
				}
				bytess := make([]byte, 2048)
				count, err := _conn.Read(bytess)
				if err != nil {
					logger.Info("read bytes", _conn.RemoteAddr, err)
					_conn.Close()
					return
				}
				if count > 2048 {
					logger.Info("message size too big", _conn.RemoteAddr, err)
					_conn.Close()
					return
				}
				firstMsg, err := BytesToMessage(bytess)
				if err != nil || firstMsg == nil {
					logger.Info("fist message is wrong", _conn.RemoteAddr, err)
					_conn.Close()
					return
				}
				logger.Info("get first message", firstMsg)
				uname, err := AuthCheck(firstMsg)
				if err != nil {
					logger.Info(err.Error() + ";" + uname)
					_conn.Close()
					return
				}
				client := Client{
					Address:    _conn.RemoteAddr(),
					Connection: _conn,
					Name:       uname,
				}
				pool.Insert(client)

				//broadcast
			}(conn)
		}
	}(listener)

	fmt.Println("application started listen on", port)

	go func() {
		signal.Notify(stopIt, os.Interrupt, os.Kill)
	}()

	_ = <-stopIt
	fmt.Print("application stopped")
}

func MessageDelivery(client Client) {
	bytess := make([]byte, 0)
	for {
		count, err := client.Connection.Read(bytess)
		if err != nil || count == 0 {
			logger.Info("read bytes :", err)
			client.Connection.Close()
			pool.Remove(client)
			return
		}
		message, err := BytesToMessage(bytess[:count])
		fmt.Println("get message:", message)
		if err != nil {
			logger.Info("can not read bytes from client", err)
			client.Connection.Close()
			pool.Remove(client)
			return
		}
		formattedMessage := MessageFormatter(client.Name, message.Content)
		message.Content = formattedMessage
		messageBytes, err := MessageToBytes(message)
		if err != nil {
			client.Connection.Close()
			pool.Remove(client)
			return
		}
		pool.Locker.Lock()
		for i := 0; i < len(pool.Clients); i++ {
			currentClient := pool.Clients[i]
			if currentClient.Name == client.Name {
				continue
			}
			currentClient.Connection.Write(messageBytes)
		}
		pool.Locker.Unlock()
	}
}

func BroadcastMessage() {}

func MessageFormatter(uname, content string) string {
	return fmt.Sprintf("[%s]:%s", uname, content)
}

func MessageToBytes(msg *Message) ([]byte, error) {
	buf := &bytes.Buffer{}
	err := binary.Write(buf, binary.BigEndian, *msg)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func BytesToMessage(bytess []byte) (*Message, error) {
	msg := &Message{}
	buf := &bytes.Buffer{}
	buf.Read(bytess)
	err := binary.Read(buf, binary.BigEndian, msg)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

func AuthCheck(msg *Message) (string, error) {
	s := strings.Split(msg.Content, "\\n")
	if len(s) != 2 {
		return "", errors.New("wrong message type")
	}
	uName := s[0]
	uKey := s[1]
	if uName == "wu-xian" && uKey == "123" {
		return "wu-xian", nil
	}
	return "", errors.New("Authentiaction failure")
}

func Init() {
	InitConnectionPool()
	InitLogger()
}

func InitConnectionPool() {
	pool = &ConnectionPool{}
}

func InitLogger() {
	logger = logs.NewLogger(1000)
	logger.SetLogger("file", `{"filename":"test.log"}`)
}
