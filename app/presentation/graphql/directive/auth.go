package directive

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/amasok/sample-graphql/app/domain/authToken"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func Auth(ctx context.Context, _ interface{}, next graphql.Resolver) (res interface{}, err error) {
	tokenData, err := authToken.GetToken(ctx)
	if err != nil {
		return nil, &gqlerror.Error{
			Message: "Access Denied",
		}
	}
	// ログインユーザーIDを見てみる
	fmt.Printf("login user_id: %s", tokenData.GetUserID())
	return next(ctx)
}
