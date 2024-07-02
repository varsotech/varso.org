package log

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/luminancetech/varso/src/common/api"
	"github.com/luminancetech/varso/src/services/analytics/internal/modules/accesslog"
)

func AccessLog(_ *api.Writer, r *http.Request, params httprouter.Params, _ *api.JWT) (*any, *api.Error) {
	ip := params.ByName("ip")
	uri := params.ByName("uri")

	accesslog.Create(
		r.Context(),
		ip,
		uri,
		r.Header.Get("X-Forwarded-For"),
		r.Header.Get("X-Forwarded-Proto"),
		r.Header.Get("X-Forwarded-Host"),
		r.Header.Get("X-Forwarded-Port"),
		r.Header.Get("X-Forwarded-Server"),
		r.Header.Get("X-Request-ID"),
		r.Header.Get("User-Agent"),
	)

	return nil, api.NewNotFoundError(fmt.Errorf("returning not found to client"), "returning not found to client")
}
