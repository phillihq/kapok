package app

import (
	"github.com/codegangsta/cli"
	"github.com/domac/kapok/util"
)

//参数初始化
func flagsInit() {

	//debug开关
	util.AddFlagBool(cli.BoolFlag{
		Name:  "debug",
		Usage: "open the debug mode",
	})

	//设置并发数
	util.AddFlagInt(cli.IntFlag{
		Name:  "c",
		Value: 10,
		Usage: "number of concurrent connections to use",
	})

	//测试持续时间
	util.AddFlagInt(cli.IntFlag{
		Name:  "d",
		Value: 10,
		Usage: "duration of test in seconds",
	})

	//调用http超时时间
	util.AddFlagInt(cli.IntFlag{
		Name:  "t",
		Value: 1000,
		Usage: "socket/request timeout in (ms)",
	})

	//http 方法 GET/POST
	util.AddFlagString(cli.StringFlag{
		Name:  "m",
		Value: "GET",
		Usage: "http method",
	})

	//设置header
	util.AddFlagString(cli.StringFlag{
		Name:  "H",
		Value: "",
		Usage: "the http headers sent to the target url",
	})

	//是否开启 keep-alived
	util.AddFlagBool(cli.BoolFlag{
		Name:  "disableka",
		Usage: "disable keep-alives",
	})

	//是否压缩
	util.AddFlagBool(cli.BoolFlag{
		Name:  "compress",
		Usage: "if prevents sending the \"Accept-Encoding: gzip\" header",
	})

	util.AddFlagString(cli.StringFlag{
		Name:  "dataFile",
		Value: "",
		Usage: "load the par which store in the file",
	})
}
