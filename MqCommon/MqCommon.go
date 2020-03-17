package MqCommon

import "strings"

type MessageBody struct {
	Id         int64  `json:id`
	Body       string `json:"text"`
	NodeNum    uint64 `json:nodeNum`
	ConnectNum uint64 `json:connectNum`
}
type ConnectOpt struct {
	Host           string //主机地址
	Port           string //端口
	MqType         string //mq 类型,nsq;mqtt;rabbit
	MessageConnect string //发送的内容
	UserName       string //用户名
	PassWord       string //密码
	Qos            int    // Quality of Service,服务质量 默认为0
	LineCommand    bool   //是否是命令行模式,等待用户输入后,发送
	ConnectNum     uint64 //每个并发需发送次数
	ClientNum      uint64 //并发个数
	TopicName      string // 话题/队列名称
	ChannelName    string // 通道/RtKey 值
}

func (opt ConnectOpt) Connect() {

	if strings.ToUpper(opt.MqType) == "MQTT" {

	} else if strings.ToUpper(opt.MqType) == "RABBIT" {

	} else {

	}

}
