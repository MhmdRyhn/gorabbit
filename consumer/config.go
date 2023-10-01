package consumer

import (
	"os"
)

type ConsumerConf struct {
	MQServerUrl string
}

func NewConsumerConf() ConsumerConf {
	conf := ConsumerConf{}
	conf.load()
	return conf
}

func (conf *ConsumerConf) load() {
	conf.MQServerUrl = os.Getenv("MQ_SERVER_URL")
}
