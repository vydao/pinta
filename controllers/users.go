package controllers

import (
	"fmt"
	"net/http"

	"pinta.com/views"
)

type Users struct {
	NewView *views.View
}

func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("views/users/new.gohtml"),
	}
}

// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is a temporary response.")
}
