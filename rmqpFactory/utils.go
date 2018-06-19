package rmqpFactory
import (
	"github.com/streadway/amqp"
	"os"
	// "log"

)
type rmqConn struct {
	Conn *amqp.Connection
	Channel *amqp.Channel
}

var connection rmqConn

func GetNewRMQConn() rmqConn {
	// connection.getNewConnection()
	// connection.getNewChannel()
	rmqUrl, defaultUrl := os.LookupEnv("CLOUDAMQP_URL")
	if !defaultUrl {
		rmqUrl = "amqp://"
	}
	conn, err := amqp.Dial(rmqUrl)
	if err != nil {
		panic("cannot connect")
	}
	failOnError(err, "Failed to connect to RabbitMQ")
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	connection := rmqConn{conn,ch}
	
	return connection
}

func (rmq rmqConn) GetConnection() *amqp.Connection {
	return rmq.Conn
}

func (rmq rmqConn) GetChannel()*amqp.Channel {
	return rmq.Channel
}

func (rmq rmqConn) getNewConnection() {
	rmqUrl, defaultUrl := os.LookupEnv("CLOUDAMQP_URL")
	if !defaultUrl {
		rmqUrl = "amqp://"
	}
	var err error
	rmq.Conn, err = amqp.Dial(rmqUrl)
	if err != nil {
		panic("cannot connect")
	}
}

func (rmq rmqConn) getNewChannel() {
	var err error
	rmq.Channel, err = rmq.Conn.Channel()
	failOnError(err, "Failed to open a channel")
}
