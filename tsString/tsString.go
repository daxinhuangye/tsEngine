package tsString

import (
	"fmt"
	"github.com/axgle/mahonia"
	"strconv"
	"strings"
)

func And(src int64, data int64) int64 {
	return src & data
}
func Or(src int64, data int64) int64 {
	return src | data
}

func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}

func ToInt(str string, start int, length int) int {
	num, err := strconv.ParseInt(Substr(str, start, length), 10, 64)
	if err != nil {
		return 0
	}
	return int(num)
}

func ToInt64(str string) int64 {
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return num
}

func FromInt64(v int64) string {
	return fmt.Sprintf("%d", v)
}

func Split(str string, sep string) []string {
	return strings.Split(str, sep)
}

//删除 byte为32的（空格），和左右的（空格）
func TrimSpace(str string) string {
	str = strings.TrimSpace(str)
	str = strings.Replace(str, " ", "", -1)
	return str
}

//删除左右空格
func TrimLrSpace(str string) string {
	str = strings.TrimLeft(str, " ")
	str = strings.TrimRight(str, " ")
	return str
}

func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}
