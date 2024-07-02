package accesslog

import (
	"context"

	"github.com/luminancetech/varso/src/services/analytics/internal/ent"
	"github.com/pkg/errors"
)

/*
	Users can submit posts in either the main chart or the friends chart.
*/

func Create(ctx context.Context, ip, uri, forwardedFor, proto, host, port, server, id, userAgent string) error {
	err := ent.Database.AccessLog.Create().
		SetIP(ip).
		SetURI(uri).
		SetForwardedFor(forwardedFor).
		SetForwardedHost(host).
		SetForwardedPort(port).
		SetForwardedServer(server).
		SetForwardedProto(proto).
		SetRequestID(id).
		SetUserAgent(userAgent).
		Exec(ctx)
	if err != nil {
		return errors.Wrapf(err, "failed creating access log in db")
	}

	return nil
}
