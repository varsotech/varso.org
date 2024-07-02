package client

import (
	"fmt"

	"github.com/luminancetech/varso/src/common/api"
	"github.com/luminancetech/varso/src/services/auth/client/models"
)

func (c *Client) GetUser(userUUID string) (*models.GetUserResponse, *api.Error) {
	return api.SendRequest[any, models.GetUserResponse](c.Client, c.AuthHeader, "GET", fmt.Sprintf("auth/user/adminget/%s", userUUID), nil)
}
