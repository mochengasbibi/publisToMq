package pak

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
)

//返回一个mqtt Client
func getMqttContent(opt *ConnectOpt) (*mqtt.Client, error) {
	taskId := 0
	//设置一个调handler
	var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("Pub Client Topic : %s \n", msg.Topic())
		fmt.Printf("Pub Client msg : %s \n", msg.Payload())
	}
	//设置连接参数
	clinetOptions := mqtt.NewClientOptions().AddBroker("tcp://127.0.0.1:1883").SetUsername("pub").SetPassword("pub")
	//设置客户端ID
	clinetOptions.SetClientID(fmt.Sprintf("go Publish client example： %d-%d", taskId, time.Now().Unix()))
	//设置handler
	clinetOptions.SetDefaultPublishHandler(messagePubHandler)
	//设置连接超时
	clinetOptions.SetConnectTimeout(time.Duration(60) * time.Second)
	//创建客户端连接
	client := mqtt.NewClient(clinetOptions)
	return &client, nil
}
