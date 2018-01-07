package tsBT

import (
	"github.com/astaxie/beego"
)

type BTNodeFalse struct {
	BTNode
}

func NewBTNodeFalse(bt *BT)(*BTNodeFalse) {
	res := new(BTNodeFalse)
	res.Bt = bt
	return res
}

func (this *BTNodeFalse)Execute(black_board *BTBlackboard)(bool) {
	if this.IsLog() {
		beego.Trace("------------------>", this.Name)
	}
	
	return false
}