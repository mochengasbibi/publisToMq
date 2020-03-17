package pak

import (
	"bufio"
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	uuid "github.com/satori/go.uuid"
	"os"
	"sync"
	"time"
)

var mqttUrl string

func MqttDo(opt *ConnectOpt) {
	if opt.LineCommand == true {
		client, _ := getMqttContent(opt)
		defer client.Disconnect(250)
		running := true
		//读取控制台输入
		reader := bufio.NewReader(os.Stdin)
		for running {

			data, _, _ := reader.ReadLine()
			command := string(data)
			if command == "stop" {
				running = false
				fmt.Println("^_^ 88")
				return
			} else {
				if len(command) == 0 {
					command = uuid.NewV4().String()
				}
				msg := MessageBody{time.Now().UnixNano(), command, 0, 0}
				sendByte, _ := json.Marshal(msg)
				client.Publish(opt.TopicName, byte(opt.Qos), false, sendByte)
			}

		}
	} else {
		waitGroup := sync.WaitGroup{}
		for i := uint64(0); i < opt.ClientNum; i++ {
			waitGroup.Add(1)
			go mqttConnPubMsgTask(opt, i, &waitGroup)
		}
		waitGroup.Wait()
		return
	}
}

//返回一个mqtt Client
func getMqttContent(opt *ConnectOpt) (mqtt.Client, error) {
	mqttUrl = fmt.Sprintf("tcp://%s:%s", opt.Host, opt.Port)
	//设置一个调handler
	//var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	//
	//}
	//设置连接参数
	clinetOptions := mqtt.NewClientOptions()
	//设置 用户名 密码
	if opt.UserName != "" && opt.PassWord != "" {
		clinetOptions.SetUsername(opt.UserName).SetPassword(opt.PassWord)
	}
	//设置客户端ID
	clinetOptions.SetClientID(uuid.NewV4().String())
	//设置handler
	//clinetOptions.SetDefaultPublishHandler(messagePubHandler)
	//设置连接超时
	clinetOptions.SetConnectTimeout(time.Duration(60) * time.Second)
	//创建客户端连接
	client := mqtt.NewClient(clinetOptions)
	return client, nil
}

func mqttConnPubMsgTask(opt *ConnectOpt, nodeNum uint64, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	client, _ := getMqttContent(opt)
	maxNum := opt.ConnectNum
	for i := uint64(0); i < maxNum; i++ {
		msg := MessageBody{Id: time.Now().UnixNano(), Body: opt.MessageConnect, ConnectNum: i, NodeNum: nodeNum}
		text, _ := json.Marshal(msg)
		token := client.Publish(opt.TopicName, byte(opt.Qos), false, text)
		//fmt.Printf("[Pub] end publish msg to mqtt broker, taskId: %d, count: %d, token : %s \n", taskId, i, token)
		token.Wait()
	}
	client.Disconnect(250)
}
