package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fi, _ := os.Open("../comlast.json")
	fii, _ := os.Open("../comslast.json")

	defer fi.Close()
	defer fii.Close()

	br := bufio.NewReader(fi)
	brr := bufio.NewReader(fii)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		str := fmt.Sprintf("%s", a)
		for {
			aa, _, cc := brr.ReadLine()
			if cc == io.EOF {
				break
			}
			strr := fmt.Sprintf("%s", aa)
			if str == strr {
				break
			}

		}

		if strings.Contains(str, "/command/") {
			fpp.WriteString(string(str))
			fpp.WriteString("\n")
		}

		fmt.Println(string(a))
	}
}
