package main

import (
	"fmt"
	"net/url"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	OssEndpoint        = "http://10.4.208.193:30111"
	OssSecretAccessKey = "mIyKT0ArkDr3ISZ8k3esTz2qiJTloVRp9a8LHdtp"
	OssAccessKeyId     = "SDS4X5Q3L4NRLIP8FWAC"
	OssRegion          = "sensetime-5"
	OssBucket          = "test-bkt-1"
)

func main() {
	creds := credentials.NewStaticCredentials(OssAccessKeyId, OssSecretAccessKey, "")
	sess, err := session.NewSession(&aws.Config{
		Region:           aws.String(OssRegion),
		Endpoint:         &OssEndpoint,
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
		Credentials:      creds,
	})
	if err != nil {
		panic(err)
	}
	service := s3.New(sess)

	// List all buckets
	result, err := service.ListBuckets(nil)
	if err != nil {
		panic(err)
	}
	for _, b := range result.Buckets {
		fmt.Printf("* %s created on %s\n",
			aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}

	// Upload to bucket
	//filename := "test.jpeg"
	//cwd, err := os.Getwd()
	//if err != nil {
	//	panic(err)
	//}
	//fp := filepath.Join(cwd, "oss", filename)
	//uploader := s3manager.NewUploader(sess)
	//file, err := os.Open(fp)
	//if err != nil {
	//	panic(err)
	//}
	//defer file.Close()
	//_, err = uploader.Upload(&s3manager.UploadInput{
	//	Bucket: aws.String(OssBucket),
	//	Key:    aws.String(filename),
	//	Body:   file,
	//})
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Successfully uploaded %q to %q\n", filename, OssBucket)

	// List objetcs in bucket
	resp, err := service.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(OssBucket),
	})
	if err != nil {
		panic(err)
	}
	for _, item := range resp.Contents {
		fmt.Println("Name:         ", *item.Key)
		fmt.Println("Last modified:", *item.LastModified)
		fmt.Println("Size:         ", *item.Size)
		fmt.Println("Storage class:", *item.StorageClass)
		fmt.Println("")
	}

	// download object from bucket
	//file, err = os.Create("test_download.jpeg")
	//if err != nil {
	//	panic(err)
	//}
	//defer file.Close()
	//downloader := s3manager.NewDownloader(sess)
	//numBytes, err := downloader.Download(file,
	//	&s3.GetObjectInput{
	//		Bucket: aws.String(OssBucket),
	//		Key:    aws.String("test.jpeg"),
	//	})
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("Downloaded", file.Name(), numBytes, "bytes")

	// Grant public-read access to object in bucket
	acl := "public-read"
	params := &s3.PutObjectAclInput{
		ACL:    &acl,
		Bucket: aws.String(OssBucket),
		Key:    aws.String("siluan.json"),
	}
	_, err = service.PutObjectAcl(params)
	if err != nil {
		panic(err)
	}
	download_link, err := url.Parse(OssEndpoint)
	if err != nil {
		panic(err)
	}
	download_link.Path = OssBucket + "/siluan.json"
	fmt.Println("Object " + *params.Key + " is now public")
	fmt.Printf("Download URL is: %s\n", download_link)
}
