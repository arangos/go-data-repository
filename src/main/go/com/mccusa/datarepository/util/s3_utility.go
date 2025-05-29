package util

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"strings"
)

// S3Utility provides helper methods for AWS S3 operations.
type S3Utility struct {
	Client *s3.Client
	Bucket string
}

// NewS3Utility constructs a new S3Utility with the given S3 client and bucket name.
func NewS3Utility(client *s3.Client, bucket string) *S3Utility {
	return &S3Utility{Client: client, Bucket: bucket}
}

// GetUniqueFileKeyName returns a file key that does not exist in S3 by appending a counter if needed.
func (u *S3Utility) GetUniqueFileKeyName(ctx context.Context, fileKeyName string) (string, error) {
	counter := 1
	unique := fileKeyName
	for {
		exists, err := u.DoesObjectExist(ctx, unique)
		if err != nil {
			return "", err
		}
		if !exists {
			break
		}
		base := strings.TrimSuffix(fileKeyName, ".pdf")
		unique = fmt.Sprintf("%s-%d.pdf", base, counter)
		counter++
	}
	return unique, nil
}

// DoesObjectExist checks if an object exists in the configured S3 bucket.
func (u *S3Utility) DoesObjectExist(ctx context.Context, key string) (bool, error) {
	_, err := u.Client.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: &u.Bucket,
		Key:    &key,
	})
	if err != nil {
		var nsk *types.NotFound
		if errors.As(err, &nsk) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
