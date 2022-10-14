package aws

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"hometown.com/hometown-serverless-go/modules/common"
)


func GetNewS3() (*s3.S3, error) {
	if AwsSession != nil {
		return s3.New(AwsSession), nil
	}
	sess, sessionErr := session.NewSessionWithOptions(SessionOption)
	if sessionErr != nil {
		return nil, sessionErr
	}
	return s3.New(sess), nil
}

func UploadObjectToS3(bucket string, fileName string, file io.Reader, output *s3manager.UploadOutput) error {
	var uploader *s3manager.Uploader
	if AwsSession != nil {
		uploader = s3manager.NewUploader(AwsSession)
	} else {
		sess, sessionErr := session.NewSessionWithOptions(SessionOption)
		if sessionErr != nil {
			return sessionErr
		}
		uploader = s3manager.NewUploader(sess)
	}
	getOutput, uploadErr := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fileName),
		Body:   file,
	})

	if uploadErr != nil {
		return uploadErr
	}
	
	common.UnmarshalFromObject(getOutput,output)

	return nil
}

func DeleteObjectToS3(bucket *string, objectName *string, output *s3.DeleteObjectOutput) error {
	getS3, getErr := GetNewS3()
	if getErr != nil {
		return getErr
	}

	getOutput, deleteErr := getS3.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(*bucket),
		Key:    aws.String(*objectName),
	})

	if deleteErr != nil {
		return deleteErr
	}

	common.UnmarshalFromObject(getOutput,output)

	return nil
}
