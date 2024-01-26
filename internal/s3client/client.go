// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package s3client

import (
	"context"
	"io"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go"
	"github.com/aws/smithy-go/logging"
	"github.com/juju/errors"
)

// Logger represents the logging methods called.
type Logger interface {
	Errorf(message string, args ...any)
	Warningf(message string, args ...any)
	Infof(message string, args ...any)
	Debugf(message string, args ...any)
	Tracef(message string, args ...any)

	IsTraceEnabled() bool
}

// HTTPClient represents the http client used to access the object store.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// CredentialsKind represents the kind of credentials used to access the object
// store.
type CredentialsKind string

const (
	// AnonymousCredentialsKind represents anonymous credentials.
	AnonymousCredentialsKind CredentialsKind = "anonymous"
	// StaticCredentialsKind represents static credentials.
	StaticCredentialsKind CredentialsKind = "static"
)

// Credentials represents the credentials used to access the object store.
type Credentials interface {
	Kind() CredentialsKind
}

// AnonymousCredentials represents anonymous credentials.
type AnonymousCredentials struct {
	Credentials
}

// Kind returns the kind of credentials.
func (AnonymousCredentials) Kind() CredentialsKind {
	return AnonymousCredentialsKind
}

// S3Client is a Juju shim around the AWS S3 client,
// which Juju uses to drive its object store requirements.
// StaticCredentials represents static credentials.
type StaticCredentials struct {
	Key     string
	Secret  string
	Session string
}

// Kind returns the kind of credentials.
func (StaticCredentials) Kind() CredentialsKind {
	return StaticCredentialsKind
}

// objectsClient is a Juju shim around the AWS S3 client,
// which Juju uses to drive it's object store requirents
type S3Client struct {
	logger Logger
	client *s3.Client
}

// NewS3Client returns a new s3Caller client for accessing the object store.
func NewS3Client(baseURL string, httpClient HTTPClient, credentials Credentials, logger Logger) (*S3Client, error) {
	credentialsProvider, err := getCredentialsProvider(credentials)
	if err != nil {
		return nil, errors.Annotate(err, "cannot get credentials provider")
	}

	awsLogger := &awsLogger{
		logger: logger,
	}

	cfg, err := config.LoadDefaultConfig(
		context.Background(),
		config.WithLogger(awsLogger),
		config.WithHTTPClient(httpClient),
		config.WithEndpointResolverWithOptions(&awsEndpointResolver{endpoint: baseURL}),
		// Standard retryer with custom max attempts. Will retry at most
		// 10 times with 20s backoff time.
		config.WithRetryer(func() aws.Retryer {
			return retry.NewStandard(
				func(o *retry.StandardOptions) {
					o.MaxAttempts = 10
					o.RateLimiter = unlimitedRateLimiter{}
				},
			)
		}),

		// The anonymous credentials are needed by the aws sdk to
		// perform anonymous s3 access.
		config.WithCredentialsProvider(credentialsProvider),
	)
	if err != nil {
		return nil, errors.Annotate(err, "cannot load default config for s3 client")
	}

	return &S3Client{
		client: s3.NewFromConfig(cfg, func(o *s3.Options) {
			o.UsePathStyle = true
		}),
		logger: logger,
	}, nil
}

// GetObject gets an object from the object store based on the bucket name and
// object name.
func (c *S3Client) GetObject(ctx context.Context, bucketName, objectName string) (io.ReadCloser, int64, string, error) {
	c.logger.Tracef("getting bucket %s object %s from s3 storage", bucketName, objectName)

	obj, err := c.client.GetObject(ctx,
		&s3.GetObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(objectName),
		})
	if err != nil {
		if err := handleError(err); err != nil {
			return nil, -1, "", errors.Trace(err)
		}
		return nil, -1, "", errors.Annotatef(err, "getting object %s on bucket %s using S3 client", objectName, bucketName)
	}
	var size int64
	if obj.ContentLength != nil {
		size = *obj.ContentLength
	}
	var hash string
	if obj.ChecksumSHA256 != nil {
		hash = *obj.ChecksumSHA256
	}
	return obj.Body, size, hash, nil
}

