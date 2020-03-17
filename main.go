//Nsq发送测试
package main

import (
	"flag"
	"fmt"
	uuid "github.com/satori/go.uuid"
	mq "ptq/MqCommon"
	"strings"
)

var host = flag.String("host", "127.0.0.1", "NSQ服务地址")
var port = flag.String("port", "4150", "NSQ服务端口")
var msgText = flag.String("text", "", "uuid.string()")
var userName = flag.String("user", "", "用户名")
var passWord = flag.String("password", "", "密码")
var connectNum = flag.Uint64("c", 1, "每个并发需发送次数")
var clientNum = flag.Uint64("n", 1, "并发个数")
var qosNum = flag.Uint("q", 0, "Qos:服务质量")
var helpMsg bool
var lineCommand bool
var nsqUrl string
var topicName = flag.String("t", "test", "topic(话题)名称")
var channelName = flag.String("channel", "test", "channel(key)名称")
var mqType = flag.String("mq", "NSQ", "mq类型:nsq;mqtt;rabbit")

func init() {
	flag.BoolVar(&helpMsg, "help", false, "帮助")
	flag.BoolVar(&helpMsg, "h", false, "帮助")
	flag.BoolVar(&lineCommand, "line", false, "发送用户输入的信息,输入 stop 结束")
	flag.BoolVar(&lineCommand, "l", false, "发送用户输入的信息,输入 stop 结束")
	if strings.Trim(*msgText, " ") == " " {
		*msgText = uuid.NewV4().String()
	}
	if *qosNum > uint(2) {
		*qosNum = 0
	}
}

// 主函数
func main() {
	flag.Parse()
	flagNum := flag.NFlag()
	if flagNum == 0 || helpMsg == true {
		//TODO 多类型的 MQ
		//TODO channel 获得 goroutine 的返回  (主要是 channel goroutine 和 WaitGroup)
		//TODO 可以做成 接口并发测试 类似apache 的ab(压力测试 )
		flag.PrintDefaults()
		return
	} else {
		mqTypeStr := "NSQ"
		if strings.ToUpper(*mqType) == "RABBIT" {
			mqTypeStr = "RABBIT"
		} else if strings.ToUpper(*mqType) == "MQTT" {
			mqTypeStr = "MQTT"
		}
		url := fmt.Sprintf("%s:%s", *host, *port)
		opt := mq.ConnectOpt{
			Host:           *host,
			Port:           *port,
			MessageConnect: *msgText,
			ConnectNum:     *connectNum,
			ClientNum:      *clientNum,
			Qos:            *qosNum,
			UserName:       *userName,
			PassWord:       *passWord,
			TopicName:      *topicName,
			ChannelName:    *channelName,
			LineCommand:    lineCommand,
			MqType:         mqTypeStr,
			Url:            url,
		}
		opt.Connect()
	}

}
