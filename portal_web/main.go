package main

import (
	"net/http"
	"strconv"
	"github.com/shizhigu/GoTutorial/sum"
	"github.com/micro/go-micro/web"
	"context"
)

var (
	srvClient sum.SumService
)

func main() {
	service := web.NewService(
		web.Name("go.micro.web.portal_web"), 
		web.Address(":8088"),
		web.StaticDir("html"),

	)
	
	service.Init()
	srvClient = sum.NewSumService("go.micro.srv.sum_service", service.Options().Service.Client() )
	service.HandleFunc("/portal_web/sum", Sum)
	service.Run()
}

//接口
func Sum(w http.ResponseWriter, r *http.Request) {
	inputString := r.URL.Query().Get("input")
	input, _ := strconv.ParseInt(inputString, 10, 10)
	
	req := &sum.SumRequest{
		Input: int32(input),
	}


	rsp, err := srvClient.GetSum(context.Background(), req)
	if err != nil {
		//ignore
	}
	w.Write([]byte(strconv.Itoa(int(rsp.Output))))

}