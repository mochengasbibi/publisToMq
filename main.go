//Nsq发送测试
package main

import (
	"flag"
	"fmt"
	"github.com/nsqio/go-nsq"
	mq "ptq/MqCommon"
)

var host = flag.String("host", "127.0.0.1", "NSQ服务地址")
var port = flag.String("port", "4150", "NSQ服务端口")
var msgText = flag.String("text", "this is test", "发送的文本")
var connectNum = flag.Uint64("c", 1, "每个并发需发送次数")
var clientNum = flag.Uint64("n", 1, "并发个数")
var helpMsg bool
var lineCommand bool
var producer *nsq.Producer
var nsqUrl string
var topicName = flag.String("t", "test", "topic(话题)名称")

func init() {
	flag.BoolVar(&helpMsg, "help", false, "帮助")
	flag.BoolVar(&helpMsg, "h", false, "帮助")
	flag.BoolVar(&lineCommand, "line", false, "发送用户输入的信息,输入 stop 结束")
	flag.BoolVar(&lineCommand, "l", false, "发送用户输入的信息,输入 stop 结束")

}

// 主函数
func main() {
	flag.Parse()
	flagNum := flag.NFlag()
	if flagNum == 0 || helpMsg == true {
		//TODO 多类型的 MQ
		//TODO channel 获得 goroutine 的返回  (主要是 channel goroutine 和 WaitGroup)
		//TODO 可以做成 接口并发测试 类似apache 的ab(压力测试 )
		//fmt.Println("一个发送消息到NSQ的命令行工具,练习一下flag和go-nsq包的使用 \n  ")
		flag.PrintDefaults()
		return
	} else {
		opt := mq.ConnectOpt{
			Host: *host,
			Port: *port,
		}
	}

}
