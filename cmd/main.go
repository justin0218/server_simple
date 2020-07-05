package main

import (
	"fmt"
	"server_simple/api"
	"server_simple/configs"
	"server_simple/internal/routers"
	"server_simple/job"
)

func main() {
	api.Log.Get().Debug("starting...")
	jobs := job.NewJob()
	go jobs.BlogChatRecord()
	go jobs.GdrawData()
	err := api.Rds.Get().Ping().Err()
	if err != nil {
		panic(err)
	}
	api.Mysql.Get()
	fmt.Println("server started!!!")
	err = routers.Init().Run(fmt.Sprintf(":%d", configs.Dft.Get().Http.Port))
	if err != nil {
		panic(err)
	}
}
