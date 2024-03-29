package request

import (
	"github.com/hashicorp/go-retryablehttp"
	"go.uber.org/zap"
	"time"
)

func NewDefault() *retryablehttp.Client {
	client := retryablehttp.NewClient()
	client.RetryMax = 5
	client.RetryWaitMin = time.Millisecond * 100
	client.RetryWaitMax = time.Second * 2
	client.Logger = newLoggerWrapper(zap.L())
	return client
}

func NewWithCallerSkip(skip int) *retryablehttp.Client {
	client := NewDefault()
	client.Logger = newLoggerWrapper(zap.L().WithOptions(zap.AddCallerSkip(skip)))
	return client
}
