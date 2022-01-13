package service

import "ToDo-app/pkg/repository"

type Authorization interface {

}

type ToDoList interface {

}

type ToDoItem interface {

}

type Service struct {
	Authorization
	ToDoList
	ToDoItem
}

func NewService(repository *repository.Repository) *Service {
	return &Service{}
}
