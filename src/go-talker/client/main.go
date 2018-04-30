package main

import (
	"errors"
	"fmt"
	"learn/src/go-talker/proto"
	"net"
	"os"
	"os/signal"
	"strconv"
	"time"

	"learn/src/go-talker/log"

	"gopkg.in/ini.v1"
)

var (
	IP         string
	Port       int
	UName      string
	Key        string
	connection *net.TCPConn
	stopIt     chan os.Signal = make(chan os.Signal, 1)
)

func main() {
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
	}
	go MessageReceiver(connection)
	go MessagePublisher(connection)

	go func() {
		signal.Notify(stopIt, os.Interrupt, os.Kill)
	}()

	_ = <-stopIt
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
		message, err := proto.BytesToMessage(bytes[:count])
		if err != nil {
			log.Logger.Info("invalid message:", err)
			return
		}
		fmt.Println(message.Content)
	}
}

func MessagePublisher(conn *net.TCPConn) {
	content := ""
	for {
		_, err := fmt.Scanln(&content)
		if err != nil {
			fmt.Println("errors:", err)
			return
		}
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

func Init() error {
	log.InitLogger()
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
