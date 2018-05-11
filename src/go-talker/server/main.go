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

	"github.com/astaxie/beego/logs"
	cli "github.com/urfave/cli"
)

var (
	logger            *logs.BeeLogger
	stopIt            chan os.Signal = make(chan os.Signal, 1)
	port              string         = ":34567"
	pool              *ConnectionPool
	clientIndex       int32 = 1
	clientIndexLocker sync.Mutex
)

const VERSION = "0.0.1"
const MESSAGE_MAX_LENGTH = 2048

type Client struct {
	Id         int32
	Connection *net.TCPConn
	Address    net.Addr
	Name       string
}

type ConnectionPool struct {
	Locker  sync.Mutex
	Clients []Client
}

func getClientIndex() int32 {
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
				//broadcast
				MessageDelivery(client)

				Logout(client)
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

//MessageDelivery 消息分发
func MessageDelivery(client *Client) {
	bytes := make([]byte, MESSAGE_MAX_LENGTH)
	for {
		count, err := client.Connection.Read(bytes)
		log.Logger.Info("get bytes from client ", bytes[:count])
		if err != nil {
			logger.Info("read bytes :", err)
			pool.Remove(client)
			return
		}
		message := MessageInterpreter(bytes[:count])
		switch message.Type {
		case proto.COMMUNICATION_TYPE_LogoutRequest:
			{
				Logout(client)
				return
			}
		case proto.COMMUNICATION_TYPE_ClientSend:
			{
				logger.Info("get bytes from client", message.MessageClientSend.Content)
				messageReceive := &proto.MessageWarpper{
					Type: proto.COMMUNICATION_TYPE_ClientReceived,
					MessageClientReceived: &proto.MessageClientReceived{
						Id:      client.Id,
						Name:    client.Name,
						Content: message.MessageClientSend.Content,
					},
				}
				bytesTmp, err := messageReceive.Marshal()
				if err != nil {
					log.Logger.Info("err", err)
					return
				}
				pool.Locker.Lock()
				for i := 0; i < len(pool.Clients); i++ {
					currentClient := pool.Clients[i]
					if currentClient.Id == client.Id {
						continue
					}
					currentClient.Connection.Write(bytesTmp)
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

//BroadcastMessage 广播消息
func BroadcastMessage(message *proto.MessageWarpper) {
	bytes, err := message.Marshal()
	if err != nil {
		logger.Info("err", err)
		return
	}

	messageType := message.Type
	if messageType != proto.COMMUNICATION_TYPE_ClientLogin ||
		messageType != proto.COMMUNICATION_TYPE_ClientLogout ||
		messageType != proto.COMMUNICATION_TYPE_ClientReceived {
		logger.Info("invalid message type")
		return
	}

	pool.Locker.Lock()
	for i := 0; i < len(pool.Clients); i++ {
		pool.Clients[i].Connection.Write(bytes)
	}
	pool.Locker.Unlock()
}

//MessageInterpreter 获取包装壳
func MessageInterpreter(bytes []byte) (msg *proto.MessageWarpper) {
	warpper := &proto.MessageWarpper{}
	err := warpper.Unmarshal(bytes)
	if err != nil {
		logger.Info("", err)
		return nil
	}
	return warpper
}

//Login 用户登入
func Login(conn *net.TCPConn) (client *Client, message *proto.MessageWarpper, err error) {
	if err := conn.SetDeadline(time.Now().Add(10 * time.Second)); err != nil {
		return nil, nil, err
	}
	bytes := make([]byte, MESSAGE_MAX_LENGTH)
	var loginMessage proto.MessageWarpper
	for i := 0; i < 3; i++ {
		count, err := conn.Read(bytes)
		if err != nil {
			return nil, nil, err
		}
		if count >= MESSAGE_MAX_LENGTH {
			return nil, nil, errors.New("message too large")
		}
		msg := MessageInterpreter(bytes[:count])
		if msg.Type == proto.COMMUNICATION_TYPE_LoginRequest {
			break
		}
		if i == 2 {
			LoginFault(client)
			return nil, nil, errors.New("wrong login message")
		}
		time.Sleep(time.Second * 3)
		continue
	}
	clientIndex := getClientIndex()
	client = &Client{
		Address:    conn.RemoteAddr(),
		Connection: conn,
		Id:         clientIndex,
		Name:       loginMessage.MessageLoginRequest.Name,
	}
	pool.Insert(client)
	LoginSucceed(client)
	BroadcastClientLogin(client)
	return client, &loginMessage, nil
}

//BroadcastClientLogin 广播用户登入
func BroadcastClientLogin(client *Client) error {
	clientLoginMessage := &proto.MessageWarpper{
		Type: proto.COMMUNICATION_TYPE_ClientLogin,
		MessageClientLogin: &proto.MessageClientLogin{
			Id:   client.Id,
			Name: client.Name,
		},
	}
	BroadcastMessage(clientLoginMessage)
	return nil
}

//LoginSucceed 用户登入成功
func LoginSucceed(client *Client) error {
	loginResponseMessage := &proto.MessageWarpper{
		Type: proto.COMMUNICATION_TYPE_LoginResponse,
		MessageLoginResponse: &proto.MessageLoginResponse{
			Succeed: true,
		},
	}
	bytes, err := loginResponseMessage.Marshal()
	if err != nil {
		return err
	}
	client.Connection.Write(bytes)
	return nil
}

//LoginFault 用户登入失败
func LoginFault(client *Client) error {
	loginResponseMessage := &proto.MessageWarpper{
		Type: proto.COMMUNICATION_TYPE_LoginResponse,
		MessageLoginResponse: &proto.MessageLoginResponse{
			Succeed: false,
		},
	}
	bytes, err := loginResponseMessage.Marshal()
	if err != nil {
		return err
	}
	client.Connection.Write(bytes)
	return nil
}

//Logout 用户登出
func Logout(client *Client) {
	BroadcastClientLogout(client)
	LogoutSucceed(client)
	pool.Remove(client)
}

//BroadcastClientLogout 广播用户登出
func BroadcastClientLogout(client *Client) error {
	logoutMessage := &proto.MessageWarpper{
		Type: proto.COMMUNICATION_TYPE_ClientLogout,
		MessageClientLogout: &proto.MessageClientLogout{
			Id:   client.Id,
			Name: client.Name,
		},
	}
	BroadcastMessage(logoutMessage)
	return nil
}

//LogoutSucceed 用户登出成功
func LogoutSucceed(client *Client) error {
	logoutMessage := &proto.MessageWarpper{
		Type: proto.COMMUNICATION_TYPE_LogoutResponse,
		MessageLogoutResponse: &proto.MessageLogoutResponse{
			Succeed: true,
		},
	}
	bytes, err := logoutMessage.Marshal()
	if err != nil {
		return err
	}
	client.Connection.Write(bytes)
	return nil
}

//LogoutFault 用户登出失败
func LogoutFault(client *Client) error {
	logoutMessage := &proto.MessageWarpper{
		Type: proto.COMMUNICATION_TYPE_LogoutResponse,
		MessageLogoutResponse: &proto.MessageLogoutResponse{
			Succeed: false,
		},
	}
	bytes, err := logoutMessage.Marshal()
	if err != nil {
		return err
	}
	client.Connection.Write(bytes)
	return nil
}

func Init() {
	InitConnectionPool()
	logger = log.Logger
}

func InitConnectionPool() {
	pool = &ConnectionPool{}
}
