package tsBT

import (
	"github.com/astaxie/beego"
)

type BTNodeAnd struct {
	BTNode
	Childs []IBTNode
}

func NewBTNodeAnd(bt *BT)(*BTNodeAnd) {
	res := new(BTNodeAnd)
	res.Bt = bt
	return res
}

func (this *BTNodeAnd)AddChild(n IBTNode)(bool) {
	this.Childs = append(this.Childs, n)
	return true
}

func (this *BTNodeAnd)Execute(black_board *BTBlackboard)(bool) {
	if this.IsLog() {
		beego.Trace("------------------>", this.Name)
	}
		
	for i:=0; i<len(this.Childs); i++ {
		if this.Childs[i].Execute(black_board)==false {
			return false
		}
	}
	return true
}