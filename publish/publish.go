package publish

import (
	"fmt"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	_ "github.com/lib/pq"
)

// Connect to MQTT service
func connMQTT(broker, user, passwd string) (bool, MQTT.Client) {
	opts := MQTT.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetUsername("mqttbroker")
	opts.SetPassword("123")

	mc := MQTT.NewClient(opts)
	if token := mc.Connect(); token.Wait() && token.Error() != nil {
		return false, mc
	}

	return true, mc
}

// make an announcement
func Publish(devideid, value string) {
	// pub username and password
	b, mc := connMQTT("tcp://192.168.15.128:1883", "pub", "aaabbb")
	if !b {
		fmt.Println("pub connMQTT failed")
		return
	}
	mc.Publish("post", 0x00, true, devideid+"Temperature: "+value)
}
