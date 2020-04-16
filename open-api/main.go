package main



import (
	"context"
	api "github.com/micro/go-micro/api/proto"
	"github.com/shizhigu/GoTutorial/sum"
	"strconv"
	"encoding/json"
	"github.com/micro/go-micro"

)

var (
	srvClient sum.SumService
)

type Open struct {

}

func (open Open) Fetch(ctx context.Context, req *api.Request, rsp *api.Response) error {
	sumInputStr := req.Get["sum"].Values[0]

	sumInput, _ := strconv.ParseInt(sumInputStr, 10, 10)
	// sumInput = int32(sumInput)
	

	sumReq := &sum.SumRequest{
		Input: int32(sumInput),
	}


	sumRsp, _ := srvClient.GetSum(ctx, sumReq)

	ret, _ := json.Marshal(map[string]interface{} {
		"sum":	sumRsp,
	})

	rsp.Body = string(ret)
	return nil
}



func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.open"),
	)

	srvClient = sum.NewSumService("github.com/shizhigu/GoTutorial/sum", service.Client())
	
	if err:=service.Run(); err!=nil {
		panic(err)
	}
}