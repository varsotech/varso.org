package user

import (
	"context"

	"github.com/luminancetech/varso/src/services/auth/internal/config"
	"github.com/luminancetech/varso/src/services/auth/internal/ent/build"
	"github.com/luminancetech/varso/src/services/auth/internal/modules/role"
	"github.com/pkg/errors"
)

func Initialize(ctx context.Context) error {
	exist, err := Exist(ctx, config.AuthAdminEmail)
	if err != nil {
		return err
	}

	if exist {
		return nil
	}

	_, err = Register(ctx, config.AuthAdminEmail, config.AuthAdminPassword)
	if err != nil {
		return errors.Wrapf(err, "failed to register admin user")
	}

	adminUser, err := Get(ctx, config.AuthAdminEmail)
	if err != nil {
		return err
	}

	adminRole, err := role.Get(ctx, role.KeyAdmin)
	if err != nil {
		return err
	}

	err = SetRoles(ctx, adminUser, []*build.Role{adminRole})
	if err != nil {
		return err
	}

	return nil
}
