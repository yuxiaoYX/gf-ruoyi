package task

// import (
// 	"os/exec"

// 	"github.com/gogf/gf/v2/frame/g"
// 	"github.com/gogf/gf/v2/os/gfile"
// 	"github.com/gogf/gf/v2/os/gtime"
// 	"github.com/gogf/gf/v2/util/gconv"
// )

// func init() {
// 	var task1 Entity
// 	task1.FuncName = "test1"
// 	task1.Param = nil
// 	task1.Run = Test1
// 	Add(task1)

// 	var task2 Entity
// 	task2.FuncName = "test2"
// 	task2.Param = nil
// 	task2.Run = Test2
// 	Add(task2)

// 	var shellScript Entity
// 	shellScript.FuncName = "shellScript"
// 	shellScript.Param = nil
// 	shellScript.Run = ShellScript
// 	Add(shellScript)

// 	var terminalExecMysql Entity
// 	terminalExecMysql.FuncName = "terminalExecMysql"
// 	terminalExecMysql.Param = nil
// 	terminalExecMysql.Run = TerminalExecMysql
// 	Add(terminalExecMysql)

// 	var bingImg Entity
// 	bingImg.FuncName = "bingImg"
// 	bingImg.Param = nil
// 	bingImg.Run = BingImg
// 	Add(bingImg)
// }

// //无参测试
// func Test1() {
// 	println("无参测试")
// }

// //传参测试
// func Test2() {
// 	//获取参数
// 	task := GetByName("test2")
// 	if task == nil {
// 		return
// 	}
// 	for _, v := range task.Param {
// 		println(v)
// 	}
// }

// // 运行shell脚本
// func ShellScript() {
// 	// 获取参数
// 	task := GetByName("shellScript")
// 	// exec.Command第一个参数是可执行文件的名称，mysqldump并不是运行在shell环境中
// 	cmd := exec.Command("bash", "-c", task.Param[0])
// 	if err := cmd.Start(); err != nil {
// 		g.Log().Error(err)
// 	}
// 	g.Log().Info("shellScript:运行Shell脚本！")
// }

// // 运行shell脚本,备份oa-gf数据库
// func TerminalExecMysql() {
// 	// 获取参数
// 	task := GetByName("terminalExecMysql")
// 	param := task.Param[0] + gconv.String(gtime.Timestamp()) + "oa-gf.sql"
// 	// exec.Command第一个参数是可执行文件的名称，mysqldump并不是运行在shell环境中
// 	cmd := exec.Command("bash", "-c", param)
// 	if err := cmd.Start(); err != nil {
// 		g.Log().Error(err)
// 	}
// 	g.Log().Info("terminalExecMysql:备份数据库！")
// }

// // 下载必应每日图片
// func BingImg() {
// 	// 获取必应每日图片地址
// 	content := g.Client().GetContent("https://cn.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1")
// 	contentMap := gconv.Map(content)
// 	url := contentMap["images"].([]interface{})[0].(map[string]interface{})["url"]
// 	bingUrl := "https://www.bing.com" + url.(string)
// 	// 下载必应图片
// 	if r, err := g.Client().Get(bingUrl); err != nil {
// 		panic(err)
// 	} else {
// 		defer r.Close()
// 		gfile.PutBytes(g.Cfg().GetString("server.ServerRoot")+"/login-background.jpg", r.ReadAll())
// 	}
// 	g.Log().Info("bingImg:下载必应每日图片！")
// }
