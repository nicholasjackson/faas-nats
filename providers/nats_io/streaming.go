package nats

import (
	"fmt"

	"github.com/DataDog/datadog-go/statsd"
	hclog "github.com/hashicorp/go-hclog"
	stan "github.com/nats-io/go-nats-streaming"
	"github.com/nicholasjackson/faas-nats/providers"
)

type StreamingProvider struct {
	Name      string
	Server    string    `hcl:"server"`
	ClusterID string    `hcl:"cluster_id"`
	Queue     string    `hcl:"queue"`
	AuthBasic AuthBasic `hcl:"auth_basic,block"`
	AuthMTLS  AuthMTLS  `hcl:"auth_mtls,block"`

	connection   Connection
	subscription stan.Subscription
	stats        *statsd.Client
	logger       hclog.Logger
	msgChannel   chan *providers.Message
}

type AuthBasic struct {
	User     string `hcl:"user"`
	Password string `hcl:"password"`
}

type AuthMTLS struct {
	TLSClientKey    string `hcl:"tls_client_key"`
	TLSClientCert   string `hcl:"tls_client_cert"`
	TLSClientCACert string `hcl:"tls_client_cacert"`
}

func (sp *StreamingProvider) Type() string {
	return "nats_queue"
}

func (sp *StreamingProvider) Setup(cp providers.ConnectionPool, logger hclog.Logger, stats *statsd.Client) error {
	pool := cp.(ConnectionPool)

	conn, err := pool.GetConnection(sp.Server, sp.ClusterID)
	if err != nil {
		stats.Incr("connection.nats.failed", nil, 1)
		logger.Error("Unable to connect to nats server", "error", err)
	}

	stats.Incr("connection.nats.created", nil, 1)
	logger.Debug("Created connection for", sp.Server, sp.ClusterID)

	sp.connection = conn
	sp.stats = stats
	sp.logger = logger

	return nil
}

func (sp *StreamingProvider) Listen() (<-chan *providers.Message, error) {
	qGroup := fmt.Sprintf("%s-%s", sp.Queue, sp.Name)
	subscription, err := sp.connection.QueueSubscribe(sp.Queue, qGroup, sp.messageHandler)
	if err != nil {
		sp.stats.Incr("subscription.nats.failed", nil, 1)
		sp.logger.Error("Failed to create subscription for", sp.Queue)
		return nil, err
	}

	sp.stats.Incr("subscription.nats.created", nil, 1)
	sp.logger.Debug("Created subscription for", sp.Queue)
	sp.subscription = subscription

	sp.msgChannel = make(chan *providers.Message)

	return sp.msgChannel, nil
}

func (sp *StreamingProvider) Stop() error {
	sp.subscription.Close()
	sp.subscription.Unsubscribe()

	return nil
}

func (sp *StreamingProvider) messageHandler(msg *stan.Msg) {
	m := &providers.Message{}
	m.Data = msg.Data
	m.Redelivered = msg.Redelivered
	m.Timestamp = msg.Timestamp
	m.Sequence = msg.Sequence

	sp.msgChannel <- m
}
