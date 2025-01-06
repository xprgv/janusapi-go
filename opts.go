package janusapi

import (
	"time"

	"github.com/xprgv/janusapi-go/pkg/logger"
	"github.com/xprgv/janusapi-go/pkg/model"
)

func getDefaultOptions() *Options {
	return &Options{
		JanusConnectTimeout:    5 * time.Second,
		JanusReconnectInterval: 5 * time.Second,
		TransactionTTL:         10 * time.Second,
		EventChannel:           nil,
		Logger:                 logger.Null(),
	}
}

type Options struct {
	JanusConnectTimeout    time.Duration
	JanusReconnectInterval time.Duration
	TransactionTTL         time.Duration
	EventChannel           chan model.LibraryEvent
	Logger                 logger.Logger
}

type Option func(*Options) error

func WithJanusConnectTimeout(d time.Duration) Option {
	return func(o *Options) error {
		o.JanusConnectTimeout = d
		return nil
	}
}

func WithJanusReconnectInterval(d time.Duration) Option {
	return func(o *Options) error {
		o.JanusReconnectInterval = d
		return nil
	}
}

func WithTransactionTTL(d time.Duration) Option {
	return func(o *Options) error {
		o.TransactionTTL = d
		return nil
	}
}

func WithEventChannel(ch chan model.LibraryEvent) Option {
	return func(o *Options) error {
		o.EventChannel = ch
		return nil
	}
}

func WithLogger(logger logger.Logger) Option {
	return func(o *Options) error {
		o.Logger = logger
		return nil
	}
}
