package restdto

type Result struct {
	IsSuccess bool
	Msg       string
	Data      interface{}
}

func Success() *Result {
	return &Result{
		IsSuccess: true,
		Msg:       "Success",
	}
}
func Fail500() *Result {
	return &Result{
		IsSuccess: false,
		Msg:       "Server Error",
	}
}

func Fail403() *Result {
	return &Result{
		IsSuccess: false,
		Msg:       "Data Error",
	}
}
