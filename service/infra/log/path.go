package log

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// 当前项目根目录
var API_ROOT string

// 获取项目路径
func GetPath() string {

	if API_ROOT != "" {
		return API_ROOT
	}

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		print(err.Error())
	}

	API_ROOT = strings.Replace(dir, "\\", "/", -1)
	return API_ROOT
}

// 判断文件目录否存在
func IsDirExists(path string) bool {
	fi, err := os.Stat(path)

	if err != nil {
		return os.IsExist(err)
	} else {
		return fi.IsDir()
	}

}

func MkFile(fileName string) error {
	file, err := os.Create(fileName)
	defer file.Close()
	return err
}

func WriteFile(filePath string, byteValue []byte) error {
	os.Truncate(filePath, 0)
	return ioutil.WriteFile(filePath, byteValue, 0644)
}

func ReadFile(filePath string) ([]byte, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

// 创建文件目录
func MkdirFile(path string) error {

	err := os.Mkdir(path, os.ModePerm) //在当前目录下生成md目录
	if err != nil {
		return err
	}
	return nil
}

func EnsureDir(path string) error {
	err := MkdirFile(path)

	if err == nil || os.IsExist(err) {
		return nil
	}
	return err
}

func IsFileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}
	return false
}

func CopyFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func AssemblePath(args ...string) string {
	return strings.Join(args, "/")
}

func RemoveTmpFile(filePath string) error {
	var err error
	if IsFileExists(filePath) {
		err = os.Remove(filePath)
	}
	return err
}
