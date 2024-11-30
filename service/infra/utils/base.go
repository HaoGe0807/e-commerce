package utils

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"crypto/sha256"
	"e-commerce/service/infra/errors"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const (
	SMALLER = -1
	EQUAL   = 0
	LARGER  = 1
)

const (
	STATUS_CONNECTED    = 1
	STATUS_DISCONNECTED = 0
)

const SUB_VERSION_MAX_BITS = 5
const SUB_VERSION_NUM = 3
const INIT_VERSION = "0.0.0"

// 获取当前系统环境
func GetRunTime() string {
	//获取系统环境变量
	RUN_TIME := os.Getenv("RUN_TIME")
	if RUN_TIME == "" {
		return "local"
	}
	return RUN_TIME
}

func GetNamespace() string {
	NAMESPACE := os.Getenv("NAMESPACE")
	if NAMESPACE == "" {
		return "native"
	}
	return NAMESPACE
}

func ConvertPageInfo(pageNum, pageSize int32) (int32, int32, error) {
	if !IsValidPage(pageNum, pageSize) {
		return 0, 0, errors.ErrorEnum(errors.ERR_INVALID_PARAM, "page is invalid")
	}

	offset := (pageNum - 1) * pageSize
	limit := pageSize
	return offset, limit, nil
}

func ConvertBinVersion(srcVersion string) (string, error) {
	var dstVersion string

	arrSrcVersion := strings.Split(srcVersion, ".")
	stringLen := len(arrSrcVersion)
	if SUB_VERSION_NUM != stringLen {
		return dstVersion, errors.ErrorEnum(errors.ERR_INVALID_FORMAT, "format is wrong")
	}

	for index := 0; index < stringLen; index++ {
		srcString := arrSrcVersion[index]
		strLen := len(srcString)
		var tmpString string
		for num := 0; num < SUB_VERSION_MAX_BITS-strLen; num++ {
			tmpString += "0"
		}
		dstVersion += (tmpString + srcString)
	}
	return dstVersion, nil
}

func CompareBinVersion(srcVersion string, dstVersion string) int {
	arrSrcVersion := strings.Split(srcVersion, ".")
	arrDstVersion := strings.Split(dstVersion, ".")

	if SUB_VERSION_NUM != len(arrSrcVersion) {
		arrSrcVersion = strings.Split(INIT_VERSION, ".")
	}

	if SUB_VERSION_NUM != len(arrDstVersion) {
		arrDstVersion = strings.Split(INIT_VERSION, ".")
	}
	for index := 0; index < len(arrSrcVersion); index++ {
		srcInt, _ := strconv.Atoi(arrSrcVersion[index])
		dstInt, _ := strconv.Atoi(arrDstVersion[index])

		if srcInt < dstInt {
			return LARGER
		} else if srcInt > dstInt {
			return SMALLER
		}
	}
	return EQUAL
}

var node, _ = NewNode(1)

func GenerateMessageId() string {
	id := node.Generate()
	return strconv.Itoa(int(id))
}

func Generate64BitId() int64 {
	return int64(node.Generate())
}

var snowBallNode, _ = NewNode(1)

func Generate32BitId() int32 {
	return int32(snowBallNode.Generate())
}

func CheckUnion(leftFields, rightFields []string) bool {
	isUnion := false
	for _, left := range leftFields {
		for _, right := range rightFields {
			if left == right {
				isUnion = true
				break
			}
		}
		if isUnion {
			break
		}
	}
	return isUnion
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func HasChinese(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}
	return false
}

const (
	SEP_COMMA = ","
)

func Int2String(value int64) string {
	return strconv.FormatInt(value, 10)
}

func Float2String(value float64, decimal int) string {
	return strconv.FormatFloat(value, 'f', decimal, 64)
}

func String2Int64(source string) int64 {
	value, err := strconv.ParseInt(source, 10, 64)
	if err != nil {
		value = 0
	}
	return value
}

func String2Int32(source string) int32 {
	value, err := strconv.ParseInt(source, 10, 32)
	if err != nil {
		value = 0
	}
	return int32(value)
}

func String2Float64(source string) float64 {
	value, err := strconv.ParseFloat(source, 64)
	if err != nil {
		value = 0.0
	}
	return value
}

func String2SliceInt32(str string, sep string) []int32 {
	if len(str) == 0 {
		return []int32{}
	}

	items := strings.Split(str, sep)
	ret := make([]int32, 0, len(items))
	for _, itemStr := range items {
		item, err := strconv.ParseInt(itemStr, 10, 32)
		if err != nil {
			fmt.Printf("[ERROR] String2SliceInt32 - failed to parse string to int for %v, err: %v \n", itemStr, err.Error())
			continue
		}

		ret = append(ret, int32(item))
	}

	return ret
}

func Int2Bytes(source int64) []byte {
	return []byte(Int2String(source))
}

func Float2Bytes(source float64, decimal int) []byte {
	return []byte(Float2String(source, decimal))
}

func GetMD5(plainText string) string {
	h := md5.New()
	h.Write([]byte(plainText))
	return hex.EncodeToString(h.Sum(nil))
}

func GetSha256(planText string) string {
	sha := sha256.New()
	sha.Write([]byte(planText))
	return hex.EncodeToString(sha.Sum(nil))
}

func DesEncryption(key, iv, plainText []byte) (string, error) {
	block, err := des.NewCipher(key)

	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	origData := PKCS5Padding(plainText, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	cryted := make([]byte, len(origData))
	blockMode.CryptBlocks(cryted, origData)
	return base64.StdEncoding.EncodeToString(cryted), nil
}

func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}
