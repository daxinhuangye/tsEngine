// tsBeego
package tsBeego

import (
	"github.com/astaxie/beego"
)

const (
	MaxPageSize = 1000000

	//****************返回值****************************
	SuccessProto 	int = 1 // 成功
	ErrorProto   	int = 2 // 错误
	ParamError   	int = 3 // http请求参数错误
	DbError      	int = 4 // 数据库错误
	NoOwner			int = 5
	NoLogin			int = 6
)

const (
	IsDel int64 = 1 // 删除
	NoDel int64 = 2 // 不删除

	Valid   int64 = 1 // 有效
	Invalid int64 = 2 // 无效

	Forbid   int64 = 1 // 禁用
	NoForbid int64 = 2 // 非禁用
)

type TsBaseController struct {
	beego.Controller
	Code   int
	Result interface{}
}

//json 输出
func (this *TsBaseController) TraceJson() {
	this.Data["json"] = &map[string]interface{}{"Code": this.Code, "Data": this.Result}
	this.ServeJSON()
	this.StopRun()
}
