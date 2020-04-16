package handler

import (
	"context"
	"github.com/shizhigu/GoTutorial/sum"
	"github.com/shizhigu/GoTutorial/service"
)

type handler struct {

}

func (h handler) GetSum(ctx context.Context, req *sum.SumRequest, res *sum.SumResponse) error {
	inputs := make([]int32, 0)

	var i int32 = 0;
	for ; i<=req.Input; i++ {
		inputs = append(inputs, i)
	}

	res.Output = service.GetSum(inputs...)

	return nil
}

func Handler() sum.SumHandler {
	return handler{}
}