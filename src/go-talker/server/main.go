package main

import (
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"

	"learn/src/go-talker/log"

	"learn/src/go-talker/proto"

	"learn/src/go-talker/common"

	"github.com/astaxie/beego/logs"
	p "github.com/golang/protobuf/proto"
	cli "github.com/urfave/cli"
)

var (
	logger            *logs.BeeLogger
	stopIt            chan os.Signal = make(chan os.Signal, 1)
	port              string         = ":34567"
	pool              *ConnectionPool
	clientIndex       uint8 = 1
	clientIndexLocker sync.Mutex
)

const VERSION = "0.0.1"
const MESSAGE_MAX_LENGTH = 2048

type Client struct {
	Id         uint8
	Connection *net.TCPConn
	Address    net.Addr
	Name       string
}

type ConnectionPool struct {
	Locker  sync.Mutex
	Clients []Client
}

func getClientIndex() uint8 {
	clientIndexLocker.Lock()
	clientIndex++
	clientIndexLocker.Unlock()
	return clientIndex
}

func (self *ConnectionPool) Insert(client *Client) {
	self.Locker.Lock()
	self.Clients = append(pool.Clients, *client)
	self.Locker.Unlock()
}

func (self *ConnectionPool) Remove(client *Client) {
	self.Locker.Lock()
	for i := 0; i < len(self.Clients); i++ {
		self.Clients[i].Connection.Write([]byte("client" + strconv.Itoa(int(self.Clients[i].Id))))
		if self.Clients[i].Id == client.Id {
			self.Clients = append(self.Clients[:i], self.Clients[i+1:]...)
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
			rawConn, _ := _listener.Accept()
			conn := rawConn.(*net.TCPConn)
			fmt.Println("open connection:", conn.RemoteAddr())
			go func(_conn *net.TCPConn) {
				defer log.Logger.Info("connection closed : ", _conn.RemoteAddr().String())
				defer _conn.Close()
				client, loginMessage, err := Login(_conn)
				if err != nil {
					log.Logger.Info("connection closed  ", err)
					return
				}
				if loginMessage == nil {
					log.Logger.Info("wrong message")
					return
				}
				BroadcastMessage(loginMessage)
				//broadcast
				MessageDelivery(client)

				logoutMessage := proto.LogoutMessage{
					Id:   int32(client.Id),
					Name: client.Name,
					Head: &proto.Header{
						Type:   1,
						Length: 0,
					},
				}

				logoutMessage.Head.Length = int64(p.Size(&logoutMessage))
				BroadcastMessage(logoutMessage)
				Logout(client, logoutMessage)
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

func MessageDelivery(client *Client) {
	bytess := make([]byte, MESSAGE_MAX_LENGTH)
	for {
		count, err := client.Connection.Read(bytess)
		//client.Connection.CloseRead()
		log.Logger.Info("get bytes from client ", bytess[:count])
		if err != nil {
			logger.Info("read bytes :", err)
			pool.Remove(client)
			return
		}
		message := MessageInterpreter(bytess[:count])
		switch message.(type) {
		case proto.LoginMessage:
			{
				log.Logger.Info("so much login message")
				return
			}
		case proto.LogoutMessage:
			{
				return
			}
		case proto.Content:
			{
				content := message.(proto.Content)
				logger.Info("get bytes from client", content)
				if err != nil {
					logger.Info("can not read bytes from client", err)
					client.Connection.Close()
					pool.Remove(client)
					return
				}
				formattedMessage := MessageFormatter(client.Name, content.Content)
				content.Content = formattedMessage
				content.Unmarshal(bytess)
				if err != nil {
					client.Connection.Close()
					pool.Remove(client)
					return
				}
				pool.Locker.Lock()
				for i := 0; i < len(pool.Clients); i++ {
					currentClient := pool.Clients[i]
					if currentClient.Id == client.Id {
						continue
					}
					currentClient.Connection.Write(bytess)
				}
				pool.Locker.Unlock()
			}
		default:
			{
				log.Logger.Info("unknown message type")
				return
			}
		}
	}
}

func BroadcastMessage(message interface{}) {
	bytes := make([]byte, MESSAGE_MAX_LENGTH)
	switch message.(type) {
	case proto.LoginMessage:
		{
			m := message.(proto.LoginMessage)
			bytes1, err := m.Marshal()
			bytes = bytes1
			common.CheckError(err)
		}
	case proto.LogoutMessage:
		{
			m := message.(proto.LogoutMessage)
			bytes1, err := m.Marshal()
			bytes = bytes1
			common.CheckError(err)
		}
	case proto.Content:
		{
			m := message.(proto.Content)
			bytes1, err := m.Marshal()
			bytes = bytes1
			common.CheckError(err)
		}
	default:
		{
			return
		}
	}
	pool.Locker.Lock()
	for i := 0; i < len(pool.Clients); i++ {
		pool.Clients[i].Connection.Write(bytes)
	}
	pool.Locker.Unlock()
}

func MessageFormatter(uname, content string) string {
	return fmt.Sprintf("[%s]:%s", uname, content)
}

func MessageInterpreter(bytes []byte) (msg interface{}) {
	header := proto.Header{}
	err := p.Unmarshal(bytes, &header)
	if err != nil {
		log.Logger.Info("MessageInterpreter", err)
		return
	}
	switch header.Type {
	case 0:
		{
			lm := proto.LoginMessage{}
			lm.Unmarshal(bytes)
			msg = lm
		}
	case 1:
		{
			lm := proto.LogoutMessage{}
			lm.Unmarshal(bytes)
			msg = lm
		}
	case 2:
		{
			c := proto.Content{}
			c.Unmarshal(bytes)
			msg = c
		}
	default:
		{
			return nil
		}
	}
	return msg
}

func Login(conn *net.TCPConn) (client *Client, message *proto.LoginMessage, err error) {
	if err := conn.SetDeadline(time.Now().Add(10 * time.Second)); err != nil {
		return nil, nil, err
	}
	bytes := make([]byte, MESSAGE_MAX_LENGTH)
	var loginMessage proto.LoginMessage
	for i := 0; i < 3; i++ {
		count, err := conn.Read(bytes[:])
		common.CheckError(err)
		if count == MESSAGE_MAX_LENGTH {
			return nil, nil, errors.New("message too large")
		}
		msg := MessageInterpreter(bytes[:count])
		switch msg.(type) {
		case proto.LoginMessage:
			{
				loginMessage = msg.(proto.LoginMessage)
				break
			}
		default:
			{
				if i == 2 {
					return nil, nil, errors.New("wrong login message")
				}
				time.Sleep(time.Second * 3)
				continue
			}
		}
	}
	clientIndex := getClientIndex()
	client = &Client{
		Address:    conn.RemoteAddr(),
		Connection: conn,
		Id:         clientIndex,
		Name:       loginMessage.Name,
	}
	pool.Insert(client)
	return client, &loginMessage, nil
}

func Logout(client *Client, logoutMessage proto.LogoutMessage) {
	err := client.Connection.Close()
	if err != nil {
		logger.Info("can not read bytes from client", err)
	}
	pool.Remove(client)
}
func Init() {
	InitConnectionPool()
	logger = log.Logger
}

func InitConnectionPool() {
	pool = &ConnectionPool{}
}
