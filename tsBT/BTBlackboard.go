package tsBT

import (
	"tsEngine/tsAttr"
	"tsEngine/tsLua"
)

type BTBlackboard struct {
	*tsAttr.Attrs
	Lua *tsLua.TsLua
	Ws interface{}
	User interface{}
	Msg interface{}
	Service interface{}
	Result int32
}

func NewBTBlackboard()(*BTBlackboard) {
	res := new(BTBlackboard)
	res.Attrs = tsAttr.NewAttrs()
	res.Ws = nil
	res.User = nil
	res.Msg = nil
	res.Service = nil
	res.Result = 0
	return res
}

func (this *BTBlackboard)Clear()() {
	this.Lua = nil
	this.Ws = nil
	this.User = nil
	this.Msg = nil
	this.Result = 0
	this.Attrs.Clear()
}