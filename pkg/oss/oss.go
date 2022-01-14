package oss

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"os"
)

var (
	endpoint        = "oss-cn-beijing.aliyuncs.com"
	accessKeyID     = "xxx"
	accessKeySecret = "xxx"
	bucketName      = "xxx"
)

func Upload(fd io.Reader, objName string) {
	// 创建OSSClient实例。
	client, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		fmt.Printf("OSSUpload: %v", err)
		os.Exit(1)
	}

	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Printf("OSSUpload: %v", err)
		os.Exit(1)
	}

	// 上传文件流。
	err = bucket.PutObject(objName, fd)
	if err != nil {
		fmt.Printf("OSSUpload: %v", err)
		os.Exit(1)
	}
	//fmt.Println("[*] Upload success")
}

func UploadFile(localFile string, objName string) {
	// 创建OSSClient实例。
	client, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		fmt.Printf("OSSUpload: %v", err)
		os.Exit(1)
	}

	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Printf("OSSUpload: %v", err)
		os.Exit(1)
	}

	// 读取本地文件。
	fd, err := os.Open(localFile)
	if err != nil {
		fmt.Printf("OSSUpload: %v", err)
		os.Exit(1)
	}
	defer fd.Close()

	// 上传文件流。
	err = bucket.PutObject(objName, fd)
	if err != nil {
		fmt.Printf("OSSUpload: %v", err)
		os.Exit(1)
	}
	//fmt.Println("[*] Upload success")
}

func Download(objName string) io.Reader {
	var fd io.Reader
	// 创建OSSClient实例。
	client, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		fmt.Printf("OSSDownload: %v", err)
		os.Exit(1)
	}

	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Printf("OSSDownload: %v", err)
		os.Exit(1)
	}

	// 下载文件到本地文件。
	fd, err = bucket.GetObject(objName)
	if err != nil {
		fmt.Printf("OSSDownload: %v", err)
		os.Exit(1)
	}
	//fmt.Println("[*] Download success")
	return fd
}

func DownloadFile(savePath string, objName string) {
	// 创建OSSClient实例。
	client, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		fmt.Printf("OSSDownload: %v", err)
		os.Exit(1)
	}

	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Printf("OSSDownload: %v", err)
		os.Exit(1)
	}

	// 下载文件到本地文件。
	err = bucket.GetObjectToFile(objName, savePath)
	if err != nil {
		fmt.Printf("OSSDownload: %v", err)
		os.Exit(1)
	}
	//fmt.Println("[*] Download success")
}
