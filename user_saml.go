package user_xsaml_idp

import (
	ect_samlidp "github.com/moisespsena/go-ecletus-samlidp"
)

type UserFactory struct {
}

func (this UserFactory) SamlIDPUserFactory(ctx *ect_samlidp.UserFactoryContext) (samlUser *ect_samlidp.User, err error) {
	samlUser = ect_samlidp.NewUser(
		ect_samlidp.GetNameId(ctx),
		ctx.User.GetName(),
		ctx.User.GetEmail(),
	)
	return
}
