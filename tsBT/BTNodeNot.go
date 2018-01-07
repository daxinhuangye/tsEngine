package tsBT

import (
	"github.com/astaxie/beego"
)

type BTNodeNot struct {
	BTNode
	Childs []IBTNode
}

func NewBTNodeNot(bt *BT)(*BTNodeNot) {
	res := new(BTNodeNot)
	res.Bt = bt
	return res
}

func (this *BTNodeNot)AddChild(n IBTNode)(bool) {
	this.Childs = append(this.Childs, n)
	return true
}

func (this *BTNodeNot)Execute(black_board *BTBlackboard)(bool) {
	if this.IsLog() {
		beego.Trace("------------------>", this.Name)
	}
		
	for i:=0; i<len(this.Childs); i++ {
		if this.Childs[i].Execute(black_board)==true {
			return false
		} else {
			return true
		}
	}
	return false
}