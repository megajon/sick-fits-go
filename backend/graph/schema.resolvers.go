package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/megajon/sick-fits-go/graph/generated"
	"github.com/megajon/sick-fits-go/graph/model"
)

func (p *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := &model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}
	return p.UsersRepo.CreateUser(user)
}

func (u *mutationResolver) CreateProduct(ctx context.Context, input model.NewProduct) (*model.Product, error) {
	product := &model.Product{
		Name:        input.Name,
		Description: input.Description,
	}
	return u.ProductsRepo.CreateProduct(product)
}

func (u *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return u.UsersRepo.GetUsers()
}

func (u *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	return u.ProductsRepo.GetProducts()
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
