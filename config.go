package main

import (
	"errors"
	"github.com/asaskevich/govalidator"
)

type Config interface {
	Validate() error
	API() string
	URL() string
}

type Flag struct {
	APIKey  string
	WebSite string
}

func NewConfig(APIkey, WebSite string) Config {
	return &Flag{
		APIKey:  APIkey,
		WebSite: WebSite,
	}
}

func (f *Flag) Validate() error {
	if len(f.APIKey) == 0 {
		return errors.New("Invalid API Key")
	}
	if len(f.WebSite) == 0 {
		return errors.New("Invalid WebSite")
	}
	if !govalidator.IsURL(f.WebSite) {
		return errors.New("Invalid WebSite")
	}

	return nil
}

func (f *Flag) API() string {
	return f.APIKey
}

func (f *Flag) URL() string {
	return f.WebSite
}
