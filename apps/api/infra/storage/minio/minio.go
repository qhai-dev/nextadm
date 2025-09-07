package minio

import (
	"context"
	"io"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Client struct {
	host            string
	client          *minio.Client
	accessKeyID     string
	secretAccessKey string
	bucketName      string
	endpoint        string
}

func NewClient(ctx context.Context, endpoint, accessKeyID, secretAccessKey string) (*Client, error) {
	client, err := minio.New(endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		//Secure: true,
	})
	if err != nil {
		return nil, err
	}
	m := &Client{
		client: client,
	}
	return m, nil
}

func (c *Client) CreateBucketIfNeed(ctx context.Context, bucketName, region string) (string, error) {
	// 创建一个新的存储桶
	err := c.client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: region})
	if err != nil {
		// 检查存储桶是否存在
		exists, errBucketExists := c.client.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			return bucketName, nil
		}

		return "", errBucketExists
	}
	return bucketName, nil
}

func (c *Client) GetObject(ctx context.Context, bucketName, objectName string) (io.ReadCloser, error) {
	object, err := c.client.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	return object, nil
}

func (c *Client) PutObject() {

}

func (c *Client) RemoveObject() {

}

func (c *Client) SetBucketPolicy(ctx context.Context, bucketName string) error {
	policy := `{"Version": "2012-10-17","Statement": [{"Action": ["s3:GetObject"],"Effect": "Allow","Principal": {"AWS": ["*"]},"Resource": ["arn:aws:s3:::nextadm/*"],"Sid": ""}]}`

	err := c.client.SetBucketPolicy(context.Background(), bucketName, policy)
	if err != nil {
		return err
	}
	return nil
}
