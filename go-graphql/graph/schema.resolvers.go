package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/fabio/graphql/database"
	"github.com/fabio/graphql/graph/model"
	"github.com/fabio/graphql/pkg/auth"
	"github.com/fabio/graphql/pkg/jwt"
	"github.com/fabio/graphql/pkg/utils"

	"github.com/jinzhu/gorm"
)

type contextKey struct {
	name string
}

var userCtxKey = &contextKey{"user"}

// CreateLink is the resolver for the createLink field.
func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {

	user := auth.FindUser(ctx)
	if user == nil {
		return &model.Link{}, fmt.Errorf("access denied")
	}

	db := database.Db

	graphqlUser := &model.User{
		ID:   user.ID,
		Name: user.Name,
	}

	// Create a new link
	link := &model.Link{
		Address: input.Address,
		Title:   input.Title,
		User:    graphqlUser,
	}

	// Save the link to the database
	if err := db.Create(&link).Error; err != nil {
		return nil, err
	}

	return link, nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	db := database.Db

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		log.Fatal(err)
	}

	// Create a new user
	user := &model.User{
		Name:     input.Username,
		Password: hashedPassword,
	}

	// Save the user to the database
	if err := db.Create(&user).Error; err != nil {
		return "", err
	}

	token, err := jwt.GenerateToken(user.Name)
	if err != nil {
		return "", err
	}

	return token, nil
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	db := database.Db

	var foundUser model.User
	if err := db.Where("Name = ?", input.Username).First(&foundUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", errors.New("wrong username or password")
		} else {
			return "", err
		}
	}

	correct := utils.CheckPasswordHash(input.Password, foundUser.Password)
	if !correct {
		return "", errors.New("wrong username or password")
	}

	token, err := jwt.GenerateToken(foundUser.Name)
	if err != nil {
		return "", err
	}

	return token, nil
}

// Links is the resolver for the links field.
func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	db := database.Db
	var links []*model.Link
	db.Find(&links)
	return links, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
