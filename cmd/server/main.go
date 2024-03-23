package main

import (
	"fmt"

	config "github.com/eduardogomesf/shopping/configs"
)

func main() {
	conf := config.LoadConfig(".")
	fmt.Printf("%+v\n", conf)
}
