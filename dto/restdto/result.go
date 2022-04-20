package restdto

type Result struct {
	IsSuccess bool
	Msg       string
}

func Success()*Result  {
	return &Result{
		IsSuccess: true,
		Msg: "Success",
	}
}
