package tsBT

import (
)

type IBTNode interface {
	GetIndex()(string)
	GetName()(string)
	AddChild(n IBTNode)(bool)
	Execute(black_board *BTBlackboard)(bool)
	FromJson(data map[string]interface{})(bool)
}

type BTNode struct {
	Name string
	Index string
	Log string
	Bt	*BT
}

func (this *BTNode)GetIndex()(string) {
	return this.Index
}

func (this *BTNode)GetName()(string) {
	return this.Name
}

func (this *BTNode)AddChild(n IBTNode)(bool) {
	return false
}

func (this *BTNode)IsLog()(bool) {
	if this.Log=="1" {
		this.Bt.IsLog = true
		return true
	}
	
	return this.Bt.IsLog
}

func (this *BTNode)FromJson(data map[string]interface{})(bool) {
	var ok bool
	i_Name,ok := data["Name"]
	if !ok {
		return false
	}
	i_Index, ok := data["Index"]
	if !ok {
		return false
	}
	i_Log, ok := data["Log"]
	if !ok {
		this.Log = ""
	} else {
		this.Log = i_Log.(string)
	}
	this.Name = i_Name.(string)
	this.Index = i_Index.(string)
	
	return true
}