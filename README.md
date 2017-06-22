# AWS S3 upload and download

This is a simple cli-tool to upload and download files to AWS S3.
It also supports server side encryption via aws kms.

## Rationale

You can use this script to manage your environment variables and your secrets.
Simply run *upload* to upload your upload your environment variables to AWS S3.
During your service startup you run *download* and source the downloaded file.

This works also with secrets because you can use *server-side-encryption* with
AWS KMS.


## Installation

You have to install and configure awscli or export the aws environment
variables.

   + AWS_ACCESS_KEY_ID
   + AWS_SECRET_ACCESS_KEY

[AWS Configuration Reference](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html)

```bash
# install go dependencies
go get -u github.com/aws/aws-sdk-go/...
```

In order to use this script with AWS S3 and AWS KMS encryption you need to have
access to the specified S3 bucket and AWS KMS key via IAM roles and policies.

## Usage

```bash
# download to filename
go run cmd/s3/main.go -download -region eu-west-1 -bucket bucket-foobar-2 -objectkey test.txt -filename test.txt

# upload filename without encryption
go run cmd/s3/main.go -upload -region eu-west-1 -bucket bucket-foobar-2 -objectkey test.txt -filename test.txt

# upload with server side encryption with kms
# specified kms-key is used for encryption and stored as metadata. Hence downloading and decryption is transparent.
go run cmd/s3/main.go -upload -region eu-west-1 -bucket bucket-foobar-2 -objectkey test.txt -filename test.txt -kmskey arn:aws:kms:eu-west-1:kms-key-id
```
