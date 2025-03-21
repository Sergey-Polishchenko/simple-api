// Package domain contains core business entities.
package domain

import "strings"

// User represents a system user.
type User struct {
	id   string
	name string
}

// NewUser creates a new User instance.
func NewUser(id, name string) *User {
	name = strings.Trim(name, " ")
	return &User{id: id, name: name}
}

// ID returns the user ID.
func (u *User) ID() string {
	return u.id
}

// Name returns the user name.
func (u *User) Name() string {
	return u.name
}
