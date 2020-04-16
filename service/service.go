package service

func GetSum(inputs ...int32) (res int32) {
	for _, v:=range inputs {
		res += v
	}
	return res
}
