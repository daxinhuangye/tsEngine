package tsBT

import (
	"tsEngine/tsFile"
)

type BTMgr struct {
	BTs map[string]*BT
	DirPath string
}

func NewBTMgr()(res *BTMgr) {
	res = new(BTMgr)
	res.BTs = make(map[string]*BT)
	return
}

func (this *BTMgr)Load(dir string)(err error) {
	this.DirPath = dir
	
	file_list, err := tsFile.Filelist(this.DirPath)
	if err!=nil {
		return
	}
	
	for i:=0; i<len(file_list); i++ {
		data, err := tsFile.ReadFile(file_list[i])
		if err!=nil || len(data)<=0 {
			continue
		}
		
		bt := NewBT()
		bt.FromJson(data)
		bt.Name = bt.Root.GetName()
		this.BTs[bt.Root.GetName()] = bt
	}
	return
}

func (this *BTMgr)Reload()(err error) {
	return this.Load(this.DirPath)
}

func (this *BTMgr)GetBt(key string)(bt *BT){
	bt, _ = this.BTs[key]
	return
}