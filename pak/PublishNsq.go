package pak

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/nsqio/go-nsq"
	"github.com/satori/go.uuid"
	"os"
	"sync"
	"time"
)

var producer *nsq.Producer
var nsqUrl string

func NsqDo(opt *ConnectOpt) {
	nsqUrl = fmt.Sprintf("%s:%s", opt.Host, opt.Port)
	if opt.LineCommand == true {
		producer, _ = nsq.NewProducer(nsqUrl, nsq.NewConfig())
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

				producer.Publish(opt.TopicName, sendByte)
			}

		}
	} else {
		waitGroup := sync.WaitGroup{}
		for i := uint64(0); i < opt.ClientNum; i++ {
			waitGroup.Add(1)
			go publishToNSQ(opt.TopicName, i, opt.ConnectNum, opt.MessageConnect, &waitGroup)
		}
		waitGroup.Wait()
		return
	}

}

// 初始化生产者
func InitProducer(str string) nsq.Producer {
	var err error
	producer, err = nsq.NewProducer(str, nsq.NewConfig())
	if err != nil {
		fmt.Println("NSQ 链接错误")
		panic(err)
	}
	return *producer
}

//发布消息
func publishToNSQ(topic string, i uint64, maxNum uint64, msgText string, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	prod := InitProducer(nsqUrl)
	for cNum := uint64(0); cNum < maxNum; cNum++ {
		msg := MessageBody{Id: time.Now().UnixNano(), Body: msgText, ConnectNum: cNum, NodeNum: i}
		message, _ := json.Marshal(msg)
		prod.Publish(topic, message)
	}

}
