package tsBT

import (
	"github.com/astaxie/beego"
)

type BTNodeMap struct {
	BTNode
	Childs map[string]IBTNode
	BlackboardKey string
}

func NewBTNodeMap(bt *BT)(*BTNodeMap) {
	res := new(BTNodeMap)
	res.Childs = make(map[string]IBTNode)
	res.Bt = bt
	return res
}

func (this *BTNodeMap)AddChild(n IBTNode)(bool) {
	this.Childs[n.GetIndex()] = n
	return true
}

func (this *BTNodeMap)Execute(black_board *BTBlackboard)(bool) {
	if this.IsLog() {
		beego.Trace("------------------>", this.Name)
	}
		
	i_key,ok := black_board.Field[this.BlackboardKey]
	if !ok {
		beego.Error("BTNodeMap ", this.BlackboardKey, " no find on black_board")
		return false
	}
	num_key := i_key.(string)
	node, ok:=this.Childs[num_key]
	if !ok {
		//beego.Error("BTNodeMap no find child node")
		return false
	}
	node.Execute(black_board)
	return true
}

func (this *BTNodeMap)FromJson(data map[string]interface{})(bool) {
	i_BlackboardKey,ok := data["BlackboardKey"]
	if !ok {
		return false
	}
	this.BlackboardKey = i_BlackboardKey.(string)
	return this.BTNode.FromJson(data)
}