package main

import (
	"fmt"
	"os"

	"github.com/buger/jsonparser"
)

func main() {
	fmt.Println("tses")
	data, err := os.ReadFile("./auth.json")
	if err != nil {
		fmt.Println(err)
	}
	//datalsit, _, _, _ := jsonparser.Get(data, "result")
	//fmt.Println(string(datalsit))

	// for _, v := range datalsit {
	// 	jsonparser.Get([]byte(v), "names")
	// }

	fp, _ := os.OpenFile("./com.json", os.O_RDWR|os.O_CREATE, 0666)
	defer fp.Close()

	jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		datalong, _, _, _ := jsonparser.Get(value, "perms")
		fmt.Println(string(datalong))
		fp.WriteString(string(datalong))
		fp.WriteString("\n")
	}, "result")
}
