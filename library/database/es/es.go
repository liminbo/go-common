package es

import (
	"fmt"
	elastic "gopkg.in/olivere/elastic.v5"
	"time"
)

type Config struct {
	Addr          string          // write data source name.
	Timeout time.Duration
}

type ElasticSearch struct {
	Client *elastic.Client
	Timeout time.Duration
}

func New(c *Config) (es *ElasticSearch){
	client, err := elastic.NewClient(elastic.SetURL(c.Addr), elastic.SetSniff(false))
	if err != nil {
		panic(fmt.Sprintf("es:集群连接失败, cluster: %v", err))
	}
	es = &ElasticSearch{
		Client:  client,
		Timeout: time.Duration(c.Timeout),
	}
	return
}