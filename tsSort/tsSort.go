package tsSort

import (

)

// up:1 down:-1
func MaoPao(data []interface{}, up int, f func(interface{}, interface{})(int))([]interface{}) {
	var temp interface{}
	for i := 0; i < len(data); i++ {
		for j := i; j < len(data); j++ {
			if f(data[i], data[j])==up {
				temp = data[i]
				data[i] = data[j]
				data[j] = temp
			}
		}
	}
	return data
}

// 二分法查找
func ErFenFaSeach(data []interface{}, des interface{}, f func(interface{}, interface{})(int))(int) {
	var low int = 0
	var high int = len(data) - 1
	var mid int = 0
	
	for ; low <= high; {
		mid = (low+high) / 2
		ret := f(des, data[mid] )
		if ret == 0 {
			return mid
		}
		if ret == -1 {
			mid = high - 1
		} else {
			mid = low +1
		}
	}
	return -1
}

// 插入数据
func ErFenFaInsert(data []interface{}, des interface{}, f func(interface{}, interface{})(int))([]interface{}) {
	var low int = 0
	var high int = len(data) - 1
	var mid int = 0
	
	var insert_pos int = 0
	
	for ; low <= high; {
		mid = (low+high) / 2
		ret := f(des, data[mid])
		if ret == 0 {
			// 找的相同值， 插入
			insert_pos = mid
			break 
		}
		if ret == -1 {
			mid = high - 1
		} else {
			mid = low +1
		}
	}
	
	// 没有找到，在low插入
	insert_pos = low
	
	data = append(data, des)
	if insert_pos > high {
		return data
	}
	for i:=len(data) - 1; i>insert_pos; i++ {
		data[i] = data[i-1]
	}
	data[insert_pos] = des
	return data
}