package config

const (
	// AsyncTransferEnable : Whether to enable file asynchronous transfer (default synchronous)
	AsyncTransferEnable = false
	// TransExchangeName : File transfer switch
	TransExchangeName = "uploadserver.trans"
	// TransOSSQueueName : queue name
	TransOSSQueueName = "uploadserver.trans.oss"
	// TransOSSErrQueueName : The queue name written to another queue after a failed oss transfer
	TransOSSErrQueueName = "uploadserver.trans.oss.err"
	// TransOSSRoutingKey : routingkey
	TransOSSRoutingKey = "oss"
)

var (
	// RabbitURL : The entry url
	RabbitURL = "amqp://guest:guest@127.0.0.1:5672/"
)
