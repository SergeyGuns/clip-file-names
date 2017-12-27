package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/atotto/clipboard"
)

func main() {
	var outPut bytes.Buffer
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		outPut.WriteString(file.Name())
		outPut.WriteString("\n")
	}
	fmt.Println(outPut.String())
	clipboard.WriteAll(outPut.String())
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