// PutObject puts an object into the object store based on the bucket name and
// object name.
func (c *S3Client) PutObject(ctx context.Context, bucketName, objectName string, body io.Reader, hash string) error {
	c.logger.Tracef("putting bucket %s object %s to s3 storage", bucketName, objectName)

	obj, err := c.client.PutObject(ctx,
		&s3.PutObjectInput{
			Bucket:            aws.String(bucketName),
			Key:               aws.String(objectName),
			Body:              body,
			ChecksumAlgorithm: types.ChecksumAlgorithmSha256,
		})
	if err != nil {
		if err := handleError(err); err != nil {
			return errors.Trace(err)
		}
		return errors.Annotatef(err, "putting object %s on bucket %s using S3 client", objectName, bucketName)
	}
	if hash != "" && obj.ChecksumSHA256 != nil && hash != *obj.ChecksumSHA256 {
		return errors.Errorf("hash mismatch, expected %q got %q", hash, *obj.ChecksumSHA256)
	}
	return nil
}

// DeleteObject deletes an object from the object store based on the bucket name
// and object name.
func (c *S3Client) DeleteObject(ctx context.Context, bucketName, objectName string) error {
	c.logger.Tracef("deleting bucket %s object %s from s3 storage", bucketName, objectName)

	_, err := c.client.DeleteObject(ctx,
		&s3.DeleteObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(objectName),
		})
	if err != nil {
		if err := handleError(err); err != nil {
			return errors.Trace(err)
		}
		return errors.Annotatef(err, "deleting object %s on bucket %s using S3 client", objectName, bucketName)
	}
	return nil
}

// CreateBucket creates a bucket in the object store based on the bucket name.
func (c *S3Client) CreateBucket(ctx context.Context, bucketName string) error {
	c.logger.Tracef("creating bucket %s in s3 storage", bucketName)

	_, err := c.client.CreateBucket(ctx,
		&s3.CreateBucketInput{
			Bucket: aws.String(bucketName),
		})
	if err != nil {
		if err := handleError(err); err != nil {
			return errors.Trace(err)
		}
		return errors.Annotatef(err, "unable to create bucket %s using S3 client", bucketName)
	}
	return nil
}

// forbiddenErrorCodes is a list of error codes that are returned when the
// credentials are invalid.
// https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html#ErrorCodeList
var forbiddenErrorCodes = map[string]struct{}{
	"AccessDenied":          {},
	"InvalidAccessKeyId":    {},
	"InvalidSecurity":       {},
	"SignatureDoesNotMatch": {},
}

var alreadyExistCodes = map[string]struct{}{
	"BucketAlreadyExists":     {},
	"BucketAlreadyOwnedByYou": {},
}

var notFoundCodes = map[string]struct{}{
	"NoSuchBucket": {},
	"NoSuchKey":    {},
}

func handleError(err error) error {
	if err == nil {
		return nil
	}

	var ae smithy.APIError
	if errors.As(err, &ae) {
		if _, ok := notFoundCodes[ae.ErrorCode()]; ok {
			return errors.NewNotFound(err, ae.ErrorMessage())
		}
		if _, ok := forbiddenErrorCodes[ae.ErrorCode()]; ok {
			return errors.NewForbidden(err, ae.ErrorMessage())
		}
		if _, ok := alreadyExistCodes[ae.ErrorCode()]; ok {
			return errors.NewAlreadyExists(err, ae.ErrorMessage())
		}

	}

	return errors.Trace(err)
}

type awsEndpointResolver struct {
	endpoint string
}

// ResolveEndpoint returns the endpoint for the given service and region.
func (a *awsEndpointResolver) ResolveEndpoint(_, _ string, options ...any) (aws.Endpoint, error) {
	return aws.Endpoint{
		URL: a.endpoint,
	}, nil
}

type awsLogger struct {
	logger Logger
}

func (l *awsLogger) Logf(classification logging.Classification, format string, v ...any) {
	switch classification {
	case logging.Warn:
		l.logger.Warningf(format, v)
	case logging.Debug:
		l.logger.Debugf(format, v)
	default:
		l.logger.Tracef(format, v)
	}
}

func getCredentialsProvider(creds Credentials) (aws.CredentialsProvider, error) {
	switch creds.Kind() {
	case AnonymousCredentialsKind:
		return aws.AnonymousCredentials{}, nil
	case StaticCredentialsKind:
		s := creds.(StaticCredentials)
		return credentials.NewStaticCredentialsProvider(s.Key, s.Secret, s.Session), nil
	default:
		return nil, errors.Errorf("unknown credentials kind %q", creds.Kind())
	}
}

type unlimitedRateLimiter struct{}

func (unlimitedRateLimiter) AddTokens(uint) error { return nil }
func (unlimitedRateLimiter) GetToken(context.Context, uint) (func() error, error) {
	return noOpToken, nil
}
func noOpToken() error { return nil }
