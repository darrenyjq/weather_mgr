package key

import "fmt"

//用户账户
func RedisKvAccount(uid int64) string {
	return fmt.Sprintf("account_%d",uid)
}