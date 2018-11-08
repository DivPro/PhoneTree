package main

import (
	"fmt"
	"mobile/operator"
	"os"
)

var tree = operator.NewTree(operator.NewTreeNode("9", 1, nil))
var parser = operator.NewParser(os.Args[1])

func main() {
fmt.Println(os.Args)

	parser.Parse()

	tree.AddRange("9198800001", "9399299999", "tele2")
	tree.AddRange("9198800004", "9399299998", "beeline")
	tree.AddPhone("9266512666", "megafon")

	fmt.Println(tree.Find("9266512666"))
}
