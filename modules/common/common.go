package common

import (
	"encoding/json"
	"log"
	"reflect"

	"github.com/aws/aws-lambda-go/events"
	"hometown.com/hometown-serverless-go/types"
)

func IsError(err error) bool {
	return err != nil
}

func ResponseError(err error) (*events.APIGatewayProxyResponse, error){
	return &events.APIGatewayProxyResponse{
		Body: "{\"error\":\""+err.Error()+"\"}",
		StatusCode: 400,
	},nil
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

func IsValidationKey [T map[string]interface{} | map[string]string](param *T, key *string, keyType *string) bool {
	body := map[string]interface{}{}
	bodyBin,_ := json.Marshal(*param)

	json.Unmarshal(bodyBin,&body)
	
	validValue := body[*key]
	
	log.Println(reflect.TypeOf(validValue).String())
	
	return reflect.TypeOf(validValue).String() != *keyType
}



func RequestValid(event *events.APIGatewayProxyRequest, keyList *[]types.ValidKey) bool {
	result := true
	for _,validKey := range *keyList {
		body := map[string]interface{}{}
		json.Unmarshal([]byte(event.Body),&body)

		if !IsValidationKey(&body,&validKey.Key,&validKey.KeyType) {
			result = false
			break
		}
	}
	return result
}