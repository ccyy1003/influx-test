package main

import (
	"fmt"
	_ "github.com/influxdata/influxdb/client/v2"
	"influx-test/common"
	"influx-test/data"
	"influx-test/tasks"
	"log"
	"os"
	//"sync"
)

const (
	mydb = "mydb"
	addr = "http://localhost:8086"
)

const (
	maxChanLen = 7
)

type Client struct {
	HttpClnt *common.ClntCtx
	//Chan     chan common.TestRes
	//WG       sync.WaitGroup
}

func (c *Client) Init() {
	c.HttpClnt.Init(addr, os.Getenv("INFLUX_USER"), os.Getenv("INFLUX_PWD"))
}

func (c *Client) Close() {
	// Close client resources
	if err := c.HttpClnt.Client.Close(); err != nil {
		log.Fatal(err)
	}
	//close(c.Chan)
}
func (c *Client) TestCQ() {

	data.Init(c.HttpClnt.Client)
	tr := tasks.DoTask(c.HttpClnt.Client, &tasks.CQTask{})
	tr.Print()

}

func (c *Client) TestFunc() {

	data.Init(c.HttpClnt.Client)
	tr := tasks.DoTask(c.HttpClnt.Client, &tasks.FuncTask{})
	tr.Print()

}

func (c *Client) TestHint() {

	data.Init(c.HttpClnt.Client)
	tr := tasks.DoTask(c.HttpClnt.Client, &tasks.HintTask{})
	tr.Print()

}
func (c *Client) TestMathOpt() {

	data.Init(c.HttpClnt.Client)
	tr := tasks.DoTask(c.HttpClnt.Client, &tasks.MathOptTask{})
	tr.Print()

}

func (c *Client) TestInterface() {

	data.Init(c.HttpClnt.Client)
	tr := tasks.DoTask(c.HttpClnt.Client, &tasks.OtherTask{})
	tr.Print()

}

func (c *Client) TestMgDb() {

	data.Init(c.HttpClnt.Client)
	tr := tasks.DoTask(c.HttpClnt.Client, &tasks.MgDbTask{})
	tr.Print()

}
func (c *Client) TestQuery() {

	data.Init(c.HttpClnt.Client)
	tr := tasks.DoTask(c.HttpClnt.Client, &tasks.QueryTask{})
	tr.Print()

}
func (c *Client) TestShow() {

	data.Init(c.HttpClnt.Client)
	tr := tasks.DoTask(c.HttpClnt.Client, &tasks.ShowTask{})
	tr.Print()

}

func (c *Client) TestAll() {
	c.TestCQ()
	c.TestFunc()
	c.TestHint()
	c.TestMathOpt()
	c.TestInterface()
	c.TestMgDb()
	c.TestQuery()
	c.TestShow()
}

func main() {
	var c = Client{
		HttpClnt: common.HttpClnt,
		//Chan:     make(chan common.TestRes, maxChanLen),
	}
	c.HttpClnt.Init(addr, os.Getenv("INFLUX_USER"), os.Getenv("INFLUX_PWD"))
	defer c.Close()
	//c.WG.Add(1)

	//finished := make(chan bool)
	// go func() {
	// 	for {
	// 		select {
	// 		case tr := <-c.Chan:
	// 			tr.Print()
	// 			//finished <- true
	// 		}
	// 	}
	// }()

	var opt string
	for {
		fmt.Println("Test Module(cq, func, hint, interface, mathopt, mgdb, query, show, all, quit) : ")
		fmt.Scanln(&opt)
		switch opt {
		case "cq":
			c.TestCQ()
		case "func":
			c.TestFunc()
		case "hint":
			c.TestHint()
		case "interface":
			c.TestInterface()
		case "mathopt":
			c.TestMathOpt()
		case "mgdb":
			c.TestMgDb()
		case "query":
			c.TestQuery()
		case "show":
			c.TestShow()
		case "all":
			c.TestAll()
		case "quit":
			//c.WG.Done()
			return
		default:
			fmt.Println("invalid input,expect cq, func, hint, interface, mathopt, mgdb, query, show, all, quit")
		}
		//<-finished
	}

	//c.WG.Wait()

}
