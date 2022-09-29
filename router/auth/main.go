package main

import (
	"encoding/json"
	"os"
	"runtime"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"hometown.com/hometown-serverless-go/modules/database"
	"hometown.com/hometown-serverless-go/modules/validation"
)

var Effect string;

func Handler(event events.APIGatewayCustomAuthorizerRequestTypeRequest) (events.APIGatewayCustomAuthorizerResponse , error){
	headers := event.Headers;
	// queryStringParameters := event.QueryStringParameters;
	// stageVariables := event.StageVariables;
	// requestContext := event.RequestContext;
	// tmp := strings.Split(event.MethodArn, ":")
	// apiGatewayArnTmp := strings.Split(tmp[5], "/")
	// awsAccountId := tmp[4]
	// region := tmp[3]
	// ApiId := apiGatewayArnTmp[0];
	// stage := apiGatewayArnTmp[1];
	// route := apiGatewayArnTmp[2];
	principalId := os.Getenv("PRINCIPAL_ID")
	if headers["x-api-key"] != os.Getenv("API_KEY") {
		return *GenerateDeny(&principalId,&event.MethodArn) , nil
	}


	accessToken := headers["accesstoken"];

	userInfo, validErr := validation.UserValidation(&accessToken)
	

	if validErr != nil {
		return *GenerateDeny(&principalId,&event.MethodArn) , nil
	} else if userInfo != nil {
		userBin,_ := json.Marshal(*userInfo)
		// 테스트 진행 필요
		event.StageVariables["userData"] = string(userBin)
	}

	return *GenerateAllow(&principalId,&event.MethodArn), nil
}

func GeneratePolicy(principalId *string, resource *string) *events.APIGatewayCustomAuthorizerResponse {
	var AuthResponse events.APIGatewayCustomAuthorizerResponse;
	AuthResponse.PrincipalID = *principalId
	var PolicyDocument events.APIGatewayCustomAuthorizerPolicy

	PolicyDocument.Version = "2012-10-17"
	PolicyDocument.Statement = make([]events.IAMPolicyStatement,1)

	var statementOne events.IAMPolicyStatement
	statementOne.Action = make([]string, 1)
	statementOne.Effect = Effect
	statementOne.Resource = make([]string, 1)
	statementOne.Action[0] = "execute-api:Invoke"
	statementOne.Resource[0] = *resource
	PolicyDocument.Statement[0] = statementOne

	AuthResponse.PolicyDocument = PolicyDocument
	return &AuthResponse
}	

func GenerateAllow(principalId *string, resource *string) *events.APIGatewayCustomAuthorizerResponse {
	Effect = "Allow"
	return GeneratePolicy(principalId,resource)
}

func GenerateDeny(principalId *string, resource *string) *events.APIGatewayCustomAuthorizerResponse {
	Effect = "Deny"
	return GeneratePolicy(principalId,resource)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	defer database.MasterDatabase.Close()
	defer database.SlaveDatabase.Close()

	lambda.Start(Handler)
}