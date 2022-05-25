package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/buger/jsonparser"
)

func main() {
	fmt.Println("tses")
	data, err := os.ReadFile("./d37.json")
	if err != nil {
		fmt.Println(err)
	}
	//datalsit, _, _, _ := jsonparser.Get(data, "result")
	//fmt.Println(string(datalsit))

	// for _, v := range datalsit {
	// 	jsonparser.Get([]byte(v), "names")
	// }

	fp, _ := os.OpenFile("./d37s.json", os.O_RDWR|os.O_CREATE, 0666)
	defer fp.Close()

	jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		datalong, _, _, _ := jsonparser.Get(value, "perms")
		fmt.Println(string(datalong))
		fp.WriteString(string(datalong))
		fp.WriteString("\n")
	}, "result")

	fi, err := os.Open("./d37s.json")
	fpp, _ := os.OpenFile("./d37last.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		str := fmt.Sprintf("%s", a)
		if strings.Contains(str, "/command/") {
			fpp.WriteString(string(str))
			fpp.WriteString("\n")
		}

		fmt.Println(string(a))
	}

}
