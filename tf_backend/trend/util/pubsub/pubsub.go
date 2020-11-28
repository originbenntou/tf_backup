package pubsub

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"time"

	"github.com/TrendFindProject/tf_backend/trend/constant"
)

func PublishSingleGoroutine(ctx context.Context, msg string, sid uint64) error {
	client, err := pubsub.NewClient(ctx, constant.ProjectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}

	t := client.Topic(constant.TopicID)

	t.PublishSettings.Timeout = 300 * time.Second
	t.PublishSettings.NumGoroutines = 1

	result := t.Publish(ctx, &pubsub.Message{
		// 何故かAttributesが送れない...
		Data: []byte(fmt.Sprintf("%s,%d", msg, sid)),
	})

	t.Stop()

	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		return fmt.Errorf("Get: %v", err)
	}

	fmt.Printf("Published a message; msg ID: %v, msg: %s, search_id: %d\n", id, msg, sid)

	return nil
}
