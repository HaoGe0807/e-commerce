package errors

type Err struct {
	Code int32
	Msg  string
	Data []byte
}

func (e Err) Error() string {
	//TODO implement me
	panic("implement me")
}

func ErrorEnum(code int32, msg string) error {
	err := new(Err)

	if msg == "" {
		if errorMessages[ErrorCode(code)] != "" {
			msg = errorMessages[ErrorCode(code)]
		} else {
			code = DOMAIN_ERROR
			msg = errorMessages[ErrorCode(code)]
		}
	}

	err.Code = code
	err.Msg = msg
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

	// product

)

var errorMessages = map[ErrorCode]string{
	PARAMS_ERROR: "参数异常",
	BIZ_ERROR:    "业务服务异常",
	DOMAIN_ERROR: "领域服务异常",
}
