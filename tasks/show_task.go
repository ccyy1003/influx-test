package tasks

type ShowTask struct {
	Task
}

var show_list = map[string]string{
	"SHOW DATABASES":          "SHOW DATABASES",
	"SHOW SERIES":             "SHOW SERIES FROM yottadb_partition_replicas_num_lzl where cluster_name = 'yottadb-capd-gz-dataplane6' ORDER BY ASC LIMIT 5 OFFSET 5",
	"SHOW MEASUREMENTS":       "SHOW MEASUREMENTS WHERE cluster_name = 'yottadb-capd-gz-dataplane6' LIMIT 3 OFFSET 0",
	"SHOW TAG KEYS":           "SHOW TAG KEYS WHERE cluster_name = 'yottadb-capd-gz-dataplane6' ORDER BY ASC LIMIT 5 OFFSET 1 SLIMIT 2 SOFFSET 2",
	"SHOW TAG VALUES":         "SHOW TAG VALUES WITH KEY IN (account_id, cluster_display_name, cluster_name)  ORDER BY ASC LIMIT 8 OFFSET 1 SLIMIT 2 SOFFSET 1",
	"SHOW FIELD KEYS":         "SHOW FIELD KEYS FROM yottadb_partition_replicas_num ORDER BY ASC LIMIT 4 OFFSET 4",
	"SHOW SERIES CARDINALITY": "SHOW SERIES CARDINALITY from yottadb_partition_replicas_num group by partition_id",
}

func (s *ShowTask) Prepare() {
	s.TaskName = "show task"
	s.CmdList = show_list
}

func (s *ShowTask) GetTaskName() string {
	return "show"
}
