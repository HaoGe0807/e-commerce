package errors

type Err struct {
	Code int32
	Msg  string
	Data []byte
}

func (e Err) Error() string {
	//TODO implement me
	return e.Msg
}

func ErrorEnum(code int32, msg ...string) error {
	err := new(Err)
	var message string
	if len(msg) > 0 {
		message = msg[0]
	}

	if message == "" {
		if errorMessages[ErrorCode(code)] != "" {
			message = errorMessages[ErrorCode(code)]
		} else {
			code = DOMAIN_ERROR
			message = errorMessages[ErrorCode(code)]
		}
	}

	err.Code = code
	err.Msg = message
	return err
}

func Error(err *Err, msg string) error {
	if err.Code == 0 {
		err.Code = BIZ_ERROR
	}

	if msg != "" {
		err.Msg = msg
	}

	return err
}

func (err Err) GetCode() int32 {
	return err.Code
}
func (err Err) GetMsg() string  { return err.Msg }
func (err Err) GetData() []byte { return err.Data }

/*
The `errCode` is structured in three segments: First Segment + Second Segment + Third Segment
First Segment: Represents the business context (two digits)
Second Segment: Represents the domain (two digits)
Third Segment: Represents the domain service logic error code (two digits)

Value Ranges for each segment:
First Segment: [10, 99]
Second Segment: [01, 99]
Third Segment: [01, 99]
*/

type ErrorCode int32

const (
	// Common
	PARAMS_ERROR = 100000
	BIZ_ERROR    = 100001
	DOMAIN_ERROR = 100002

	// product  1001号段
	SKU_NAME_DUPLICATE           = 100100
	PRODUCT_ONLY_ONE_DEFAULT_SKU = 100101
	SKU_CODE_ERROR               = 100102

	// product category
	CATEGORY_NOT_EXIST     = 100201
	CATEGORY_NAME_EXIST    = 100202
	CATEGORY_EXIST_PRODUCT = 100203
)

var errorMessages = map[ErrorCode]string{
	PARAMS_ERROR: "参数异常",
	BIZ_ERROR:    "业务服务异常",
	DOMAIN_ERROR: "领域服务异常",

	SKU_NAME_DUPLICATE:           "规格名称重复",
	PRODUCT_ONLY_ONE_DEFAULT_SKU: "商品只能有一个默认规格",
	SKU_CODE_ERROR:               "规格条码不符合要求",

	CATEGORY_NOT_EXIST:     "分类不存在",
	CATEGORY_NAME_EXIST:    "分类名称已存在",
	CATEGORY_EXIST_PRODUCT: "分类存在关联商品，禁止删除",
}
