// This file is auto-generated by internal/cmd/genoptions/main.go. DO NOT EDIT

package httprc

import (
	"time"

	"github.com/lestrrat-go/option"
)

type Option = option.Interface

type MemoryCacheOption interface {
	Option
	memoryCacheOption()
}

type memoryCacheOption struct {
	Option
}

func (*memoryCacheOption) memoryCacheOption() {}

type RegisterOption interface {
	Option
	registerOption()
}

type registerOption struct {
	Option
}

func (*registerOption) registerOption() {}

type identHTTPClient struct{}
type identRefreshInterval struct{}
type identRefreshWindow struct{}

func (identHTTPClient) String() string {
	return "WithHTTPClient"
}

func (identRefreshInterval) String() string {
	return "WithRefreshInterval"
}

func (identRefreshWindow) String() string {
	return "WithRefreshWindow"
}

func WithHTTPClient(v HTTPClient) RegisterOption {
	return &registerOption{option.New(identHTTPClient{}, v)}
}

func WithRefreshInterval(v time.Duration) RegisterOption {
	return &registerOption{option.New(identRefreshInterval{}, v)}
}

func WithRefreshWindow(v time.Duration) MemoryCacheOption {
	return &memoryCacheOption{option.New(identRefreshWindow{}, v)}
}