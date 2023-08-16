package tasks

import (
	"fmt"
	"influx-test/common"
	"math/rand"
	"time"

	"github.com/influxdata/influxdb/client/v2"
)

// 覆盖的接口
// NewHTTPClient
// Ping
// query
// NewUDPClient
// NewBatchPoints
// SetDatabase
// SetRetentionPolicy
// SetWriteConsistency
// SetPrecision
// NewPoint
// Point_withoutTime
// QueryWithParams
// QueryWithRP

type OtherTask struct {
}

const (
	TotalInterface = 12
)

func TestClient(tr *common.TestRes) {
	// NOTE: this assumes you've setup a user and have setup shell env variables,
	// namely INFLUX_USER/INFLUX_PWD. If not just omit Username/Password below.
	_, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     common.Addr,
		Username: common.Username,
		Password: common.Password,
	})
	if err != nil {
		info := "[ NewHTTPClient ] :" + err.Error()
		tr.ErrInfos = append(tr.ErrInfos, info)
		return
	}
	tr.PassCnt++
}

// Ping the cluster using the HTTP client
func TestClient_Ping(tr *common.TestRes) {
	// Make client
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     common.Addr,
		Username: common.Username,
		Password: common.Password,
	})
	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
		return
	}
	defer c.Close()

	_, _, err = c.Ping(0)
	if err != nil {
		info := "[ Ping ] :" + err.Error()
		tr.ErrInfos = append(tr.ErrInfos, info)
		return
	}
	tr.PassCnt++
}

// Write a point using the HTTP client
func TestClient_write(tr *common.TestRes) {
	// Make client
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     common.Addr,
		Username: common.Username,
		Password: common.Password,
	})
	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
		return
	}
	defer c.Close()

	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "mydb",
		Precision: "s",
	})
	if err != nil {
		info := "[ NewBatchPoints ] :" + err.Error()
		tr.ErrInfos = append(tr.ErrInfos, info)
		return
	}
	// Create a point and add to batch
	tags := map[string]string{"cpu": "cpu-total"}
	fields := map[string]interface{}{
		"idle":   10.1,
		"system": 53.3,
		"user":   46.6,
	}
	pt, err := client.NewPoint("cpu_usage", tags, fields, time.Now())
	if err != nil {
		info := "[ NewPoint ] :" + err.Error()
		tr.ErrInfos = append(tr.ErrInfos, info)
		return
	}
	bp.AddPoint(pt)

	// Write the batch
	err = c.Write(bp)
	if err != nil {
		info := "[ Write ] :" + err.Error()
		tr.ErrInfos = append(tr.ErrInfos, info)
		return
	}
	tr.PassCnt++
}

// Make a Query
func TestClient_query(tr *common.TestRes) {
	// Make client
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     common.Addr,
		Username: common.Username,
		Password: common.Password,
	})
	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
		return
	}
	defer c.Close()

	q := client.NewQuery("SELECT count(*) FROM car", "mydb", "ns")
	if response, err := c.Query(q); err != nil || response.Error() != nil {
		info := "[ Query ] :"
		if err != nil {
			info += " " + err.Error() + " "
		}
		if response.Error() != nil {
			info += " " + response.Error().Error() + " "
		}
		tr.ErrInfos = append(tr.ErrInfos, info)
		return
	}
	tr.PassCnt++
}

// Write a point using the UDP client
func TestClient_uDP(tr *common.TestRes) {
	// Make client
	config := client.UDPConfig{Addr: "localhost:8089"}
	c, err := client.NewUDPClient(config)
	if err != nil {
		info := "[ NewUDPClient ] :" + err.Error()
		tr.ErrInfos = append(tr.ErrInfos, info)
		return
	}
	defer c.Close()

	// Create a new point batch
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "mydb",
		Precision: "s",
	})

	// Create a point and add to batch
	tags := map[string]string{"cpu": "cpu-total"}
	fields := map[string]interface{}{
		"idle":   10.1,
		"system": 53.3,
		"user":   46.6,
	}
	pt, err := client.NewPoint("uudp", tags, fields, time.Now())
	if err != nil {
		info := "[ NewPoint ] :" + err.Error()
		tr.ErrInfos = append(tr.ErrInfos, info)
		return
	}
	bp.AddPoint(pt)

	// Write the batch
	err = c.Write(bp)
	if err != nil {
		info := "[ Client_uDP Write ] :" + err.Error()
		tr.ErrInfos = append(tr.ErrInfos, info)
		return
	}
	tr.PassCnt++
}

// Create a batch and add a point
func TestBatchPoints(tr *common.TestRes) {
	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "BumbleBeeTuna",
		Precision: "s",
	})
	if err != nil {
		info := "[ NewBatchPoints ] :" + err.Error()
		tr.ErrInfos = append(tr.ErrInfos, info)
		return
	}
	// Create a point and add to batch
	tags := map[string]string{"cpu": "cpu-total"}
	fields := map[string]interface{}{
		"idle":   10.1,
		"system": 53.3,
		"user":   46.6,
	}
	pt, err := client.NewPoint("cpu_usage", tags, fields, time.Now())
	if err != nil {
		info := "[ NewPoint ] :" + err.Error()
		tr.ErrInfos = append(tr.ErrInfos, info)
		return
	}
	bp.AddPoint(pt)
	tr.PassCnt++
}

