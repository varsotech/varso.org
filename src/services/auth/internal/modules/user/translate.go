package user

import (
	"github.com/luminancetech/varso/src/services/auth/client/models"
	"github.com/luminancetech/varso/src/services/auth/internal/ent/build"
)

func TranslateUser(user *build.User) *models.User {
	return &models.User{
		Uuid:   user.ID.String(),
		Banned: !user.BanTime.IsZero(),
	}
}
