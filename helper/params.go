package helper

import "base/cootek/pgd/account"

type (
	ParamUserInfo struct {
		Account	*account.Account `json:"account"`
	}
	
	ParamEncryptData struct {
		Data string `json:"data"`
		SignType string `json:"sign_type"`
	}
)
