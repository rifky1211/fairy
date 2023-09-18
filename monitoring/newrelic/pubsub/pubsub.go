package pubsub

import (
	"context"

	"github.com/newrelic/go-agent/v3/newrelic"
	"gitlab.com/nabati/superapp/fairy/pubsub"
)

type client struct {
	dialect string
	pubsub  pubsub.PubSub
}

// New to create new newrelic plugin for pubsub.
func New(d string, ps pubsub.PubSub) pubsub.PubSub {
	return &client{
		dialect: d,
		pubsub:  ps,
	}
}

// Publish to publish message.
func (c *client) Publish(ctx context.Context, topic string, data interface{}, durable bool) error {
	segment := newrelic.MessageProducerSegment{
		StartTime:       newrelic.FromContext(ctx).StartSegmentNow(),
		Library:         c.dialect,
		DestinationType: newrelic.MessageTopic,
		DestinationName: topic,
	}
	defer segment.End()

	return c.pubsub.Publish(ctx, topic, data, false)
}

// Subscribe to subscribe.
func (c *client) Subscribe(ctx context.Context, topic string) (interface{}, error) {
	return c.pubsub.Subscribe(ctx, topic)
}

// Close to close connection.
func (c *client) Close() error {
	return c.pubsub.Close()
}
