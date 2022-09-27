package main

import (
	"errors"
	"runtime"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

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

	if headers["x-api-key"] != "my-api-key" {
		return GenerateDeny("nspark",event.MethodArn), errors.New("Unauthorized")
	}
	return GenerateAllow("nspark",event.MethodArn), nil
}

func GeneratePolicy(principalId string, effect string, resource string) events.APIGatewayCustomAuthorizerResponse {
	var AuthResponse events.APIGatewayCustomAuthorizerResponse;
	AuthResponse.PrincipalID = principalId
	var PolicyDocument events.APIGatewayCustomAuthorizerPolicy

	PolicyDocument.Version = "2012-10-17"
	var statement []events.IAMPolicyStatement
	PolicyDocument.Statement =  statement

	var statementOne events.IAMPolicyStatement
	statementOne.Action = make([]string, 1)
	statementOne.Effect = effect
	statementOne.Resource = make([]string, 1)
	statementOne.Action[0] = "execute-api:Invoke"
	statementOne.Resource[0] = resource
	PolicyDocument.Statement[0] = statementOne

	AuthResponse.PolicyDocument = PolicyDocument

	return AuthResponse
}	

func GenerateAllow(principalId string, resource string) events.APIGatewayCustomAuthorizerResponse {
	return GeneratePolicy(principalId,"Allow",resource)
}

func GenerateDeny(principalId string, resource string) events.APIGatewayCustomAuthorizerResponse {
	return GeneratePolicy(principalId,"Deny",resource)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	lambda.Start(Handler)
}