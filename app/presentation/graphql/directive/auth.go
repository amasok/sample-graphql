package directive

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/amasok/sample-graphql/app/domain/authToken"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func Auth(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	tokenData := authToken.CtxValue(ctx)
	if tokenData == nil {
		return nil, &gqlerror.Error{
			Message: "Access Denied",
		}
	}

	return next(ctx)
}
