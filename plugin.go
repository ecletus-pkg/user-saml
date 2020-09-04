package user_xsaml_idp

import (
	"github.com/ecletus/plug"
)

type Plugin struct {
	UserFactoryKey string
}

func (p Plugin) ProvideOptions() []string {
	return []string{p.UserFactoryKey}
}

func (p Plugin) ProvidesOptions(options *plug.Options) {
	options.Set(p.UserFactoryKey, &UserFactory{})
}