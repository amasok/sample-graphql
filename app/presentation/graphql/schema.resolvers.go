package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/amasok/sample-graphql/app/domain/authToken"
	"github.com/amasok/sample-graphql/app/presentation/graphql/generated"
	"github.com/amasok/sample-graphql/app/presentation/graphql/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	var findUser *model.User
	for _, user := range r.users {
		if user.ID == input.UserID {
			findUser = user
			break
		}
	}

	if findUser == nil {
		return nil, fmt.Errorf("not found user_id: %s", input.UserID)
	}
	todo := &model.Todo{
		Text: input.Text,
		ID:   fmt.Sprintf("T%d", rand.Int()),
		User: findUser,
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := &model.User{
		ID:       fmt.Sprintf("T%d", rand.Int()),
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}
	r.users = append(r.users, user)
	return user, nil
}

func (r *mutationResolver) Authorization(ctx context.Context, input model.AuthorizationRequest) (*model.AuthorizationResponse, error) {
	// panic(fmt.Errorf("not implemented"))
	var findUser *model.User
	for _, user := range r.users {
		// まだちゃんとした実装でないため全件回してとってきてる
		if user.Email == input.Email && user.Password == input.Password {
			findUser = user
			break
		}
	}

	if findUser == nil {
		// FIXME: 本来セキュリティ的にメールアドレスをログに出すべきでない
		return nil, fmt.Errorf("not found user_email: %s", input.Email)
	}

	authToken, err := authToken.New(findUser.ID, time.Now())
	if err != nil {
		return nil, err
	}

	return &model.AuthorizationResponse{
		Token: authToken.GetToken(),
	}, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return r.users, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
