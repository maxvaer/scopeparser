package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func checkError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return err != nil
}

func createFile(Path string, Name string) *os.File {
	_ = os.Remove(Path + Name)
	file, err := os.Create(Path + Name)
	if checkError(err) {
		fmt.Println("Something went wrong while creating files:" + err.Error())
		return nil
	}
	return file
}

func parseJsonToFile(File *os.File, Json gjson.Result, addPorts bool) {
	reg, err := regexp.Compile("[^a-zA-Z0-9.]+")
	checkError(err)
	for _, currentJson := range Json.Array() {
		host := currentJson.Get("host").String()
		host = reg.ReplaceAllString(host, "")
		host = strings.Replace(host, "..", "", -1)

		port := currentJson.Get("port").String()
		port = ":" + reg.ReplaceAllString(port, "")

		protocol := currentJson.Get("protocol").String()
		protocol += "://"

		if addPorts {
			_, _ = fmt.Fprintln(File, protocol+host+port)
		} else {
			_, _ = fmt.Fprintln(File, protocol+host)
		}
	}
	_ = File.Close()
}

func getPath() string {
	dir, err := os.Getwd()
	checkError(err)
	dir += "/"

	return dir
}

func readFromStdIn() string {
	scanner := bufio.NewScanner(os.Stdin)
	file := ""
	for scanner.Scan() {
		file += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return file
}

func main() {
	var sourceFile string
	flag.StringVar(&sourceFile, "f", "scope.json", "Path to the Burp Project Configuration File.")

	var addPorts bool
	flag.BoolVar(&addPorts, "p", false, "Add Ports end the end, e.g http://example.site:80, default = false.")

	flag.Parse()

	fileRead, err :=ioutil.ReadFile(sourceFile)
	jsonString := string(fileRead)
	if checkError(err) {
		jsonString = readFromStdIn()
	}
	if !gjson.Valid(jsonString){
		fmt.Println(errors.New("invalid json"))
		os.Exit(1)
	}

	excludedJson := gjson.Get(jsonString,"target.scope.exclude")
	includedJson := gjson.Get(jsonString,"target.scope.include")

	dir := getPath()

	scope := createFile(dir, "scope.txt")
	excluded := createFile(dir, "excluded.txt")

	parseJsonToFile(scope, includedJson, addPorts)
	parseJsonToFile(excluded, excludedJson, addPorts)
}
