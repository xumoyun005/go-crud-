package todo

import "fmt"

type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required" example:"Buy groceries"`
	Description string `json:"description" db:"description" example:"Milk, Bread, Eggs"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required" example:"Item title"`
	Description string `json:"description" db:"description" example:"Item description"`
	Done        bool   `json:"done" db:"done" example:"false"`
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}

type UpdateListInput struct {
	Title       *string `json:"title" example:"list update"`
	Description *string `json:"description" example:"update description"`
}

func (i UpdateListInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return fmt.Errorf("one of title or description is required")
	}
	return nil
}

type UpdateItemInput struct {
	Title       *string `json:"title" example:"item update"`
	Description *string `json:"description" example:"update description"`
	Done        *bool   `json:"done" example:"false"`
}

func (i UpdateItemInput) Validate() error {
	if i.Title == nil && i.Description == nil && i.Done == nil {
		return fmt.Errorf("one of title or description is required")
	}
	return nil
}
