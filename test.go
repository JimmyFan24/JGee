package main

import (
	"fmt"
	"strings"
)

func ParsePattern(pttern string)  []string{
	vs := strings.Split(pttern,"/")
	parts := make([]string,0)
	for _,item := range vs {
		if item != ""{
			fmt.Println(item)
			parts =append(parts,item)
			fmt.Println("----------",item[0])
			if item[0] == '*'{
				break
			}
		}

	}

	return parts
}
type nodetest struct {
	pattern string
	part string
	children []*nodetest
	isWild bool
}

