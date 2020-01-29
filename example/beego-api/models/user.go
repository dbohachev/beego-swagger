package models

import (
	"errors"
	"time"
)

// User defines user struct
type User struct {
	// ID defines the user ID. Should be unique per users set
	ID int64 `json:"id" example:"1" format:"int64"`
	// Name is the name of a user
	Name string `json:"name" example:"Dmytro" minLength:"1" maxLength:"30"`
	// Country is the country where user lives
	Country string `json:"country" example:"Ukraine"`
}

var (
	// Users holds map of our users (ID->User). In real life, access should be protected!
	Users map[int64]*User
)

func init() {
	Init()
}

// Init provides initial data for our application
func Init() {
	Users = make(map[int64]*User)
	Users[1] = &User{1, "Dmytro", "Ukraine"}
	Users[2] = &User{2, "John", "USA"}
	Users[3] = &User{3, "Peter", "United Kingdom"}
}

// AddOne adds single user to the map
func AddOne(user User) int64 {
	// This will (for test needs only) provide "unique" values. In real life, this
	// will not be guaranteed unique
	user.ID = time.Now().Unix()
	Users[user.ID] = &user
	return user.ID
}

// GetOne gets single user by ID
func GetOne(id int64) (object *User, err error) {
	if v, ok := Users[id]; ok {
		return v, nil
	}
	return nil, errors.New("provided ID does not exist")
}

// GetAll gets all users as map
func GetAll() map[int64]*User {
	return Users
}

// Update updates single user in map (search is performed by ID)
func Update(user User) (err error) {
	if _, ok := Users[user.ID]; ok {
		Users[user.ID] = &user
		return nil
	}

	return errors.New("provided id does not exist")
}

// Delete deletes user form map, by ID
func Delete(id int64) (*User, error) {
	// Not safe, access to globally shared object not protected (skipping the fact global objects
	// often are anti-patterns
	user, ok := Users[id]
	if ok {
		delete(Users, id)
		return user, nil
	}
	return nil, errors.New("user with specified id was not found")
}
