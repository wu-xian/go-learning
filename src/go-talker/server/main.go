package main

import (
	"encoding/binary"
	"bytes"
	"time"
	"os"

	"net"

	cli "github.com/urfave/cli"
)

const VERSION = "0.0.1"

type Message struct{
	Content string
	Time uint64
}

type Client struct{
	Name string
	Connection net.Conn
}

func main() {
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

func startAction(ctx cli.Context) {
	listener, err := net.Listen("tcp", ":34567")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		//conn.
	}
}

func MessageToBytes(msg Message) (byte[],error){
	buf := &bytes.Buffer{}
	err := binary.Write(buf,binary.BigEndian,&msg)
	if err!= nil{
		return _,err
	}
	return buf.Bytes()
}

func BytesToMessage(bytess []byte) (Message,error){
	msg := &Message{}
	buf := bytess.Buffer
err:=	binary.Read(buf,binary.BigEndian,&msg)
if err!= nil{
	return  _,err
}
return msg;
}