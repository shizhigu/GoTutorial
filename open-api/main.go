package main



import (
	"context"
	api "github.com/micro/go-micro/api/proto"
	"github.com/shizhigu/GoTutorial/sum"
	"strconv"
	"encoding/json"
	"github.com/micro/go-micro"
	proto "github.com/shizhigu/GoTutorial/open-api/proto"
	"fmt"

)

var (
	srvClient sum.SumService
)

type Open struct {

}

func (open *Open) Fetch(ctx context.Context, req *api.Request, rsp *api.Response) error {
	fmt.Println(req)
	fmt.Println("hello")
	sumInputStr := req.Get["sum"].Values[0]
	fmt.Println(sumInputStr)
	
	sumInput, _ := strconv.ParseInt(sumInputStr, 10, 10)
	// sumInput = int32(sumInput)
	

	sumReq := &sum.SumRequest{
		Input: int32(sumInput),
	}
	

	sumRsp, err := srvClient.GetSum(context.Background(), sumReq)
	if err != nil {
		panic(err)
	}
	
	fmt.Println(sumRsp.Output)
	fmt.Println("world")
	ret, _ := json.Marshal(map[string]interface{} {
		"sum":	sumRsp.Output,
	})
	
	rsp.Body = string(ret)
	return nil
}



func main() {
	service := micro.NewService(
		micro.Name("go.micro.haha.api.open"),
	)
	service.Init()

	srvClient = sum.NewSumService("github.com/shizhigu/GoTutorial/sum", service.Client())

	proto.RegisterOpenHandler(service.Server(), new(Open))
	// service.HandleFunc("/sum", Sum)
	fmt.Println("HIHIHIHI")
	_ = service.Run()
	// if err:=service.Run(); err!=nil {
	// 	panic(err)
	// }
}