package main

import (
	"flag"

	ms3 "github.com/gerlacdt/aws-s3/pkg/s3"

	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func getSession(region string) *session.Session {
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String(region)}))
	return sess
}

func parametersValid(bucket, objectKey, filename string) bool {
	if bucket == "" || objectKey == "" || filename == "" {
		return false
	}
	return true
}

func main() {
	// parse cli parameters
	bucketname := flag.String("bucket", "", "s3 bucketname")
	objectKey := flag.String("objectkey", "", "bucket object key")
	filename := flag.String("filename", "", "filename for input or output")
	region := flag.String("region", "eu-west-1", "aws region")
	download := flag.Bool("download", false, "download file from aws s3")
	upload := flag.Bool("upload", false, "upload file from aws s3")
	kmskey := flag.String("kmskey", "", "kmskey to use for encryption")

	flag.Parse()

	if parametersValid(*bucketname, *objectKey, *filename) == false {
		log.Fatal("Input parameter missing")
	}

	if *download == false && *upload == false {
		log.Fatal("Specify -download or -upload")
	}

	sess := getSession(*region)

	if *download {
		ms3.Download(sess, *bucketname, *objectKey, *filename)
		return
	}

	if *upload {
		ms3.Upload(sess, *bucketname, *objectKey, *filename, *kmskey)
		return
	}
}
