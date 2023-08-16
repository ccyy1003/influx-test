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

func (c *ClntCtx) Init() error {
	c.Addr = Addr
	c.Username = Username
	c.Password = Password

	cc, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     c.Addr,
		Username: c.Username,
		Password: c.Password,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(c.Addr, c.Username, c.Password)
	c.Client = cc
	return nil
}
