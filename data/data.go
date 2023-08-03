package data

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	"github.com/influxdata/influxdb/client/v2"
)

func getVerifyData() {

}

// func DownLoad(fileUrl string) string {
// 	url := fileUrl
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer resp.Body.Close()
// 	parts := strings.Split(fileUrl, "/")
// 	file, err := os.Create("../data/txt/" + parts[len(parts)-1])
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()
// 	_, err = io.Copy(file, resp.Body)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("File downloaded successfully")
// 	return parts[len(parts)-1]
// }

// func InitByUrl(fileUrl string) {
// 	file := "../data/txt/" + DownLoad(fileUrl)
// 	fmt.Println(file)
// 	url := "http://localhost:8086/write?db=NOAA_water_database"
// 	method := "POST"
// 	dataFile := file
// 	// 读取数据文件内容
// 	data, err := os.ReadFile(dataFile)
// 	if err != nil {
// 		fmt.Println("读取数据文件失败：", err)
// 		return
// 	}
// 	// 创建HTTP请求
// 	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
// 	if err != nil {
// 		fmt.Println("创建HTTP请求失败:", err)
// 		return
// 	}
// 	req.Header.Set("Content-Type", "application/octet-stream")
// 	// 发送HTTP请求
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println("发送HTTP请求失败:", err)
// 		return
// 	}
// 	defer resp.Body.Close()
// }

func InitDb(clnt client.Client, db string, filename string) {
	q := client.NewQuery("create database "+db, "", "ns")
	if response, err := clnt.Query(q); err == nil && response.Error() != nil {
		fmt.Println("error")
	}
	url := "http://localhost:8086/write?db=" + db
	method := "POST"
	dataFile := "../data/txt/" + filename
	// 读取数据文件内容
	data, err := os.ReadFile(dataFile)
	if err != nil {
		fmt.Println("读取数据文件失败：", err)
		return
	}
	// 创建HTTP请求
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("创建HTTP请求失败：", err)
		return
	}
	req.Header.Set("Content-Type", "application/octet-stream")
	// 发送HTTP请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送HTTP请求失败：", err)
		return
	}
	defer resp.Body.Close()
}

// init test data
func Init(clnt client.Client) {
	InitDb(clnt, "mydb", "test_data.txt")
	InitDb(clnt, "mydb", "NOAA_data.txt")
}
