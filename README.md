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
| cq                 	|                                                                                                                                                                                  	|
|--------------------	|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------	|
| CREATE CONTINUOUS  	| create continuous query min_car_cq on mydb begin select mean(speed) as speed, mean(temp) as temp into min_car from car group by time(1m), * end                                  	|
| CREATE CONTINUOUS1 	| create continuous query hour_car_cq on mydb resample every 15m for 1h begin select mean(speed) as speed, mean(temp) as temp into hour_car from car group by time(1h, 30m), * end 	|
| SHOW CONTINUOUS    	| show continuous queries                                                                                                                                                          	|
| DROP CONTINUOUS    	| drop continuous query hour_car_cq on mydb                                                                                                                                        	|

| func（Transformations） 	|                                                                                                                                                                                                             	|
|-------------------------	|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------	|
| ABS                     	| SELECT ABS(speed) FROM car                                                                                                                                                                                  	|
| ACOS                    	| SELECT ACOS(of_capacity) FROM park_occupancy                                                                                                                                                                	|
| ASIN                    	| SELECT ASIN(of_capacity) FROM park_occupancy                                                                                                                                                                	|
| ATAN                    	| SELECT ATAN(of_capacity) FROM park_occupancy                                                                                                                                                                	|
| ATAN2                   	| SELECT ATAN2(altitude_ft, distance_ft) FROM flight_data                                                                                                                                                     	|
| CEIL                    	| SELECT CEIL(water_level) FROM h2o_feet WHERE location = 'santa_monica' LIMIT 10                                                                                                                             	|
| COS                     	| SELECT COS(water_level) FROM h2o_feet WHERE location = 'santa_monica' LIMIT 10                                                                                                                              	|
| CUMULATIVE_SUM          	| SELECT CUMULATIVE_SUM(water_level) FROM h2o_feet WHERE location = 'santa_monica' LIMIT 10                                                                                                                   	|
| DERIVATIVE              	| SELECT DERIVATIVE(water_level) FROM h2o_feet WHERE location = 'santa_monica' LIMIT 10                                                                                                                       	|
| DIFFERENCE              	| SELECT DIFFERENCE(water_level) FROM h2o_feet WHERE location = 'santa_monica' LIMIT 10                                                                                                                       	|
| ELAPSED                 	| SELECT ELAPSED(water_level) FROM h2o_feet WHERE location = 'santa_monica' LIMIT 10                                                                                                                          	|
| EXP                     	| SELECT EXP(water_level) FROM h2o_feet WHERE location = 'santa_monica' LIMIT 10                                                                                                                              	|
| FLOOR                   	| SELECT FLOOR(water_level) FROM h2o_feet WHERE location = 'santa_monica' LIMIT 10                                                                                                                            	|
| LN                      	| SELECT LN(water_level) FROM h2o_feet WHERE location = 'santa_monica' LIMIT 10                                                                                                                               	|
| LOG                     	| SELECT LOG(water_level, 4) FROM h2o_feet WHERE location = 'santa_monica' LIMIT 10                                                                                                                           	|
| LOG2                    	| SELECT LOG2(water_level) FROM h2o_feet WHERE location = 'santa_monica' LIMIT 10                                                                                                                             	|
| LOG10                   	| SELECT LOG10(water_level) FROM h2o_feet WHERE location = 'santa_monica' LIMIT 10                                                                                                                            	|
| MOVING_AVERAGE          	| SELECT MOVING_AVERAGE(water_level, 2) FROM h2o_feet WHERE location = 'santa_monica' LIMIT 10                                                                                                                	|
| NON_NEGATIVE_DERIVATIVE 	| SELECT NON_NEGATIVE_DERIVATIVE(water_level) FROM h2o_feet WHERE location = 'santa_monica' LIMIT 10                                                                                                          	|
| NON_NEGATIVE_DIFFERENCE 	| SELECT NON_NEGATIVE_DIFFERENCE(water_level) FROM h2o_feet WHERE location = 'santa_monica' LIMIT 10                                                                                                          	|
| POW                     	| SELECT POW(water_level, 4) FROM h2o_feet WHERE location = 'santa_monica' LIMIT 10                                                                                                                           	|
| ROUND                   	| SELECT ROUND(water_level) FROM h2o_feet WHERE location = 'santa_monica' LIMIT 10                                                                                                                            	|
| SIN                     	| SELECT SIN(water_level) FROM h2o_feet WHERE location = 'santa_monica' LIMIT 10                                                                                                                              	|
| SQRT                    	| SELECT SQRT(water_level) FROM h2o_feet WHERE location = 'santa_monica' LIMIT 10                                                                                                                             	|
| TAN                     	| SELECT TAN(water_level) FROM h2o_feet WHERE location = 'santa_monica' LIMIT 10                                                                                                                              	|
|                         	|                                                                                                                                                                                                             	|
| func(Aggregations)      	|                                                                                                                                                                                                             	|
| COUNT                   	| SELECT COUNT(water_level) FROM h2o_feet                                                                                                                                                                     	|
| DISTINCT                	| SELECT DISTINCT(\"level description\") FROM h2o_feet                                                                                                                                                        	|
| INTEGRAL                	| SELECT INTEGRAL(water_level) FROM h2o_feet WHERE location = 'santa_monica'                                                                                                                                  	|
| MEAN                    	| SELECT MEAN(water_level) FROM h2o_feet                                                                                                                                                                      	|
| MEDIAN                  	| SELECT MEDIAN(water_level) FROM h2o_feet                                                                                                                                                                    	|
| MODE                    	| SELECT MODE(\"level description\") FROM h2o_feet                                                                                                                                                            	|
| SPREAD                  	| SELECT SPREAD(water_level) FROM h2o_feet                                                                                                                                                                    	|
| STDDEV                  	| SELECT STDDEV(water_level) FROM h2o_feet                                                                                                                                                                    	|
| SUM                     	| SELECT SUM(water_level) FROM h2o_feet                                                                                                                                                                       	|
|                         	|                                                                                                                                                                                                             	|
| func(Selectors)         	|                                                                                                                                                                                                             	|
| BOTTOM                  	| SELECT BOTTOM(water_level,3) FROM h2o_feet                                                                                                                                                                  	|
| FIRST                   	| SELECT FIRST(\"level description\") FROM h2o_feet                                                                                                                                                           	|
| LAST                    	| SELECT LAST(\"level description\") FROM h2o_feet                                                                                                                                                            	|
| MAX                     	| SELECT MAX(water_level) FROM h2o_feet                                                                                                                                                                       	|
| MIN                     	| SELECT MIN(water_level) FROM h2o_feet                                                                                                                                                                       	|
| PERCENTILE              	| SELECT PERCENTILE(water_level,5) FROM h2o_feet                                                                                                                                                              	|
| SAMPLE                  	| SELECT SAMPLE(water_level,2) FROM h2o_feet                                                                                                                                                                  	|
| TOP                     	| SELECT TOP(water_level,3) FROM h2o_feet                                                                                                                                                                     	|
|                         	|                                                                                                                                                                                                             	|
| func(Predictors)        	|                                                                                                                                                                                                             	|
| HOLT_WINTERS            	| SELECT HOLT_WINTERS_WITH_FIT(FIRST(water_level),10,4) FROM mydb.autogen.h2o_feet WHERE location='santa_monica' AND time >= '2019-09-15 22:12:00' AND time <= '2019-09-28 03:00:00' GROUP BY time(379m,348m) 	|

| hint             	|                                                                                                                                                                                                                                                                                                                                                                                             	|
|------------------	|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------	|
| FULL SERIES CASE 	| select count(*) from \"250_20963_apm_calculate_prod_default_a82d\" where \"single_num\"::field > 0 and app::tag='1622-admin' and area::tag='qy' and cluster::tag='prodOpenTelemtry' and \"container_id\"::tag='_' and \"zyx_agg_type\"::tag='ins' and \"zyx_data_grain\"::tag='60' and \"zyx_instance_mark\"::tag='11.149.48.50' and \"zyx_version\"::tag='0' and \"single_num\"::field > 0 	|
