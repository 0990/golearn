package main

import (
	"fmt"
	"github.com/nats-io/go-nats"
	"time"
)

const (
	testStr = `{"area_id":10,"log_info":"eyJsb2dvbklQIjoiMTAuMjI1LjEwLjI0NSIsInVzZXJJRCI6MjN9","log_time":1533645769,"login_from":1,"op_type":101,"param1":1,"param2":10,"tbl_name":"tbl_login","user_account":"hello","user_level":1}`
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}

	nc.Subscribe("log_topic", func(m *nats.Msg) {
		fmt.Printf("Received hello message:%s\n", string(m.Data))
	})

	nc.Publish("log_topic", []byte("sdfsdf"))

	time.Sleep(time.Second * 5)
}