package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	c := strings.Split(os.Args[1], "-")
	re := regexp.MustCompile("..")
	d := re.FindAllString(c[3]+c[4], -1)
	for i, b := range d {
		d[i] = "0x" + b
	}
	fmt.Printf("&ole.GUID{0x%s, 0x%s, 0x%s, [8]byte{%s}}\n", c[0], c[1], c[2], strings.Join(d, ", "))
}
