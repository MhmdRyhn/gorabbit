package producer

import (
	"os"
)

type ProducerConf struct {
	MQServerUrl string
}

func NewProducerConf() ProducerConf {
	conf := ProducerConf{}
	conf.load()
	return conf
}

func (conf *ProducerConf) load() {
	conf.MQServerUrl = os.Getenv("MQ_SERVER_URL")
}
