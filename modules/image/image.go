package image

import (
	"bytes"
	"io"

	"github.com/aws/aws-lambda-go/events"
	"github.com/google/uuid"
	"github.com/grokify/go-awslambda"
)

func GetByteReaderByMultipart(event *events.APIGatewayProxyRequest) ([]*bytes.Reader, []string, error) {
	r , readErr := awslambda.NewReaderMultipart(*event)
	if readErr != nil {
		return nil, nil, readErr
	}
	var err error
	var readerArr []*bytes.Reader
	var nameArr []string
	for {
		part, nexterr := r.NextPart()
		
		if part != nil {
			defer part.Close()
		}
		
		if nexterr == io.EOF {
			break
		} else if nexterr != nil {
			readerArr = nil
			nameArr = nil
			err = nexterr
			break
		}
		content, ioErr := io.ReadAll(part)
		if ioErr != nil {
			readerArr = nil
			nameArr = nil
			err = ioErr
			break
		}
		readerArr = append(readerArr,bytes.NewReader(content))
		nameArr = append(nameArr,uuid.NewString() + "-" + part.FileName())
	}
	return readerArr, nameArr, err
}