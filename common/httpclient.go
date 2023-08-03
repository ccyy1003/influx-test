package common

import (
	_ "encoding/json"

	"log"

	"github.com/influxdata/influxdb/client/v2"
)

type ClntCtx struct {
	Addr     string
	Username string
	Password string
	Client   client.Client
}

var HttpClnt = new(ClntCtx)

func (c *ClntCtx) Init(addr string, username string, password string) error {
	c.Addr = addr
	c.Username = username
	c.Password = password

	cc, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     addr,
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Fatal(err)
	}
	c.Client = cc
	return nil
}
