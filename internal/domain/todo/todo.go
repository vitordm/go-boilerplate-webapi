package todo

import "time"

type Todo struct {
	Id   uint64
	UUID string

	Name  string
	Items *[]TodoItem

	IsActive  bool
	IsDeleted bool

	CreatedAt *time.Time
	CreatedBy *string
	UpdateAt  *time.Time
	UpdateBy  *string
}

type TodoItem struct {
	Id      uint64
	UUID    string
	Item    string
	StartAt *time.Time
	EndsAt  *time.Time

	IsActive  bool
	IsDeleted bool

	CreatedAt *time.Time
	CreatedBy *string
	UpdateAt  *time.Time
	UpdateBy  *string
}
