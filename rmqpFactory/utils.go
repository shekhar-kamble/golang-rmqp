package rmqpFactory
import (
	"github.com/streadway/amqp"
	"os"
	// "log"

)
type rmqConn struct {
	conn *amqp.Connection
	channel *amqp.Channel
}

// var connection rmqConn

func GetNewRMQConn() (connection rmqConn) {
	connection.getNewConnection()
	connection.getNewChannel()
	return connection
}

func (rmq rmqConn) GetConnection() *amqp.Connection {
	return rmq.conn
}

func (rmq rmqConn) GetChannel()*amqp.Channel {
	return rmq.channel
}

func (rmq rmqConn) getNewConnection() {
	rmqUrl, defaultUrl := os.LookupEnv("CLOUDAMQP_URL")
	if !defaultUrl {
		rmqUrl = "amqp://"
	}
	var err error
	rmq.conn, err = amqp.Dial(rmqUrl)
	if err != nil {
		panic("cannot connect")
	}
}

func (rmq rmqConn) getNewChannel() {
	var err error
	rmq.channel, err = rmq.conn.Channel()
	failOnError(err, "Failed to open a channel")
	exName, ok := os.LookupEnv("EXCHANGE_NAME") 
	if !ok {
		exName = "common-exchange"
	}
	err = rmq.channel.ExchangeDeclare(
		exName,			// name
		"direct",	// type
		true,			// durable
		false,		// auto-deleted
		false,		// internal
		false,		// no-wait
		nil,			// arguments
	)

	if err!=nil{
		panic(err)
	}
}
