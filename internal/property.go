package internal

import (
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"strings"
	"time"
)

func SubscribeProperty(client mqtt.Client) {
	//订阅消息
	client.Subscribe("up/property/+/+", 0, func(client mqtt.Client, message mqtt.Message) {
		topics := strings.Split(message.Topic(), "/")
		pid := topics[2]
		id := topics[3]

		var properties map[string]interface{}
		err := json.Unmarshal(message.Payload(), &properties)
		if err != nil {
			//log
			return
		}

		tm := time.Now()
		_ = Write(pid, id, properties, tm.UnixMilli())
		//influx.Insert(pid, id, properties, tm)
	})
}
