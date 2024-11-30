package utils

func GetCommomInterval(iv []int64, iv2 []int64) []int64 {
	var (
		commonIvL, commonIvR int64
		commonIv             []int64
	)
	ivL, ivR := iv[0], iv[1]
	iv2L, iv2R := iv2[0], iv2[1]
	if ivL <= iv2L {
		if iv2L >= ivR {
			return commonIv
		}
		commonIvL = Min(iv2L, ivR)
		commonIvR = Min(ivR, iv2R)
	} else {
		if ivL >= iv2R {
			return commonIv
		}
		commonIvL = Min(ivL, iv2R)
		commonIvR = Min(iv2R, ivR)
	}

	commonIv = append(commonIv, commonIvL, commonIvR)
	return commonIv
}

func GetCommomIntervalList(iv []int64, list [][]int64) [][]int64 {
	result := make([][]int64, 0)

	for _, iv2 := range list {
		commonIv := GetCommomInterval(iv, iv2)
		if len(commonIv) == 2 {
			result = append(result, commonIv)
		}
	}
	return result
}

func GetCommomIntervalList2(list1 [][]int64, list2 [][]int64) [][]int64 {
	result := make([][]int64, 0)

	for _, iv := range list1 {
		result = append(result, GetCommomIntervalList(iv, list2)...)
	}
	return result
}
