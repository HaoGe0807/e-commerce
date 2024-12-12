package utils

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func Max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func PagingList(pageNum, pageSize, totalCount int32) (int32, int32, error) {
	var startIndex int32
	var endIndex int32
	startIndex = pageSize * (pageNum - 1)
	if totalCount >= pageSize*pageNum {
		endIndex = pageSize * pageNum
	} else {
		if startIndex > totalCount {
			return 0, 0, errors.New("start index exceed total count")
		}
		endIndex = totalCount
	}
	return startIndex, endIndex, nil
}

func GetJsonString(v interface{}) string {
	str, err := json.MarshalToString(v)
	if err != nil {
		return ""
	}
	return str
}

// string list 去重
func UniqueForStringList(list []string) []string {
	result := make([]string, 0, len(list))

	m := make(map[string]struct{})
	for _, item := range list {
		m[item] = struct{}{}
	}
	for k := range m {
		if IsValidField(k) {
			result = append(result, k)
		}
	}
	return result
}

func ContainsString(slice []string, target string) bool {
	sort.Strings(slice)
	index := sort.SearchStrings(slice, target)
	return index < len(slice) && slice[index] == target
}

func ModelIdNext(modelName string) string {
	// 截取modelName的第一个字符
	firstChar := modelName[0:1]
	// 将第一个字符转换为大写
	firstChar = strings.ToUpper(firstChar)
	// 获取当前时间戳
	currentTime := time.Now().Unix()
	// 将当前时间戳转为20241212这种结构
	currentTimeStr := time.Unix(currentTime, 0).Format("20060102")
	// 当前时间戳后跟五位随机数
	randomNum := rand.Intn(90000) + 10000
	// 拼接字符串
	//id := firstChar + strconv.FormatInt(currentTime, 10) + strconv.Itoa(randomNum)
	id := firstChar + currentTimeStr + strconv.Itoa(randomNum)
	return id
}
