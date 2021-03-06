package nats

import stan "github.com/nats-io/go-nats-streaming"

//go:generate moq -out mock_connection.go . Connection
// Connection  defines the behaviour required for the Nats connection
type Connection interface {
	QueueSubscribe(
		subject,
		qgroup string,
		cb stan.MsgHandler,
		opts ...stan.SubscriptionOption) (stan.Subscription, error)
	Publish(subj string, data []byte) error
}

//go:generate moq -out mock_connection_pool.go . ConnectionPool
type ConnectionPool interface {
	GetConnection(server, clusterID string) (Connection, error)
}

//go:generate moq -out mock_subscription.go . Subscription
type Subscription interface {
	stan.Subscription
}
