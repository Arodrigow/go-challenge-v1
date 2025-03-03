package hashmap

import (
	"errors"

	"github.com/gofrs/uuid"
)

type Id uuid.UUID

type User struct {
	FirstName string
	LastName  string
	Biography string
}

type Application struct {
	Data map[Id]User
}

type response struct {
	Id        Id
	FirstName string
	LastName  string
	FullName  string
	Biography string
}

func (a Application) FindAll() map[Id]User {
	if a.Data == nil {
		return make(map[Id]User)
	}

	return a.Data
}

func (a Application) FindById(id Id) (*User, error) {

	user, ok := a.Data[id]
	if !ok {
		return nil, errors.New("usuário não encontrado")
	}

	return &user, nil
}

func (a *Application) Insert(u User) response {
	newId, _ := uuid.NewV4()

	a.Data[Id(newId)] = u

	return userToResponse(Id(newId), u)
}

func (a *Application) Update(id Id, u User) (response, error) {
	user, err := a.FindById(id)
	if err != nil {
		return response{}, err
	}

	user.FirstName = u.FirstName
	user.LastName = u.FirstName
	user.Biography = u.Biography
	a.Data[id] = *user

	return userToResponse(id, *user), nil
}

func (a *Application) Delete(id Id) (response, error) {
	user, err := a.FindById(id)
	if err != nil {
		return response{}, err
	}

	delete(a.Data, id)

	return userToResponse(id, *user), nil
}

func userToResponse(id Id, u User) response {
	var result = response{
		Id:        Id(id),
		FirstName: u.FirstName,
		LastName:  u.FirstName,
		FullName:  u.FirstName + " " + u.LastName,
		Biography: u.Biography,
	}
	return result
}
