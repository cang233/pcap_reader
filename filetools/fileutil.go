package filetools

import (
	"io/ioutil"
	"strings"
)

//ExtractRealName extract file name from "xx/xx/xxx/aa.xx" and return "aa",return aa.xx if needSuffix is true
func ExtractRealName(absFile string, needSuffix bool) string {
	if len(absFile) == 0 {
		return ""
	}
	paths := strings.Split(strings.TrimSpace(absFile), GetOsSeperator())
	if needSuffix {
		return paths[len(paths)-1]
	}
	// aaa.xx
	lasts := strings.Split(strings.TrimSpace(paths[len(paths)-1]), ".")
	return strings.Join(lasts[:len(lasts)-1], ".")
}

//WriteFile write data string to filePath
func WriteFile(filePath string, data string) {
	ioutil.WriteFile(filePath, []byte(data), 0666)
}
