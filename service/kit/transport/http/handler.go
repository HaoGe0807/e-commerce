package http

import (
	"context"
	errors "e-commerce/service/infra/errors"
	"e-commerce/service/kit/endpoint"
	"encoding/json"
	"net/http"
	"reflect"
)

func encodeResp(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	newResponse := &Response{
		Code: errors.ERR_SUCCESS,
		Msg:  "",
		Data: response,
	}
	return json.NewEncoder(w).Encode(newResponse)

}

func ReplyWithContent(resp endpoint.Response, err error) []byte {
	result := map[string]interface{}{
		"code": 1,
		"msg":  "",
		"data": struct{}{},
	}
	if err == nil {
		if resp != nil {
			data, err := resp.ToJson()
			if err != nil {
				result["code"] = errors.ERR_JSON_ERROR
			} else {
				result["data"] = data
			}
		}
	} else {
		result["code"] = errors.ERR_ERROR
		result["msg"] = err.Error()
		_, hasMethod := reflect.TypeOf(err).MethodByName("Code")
		if hasMethod {
			err2 := err.(*errors.Err)
			result["code"] = err2.GetCode()
		}
	}
	ret, _ := json.Marshal(result)
	return ret
}

func ReplyWithoutContent(err error) []byte {
	result := map[string]interface{}{
		"code": 1,
		"msg":  "",
		"data": struct{}{},
	}
	if err != nil {
		result["code"] = errors.ERR_ERROR
		result["msg"] = err.Error()
		_, hasMethod := reflect.TypeOf(err).MethodByName("Code")
		if hasMethod {
			err2 := err.(*errors.Err)
			result["code"] = err2.GetCode()
		}
	}
	ret, _ := json.Marshal(result)
	return ret
}
