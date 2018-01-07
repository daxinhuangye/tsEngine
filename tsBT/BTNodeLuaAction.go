package tsBT

import (
	"tsEngine/tsAttr"
	"tsEngine/tsString"
	"github.com/astaxie/beego"
)

type BTNodeLuaAction struct {
	BTNode
	Fields *tsAttr.Attrs
	LuaFunc string
}

func NewBTNodeLuaAction(bt *BT)(res *BTNodeLuaAction) {
	res = new(BTNodeLuaAction)
	res.Fields = tsAttr.NewAttrs()
	res.Bt = bt
	return
}

func (this *BTNodeLuaAction)Execute(black_board *BTBlackboard)(bool) {
	if this.IsLog() {
		beego.Trace("------------------>", this.Name, "", this.LuaFunc)
	}
	
	_, err := black_board.Lua.Call(this.LuaFunc, black_board, this.Fields)
	if err!=nil {
		beego.Error(this.Name, this.LuaFunc, err)
	}
	return true
}

func (this *BTNodeLuaAction)FromJson(data map[string]interface{})(bool) {
	if this.BTNode.FromJson(data)==false {
		return false
	}
	var ok bool
	i_LuaFunc, ok := data["LuaFunc"]
	if !ok {
		return false
	}
	this.LuaFunc = i_LuaFunc.(string)
	for k,v:=range data {
		if k=="Name" || k=="Index" || k=="LuaFunc" || k=="ClassName" || k=="Style" {
			continue
		}
		param := tsString.Split(k, "-")
		if len(param)==2 {
			if param[0]=="int64" {
				this.Fields.Field[param[1]] = tsString.ToInt64(v.(string))
				continue
			}
			if param[0]=="string" {
				this.Fields.Field[param[1]] = v
				continue
			}
		}
	}
	return true
}