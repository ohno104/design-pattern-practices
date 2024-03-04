package main

import (
	"time"
)

type Connection struct {
	addr    string
	cache   bool
	timeout time.Duration
}

const (
	defaultTimeout = 10
	defaultCaching = false
)

type options struct {
	timeout time.Duration
	caching bool
}

type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
	f(o)
}

func WithTimeout(t time.Duration) Option {
	return optionFunc(func(o *options) {
		o.timeout = t
	})
}

func WithCaching(cache bool) Option {
	return optionFunc(func(o *options) {
		o.caching = cache
	})
}

func NewConnection(addr string, opts ...Option) *Connection {
	def := options{
		timeout: defaultTimeout,
		caching: defaultCaching,
	}

	for _, o := range opts {
		o.apply(&def)
	}

	return &Connection{
		addr:    addr,
		timeout: def.timeout,
		cache:   def.caching,
	}

}
