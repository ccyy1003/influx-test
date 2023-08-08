# influx-test
## 1、运行make命令打包成镜像
这里暴露的端口号为32325，也是容器启动后http监听的端口   
注：启动后会分别以os.Getenv("INFLUX_TEST_ADDR")、os.Getenv("INFLUX_TEST_USER")、os.Getenv("INFLUX_TEST_PWD")作为influxdb的访问地址，用户名和密码，若需更改，可在dockerfile文件中添加相应的环境变量
```go
make influx-test
```

## 2、以上述镜像启动容器
```go
docker run -d -p 32325:32325 influx-test:<你的镜像tag>
```


## 3、以post发送请求
```go
curl -X POST http://127.0.0.1:32325/test -d "opt=mgdb"
```
这里使用curl进行测试  
“/test”：路径  
“opt=xxx”：传递的参数，其中xxx可以为cq, func, hint, interface, mathopt, mgdb, query, show, all, quit  
   
以json格式返回结果：  
任务名称  
通过数  
测试总数  
错误信息：【语法/接口名】错误命令：错误信息  

## 测试语法及样例
| cq                 	|                                                                                                                                                                                  	|   	|   	|   	|
|--------------------	|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------	|---	|---	|---	|
| CREATE CONTINUOUS  	| create continuous query min_car_cq on mydb begin select mean(speed) as speed, mean(temp) as temp into min_car from car group by time(1m), * end                                  	|   	|   	|   	|
| CREATE CONTINUOUS1 	| create continuous query hour_car_cq on mydb resample every 15m for 1h begin select mean(speed) as speed, mean(temp) as temp into hour_car from car group by time(1h, 30m), * end 	|   	|   	|   	|
| SHOW CONTINUOUS    	| show continuous queries                                                                                                                                                          	|   	|   	|   	|
| DROP CONTINUOUS    	| drop continuous query hour_car_cq on mydb                                                                                                                                        	|   	|   	|   	|
