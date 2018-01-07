package tsFile

import (
	"encoding/base64"
	"encoding/csv"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"tsEngine/tsCrypto"
	"tsEngine/tsString"
)

/*
* 读取文本文件
* path:文件路径
* return：文件数据，错误
 */
func Filelist(path string) ([]string, error) {

	var list = []string{}

	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		list = append(list, path)
		return nil
	})
	if err != nil {
		return list, err
	}

	return list, nil
}

func ReadFile(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

//自动创建文件夹 并写入文件
func WriteFile(path string, data string) error {
	return WriteFileByte(path, []byte(data))
}

func WriteFileByte(path string, data []byte) error {
	// 查询文件是否存在
	_, err := os.Stat(path)
	//如果文件不存在则创建文件夹和文件
	if err != nil {
		path = strings.Replace(path, "./", "", -1)
		temp := strings.Split(path, "/")
		dir := "./"
		for i := 0; i < len(temp)-1; i++ {
			dir += temp[i] + "/"
		}
		//判断文件夹是否存在如果不存在则创建文件夹
		_, err = os.Stat(dir)
		if err != nil {
			err = os.MkdirAll(dir, 0777)
			if err != nil {
				return err
			}
		}
	}
	return ioutil.WriteFile(path, data, os.ModeAppend)
}

//创建base64图片文件
func WriteImgFile(path string, filename string, data string) (string, error) {

	ext, img_data := ".jpg", ""

	if strings.Contains(data, ",") {
		temp := strings.Split(data, ",")
		if strings.Contains(temp[0], "png") {
			ext = ".png"
		} else if strings.Contains(temp[0], "gif") {
			ext = ".gif"
		}
		img_data = temp[1]

	} else {
		img_data = data
	}

	img, err := base64.StdEncoding.DecodeString(img_data)
	if err != nil {
		return "", err
	}

	//img := img_data
	filename = tsCrypto.GetMd5([]byte(filename))

	for i := 1; i <= 3; i++ {
		path += tsString.Substr(filename, 0, i) + "/"
	}
	path += filename + ext
	err = WriteFile(path, string(img))

	if err != nil {
		return "", err
	}
	path = strings.Replace(path, "./", "/", -1)
	return path, nil
}

/*
 * @brief 根据filename的md5值，创建文件路径并保存文件。文件路径（path+md5[1]/md5[2]/md5[3]/md5+ext）
 * @param path 文件保存路径(相对路径 ./static/)
 * @param filename 文件名称
 * @param ext 文件扩展名
 * @param data 文件数据
 * @return 文件路径、错误
 */
func WriteFileFromBase64(path string, filename string, ext string, data string) (string, error) {

	file_string := ""

	if strings.Contains(data, ",") {
		temp := strings.Split(data, ",")
		file_string = temp[1]
	} else {
		file_string = data
	}

	file_data, err := base64.StdEncoding.DecodeString(file_string)
	if err != nil {
		return "", err
	}

	//img := img_data
	filename = tsCrypto.GetMd5([]byte(filename))

	for i := 1; i <= 3; i++ {
		path += tsString.Substr(filename, 0, i) + "/"
	}
	path += filename + ext
	err = WriteFile(path, string(file_data))

	if err != nil {
		return "", err
	}
	path = strings.Replace(path, "./", "/", -1)
	return path, nil
}

func ReadJsonFile(path string, ob interface{}) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, ob)
	if err != nil {
		return err
	}
	return nil
}

func WriteJsonFile(path string, ob interface{}) error {
	data, err := json.Marshal(ob)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, data, os.ModeAppend)
}

func ReadCsvFile(path string) ([][]string, error) {
	data, err := ReadFile(path)
	if err != nil {
		return nil, err
	}
	str_read := strings.NewReader(data)
	csv_read := csv.NewReader(str_read)
	return csv_read.ReadAll()
}

func DelFile(path string) error {
	return os.Remove(path) //删除文件
}

func Md5FromFile(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	md5 := tsCrypto.GetMd5(data)
	return md5, nil
}

func FileNameAndExt(file string) (string, string) {
	pos := strings.LastIndex(file, ".")
	if pos < 0 {
		return "", ""
	}
	data := []byte(file)
	return string(data[:pos]), string(data[pos+1:])
}

func CopyFile(src string, des string) (err error) {
	out, err := os.Create(des)
	if err != nil {
		return
	}
	defer out.Close()
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()

	_, err = io.Copy(out, in)
	return
}

func DownLoadFile(url string, filename string) bool {
	res, err := http.Get(url)
	defer res.Body.Close()
	if err != nil {
		return false
	}
	// 创建文件
	dst, err := os.Create(filename)
	if err != nil {
		return false
	}
	// 生成文件
	io.Copy(dst, res.Body)
	return true
}
