package tasks

type HintTask struct {
	Task
}

var hint_list = map[string]string{
	"FULL SERIES CASE": "select count(*) from \"250_20963_apm_calculate_prod_default_a82d\" where \"single_num\"::field > 0 and app::tag='1622-admin' and area::tag='qy' and cluster::tag='prodOpenTelemtry' and \"container_id\"::tag='_' and \"zyx_agg_type\"::tag='ins' and \"zyx_data_grain\"::tag='60' and \"zyx_instance_mark\"::tag='11.149.48.50' and \"zyx_version\"::tag='0' and \"single_num\"::field > 0",
}

func (h *HintTask) Prepare() {
	h.TaskName = "hint task"
	h.CmdList = hint_list
}

func (h *HintTask) GetTaskName() string {
	return "hint task"
}
