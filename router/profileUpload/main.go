package main

import (
	"encoding/json"
	"runtime"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"hometown.com/hometown-serverless-go/modules/aws"
	"hometown.com/hometown-serverless-go/modules/image"
)

// func parseForm(key, mpheader string, body io.Reader) (string, string, io.Reader, error) {
// 	var buf []byte
// 	filename := "unknown"
// 	contentType := "*"
// 	mediaType, params, err := mime.ParseMediaType(mpheader)
// 	fmt.Println(mediaType)
// 	if err != nil {
// 		return "", "", nil, err
// 	}
// 	if strings.HasPrefix(mediaType, "multipart/") {
// 		mr := multipart.NewReader(body, params["boundary"])
// 		fmt.Printf("DEBUG:: boundary: %v\n", params["boundary"])
// 		for {
// 			p, err := mr.NextPart()

// 			fmt.Println(p)
// 			fmt.Println(err)
// 			if err == io.EOF {
// 				fmt.Printf("DEBUG:: parse EOF\n")
// 				break
// 			}
// 			if err != nil {
// 				return "", "", nil, err
// 			}
// 			buf, err = ioutil.ReadAll(p)
// 			if err != nil {
// 				return "", "", nil, err
// 			}
// 			fmt.Printf("DEBUG:: part: %v\n", p.Header)
// 			filename = p.FileName()
// 			contentType = p.Header.Get("Content-Type")
// 		}
// 	}
// 	fmt.Printf("DEBUG:: parsed file name: %v\n", filename)
// 	fmt.Printf("DEBUG:: parsed content type: %v\n", contentType)
// 	fmt.Printf("DEBUG:: parsed size in bytes: %v\n", len(buf))
// 	// fmt.Printf("DEBUG:: parsed content: %v\n", string(buf))
// 	return filename, contentType, strings.NewReader(string(buf)), nil
// }

func Handler(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	res := events.APIGatewayProxyResponse{
		Body: "",
		StatusCode: 200,
	}
	byteReaderArr, imgNameArr, readErr := image.GetByteReaderByMultipart(&event)
	var output s3manager.UploadOutput
	
	if readErr != nil && byteReaderArr == nil && imgNameArr == nil {
		errBin,_ := json.Marshal(readErr)
		res.Body = string(errBin)
		return res, readErr
	}

	for i := 0; i < len(byteReaderArr); i++ {
		awsErr := aws.UploadObjectToS3("hometown-user-bucket","profile/" + imgNameArr[i], byteReaderArr[i],&output)
		if awsErr != nil {
			errBin,_ := json.Marshal(awsErr)
			res.Body = string(errBin)
			return res, awsErr
		}	
	} 
	
	outputBin, _ := json.Marshal(output)
	res.Body = string(outputBin)
	return res, nil
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	lambda.Start(Handler)
}