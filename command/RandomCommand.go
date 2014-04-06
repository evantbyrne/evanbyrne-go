package main

import (
	"fmt"
	"github.com/evantbyrne/evanbyrne-go/util"
	"os"
	"strconv"
)

func main() {
	if i, err := strconv.Atoi(os.Args[1]); err == nil {
		fmt.Println(util.Random(i));
	}
}