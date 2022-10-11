package main

import (
	"fmt"
	
	config "assignment-2/config/database"
	"assignment-2/router"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`.env`)
	
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool("DEBUG") {
		fmt.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	db := config.DBInit()
	route := router.Route(db)

	route.Run()
}
