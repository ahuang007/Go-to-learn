#### 编译leafServer

* git clone https://github.com/name5566/leaf.git
	- GoPath\src\github.com\name5566\leaf

* git clone https://github.com/gorilla/websocket.git
	- GoPath\src\github.com\gorilla\websocket
	
* git clone https://github.com/name5566/leafserver.git
	- GoPath\src\server

* 编译GoPath\src\server\main.go
	- 生成 server.exe
	
* copy GoPath\src\server\server.exe > bin/ 
	
* 执行server.exe
	- 成功： 2018/12/09 13:57:35 [release] Leaf 1.1.3 starting up
	
* bin/conf/server.json 增加 "ConsolePort": 3333
	- telnet localhost 3333 可以进入控制台
	```
		Leaf# help
		Commands:
		help - this help text
		cpuprof - CPU profiling for the current process
		prof - writes a pprof-formatted snapshot
		echo - echo user inputs
		quit - exit console
	```