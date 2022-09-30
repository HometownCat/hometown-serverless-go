package common

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/aws/aws-lambda-go/events"
	"github.com/imdario/mergo"
	"hometown.com/hometown-serverless-go/types"
)

func IsError(err error) bool {
	return err != nil
}

func ResponseError(err error) (*events.APIGatewayProxyResponse, error) {
	return &events.APIGatewayProxyResponse{
		Body:       "{\"error\":\"" + err.Error() + "\"}",
		StatusCode: 400,
	}, nil
}

func ReturnNotNil(arg1 interface{}, arg2 interface{}) *interface{} {
	if arg1 != nil {
		return &arg1
	} else if arg2 != nil {
		return &arg2
	} else {
		return nil
	}
}

// func GetMapFromString(parseStr *string) (){

// }

func isValidationKey(param *map[string]interface{}, key *string, keyType *string) bool {
	var body map[string]interface{} = *param
	validValue := body[*key]

	return reflect.TypeOf(validValue).String() == *keyType
}

func RequestValid(event *events.APIGatewayProxyRequest, keyList *[]types.ValidKey) (*map[string]interface{}, error) {
	result := true
	notMatchedParam := map[string]interface{}{}
	param := map[string]interface{}{}

	param["userIp"] = event.RequestContext.Identity.SourceIP

	json.Unmarshal([]byte(event.Body), &param)

	mergo.Merge(&param, event.Headers)
	mergo.Merge(&param, event.QueryStringParameters)
	mergo.Merge(&param, event.PathParameters)

	for _, validKey := range *keyList {
		if !isValidationKey(&param, &validKey.Key, &validKey.KeyType) {
			notMatchedParam[validKey.Key] = param[validKey.Key]
			result = false
			break
		}
	}
	if !result {
		notMatchedBin, _ := json.Marshal(notMatchedParam)
		return nil, errors.New("not matched : " + string(notMatchedBin))
	}
	return &param, nil
}

func IsExistKey(param map[string]string, key string) bool {

	if _, exist := param[key]; exist {
		return true
	}
	return false
}
