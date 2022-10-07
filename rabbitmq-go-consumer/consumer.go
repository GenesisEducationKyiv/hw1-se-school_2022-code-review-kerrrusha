package main

import (
	"fmt"
	"strings"

	"github.com/streadway/amqp"
)

func isErrorLog(logMsg string) bool {
	return strings.Contains(logMsg, "[ERROR]")
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer func(ch *amqp.Channel) {
		err := ch.Close()
		if err != nil {
			panic(err)
		}
	}(ch)

	msgs, err := ch.Consume(
		"LogQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			msg := string(d.Body)
			if isErrorLog(msg) {
				fmt.Println(msg)
			}
		}
	}()
	<-forever
}
