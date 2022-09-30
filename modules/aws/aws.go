package aws

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)
var AwsConfig aws.Config = aws.Config{
	Region: aws.String("ap-northeast-2"),
}

var SessionOption session.Options = session.Options{
	Config: AwsConfig,
	Profile: "hometown",
}

var AwsSession *session.Session

func SetSession() {
	if AwsSession != nil {
		sess, sessionErr := session.NewSessionWithOptions(SessionOption)
		if sessionErr != nil {
			log.Panic(sessionErr)
		}
		AwsSession = sess
	}
}

