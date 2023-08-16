package main

import (
	"encoding/json"
	"influx-test/common"
	"influx-test/data"
	"influx-test/tasks"
	"log"
	"net/http"

	_ "github.com/influxdata/influxdb/client/v2"
	//"sync"
)

const (
	mydb        = "mydb"
	server_port = "32325"
)

const (
	maxChanLen = 7
)

type Client struct {
	HttpClnt *common.ClntCtx
	//Chan     chan common.TestRes
	//WG       sync.WaitGroup
}

func (c *Client) Close() {
	// Close client resources
	if err := c.HttpClnt.Client.Close(); err != nil {
		log.Fatal(err)
	}
	//close(c.Chan)
}
func (c *Client) TestCQ() common.TestRes {

	data.Init(c.HttpClnt.Client)
	tr := tasks.DoTask(c.HttpClnt.Client, &tasks.CQTask{})
	return tr
}

func (c *Client) TestFunc() common.TestRes {

	data.Init(c.HttpClnt.Client)
	tr := tasks.DoTask(c.HttpClnt.Client, &tasks.FuncTask{})
	return tr

}

func (c *Client) TestHint() common.TestRes {

	data.Init(c.HttpClnt.Client)
	tr := tasks.DoTask(c.HttpClnt.Client, &tasks.HintTask{})
	return tr

}
func (c *Client) TestMathOpt() common.TestRes {

	data.Init(c.HttpClnt.Client)
	tr := tasks.DoTask(c.HttpClnt.Client, &tasks.MathOptTask{})
	return tr

}

func (c *Client) TestInterface() common.TestRes {

	data.Init(c.HttpClnt.Client)
	tr := tasks.DoTask(c.HttpClnt.Client, &tasks.OtherTask{})
	return tr

}

func (c *Client) TestMgDb() common.TestRes {

	data.Init(c.HttpClnt.Client)
	tr := tasks.DoTask(c.HttpClnt.Client, &tasks.MgDbTask{})
	return tr

}
func (c *Client) TestQuery() common.TestRes {

	data.Init(c.HttpClnt.Client)
	tr := tasks.DoTask(c.HttpClnt.Client, &tasks.QueryTask{})
	return tr

}
func (c *Client) TestShow() common.TestRes {

	data.Init(c.HttpClnt.Client)
	tr := tasks.DoTask(c.HttpClnt.Client, &tasks.ShowTask{})
	return tr

}

func (c *Client) TestAll() []common.TestRes {
	var trs []common.TestRes
	trs = append(trs, c.TestCQ())
	trs = append(trs, c.TestFunc())
	trs = append(trs, c.TestHint())
	trs = append(trs, c.TestMathOpt())
	trs = append(trs, c.TestInterface())
	trs = append(trs, c.TestMgDb())
	trs = append(trs, c.TestQuery())
	trs = append(trs, c.TestShow())

	return trs
}

func HttpHandle(c *Client) {
	server := &http.Server{Addr: ":" + server_port}
	for {
		http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
			if err := r.ParseForm(); err != nil {
				http.Error(w, "ParseForm error", http.StatusBadRequest)
				return
			}
			if r.Method != "POST" {
				return
			}

			opt := r.FormValue("opt")

			w.Header().Set("Content-Type", "application/json")

			switch opt {
			case "cq":
				resdata, _ := json.MarshalIndent(c.TestCQ(), "", " ")
				w.Write(resdata)
			case "func":
				resdata, _ := json.MarshalIndent(c.TestFunc(), "", " ")
				w.Write(resdata)

			case "hint":
				resdata, _ := json.MarshalIndent(c.TestHint(), "", " ")
				w.Write(resdata)

			case "interface":
				resdata, _ := json.MarshalIndent(c.TestInterface(), "", " ")
				w.Write(resdata)

			case "mathopt":
				resdata, _ := json.MarshalIndent(c.TestMathOpt(), "", " ")
				w.Write(resdata)

			case "mgdb":
				resdata, _ := json.MarshalIndent(c.TestMgDb(), "", " ")
				w.Write(resdata)

			case "query":
				resdata, _ := json.MarshalIndent(c.TestQuery(), "", " ")
				w.Write(resdata)

			case "show":
				resdata, _ := json.MarshalIndent(c.TestShow(), "", " ")
				w.Write(resdata)

			case "all":
				resdata, _ := json.MarshalIndent(c.TestAll(), "", " ")
				w.Write(resdata)
			case "quit":
				//c.WG.Done()
				return
			default:
				http.Error(w, "invalid input,expect cq, func, hint, interface, mathopt, mgdb, query, show, all, quit", http.StatusBadRequest)
			}

		})
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}

	}

}

func main() {
	common.CheckEnv()
	log.Println("check success")
	var c = Client{
		HttpClnt: common.HttpClnt,
	}
	c.HttpClnt.Init()
	log.Println("connected influxdb success")
	defer c.Close()

	data.Init(c.HttpClnt.Client)
	log.Println("init influxdb success")
	HttpHandle(&c)
}
