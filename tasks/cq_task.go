package tasks

type CQTask struct {
	Task
}

var cq_list = map[string]string{
	"CREATE CONTINUOUS":  "create continuous query min_car_cq on mydb begin select mean(speed) as speed, mean(temp) as temp into min_car from car group by time(1m), * end",
	"CREATE CONTINUOUS1": "create continuous query hour_car_cq on mydb resample every 15m for 1h begin select mean(speed) as speed, mean(temp) as temp into hour_car from car group by time(1h, 30m), * end",
	"SHOW CONTINUOUS":    "show continuous queries",
	"DROP CONTINUOUS":    "drop continuous query hour_car_cq on mydb",
}

func (c *CQTask) Prepare() {
	c.TaskName = "cq task"
	c.CmdList = cq_list
}

func (c *CQTask) GetTaskName() string {
	return "cq task"
}
