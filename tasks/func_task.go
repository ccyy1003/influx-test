package tasks

type FuncTask struct {
	Task
}

var dbname = "mydb"

var func_list = map[string]string{
	// Transformations
	"ABS":            "select abs(speed) from car",
	"ACOS":           "SELECT ACOS(of_capacity) FROM park_occupancy",
	"ASIN":           "SELECT ASIN(of_capacity) FROM park_occupancy",
	"ATAN":           "SELECT ATAN(of_capacity) FROM park_occupancy",
	"ATAN2":          "SELECT ATAN2(altitude_ft, distance_ft) FROM flight_data",
	"LOG":            "SELECT LOG(water_level, 4) FROM h2o_feet WHERE location = 'santa_monica' LIMIT 10",
	"MOVING_AVERAGE": "SELECT MOVING_AVERAGE(water_level, 2) FROM h2o_feet WHERE location = 'santa_monica' LIMIT 10",
	"POW":            "SELECT POW(water_level, 4) FROM h2o_feet WHERE location = 'santa_monica' LIMIT 10",

	// Aggregations
	"COUNT":    "SELECT COUNT(water_level) FROM h2o_feet",
	"DISTINCT": "SELECT DISTINCT(\"level description\") FROM h2o_feet",
	"INTEGRAL": "SELECT INTEGRAL(water_level) FROM h2o_feet WHERE location = 'santa_monica'",
	"MEAN":     "SELECT MEAN(water_level) FROM h2o_feet",
	"MEDIAN":   "SELECT MEDIAN(water_level) FROM h2o_feet",
	"MODE":     "SELECT MODE(\"level description\") FROM h2o_feet",
	"SPREAD":   "SELECT SPREAD(water_level) FROM h2o_feet",
	"STDDEV":   "SELECT STDDEV(water_level) FROM h2o_feet",
	"SUM":      "SELECT SUM(water_level) FROM h2o_feet",

	//Selectors
	"BOTTOM":     "SELECT BOTTOM(water_level,3) FROM h2o_feet",
	"FIRST":      "SELECT FIRST(\"level description\") FROM h2o_feet",
	"LAST":       "SELECT LAST(\"level description\") FROM h2o_feet",
	"MAX":        "SELECT MAX(water_level) FROM h2o_feet",
	"MIN":        "SELECT MIN(water_level) FROM h2o_feet",
	"PERCENTILE": "SELECT PERCENTILE(water_level,5) FROM h2o_feet",
	"SAMPLE":     "SELECT SAMPLE(water_level,2) FROM h2o_feet",
	"TOP":        "SELECT TOP(water_level,3) FROM h2o_feet",

	// Predictors
	"HOLT_WINTERS": "SELECT HOLT_WINTERS_WITH_FIT(FIRST(water_level),10,4) FROM mydb.autogen.h2o_feet WHERE location='santa_monica' AND time >= '2019-09-15 22:12:00' AND time <= '2019-09-28 03:00:00' GROUP BY time(379m,348m)",
}

var keys = [...]string{
	"CEIL",
	"COS",
	"CUMULATIVE_SUM",
	"DERIVATIVE",
	"DIFFERENCE",
	"ELAPSED",
	"EXP",
	"FLOOR",
	"LN",
	"LOG2",
	"LOG10",
	"NON_NEGATIVE_DERIVATIVE",
	"NON_NEGATIVE_DIFFERENCE",
	"ROUND",
	"SIN",
	"SQRT",
	"TAN",
}

func (f *FuncTask) Prepare() {
	for _, key := range keys {
		func_list[key] = "SELECT " + key + "(water_level) FROM h2o_feet WHERE location = 'santa_monica' LIMIT 10"
	}
	f.TaskName = "func task"
	f.CmdList = func_list
}

func (f *FuncTask) GetTaskName() string {
	return "func task"
}

func (f *FuncTask) DbName() string {
	return dbname
}
