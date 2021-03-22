package xaes

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAes(t *testing.T) {
	aes := NewAes("dea5a27becaa2e054b959b3b7c8c95034949e4ab")

	val,err := aes.Encrypt([]byte(`{"goods_id":1,"bank":"alipay","phone":"15611111111","name":"曾繁然","id_card":"330401199208018142","bank_account":"1234567@qq.com","device":{"is_root":true,"is_emulator":false,"is_xposed_exist":false}}`))
	assert.NoError(t,err)
	fmt.Println(val)

	dstr :="8jU81Yj4eXY4l1Y9RYz7MRE+LS+fDWesz2nYtOhkyv/DZoDjvY+Xm8rvjhzXjw/+dufWGdIIYKDHEUPYbLPyrKTVLNO2XZpQT94VgcbPJI0N8vNH6LO3dcpwb04LT0QpwzZ4JwJDvYgg5KYMIaVqlWv0GYxgROz5KiQipSbakSkb/m4VS+0OwmfdpIIUKruTbkB3OtPYpwQlxenGVKsEieUx57fZZzEFCAI6Gh/liJ+20gDpdvVKTSan/p2+zMR3EeuEehJHiPNGp3tji74A9SxE223ZEUlRl81Qgkz6hZg="
	dval,err := aes.Decrypt(dstr)
	assert.NoError(t,err)
	fmt.Println(string(dval))
}