// Using the BatchPoints setter functions
func TestBatchPoints_setters(tr *common.TestRes) {
	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{})
	if err != nil {
		info := "[ NewBatchPoints ] :" + err.Error()
		tr.ErrInfos = append(tr.ErrInfos, info)
		return
	}
	bp.SetDatabase("BumbleBeeTuna")
	bp.SetRetentionPolicy("rp")
	bp.SetWriteConsistency("wc")
	err = bp.SetPrecision("ms")
	if err != nil {
		info := "[ SetPrecision ] :" + err.Error()
		tr.ErrInfos = append(tr.ErrInfos, info)
		return
	}
	if bp.Precision() != "ms" {
		tr.ErrInfos = append(tr.ErrInfos, fmt.Sprintf("[ SetPrecision ]: expect %s, get %s", "ms", bp.Precision()))
		return
	}
	if bp.Database() != "BumbleBeeTuna" {
		tr.ErrInfos = append(tr.ErrInfos, fmt.Sprintf("[ SetDatabase ]: expect %s, get %s", "ms", bp.Precision()))
		return
	}
	if bp.RetentionPolicy() != "rp" {
		tr.ErrInfos = append(tr.ErrInfos, fmt.Sprintf("[ SetRetentionPolicy ]: expect %s, get %s", "ms", bp.Precision()))
		return
	}
	if bp.WriteConsistency() != "wc" {
		tr.ErrInfos = append(tr.ErrInfos, fmt.Sprintf("[ SetWriteConsistency ]: expect %s, get %s", "ms", bp.Precision()))
		return
	}

	tr.PassCnt++
}

// Create a new point with a timestamp
func TestPoint(tr *common.TestRes) {
	tags := map[string]string{"cpu": "cpu-total"}
	fields := map[string]interface{}{
		"idle":   10.1,
		"system": 53.3,
		"user":   46.6,
	}
	_, err := client.NewPoint("cpu_usage", tags, fields, time.Now())
	if err != nil {
		info := "[ NewPoint with a timestamp ] :" + err.Error()
		tr.ErrInfos = append(tr.ErrInfos, info)
		return
	}
	tr.PassCnt++
}

// Create a new point without a timestamp
func TestPoint_withoutTime(tr *common.TestRes) {
	tags := map[string]string{"cpu": "cpu-total"}
	fields := map[string]interface{}{
		"idle":   10.1,
		"system": 53.3,
		"user":   46.6,
	}
	_, err := client.NewPoint("cpu_usage", tags, fields)
	if err != nil {
		info := "[ NewPoint without a timestamp ] :" + err.Error()
		tr.ErrInfos = append(tr.ErrInfos, info)
		return
	}
	tr.PassCnt++
}

// Write 1000 points
func TestClient_write1000(tr *common.TestRes) {
	sampleSize := 1000

	// Make client
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     common.Addr,
		Username: common.Username,
		Password: common.Password,
	})
	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
		return
	}
	defer c.Close()

	rand.Seed(42)

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "mydb",
		Precision: "ms",
	})

	for i := 0; i < sampleSize; i++ {
		regions := []string{"us-west1", "us-west2", "us-west3", "us-east1"}
		tags := map[string]string{
			"cpu":    "cpu-total",
			"host":   fmt.Sprintf("host%d", rand.Intn(1000)),
			"region": regions[rand.Intn(len(regions))],
		}

		idle := rand.Float64() * 100.0
		fields := map[string]interface{}{
			"idle": idle,
			"busy": 100.0 - idle,
		}

		pt, err := client.NewPoint(
			"cpu_usage",
			tags,
			fields,
			time.Now(),
		)
		if err != nil {
			println("Error:", err.Error())
			continue
		}
		bp.AddPoint(pt)
	}

	err = c.Write(bp)
	if err != nil {
		info := "[ Write 1000 points ] :" + err.Error()
		tr.ErrInfos = append(tr.ErrInfos, info)
		return
	}
	tr.PassCnt++
}

func TestClient_queryWithParams(tr *common.TestRes) {
	// Make client
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     common.Addr,
		Username: common.Username,
		Password: common.Password,
	})
	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
	}
	defer c.Close()

	q := client.NewQueryWithParameters("SELECT $fn($value) FROM $m", "mydb", "ns", client.Params{
		"fn":    client.Identifier("count"),
		"value": client.Identifier("*"),
		"m":     client.Identifier("car"),
	})
	if response, err := c.Query(q); err != nil || response.Error() != nil {
		info := "[ queryWithParams ] :"
		if err != nil {
			info += " " + err.Error() + " "
		}
		if response.Error() != nil {
			info += " " + response.Error().Error() + " "
		}
		tr.ErrInfos = append(tr.ErrInfos, info)
		return
	}
	tr.PassCnt++
}

func TestClient_queryWithRP(tr *common.TestRes) {
	// Make client
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     common.Addr,
		Username: common.Username,
		Password: common.Password,
	})
	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
	}
	defer c.Close()

	q := client.NewQueryWithRP("select count(*) from car", "mydb", "autogen", "ms")
	if response, err := c.Query(q); err != nil || response.Error() != nil {
		info := "[ queryWithRP ] :"
		if err != nil {
			info += " " + err.Error() + " "
		}
		if response.Error() != nil {
			info += " " + response.Error().Error() + " "
		}
		tr.ErrInfos = append(tr.ErrInfos, info)
		return
	}
	tr.PassCnt++
}

func (o *OtherTask) Prepare() {}

func (o *OtherTask) DbName() string { return "mydb" }

func (o *OtherTask) Start(clnt client.Client) common.TestRes {
	var tr common.TestRes

	tr.TaskName = "interface task"
	tr.TotalCnt = TotalInterface
	TestClient(&tr)
	TestClient_uDP(&tr)
	TestClient_Ping(&tr)
	TestClient_write(&tr)
	TestBatchPoints(&tr)
	TestBatchPoints_setters(&tr)
	TestPoint(&tr)
	TestPoint_withoutTime(&tr)
	TestClient_write1000(&tr)
	TestClient_query(&tr)
	TestClient_queryWithRP(&tr)
	TestClient_queryWithParams(&tr)

	return tr
}
