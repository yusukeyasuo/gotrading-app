package main

import (
	"fmt"

	"github.com/yusukeyasuo/gotrading-app/config"
)

func main() {
	fmt.Println(config.Config.ApiKey)
	fmt.Println(config.Config.ApiSecret)
}