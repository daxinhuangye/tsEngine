package tsBT

import (
	"github.com/astaxie/beego"
)

type BTNodeOr struct {
	BTNode
	Childs []IBTNode
}

func NewBTNodeOr(bt *BT)(*BTNodeOr) {
	res := new(BTNodeOr)
	res.Bt = bt
	return res
}

func (this *BTNodeOr)AddChild(n IBTNode)(bool) {
	this.Childs = append(this.Childs, n)
	return true
}

func (this *BTNodeOr)Execute(black_board *BTBlackboard)(bool) {
	if this.IsLog() {
		beego.Trace("------------------>", this.Name)
	}
		
	for i:=0; i<len(this.Childs); i++ {
		if this.Childs[i].Execute(black_board)==true {
			return true
		}
	}
	return false
}