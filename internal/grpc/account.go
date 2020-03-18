package grpc

import (
	"base/cootek/pgd/base/account"
	"context"
)

type AccountService struct {
	
}

func (a *AccountService) Info(ctx context.Context, param *account.AccountInfoParam) (result *account.AccountInfoResult,err error) {
	result = &account.AccountInfoResult{
		ErrorCode: 0,
		Msg: "success",
		Data: &account.AccountInfoResult_Data{
			Amount: 10,
			TicketNum: 20,
		},
	}
	return
}

func (a *AccountService) IncrAmount(ctx context.Context, param *account.IncrAmountParam) (result *account.IncrAmountResult,err error)  {
	result = &account.IncrAmountResult{
		ErrorCode: 0,
		Msg: "success",
		Data: &account.IncrAmountResult_Data{
			Amount: 10,
		},
	}
	return
}