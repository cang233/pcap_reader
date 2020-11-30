package reader

import (
	"fmt"
	"pcap-reader/filetools"
	"pcap-reader/parser"
	"strings"
)

//Parse parse all pcap in dir and output its data into saving dir.
func Parse(pcapDir string, savingDir string) {
	files := filetools.ListFiles(pcapDir, ".pcap")
	for i := range files {
		fmt.Printf("------ Reading pcap : %s ---------------\n", files[i])
		handlePcap(files[i], savingDir)
	}
}

//createPcapSavingDir create pcap data saving dir and return abspath.
func createPcapSavingDir(mainDir, filePath string) string {
	// extract real filename from file path
	seps := strings.Split(filePath, filetools.GetOsSeperator())
	fileName := strings.TrimSpace(seps[len(seps)-1])
	fileName = strings.TrimSuffix(fileName, ".pcap")
	absPath := filetools.CheckDirSuffix(mainDir, true) + fileName
	filetools.CheckAndMakeDir(absPath, true, true)
	return absPath
}

func handlePcap(filePath string, outputDir string) {
	outDir := createPcapSavingDir(outputDir, filePath)
	parser.Handle(filePath, outDir)
}
