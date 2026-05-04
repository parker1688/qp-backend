package tool

import (
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func RealPath(f string) string {
	p, err := filepath.Abs(f)
	if err != nil {
		log.Panicln("Get absolute path error.")
	}
	p = strings.Replace(p, "\\", "/", -1)
	l := strings.LastIndex(p, "/") + 1
	return Substr(p, 0, l)
}

func ReadFileByte(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func IsExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func IsFile(f string) bool {
	b, err := os.Stat(f)
	if err != nil {
		return false
	}
	if b.IsDir() {
		return false
	}
	return true
}

func IsDir(p string) bool {
	b, err := os.Stat(p)
	if err != nil {
		return false
	}
	if b.IsDir() {
		return true
	}
	return false
}

func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}

func CreatePath(path string, perm os.FileMode) error {
	folder, _ := filepath.Split(path)
	if !IsExist(folder) {
		err := os.MkdirAll(folder, perm)
		if err != nil {
			return err
		}
	}
	return nil
}

func Base64ToFile(data, path string) error {
	folder, _ := filepath.Split(path)
	if !IsExist(folder) {
		err := os.MkdirAll(folder, os.ModePerm)
		if err != nil {
			return err
		}
	}
	decodeData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return err
	}
	err = os.WriteFile(path, decodeData, 0666)
	if err != nil {
		return err
	}
	return nil
}

// 新建文件
// data要写入的数据
// folder文件夹目录
// file文件名
func CreateFile(data, path string) error {
	folder, _ := filepath.Split(path)
	if !IsExist(folder) {
		err := os.MkdirAll(folder, os.ModePerm)
		if err != nil {
			return err
		}
	}
	err := os.WriteFile(path, []byte(data), 0666)
	if err != nil {
		return err
	}
	return nil
}

func GetAllFileName(pathname string, s []string) ([]string, error) {
	entries, err := os.ReadDir(pathname)
	if err != nil {
		fmt.Println("read dir fail:", err)
		return s, err
	}
	for _, fi := range entries {
		if !fi.IsDir() {
			s = append(s, fi.Name())
		}
	}
	return s, nil
}

func FileMerge(inDir, infileName, outfile string, perm os.FileMode) error {
	fii, err := os.OpenFile(outfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, perm)
	if fii != nil {
		defer fii.Close()
	}
	if err != nil {
		return err
	}
	fileNames := make([]string, 0)
	fileNames, err = GetAllFileName(inDir, fileNames)
	if err != nil {
		return err
	}
	for i := 0; i < len(fileNames); i++ {
		f, err := os.OpenFile(inDir+"/"+infileName+"_"+String(i), os.O_RDONLY, os.ModePerm)
		if err != nil {
			return err
		}
		b, err := io.ReadAll(f)
		if err != nil {
			return err
		}
		_, err = fii.Write(b)
		if err != nil {
			return err
		}
		err = f.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func WriteFile(filename string, data []byte, perm os.FileMode) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}
