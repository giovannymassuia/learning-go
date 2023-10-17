package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"os"
	"sync"
)

var (
	s3Client *s3.S3
	s3Bucket string
	wg       sync.WaitGroup
)

// init() is called before main()
func init() {

	// config using for aws sso credentials
	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Profile:           os.Getenv("AWS_PROFILE"),
	})
	if err != nil {
		panic(err)
	}

	s3Client = s3.New(sess)
	s3Bucket = "learning-go-upload-s3-example"
}

func main() {
	dir, err := os.Open("./tmp")
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	// create a channel with a buffer of 100
	// this channel will be used to control the number of goroutines
	// that are uploading files to S3
	uploadControl := make(chan struct{}, 100)

	for {
		files, err := dir.Readdir(1)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("Error reading directory: %s\n", err.Error())
			continue
		}

		wg.Add(1)

		// this line will block the execution of the for loop
		// until there is a free slot in the channel
		uploadControl <- struct{}{}

		go uploadFile(files[0].Name(), uploadControl)
	}

	wg.Wait()
}

func uploadFile(filename string, uploadControl <-chan struct{}) {
	defer wg.Done()

	completeFileName := fmt.Sprintf("./tmp/%s", filename)
	f, err := os.Open(completeFileName)
	if err != nil {
		fmt.Printf("Error opening file %s: %s\n", completeFileName, err.Error())
		<-uploadControl // release a slot in the channel
		return
	}
	defer f.Close()

	// upload file to S3
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: &s3Bucket,
		Key:    aws.String(filename),
		Body:   f,
	})
	if err != nil {
		fmt.Printf("Error uploading file %s: %s\n", completeFileName, err.Error())
		<-uploadControl // release a slot in the channel
		return
	}

	fmt.Printf("File %s uploaded successfully\n", completeFileName)
	<-uploadControl // release a slot in the channel
}
