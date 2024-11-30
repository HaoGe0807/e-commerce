package utils

import (
	"context"
	"e-commerce/service/infra/errors"
	"e-commerce/service/infra/log"
	"github.com/fatih/structs"
	goEndpoint "github.com/go-kit/kit/endpoint"
	"reflect"
	"strconv"
)

const (
	CONTENT_TYPE_ID          = "id"
	CONTENT_TYPE_SN          = "sn"
	CONTENT_TYPE_PRICE       = "price"
	CONTENT_TYPE_TEXT        = "text"
	CONTENT_TYPE_BIN_VERSION = "bin_version"
)

type Middleware func(goEndpoint.Endpoint) goEndpoint.Endpoint

func LoggingMiddleware() Middleware {
	return func(next goEndpoint.Endpoint) goEndpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			ts := GetTimestamp()
			reqName := reflect.TypeOf(request).Name() + "_" + strconv.FormatInt(ts, 10)

			js, err := json.Marshal(request)
			if err != nil {
				log.Info("Message in (", reqName, ") ==> ", request)
			} else {
				log.Info("Message in (", reqName, ") ==> ", string(js))
			}
			defer log.Info("message out (", reqName, ") <==")
			return next(ctx, request)
		}
	}
}

func ParameterCheckMiddleware() Middleware {
	return func(next goEndpoint.Endpoint) goEndpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			types := reflect.TypeOf(request)
			values := reflect.ValueOf(request)
			for i := 0; i < types.NumField(); i++ {
				field := types.Field(i)
				value := values.FieldByName(field.Name)
				content, hasContent := field.Tag.Lookup("check_null")
				if hasContent {
					switch content {
					case CONTENT_TYPE_ID:
						if !IsValidId(value.Int()) {
							return nil, errors.ErrorEnum(errors.ERR_INVALID_ID,
								"id is invalid "+field.Name)
						}
					case CONTENT_TYPE_SN:
						if !IsValidField(value.String()) {
							return nil, errors.ErrorEnum(errors.ERR_INVALID_SN,
								"sn is invalid "+field.Name)
						}
					case CONTENT_TYPE_TEXT:
						if !IsValidField(value.String()) {
							return nil, errors.ErrorEnum(errors.ERR_INVALID_PARAM,
								"text is invalid "+field.Name)
						}
					case CONTENT_TYPE_PRICE:
						if !IsValidPrice(value.Float()) {
							return nil, errors.ErrorEnum(errors.ERR_INVALID_PRODUCT_PRICE,
								"price is invalid "+field.Name)
						}
					default:
						log.Warn("unsupported content ", content)
					}
				}
			}
			return next(ctx, request)
		}
	}
}

func HttpRequestParameterCheck(request interface{}) (map[string]interface{}, error) {
	if !structs.IsStruct(request) {
		return nil, errors.ErrorEnum(errors.ERR_INVALID_PARAM,
			"request is invalid")
	}
	mapRequest := structs.Map(request)
	log.Info("original http request: ", mapRequest)
	types := reflect.TypeOf(request)
	values := reflect.ValueOf(request)
	for i := 0; i < types.NumField(); i++ {
		field := types.Field(i)
		value := values.FieldByName(field.Name)

		checkNullContent, hasCheckNullContent := field.Tag.Lookup("check_null")
		if hasCheckNullContent {
			switch checkNullContent {
			case CONTENT_TYPE_ID:
				if !IsValidId(value.Int()) {
					return nil, errors.ErrorEnum(errors.ERR_INVALID_ID,
						"id is invalid "+field.Name)
				}
			case CONTENT_TYPE_TEXT:
				if !IsValidField(value.String()) {
					return nil, errors.ErrorEnum(errors.ERR_INVALID_PARAM,
						"text is invalid "+field.Name)
				}
			default:
				log.Warn("unsupported content ", checkNullContent)
			}
		}

		optionalContent, hasOptionalContent := field.Tag.Lookup("optional")
		if hasOptionalContent {
			switch optionalContent {
			case CONTENT_TYPE_ID:
				if !IsValidId(value.Int()) {
					jsonField := field.Tag.Get("json")
					delete(mapRequest, jsonField)
				}
			default:
				log.Warn("unsupported content ", optionalContent)
			}
		}
	}
	return mapRequest, nil
}
