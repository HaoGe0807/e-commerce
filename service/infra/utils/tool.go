package utils

import (
	"fmt"
	"math"
	"strings"
)

func StructToMap(obj interface{}) (newMap map[string]interface{}, err error) {
	data, err := json.Marshal(obj)

	if err != nil {
		return
	}

	err = json.Unmarshal(data, &newMap)
	return
}

func EncodeVerCode(version string) int32 {
	list := strings.Split(version, ".")
	var verCode int64
	for i := 0; i < len(list); i++ {
		verCode += String2Int64(list[i]) * int64(math.Pow(10, float64(2*(int64(len(list)-i-1)))))
	}
	return int32(verCode)
}

func GetBizNo(BizId string, BizType string) string {
	const MD5LENGTH = 18
	nowTime := GetNowTime().UnixNano()
	randomStr := GetRandomString(1)
	md5 := GetMD5(fmt.Sprintf("%s%d%s", BizId, nowTime, randomStr))

	no := BizType + md5[0:MD5LENGTH]

	return strings.ToUpper(no)
}
