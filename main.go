package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

const certPath string = "ca.crt"

func main() {
	info, err := os.Stat(certPath)
	if os.IsNotExist(err) {
		return
	}
	if info.IsDir() {
		return
	}

	f, err := os.Open(certPath)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	fmt.Println(base64.StdEncoding.EncodeToString(content))
}
