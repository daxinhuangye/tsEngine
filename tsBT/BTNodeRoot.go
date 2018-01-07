package tsBT

import (
	"github.com/astaxie/beego"
)

type BTNodeRoot struct {
	BTNode
	Child IBTNode
}

func NewBTNodeRoot(bt *BT)(*BTNodeRoot) {
	res := new(BTNodeRoot)
	res.Bt = bt
	return res
}

func (this *BTNodeRoot)AddChild(n IBTNode)(bool) {
	this.Child = n
	return true
}

func (this *BTNodeRoot)Execute(black_board *BTBlackboard)(res bool) {
	if this.IsLog() {
		beego.Trace("------------------>", this.Name)
	}
	
	if this.Child==nil {
		return false
	}
	res = this.Child.Execute(black_board)
	return
}