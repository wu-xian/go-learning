package main

import (
	"errors"
	"fmt"
	"learn/src/go-talker/common"
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
	message    chan string    = make(chan string, 1)
	stopIt     chan os.Signal = make(chan os.Signal, 1)
)

const MESSAGE_MAX_LENGTH = 2048

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

	err = Login(connection)
	if err != nil {
		fmt.Println(err)
		return
	}

	go MessageReceiver(connection)
	go MessagePublisher(connection)

	go func() {
		signal.Notify(stopIt, os.Interrupt, os.Kill)
	}()

	Logout(connection)

	_ = <-stopIt
	//terminal.LoopClientUI(message)

	Logout(connection)
	connection.CloseRead()
	fmt.Println("application stopped")
}

func MessageReceiver(conn *net.TCPConn) {
	for {
		bytes := make([]byte, MESSAGE_MAX_LENGTH)
		count, err := conn.Read(bytes)
		//conn.CloseRead()
		if err != nil {
			log.Logger.Info("unable to read message", err)
			return
		}
		message := MessageInterpreter(bytes[:count])
		if err != nil {
			log.Logger.Info("invalid message:", err)
			return
		}
		switch message.Type {
		case proto.COMMUNICATION_TYPE_ClientReceived:
			{
				fmt.Println(message.MessageClientReceived.Content)
			}
		default:
			{
				fmt.Println("default message")
			}
		}
	}
}

func MessagePublisher(conn *net.TCPConn) {
	for {
		content := <-message
		message := proto.MessageWarpper{
			Type: proto.COMMUNICATION_TYPE_ClientSend,
			MessageClientSend: &proto.MessageClientSend{
				Content: content,
			},
		}
		bytess, err := message.Marshal()
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

//MessageInterpreter 获取包装壳
func MessageInterpreter(bytes []byte) (msg *proto.MessageWarpper) {
	warpper := &proto.MessageWarpper{}
	err := warpper.Unmarshal(bytes)
	if err != nil {
		log.Logger.Info("", err)
		return nil
	}
	return warpper
}

//Login 客户端登陆
func Login(conn *net.TCPConn) error {
	loginMessage := &proto.MessageWarpper{
		Type: proto.COMMUNICATION_TYPE_LoginRequest,
		MessageLoginRequest: &proto.MessageLoginRequest{
			Name:  UName,
			Token: "",
		},
	}
	bytes, err := loginMessage.Marshal()
	if err != nil {
		log.Logger.Info("", err)
		return err
	}
	for i := 0; i < 3; i++ {
		count, err := conn.Write(bytes)
		common.CheckError(err)
		if count >= MESSAGE_MAX_LENGTH {
			return errors.New("message too large")
		}

		readBytes := make([]byte, 0)
		count, err = conn.Read(readBytes)
		msg := MessageInterpreter(readBytes[:count])
		if msg.Type == proto.COMMUNICATION_TYPE_LoginResponse &&
			msg.MessageLoginResponse.Succeed {
			break
		}
		if i == 2 {
			return errors.New("wrong login message")
		}
		time.Sleep(time.Second * 3)
		continue
	}

	return nil
}

//Login 客户端登出
func Logout(conn *net.TCPConn) error {
	logoutMessage := &proto.MessageWarpper{
		Type:                 proto.COMMUNICATION_TYPE_LogoutRequest,
		MessageLogoutRequest: &proto.MessageLogoutRequest{},
	}
	bytes, err := logoutMessage.Marshal()
	if err != nil {
		log.Logger.Info("", err)
		return err
	}
	for i := 0; i < 3; i++ {
		count, err := conn.Write(bytes)
		common.CheckError(err)
		if count >= MESSAGE_MAX_LENGTH {
			return errors.New("message too large")
		}

		readBytes := make([]byte, 0)
		count, err = conn.Read(readBytes)
		msg := MessageInterpreter(readBytes[:count])
		if msg.Type == proto.COMMUNICATION_TYPE_LogoutResponse &&
			msg.MessageLogoutResponse.Succeed {
			break
		}
		if i == 2 {
			return errors.New("wrong logout message")
		}
		time.Sleep(time.Second * 3)
		continue
	}

	return nil
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
