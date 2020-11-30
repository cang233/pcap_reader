package filetools

import (
	"io/ioutil"
	"os"
	"runtime"
	"strings"
)

//ListFiles list all pcaps in dir.
func ListFiles(dir string, suffix string) []string {
	rds, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	dir = CheckDirSuffix(dir, true)
	var pcaps []string
	for i := range rds {
		if !rds[i].IsDir() && strings.HasSuffix(rds[i].Name(), suffix) {
			pcaps = append(pcaps, dir+rds[i].Name())
		}
	}
	return pcaps
}

//GetOsSeperator get file path seprator according to its system type.
func GetOsSeperator() string {
	if runtime.GOOS == "windows" {
		return "\\"
	} else {
		return "/"
	}
}

//CheckDirSuffix check dir suffix,and add or remove spliter using needSpliter
func CheckDirSuffix(dir string, needSpliter bool) string {
	//check dir suffix, linux
	if runtime.GOOS == "windows" {
		if !strings.HasSuffix(dir, "\\") {
			dir += "\\\\"
		}
	} else {
		if !strings.HasSuffix(dir, "/") {
			dir += "/"
		}
	}
	//till now ,the dir is /xx/xxx/xxxx/

	//if dont need spliter,then remove it.
	if !needSpliter {
		if runtime.GOOS == "windows" {
			dir = dir[:len(dir)-2]
		} else {
			dir = dir[:len(dir)-1]
		}
	}

	return dir
}

//CheckAndMakeDir check dir if exist,and make it if not exist, or remake it if exist.
func CheckAndMakeDir(dir string, needMake, needReMake bool) bool {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if needMake {
			os.MkdirAll(dir, os.ModePerm)
			os.Chmod(dir, os.ModePerm)
		}
		return false
	} else {
		// remake == delete all
		if needReMake {
			items, _ := ioutil.ReadDir(dir)
			for _, d := range items {
				os.RemoveAll(CheckDirSuffix(dir, true) + d.Name())
			}
		}
		return true
	}
}
