package main

import (
	"github.com/shizhigu/GoTutorial/sum"
	"github.com/shizhigu/GoTutorial/handler"
	"github.com/micro/go-micro"
)

func main() {
	srv := micro.NewService(
		micro.Name("go.micro.srv.sum_service"),
	)

	srv.Init()

	_ = sum.RegisterSumHandler(srv.Server(), handler.Handler())

	srv.Run()
	
	// if err := srv.Run(); err != nil {
	// 	panic(err)
	// }

}