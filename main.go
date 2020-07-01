package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
)

func main() {
	input := ""
	if len(os.Args) > 1 {
		input = os.Args[1]
	}

	fmt.Println(strconv.FormatInt(int64(hashcode.String(input)), 10))
}
