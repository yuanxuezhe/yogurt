package nats

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

const (
	url = "nats://127.0.0.1:4222"
)

var (
	nc  *nats.Conn
	err error
)

func init() {
	if nc, err = nats.Connect(url); checkErr(err) {
		//
	}
}

//send message to server
func StartClient(subj string) {
	nc.Publish(subj, []byte("Sun"))
	nc.Publish(subj, []byte("Rain"))
	nc.Publish(subj, []byte("Fog"))
	nc.Publish(subj, []byte("Cloudy"))
}

//receive message
func StartServer(subj, name string) {
	go sync(nc, subj, name)
	go async(nc, subj, name)
}

func async(nc *nats.Conn, subj, name string) {
	nc.Subscribe(subj, func(msg *nats.Msg) {
		fmt.Println(name, "Received a message From Async : ", string(msg.Data))
	})
}

func sync(nc *nats.Conn, subj, name string) {
	subscription, err := nc.SubscribeSync(subj)
	checkErr(err)
	if msg, err := subscription.NextMsg(10 * time.Second); checkErr(err) {
		if msg != nil {
			fmt.Println(name, "Received a message From Sync : ", string(msg.Data))
		}
	} else {
		//handle timeout
	}

}

func checkErr(err error) bool {
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
