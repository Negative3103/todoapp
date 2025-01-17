package service

import (
	"github.com/Negative3103/todoapp"
	"github.com/Negative3103/todoapp/pkg/repository"
)

type Authorization interface {
	CreateUser(user todoapp.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list todoapp.TodoList) (int, error)
	GetAll(userId int) ([]todoapp.TodoList, error)
	GetById(userId, listId int) (todoapp.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input todoapp.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int, item todoapp.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todoapp.TodoItem, error)
	GetById(userId, itemId int) (todoapp.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input todoapp.UpdateItemInput) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
