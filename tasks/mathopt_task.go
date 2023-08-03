package tasks

type MathOptTask struct {
	Task
}

var mathopt_list = map[string]string{
	"ADD":                  "SELECT A + 5 FROM math",
	"SUB":                  "SELECT * FROM math WHERE 1 - A <= 3",
	"MUL":                  "SELECT A * B * C FROM math",
	"DIV":                  "SELECT 10 / (A + B + C) FROM math",
	"MOD":                  "SELECT B FROM math WHERE B % 2 = 0",
	"Bitwise AND":          "SELECT A::integer & B::integer FROM math",
	"Bitwise OR":           "SELECT A::integer | 255 FROM math",
	"Bitwise Exclusive-OR": "SELECT A::integer ^ 255 FROM math",
}

func (m *MathOptTask) Prepare() {
	m.TaskName = "mathopt task"
	m.CmdList = mathopt_list
}

func (m *MathOptTask) GetTaskName() string {
	return "mathopt db"
}
