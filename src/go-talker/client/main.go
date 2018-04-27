package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"

	"gopkg.in/ini.v1"
)

var (
	IP         string
	Port       int
	UName      string
	Key        string
	connection net.Conn
	stopIt     chan os.Signal = make(chan os.Signal, 1)
)

type Message struct {
	Content string
	Time    time.Time
}

func main() {
	err := Init()
	if err != nil {
		fmt.Println(err)
		return
	}
	serverAddress := &net.TCPAddr{
		IP:   net.ParseIP(IP),
		Port: Port,
	}
	dialer := net.Dialer{
		Timeout: time.Now().Add(time.Duration(5) * time.Second),
	}
	connection, err := dialer.Dial("tcp", nil, serverAddress)
	defer connection.Close()
	fmt.Println("has been connected to the server...")
	if err != nil {
		fmt.Println("unable to connect to the server : %s:%d", IP, Port)
	}

	go MessagePublisher(connection)
	go MessagePublisher(connection)

	go func() {
		signal.Notify(stopIt, os.Interrupt, os.Kill)
	}()

	_ = <-stopIt
	fmt.Println("application stopped")
}

func MessageReceiver(conn *net.TCPConn) {
	for {
		bytes := []byte{}
		_, err := conn.Read(bytes)
		if err != nil {
			fmt.Println("unable to read message", err)
			continue
		}
		message, err := BytesToMessage(bytes)
		if err != nil {
			fmt.Println("invalid message")
			continue
		}
		fmt.Println(message.Content)
	}
}

func MessagePublisher(conn *net.TCPConn) {
	content := ""
	for {
		fmt.Scanln(content)
		message := Message{
			Content: content,
			Time:    time.Now(),
		}
		bytess, err := MessageToBytes(&message)
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = conn.Write(bytess)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func MessageFormatter(uname, content string) string {
	return fmt.Sprintf("[%s]:%s", uname, content)
}

func MessageToBytes(msg *Message) ([]byte, error) {
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
	UName = clientSection.Key("uname").String()
	Key = clientSection.Key("key").String()
	return nil
}
