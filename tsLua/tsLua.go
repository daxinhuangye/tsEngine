package tsLua

import (
	"github.com/aarzilli/golua/lua"
	"github.com/stevedonovan/luar"
	
	"tsEngine/tsFile"
	"errors"
)

type TsLua struct {
	LuaObj *lua.State
	
	LuaObjUnit map[string]*luar.LuaObject
	
	RegMap luar.Map
	TabelName string
}

func NewLua()(res *TsLua) {
	res = &TsLua{}
	return
}

/*
 * @brief 初始化lua
 */
func (this *TsLua)Init() {
	this.LuaObj = luar.Init()
	this.LuaObjUnit = make(map[string]*luar.LuaObject)
	this.RegMap = make(luar.Map)
}

/*
 * @brief 关闭lua
 */
func (this *TsLua)Close() {
	if this.LuaObj!=nil {
		this.LuaObj.Close()
		this.LuaObj = nil
		this.LuaObjUnit = make(map[string]*luar.LuaObject)
	}
}

/*
 * @brief 添加go中的lua绑定
 *
 * @param key 索引，lua中直接使用该索引调用函数
 * @param value go中的函数、指针
 * @return nil无错误；错误说明；
 *
 * @see DoReg
 */
func (this *TsLua)AddReg(key string, value interface{})(err error) {
	_,ok:=this.RegMap[key]
	if ok {
		return errors.New("have key")
	}
	this.RegMap[key] = value
	return
}

/*
 * @brief 完成注册go
 *
 * @see AddReg
 */
func (this *TsLua)DoReg(name string) {
	this.TabelName = name
	luar.Register(this.LuaObj, this.TabelName, this.RegMap)
}

/*
 * @brief 清空lua；重新注册go函数；加载目录中的所有lua文件
 *
 * @param dir 文件所在目录
 */
func (this *TsLua)ReloadFile(dir []string) {
	this.Close()
	
	this.LuaObj = luar.Init()
	this.LuaObjUnit = make(map[string]*luar.LuaObject)
	this.DoReg(this.TabelName)
	this.LoadFile(dir)	
}

// 加载文件夹下的lua文件
func (this *TsLua)LoadFile(dir []string) {
	for i:=0; i<len(dir); i++ {
		file_list, err := tsFile.Filelist(dir[i])
		if err!=nil {
			return
		}
		
		for j:=0; j<len(file_list); j++ {
			this.LuaObj.DoFile(file_list[j])
		}
	}
}

// 获得lua对象
func (this *TsLua)GetObjectUnit(name string)(res *luar.LuaObject) {
	res,ok := this.LuaObjUnit[name]
	if ok {
		return
	}
	
	res = luar.NewLuaObjectFromName(this.LuaObj, name)
	if res==nil {
		return 
	}
	this.LuaObjUnit[name] = res
	return
}

func (this *TsLua)Call(name string, args ... interface{})(res interface{}, err error) {
	l := this.GetObjectUnit(name)
	if l==nil {
		return 
	}
	res, err = l.Call(args ...)
	return
}