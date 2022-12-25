package main

import (
	"fmt"
	"os"

	"github.com/aland20/go-noting/app/cmd"
	"github.com/aland20/go-noting/app/utils"
)

func main() {

	utils.LoadEnv()

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
