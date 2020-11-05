## Golang 命令工具

中文翻译：

```
Go 是用于管理 Go 源代码的工具

用法:

	go <命令> [参数]

命令如下:

	bug         启动错误报告
	build       编译包和依赖项
	clean       删除目标文件和缓存文件
	doc         显示包或符号的文档
	env         打印Go环境信息
	fix         更新程序包以使用新的API
	fmt         gofmt（重新格式化）程序包
	generate    通过处理源文件生成 Go 文件
	get         添加依赖项到当前的模块并安装它们
	install     编译并安装软件包和依赖项
	list        列出包和模块
	mod         模块维护
	run         编译并运行 Go 程序
	test        测试包
	tool        运行指定的 Go 工具
	version     打印 Go 版本信息
	vet         报告包可能存在的错误

使用“ go help <命令>”可获取有关命令的更多信息。

其他帮助主题：

	buildconstraint 构建约束
	buildmode       构建模式
	c               在Go和C之间调用
	cache           构建和测试缓存
	environment     环境变量
	filetype        文件类型
	go.mod          go.mod文件
	gopath          GOPATH环境变量
	gopath-get      旧版 GOPATH go get
	goproxy         模块代理协议
	importpath      导入路径语法
	modules         模块，模块版本等
	module-get      模块获取
	module-auth     使用go.sum进行模块身份验证
	module-private  非公共模块的模块配置
	packages        包列表和模式
	testflag        测试标志
	testfunc        测试功能

使用“ go help <主题>”获得有关该主题的更多信息。

```

原文：

```

The commands are:

	bug         start a bug report
	build       compile packages and dependencies
	clean       remove object files and cached files
	doc         show documentation for package or symbol
	env         print Go environment information
	fix         update packages to use new APIs
	fmt         gofmt (reformat) package sources
	generate    generate Go files by processing source
	get         add dependencies to current module and install them
	install     compile and install packages and dependencies
	list        list packages or modules
	mod         module maintenance
	run         compile and run Go program
	test        test packages
	tool        run specified go tool
	version     print Go version
	vet         report likely mistakes in packages

Use "go help <command>" for more information about a command.

Additional help topics:

	buildconstraint build constraints
	buildmode       build modes
	c               calling between Go and C
	cache           build and test caching
	environment     environment variables
	filetype        file types
	go.mod          the go.mod file
	gopath          GOPATH environment variable
	gopath-get      legacy GOPATH go get
	goproxy         module proxy protocol
	importpath      import path syntax
	modules         modules, module versions, and more
	module-get      module-aware go get
	module-auth     module authentication using go.sum
	module-private  module configuration for non-public modules
	packages        package lists and patterns
	testflag        testing flags
	testfunc        testing functions

Use "go help <topic>" for more information about that topic.
```

