package cloud_tim

import (
	"fmt"
	"testing"
)

func Test_buildSig(t *testing.T) {
	strMobile := "13788888888"                      //tel 的 mobile 字段的内容
	strAppKey := "5f03a35d00ee52a21327ab048186a2c4" //sdkappid 对应的 appkey，需要业务方高度保密
	strRand := "7226249334"                         //URL 中的 random 字段的值
	strTime := "1457336869"                         //UNIX 时间戳
	str := buildSig(strAppKey, strMobile, strRand, strTime)
	fmt.Println(fmt.Sprintf("%v", str))
}

func Test_random(t *testing.T) {
	randInt := random()

	fmt.Println(randInt)
}
