package main

import (
	"github.com/aland20/go-noting/app/cmd"
	"github.com/aland20/go-noting/app/utils"
)

func main() {

	utils.LoadEnv()
	cmd.Execute()
}
