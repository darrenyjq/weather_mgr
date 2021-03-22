package abtest

import (
	"context"
)

type ABTest struct {
	Name  string `json:"name"`
	Val   int64 `json:"val"`
}

const (
	//大转盘实验
	_ABTEST_PARAM_LUCKY_DRAW	=	"DIV_REWARD_NEW_STYLE"

	//对照组值
	ABTEST_VALUE_DEFAULT	int64	=	1
)

func ABTestLuckyDrawValue(ctx context.Context) (int64) {
	return abTestGetValue(ctx, _ABTEST_PARAM_LUCKY_DRAW)
}

func abTestGetValue(ctx context.Context, abtestParam string) (abtestVal int64) {
	abtestVal = ABTEST_VALUE_DEFAULT
	data := getTests(ctx)
	val,ok := data[abtestParam]
	if !ok {
		return
	}
	abtestVal = val
	return
}

func getTests(ctx context.Context) (data map[string]int64) {
	data = map[string]int64{}

	//新版，新老两个不会同时存在的
	testVal := ctx.Value("tests")
	if testVal == nil {
		return
	}
	tests,ok := testVal.([]ABTest)
	if !ok {
		return
	}
	for _,v := range tests {
		data[v.Name] = v.Val
	}
	return
}