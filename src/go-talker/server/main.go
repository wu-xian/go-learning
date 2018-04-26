package main

import (
	"bytes"
	"encoding/binary"
	"os"
	"sync"
	"time"

	"net"

	"github.com/astaxie/beego/logs"
	cli "github.com/urfave/cli"
)

var (
	logger *logs.BeeLogger
	stopIt chan bool
)

const VERSION = "0.0.1"

type Message struct {
	Content string
	Time    uint64
	Type    string
}

type Client struct {
	Name       string
	Connection net.Conn
	Address    net.Addr
}

type ConnectionPool struct {
	Locker sync.Mutex
	Pool   []Client
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
			if err != nil {
				logger.Info("connection failure , ", conn.RemoteAddr, err)
				continue
			}
			err = conn.SetReadDeadline(time.Now().Add(time.Duration(5) * time.Second))
			if err != nil {
				logger.Info("connection timeout", conn.RemoteAddr, err)
				conn.Close()
				continue
			}
			bytess := make([]byte, 2048)
			count, err := conn.Read(bytess)
			if err != nil {
				logger.Info("read bytes", conn.RemoteAddr, err)
				conn.Close()
				continue
			}
			if count > 2048 {
				logger.Info("message size too big", conn.RemoteAddr, err)
				conn.Close()
				continue
			}
			firstMsg, err := BytesToMessage(bytess)
			if err != nil || firstMsg == nil {
				logger.Info("fist message is wrong", conn.RemoteAddr, err)
				conn.Close()
				continue
			}
			logger.Info("get first message", firstMsg)
		}
	}(listener)

	_ = <-stopIt
}

func MessageToBytes(msg Message) ([]byte, error) {
	buf := &bytes.Buffer{}
	err := binary.Write(buf, binary.BigEndian, &msg)
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

func Init() {
	InitLogger()
}

func InitLogger() {
	logger = logs.NewLogger(1000)
	logger.SetLogger("file", `{"filename":"test.log"}`)
}
