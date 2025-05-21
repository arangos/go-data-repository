package config

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	s3presign "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// AWS S3 folders
const (
	JobApplicationsFolder = "EB-3-JOB-APPLICATIONS/"
	ContractsFolder       = "EB-3-CONTRACTS/"
	VisaDocumentsFolder   = "IND-VISA-DOCUMENTS/"
)

// NewS3Client returns an AWS S3 client configured via Viper
func NewS3Client(ctx context.Context) (*s3.Client, error) {
	bucket := viper.GetString("mccusa.bucket.name")
	if bucket == "" {
		logrus.Warn("mccusa.bucket.name not set in config")
	}

	// Load default AWS config, using profile 'mccusa'
	awsCfg, err := config.LoadDefaultConfig(ctx,
		config.WithSharedConfigProfile("mccusa"),
		config.WithRegion("us-east-1"),
	)
	if err != nil {
		return nil, err
	}

	return s3.NewFromConfig(awsCfg), nil
}

// NewS3Presigner returns an AWS S3 presigner for generating pre-signed URLs
func NewS3Presigner(ctx context.Context) (*s3presign.PresignClient, error) {
	awsCfg, err := config.LoadDefaultConfig(ctx,
		config.WithSharedConfigProfile("mccusa"),
		config.WithRegion("us-east-1"),
	)
	if err != nil {
		return nil, err
	}

	return s3presign.NewPresignClient(s3.NewFromConfig(awsCfg)), nil
}
