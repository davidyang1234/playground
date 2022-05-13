package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"cloud.google.com/go/pubsub"
)

var (
	topic *pubsub.Topic

	// Messages received by this instance.
	messagesMu sync.Mutex
	messages   []string

	// token is used to verify push requests.
	// token = mustGetenv("PUBSUB_VERIFICATION_TOKEN")
)

func main() {
	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, mustGetenv("GOOGLE_CLOUD_PROJECT"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	topicName := mustGetenv("PUBSUB_TOPIC")
	topic = client.Topic(topicName)

	exists, err := topic.Exists(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if !exists {
		log.Printf("Topic %v doesn't exist - creating it", topicName)
		_, err = client.CreateTopic(ctx, topicName)
		if err != nil {
			log.Fatal(err)
		}
	}

	msg := &pubsub.Message{
		Data: []byte("This is a test msg"),
	}

	if _, err := topic.Publish(ctx, msg).Get(ctx); err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Message published.")
	// client, err := redis.Dial("tcp", "localhost:6379")
	// if err != nil {
	// 	log.Println(err)
	// }

	// psc := redis.PubSubConn{Conn: client}
	// client.Do("config", "set", "notify-keyspace-events", "Ex")
	// psc.PSubscribe("__key*__:*")
	// for {
	// 	switch v := psc.Receive().(type) {
	// 	case redis.Message:
	// 		fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
	// 	case redis.Subscription:
	// 		fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
	// 	case error:
	// 		fmt.Println(v)
	// 	}
	// }

	// pubsub := client.PSubscribe("*")
	// defer pubsub.Close()
	// fmt.Println("start listening")
	// for msg := range pubsub.Channel() {
	// 	fmt.Printf("channel=%s message=%s/n", msg.Channel, msg.Payload)
	// }
}

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("%s environment variable not set.", k)
	}
	return v
}
