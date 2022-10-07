package logger

import "github.com/streadway/amqp"

type RabbitMQLogger struct {
	*LogBuilder
}

func (c *RabbitMQLogger) log(logMsg string) {
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

	_, err = ch.QueueDeclare(
		"LogQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	err = ch.Publish(
		"",
		"LogQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(logMsg),
		},
	)
	if err != nil {
		panic(err)
	}
}

func (c *RabbitMQLogger) Info(msg string) {
	c.log(c.BuildInfo(msg))
}

func (c *RabbitMQLogger) Debug(msg string) {
	c.log(c.BuildDebug(msg))
}

func (c *RabbitMQLogger) Error(msg string) {
	c.log(c.BuildError(msg))
}

func CreateRabbitMQLogger() Logger {
	return &RabbitMQLogger{}
}
