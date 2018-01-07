package tsContain

import (
	"errors"
	"fmt"
)


type MapStrMgr struct {
	ObjMgr map[string]interface{}
	count int
}

func NewMapStrMgr()(*MapStrMgr) {
	res := new(MapStrMgr)
	res.ObjMgr = make(map[string]interface{})
	
	return res
}

func (this *MapStrMgr)AddInt64(key int64, obj interface{})(error) {
	k := fmt.Sprintf("%d", key)
	return this.Add(k, obj)
}

func (this *MapStrMgr)Add(key string, obj interface{})(error) {
	_, ok := this.ObjMgr[key]
	if ok {
		return errors.New("have obj")
	}
	this.ObjMgr[key] = obj
	this.count++
	return nil
}

func (this *MapStrMgr)GetByInt(key int64)(interface{}) {
	k := fmt.Sprintf("%d", key)
	return this.Get(k)
}

func (this *MapStrMgr)Get(key string)(interface{}) {
	obj, ok := this.ObjMgr[key]
	if !ok {
		return nil
	}
	if obj==nil {
		return nil
	}
	return obj
}

func (this *MapStrMgr)Have(key string)(bool) {
	_, ok := this.ObjMgr[key]
	return ok
}

func (this *MapStrMgr)Del(key string)() {
	if this.Have(key) {
		delete(this.ObjMgr, key)
		this.count--
	}
}

func (this *MapStrMgr)Count()(count int) {
	return this.count
}

func (this *MapStrMgr)ToArray()(res []interface{}) {
	res = make([]interface{}, 0)
	for _,v:=range this.ObjMgr {
		res = append(res, v)
	}
	return
}

func (this *MapStrMgr)Clear()() {
	for _,v:=range this.ObjMgr {
		if v!=nil {
			v = nil
		}
	}	
	this.ObjMgr = make(map[string]interface{})
	this.count = 0
}