package http

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// FormatParam 格式化参数
func FormatParam(_ context.Context, request interface{}) error {
	// //绑定
	// if err := e.ShouldBind(request); err != nil {
	// 	return err
	// }

	// 参数验证
	if err := ValidatorHandler.New().Struct(request); err != nil {
		buff := bytes.NewBufferString("")
		for _, err := range err.(validator.ValidationErrors) {
			// 翻译错误原因
			buff.WriteString(err.Translate(ValidatorHandler.Trans))
			buff.WriteString(",")
		}
		return errors.New(strings.Trim(buff.String(), ","))
	}

	return nil
}

var ValidatorHandler = validation{}

type validation struct {
	Trans ut.Translator
}

func (v *validation) New() *validator.Validate {

	// 验证
	zh_ch := zh.New()
	validate := validator.New()
	uni := ut.New(zh_ch)
	trans, _ := uni.GetTranslator("zh")
	v.Trans = trans
	// 验证器注册翻译器
	_ = zh_translations.RegisterDefaultTranslations(validate, trans)

	err := validate.RegisterValidation("timestamp", v.validateTimestamp) // 验证时间戳
	if err != nil {
		fmt.Println(err)
	}
	err = validate.RegisterValidation("datetime", v.validateDatetime) // 验证时间 YYYY-mm-dd HH:ii:ss
	if err != nil {
		fmt.Println(err)
	}
	err = validate.RegisterValidation("date", v.validateDate) // 验证日期 YYYY-mm-dd
	if err != nil {
		fmt.Println(err)
	}
	err = validate.RegisterValidation("required_if", v.validateRequiredIf) // required_if=field 1  当字段field=1时当前字段不能为空
	if err != nil {
		fmt.Println(err)
	}
	err = validate.RegisterValidation("phone", v.validatePhone) // 验证手机号
	if err != nil {
		fmt.Println(err)
	}
	RegisterTagTranslation("phone", "请输入正确的手机号!", trans, validate)

	err = validate.RegisterValidation("is_not_blank", v.validateIsNotBlank) // 验证空白
	if err != nil {
		fmt.Println(err)
	}
	err = validate.RegisterValidation("day", v.validateDay) // 验证YYYY-mm-dd
	if err != nil {
		fmt.Println(err)
	}

	return validate
}
func GetValidate() *RequestValidator {

	validate := ValidatorHandler.New()

	return &RequestValidator{Validator: validate}

}

// 自定义中文翻译
func RegisterTagTranslation(tag string, errMessage string, trans ut.Translator, validate *validator.Validate) {
	_ = validate.RegisterTranslation(tag, trans, func(ut ut.Translator) error {
		return ut.Add(tag, errMessage, false)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, err := ut.T(fe.Tag(), fe.Field())
		if err != nil {
			return fe.(error).Error()
		}
		return t
	})
}

type RequestValidator struct {
	Validator *validator.Validate
}

func (c *RequestValidator) Validate(i interface{}) error {
	return c.Validator.Struct(i)
}

// 手机号验证
func (v validation) validatePhone(fl validator.FieldLevel) bool {

	match, _ := regexp.MatchString("^1\\d{10}$", fl.Field().String())
	return match
}

// 时间戳验证
func (v validation) validateTimestamp(fl validator.FieldLevel) bool {

	var str string
	if fl.Field().Type().String() == "int" {
		i := fl.Field().Int()
		str = strconv.FormatInt(i, 10)
	} else {
		str = fl.Field().String()
	}
	match, _ := regexp.MatchString("^\\d{10}$", str)
	return match
}

// 时间验证 Y-m-d H:i:s
func (v validation) validateDatetime(fl validator.FieldLevel) bool {
	match, _ := regexp.MatchString("^(\\d{4})-(\\d{2})-(\\d{2}) (\\d{2}):(\\d{2}):(\\d{2})$", fl.Field().String())
	return match
}

// 日期验证 Y-m-d
func (v validation) validateDate(fl validator.FieldLevel) bool {
	match, _ := regexp.MatchString("^[0-9]\\d{3}-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1])$", fl.Field().String())
	return match
}

// 根据其他字段的值来判断是否必填
func (va validation) validateRequiredIf(fl validator.FieldLevel) bool {

	param := fl.Param()
	if param == "" {
		return false
	}
	// 获取待比较字段名
	params := strings.Split(param, " ")
	compareFieldName := params[0]
	// 获取待比较字段的值
	vals := params[1:]
	compareField := fl.Parent().Elem().FieldByName(compareFieldName)
	// 将待比较字段转化为string
	var v string
	field := compareField
	switch field.Kind() {
	case reflect.String:
		v = field.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v = strconv.FormatInt(field.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v = strconv.FormatUint(field.Uint(), 10)
	default:
		return false
	}
	for i := 0; i < len(vals); i++ {
		if vals[i] == v {
			// 验证字段必填
			return requireCheckFieldKind(fl, "")
		}
	}
	return true
}

// requireCheckField is a func for check field kind
func requireCheckFieldKind(fl validator.FieldLevel, param string) bool {
	field := fl.Field()
	if len(param) > 0 {
		if fl.Parent().Kind() == reflect.Ptr {
			field = fl.Parent().Elem().FieldByName(param)
		} else {
			field = fl.Parent().FieldByName(param)
		}
	}
	switch field.Kind() {
	case reflect.Slice, reflect.Map, reflect.Ptr, reflect.Interface, reflect.Chan, reflect.Func:
		return !field.IsNil()
	default:
		return field.IsValid() && field.Interface() != reflect.Zero(field.Type()).Interface()
	}
}

// 空白验证
func (v validation) validateIsNotBlank(fl validator.FieldLevel) bool {

	if val := strings.TrimSpace(fl.Field().String()); val == "" {
		return false
	}
	return true
}

// 日期验证 YYYY-mm-dd
func (v validation) validateDay(fl validator.FieldLevel) bool {

	val := strings.TrimSpace(fl.Field().String())

	match, _ := regexp.MatchString("((^((1[8-9]\\d{2})|([2-9]\\d{3}))([-])(10|12|0?[13578])([-])(3[01]|[12][0-9]|0?[1-9])$)|(^((1[8-9]\\d{2})|([2-9]\\d{3}))([-])(11|0?[469])([-])(30|[12][0-9]|0?[1-9])$)|(^((1[8-9]\\d{2})|([2-9]\\d{3}))([-])(0?2)([-])(2[0-8]|1[0-9]|0?[1-9])$)|(^([2468][048]00)([-])(0?2)([-])(29)$)|(^([3579][26]00)([-])(0?2)([-])(29)$)|(^([1][89][0][48])([-])(0?2)([-])(29)$)|(^([2-9][0-9][0][48])([-])(0?2)([-])(29)$)|(^([1][89][2468][048])([-])(0?2)([-])(29)$)|(^([2-9][0-9][2468][048])([-])(0?2)([-])(29)$)|(^([1][89][13579][26])([-])(0?2)([-])(29)$)|(^([2-9][0-9][13579][26])([-])(0?2)([-])(29)$))", val)

	return match
}
