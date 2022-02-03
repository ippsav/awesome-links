package hook

import (
	"context"

	"github.com/ippsav/awesome-links/api/ent"
	"golang.org/x/crypto/bcrypt"
)

func UserPasswordHook(next ent.Mutator) ent.Mutator {
	type PasswordSetter interface {
		SetPassword(value string)
	}
	return ent.MutateFunc(func(c context.Context, m ent.Mutation) (ent.Value, error) {
		if ps, ok := m.(PasswordSetter); ok {
			password, _ := m.Field("password")
			if password != "" {
				hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password.(string)), 2)
				if err != nil {
					return nil, err
				}
				ps.SetPassword(string(hashedPassword))
			}
		}
		return next.Mutate(c, m)
	})
}
