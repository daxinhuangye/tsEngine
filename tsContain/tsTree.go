package tsContain

import (

)

type TsTreeNode struct {
	Data interface{}
	Childs []*TsTreeNode
}

func NewTsTreeNode(c int64, l int64)(res *TsTreeNode) {
	res = new(TsTreeNode)
	res.Childs = make([]*TsTreeNode, c, l)
	return res
}

func (this *TsTreeNode)ChildCount()(count int64) {
	for i:=0; i<len(this.Childs); i++ {
		if this.Childs[i]!=nil {
			count++
		}
	}
	return
}

func (this *TsTreeNode)Clear()() {
	for i:=0; i<len(this.Childs); i++ {
		if this.Childs[i]!=nil {
			this.Childs[i].Clear()
			this.Childs[i] = nil
		}
	}
	this.Data = nil
}

func (this *TsTreeNode)Add(item *TsTreeNode)() {
	for i:=0; i<len(this.Childs); i++ {
		if this.Childs[i]==nil {
			this.Childs[i] = item
			return
		}
	}
	
	this.Childs = append(this.Childs, item)
}