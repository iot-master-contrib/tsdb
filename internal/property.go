package internal

import (
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
	"strings"
	"time"
)

func SubscribeProperty() {
	//订阅消息
	mqtt.SubscribeJson("up/property/+/+", func(topic string, properties map[string]interface{}) {
		topics := strings.Split(topic, "/")
		pid := topics[2]
		id := topics[3]

		tm := time.Now()
		_ = Write(pid, id, properties, tm.UnixMilli())
		//influx.Insert(pid, id, properties, tm)
	})
}
