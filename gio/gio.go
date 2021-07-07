package gio

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// Validate file exist
func IsFileExist(dir string) (bool, error) {
	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		return true, err
	}
	return false, nil
}

// Create folder
func CreateFolder(dir string) {
	_, err := IsFileExist(dir)
	checkError(err)
	err = os.Mkdir(dir, 0755)
	checkError(err)

}

func CreateDirAll(dir string) {
	_, err := IsFileExist(dir)
	checkError(err)

	err = os.MkdirAll(dir, 0755)
	checkError(err)

}

//Create file
func CreateFile(finalPath string, data string) {
	err := ioutil.WriteFile(finalPath, []byte(data), 0644)
	checkError(err)

}

// check error
func checkError(err error) {
	if err != nil {
		os.Exit(1)
	}
}

// read file and return string with data
func ReadFile(templateName string) string {
	data, err := ioutil.ReadFile(templateName)
	if err != nil {
		checkError(err)
	}
	return string(data)
}

// copy file with operator system
func Copy(srcFileDir string, destFileDir string) {
	srcFile, err := os.Open(srcFileDir)
	checkError(err)
	defer srcFile.Close()
	destFile, err := os.Create(destFileDir)
	checkError(err)

	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	checkError(err)

	err = destFile.Sync()

	checkError(err)
}

// read file and copy to other file
func ReadAndCopy(srcFileDir string, destFileDir string) {
	data, err := ioutil.ReadFile(srcFileDir)
	checkError(err)

	err = ioutil.WriteFile(destFileDir, data, 0644)
	checkError(err)
}

// modify a file looking for something
func ChangeFile(templateName string, MapForReplace map[string]string) {
	data := ReplaceTextInFile(templateName, MapForReplace)
	CreateFile(templateName, data)
}

// Create a new file starting from a template and an array of options to replace
func NewFileforTemplate(newName string, templateName string, MapForReplace map[string]string) {
	data := ReplaceTextInFile(templateName, MapForReplace)
	CreateFile(newName, data)
}

// replace info in file return string
func ReplaceTextInFile(templateName string, MapForReplace map[string]string) string {
	input := ReadFile(templateName)

	for key, value := range MapForReplace {
		input = strings.Replace(input, key, value, -1)
	}
	return input

}

// append string to end of the file
func AppEndToFile(destFileDir string, data string) {
	dataInFile, err := os.OpenFile(destFileDir, os.O_APPEND|os.O_WRONLY, 0600)
	checkError(err)

	defer dataInFile.Close()

	if _, err = dataInFile.WriteString(data); err != nil {
		checkError(err)
	}
}

// add many new lines
func AppendArrayEndToFile(destFileDir string, datas []string) {
	for _, data := range datas {
		AppEndToFile(destFileDir, data)
	}
}
