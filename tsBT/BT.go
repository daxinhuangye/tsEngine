package tsBT

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"tsEngine/tsTime"
)

type BT struct {
	Root IBTNode
	Blackboard *BTBlackboard
	IsLog bool
	Name string
}

func NewBT()(*BT) {
	res := new(BT)
	res.Blackboard = NewBTBlackboard()
	res.IsLog = false
	return res
}

func (this *BT) Execute(black_board *BTBlackboard)(bool) {
	t := tsTime.CurrMs()
	if this.IsLog {
		beego.Trace("------------------", this.Name, "------------------")
	}
	
	res := this.Root.Execute(black_board)
	
	if this.IsLog {
		beego.Trace(tsTime.CurrMs()-t, "********************", this.Name, "********************")
	}
	return res
}

func (this *BT)FromJson(data string)(bool) {
	var i_data interface{}
	err := json.Unmarshal([]byte(data), &i_data)
	if err!=nil {
		return false
	}
	i_json, ok := i_data.(map[string]interface{})
	if !ok {
		return false
	}
	
	this.Root = this.NodeFromJson(i_json)
	
	return true
}

func (this *BT)NodeFromJson(data map[string]interface{})(res IBTNode) {
	
	//***************************本节点数据
	i_node,ok:=data["data"]
	if !ok {
		return
	}
	node := i_node.(map[string]interface{})
	
	i_ClassName,ok:=node["ClassName"]
	if !ok {
		return 
	}
	
	ClassName := i_ClassName.(string)
	if ClassName=="NodeRoot" {
		res = NewBTNodeRoot(this)
		res.FromJson(node)
	} else if ClassName=="NodeAnd" {
		res = NewBTNodeAnd(this)
		res.FromJson(node)
	} else if ClassName=="NodeOr" {
		res = NewBTNodeOr(this)
		res.FromJson(node)
	}  else if ClassName=="NodeNot" {
		res = NewBTNodeNot(this)
		res.FromJson(node)
	} else if ClassName=="NodeMap" {
		res = NewBTNodeMap(this)
		res.FromJson(node)
	} else if ClassName=="NodeTrue" {
		res = NewBTNodeTrue(this)
		res.FromJson(node)
	} else if ClassName=="NodeFalse" {
		res = NewBTNodeFalse(this)
		res.FromJson(node)
	} else if ClassName=="NodeLuaAction" {
		res = NewBTNodeLuaAction(this)
		res.FromJson(node)
	} else if ClassName=="NodeLuaCondition" {
		res = NewBTNodeLuaCondition(this)
		res.FromJson(node)
	} else {
		return
	}

	//***************************************子节点
	i_childs,ok:=data["childs"]
	if !ok {
		return 
	}
	childs := i_childs.([]interface{})
	if len(childs)<=0 {
		return
	}
	for i:=0; i<len(childs); i++ {
		child := this.NodeFromJson(childs[i].(map[string]interface{}))
		if child!=nil {
			res.AddChild(child)
		}
	}
	return 
}