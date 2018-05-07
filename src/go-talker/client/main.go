package main

import (
	"errors"
	"fmt"
	"learn/src/go-talker/proto"
	"net"
	"os"
	"strconv"
	"time"

	"learn/src/go-talker/log"

	"learn/src/go-talker/terminal"

	p "github.com/golang/protobuf/proto"

	"gopkg.in/ini.v1"
)

var (
	IP         string
	Port       int
	UName      string
	Key        string
	connection *net.TCPConn
	message    chan string    = make(chan string, 1)
	stopIt     chan os.Signal = make(chan os.Signal, 1)
)

func main() {
	//terminal.LoopClientUI(message)
	//return
	err := Init()
	if err != nil {
		fmt.Println(err)
		return
	}
	dialer := net.Dialer{
		Timeout: time.Duration(5) * time.Second,
	}
	conn, err := dialer.Dial("tcp", IP+":"+strconv.Itoa(Port))
	connection = conn.(*net.TCPConn)
	defer connection.Close()
	log.Logger.Info("has been connected to the server...")
	if err != nil {
		log.Logger.Info("unable to connect to the server : %s:%d", IP, Port)
		return
	}
	go MessageReceiver(connection)
	go MessagePublisher(connection)

	// go func() {
	// 	signal.Notify(stopIt, os.Interrupt, os.Kill)
	// }()

	//_ = <-stopIt
	terminal.LoopClientUI(message)

	connection.CloseRead()
	fmt.Println("application stopped")
}

func MessageReceiver(conn *net.TCPConn) {
	for {
		bytes := make([]byte, 20480)
		count, err := conn.Read(bytes)
		//conn.CloseRead()
		if err != nil {
			log.Logger.Info("unable to read message", err)
			return
		}
		message, err := proto.Unmarshal(bytes[:count])
		if err != nil {
			log.Logger.Info("invalid message:", err)
			return
		}
		fmt.Println(message.Content)
	}
}

func MessagePublisher(conn *net.TCPConn) {
	for {
		content := <-message
		message := proto.Message{
			Content: content,
			Name:    UName,
			Time:    time.Now().Unix(),
		}
		bytess, err := proto.MessageToBytes(&message)
		log.Logger.Info("get bytes ", bytess)
		if err != nil {
			return
		}
		_, err = conn.Write(bytess)
		//err = conn.CloseWrite()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
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

func Login(conn *net.TCPConn, loginMessage proto.LoginMessage) {

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func Init() error {
	cfg, err := ini.Load("talker.conf")
	if err != nil {
		return errors.New("failure to load config file:talker.conf")
	}
	serverSection := cfg.Section("server")
	IP = serverSection.Key("ip").String()
	Port, err = serverSection.Key("port").Int()
	if err != nil {
		return errors.New("failure to load config : server.port")
	}
	clientSection := cfg.Section("client")
	UName = clientSection.Key("name").String()
	if UName == "" {
		panic("name is empty")
	}
	Key = clientSection.Key("key").String()
	return nil
}
