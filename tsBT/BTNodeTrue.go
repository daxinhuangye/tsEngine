package tsBT

import (
	"github.com/astaxie/beego"
)

type BTNodeTrue struct {
	BTNode
}

func NewBTNodeTrue(bt *BT)(*BTNodeTrue) {
	res := new(BTNodeTrue)
	res.Bt = bt
	return res
}

func (this *BTNodeTrue)Execute(black_board *BTBlackboard)(bool) {
	if this.IsLog() {
		beego.Trace("------------------>", this.Name)
	}
		
	return true
}