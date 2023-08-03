package tasks

import (
	"fmt"
	"influx-test/common"

	"github.com/influxdata/influxdb/client/v2"
)

type TaskInterface interface {
	Start(clnt client.Client) common.TestRes
	Prepare()
	DbName() string
}

func DoTask(clnt client.Client, t TaskInterface) common.TestRes {
	t.Prepare()
	return t.Start(clnt)
}

type Task struct {
	TaskName string
	CmdList  map[string]string
}

// prepare test cmd list for some task
func (t *Task) Prepare() {}

func (t *Task) DbName() string { return "mydb" }

func (t *Task) Start(clnt client.Client) common.TestRes {
	var tr common.TestRes

	tr.TaskName = t.TaskName
	tr.TotalCnt = len(t.CmdList)

	var q client.Query

	tr.PassCnt = 0
	for syntax, cmd := range t.CmdList {
		q = client.NewQuery(cmd, t.DbName(), "ns")
		if response, err := clnt.Query(q); err == nil && response.Error() == nil {
			tr.PassCnt++
		} else {
			// record errinfo
			info := fmt.Sprintf("[ %s ] %s :", syntax, cmd)
			if err != nil {
				info += " " + err.Error() + " "
			}
			if response.Error() != nil {
				info += " " + response.Error().Error() + " "
			}

			tr.ErrInfos = append(tr.ErrInfos, info)
		}
	}

	return tr
}
