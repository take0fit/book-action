package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"book-action/interface/gql/generated"
	"book-action/interface/gql/model"
	"book-action/internal/application/dto"
	"context"
	"fmt"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, username string, age int) (*model.User, error) {
	inputUser := dto.UserCreateInput{
		Name: username,
		Age:  age,
	}
	user, err := r.userUsecase.CreateUser(ctx, inputUser)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	fmt.Println(3, user)
	gqlUser := &model.User{
		ID:        user.ID,
		Name:      user.Name,
		Age:       user.Age,
		Resources: nil,
	}

	return gqlUser, err
}

// CreateResource is the resolver for the createResource field.
func (r *mutationResolver) CreateResource(ctx context.Context, title string, categoryID string) (*model.Resource, error) {
	panic(fmt.Errorf("not implemented: CreateResource - createResource"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	user, err := r.userUsecase.GetUserDetails(id)

	gqlUser := &model.User{
		ID:        user.ID,
		Name:      user.Name,
		Resources: nil,
	}

	return gqlUser, err
}

// Resources is the resolver for the resources field.
func (r *queryResolver) Resources(ctx context.Context) ([]*model.Resource, error) {
	panic(fmt.Errorf("not implemented: Resources - resources"))
}

// Resource is the resolver for the resource field.
func (r *queryResolver) Resource(ctx context.Context, id string) (*model.Resource, error) {
	panic(fmt.Errorf("not implemented: Resource - resource"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
