package routes

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/luminancetech/varso/src/common/api"
	"github.com/luminancetech/varso/src/services/app/client/models"
	"github.com/luminancetech/varso/src/services/app/internal/modules/newsitem"
	"github.com/luminancetech/varso/src/services/auth/client"
)

func BlurToggle(w *api.Writer, r *http.Request, p httprouter.Params, j *api.JWT, request models.BlurToggleRequest) (*any, *api.Error) {
	if _, ok := j.Permissions[client.KeyNewsItemImageBlur]; !ok {
		return nil, api.NewForbiddenError(nil, "no permission to toggle blur")
	}

	err := newsitem.ToggleBlur(r.Context(), uuid.MustParse(request.RssItemId))
	if err != nil {
		return nil, api.NewInternalError(err, "failed toggling blur")
	}

	return nil, nil
}
