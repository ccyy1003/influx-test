package tasks

type MgDbTask struct {
	Task
}

var mgdb_list = map[string]string{
	// syntax : sentence
	"CREATE DATABASE ":                  "CREATE DATABASE test_tsdb WITH DURATION 3d REPLICATION 3 PARTITIONS 16",
	"CREATE DATABASE WITH CTSDB_OPTION": "CREATE DATABASE test_tsdb WITH DURATION 3d REPLICATION 3 PARTITIONS 16 ctsdb_option '{\"route_tag\": {\"measurements\": { \"m1\": [ \"t1\", \"t2\", \"t3\"], \"m2\": [ \"t4\", \"t5\"]} } }'",
	"DROP DATABASE ":                    "DROP DATABASE test_car",
	"DROP MEASUREMENT":                  "DROP MEASUREMENT car",
	"DELETE":                            "delete from car where city = 'city_0'",
}

func (m *MgDbTask) Prepare() {
	m.TaskName = "mgdb task"
	m.CmdList = mgdb_list
}

func (m *MgDbTask) GetTaskName() string {
	return "manage db"
}
