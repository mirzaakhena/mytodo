package entity

import (
	"demo3/domain_core/model/errorenum"
	"time"
)

type Todo struct {
	ID      string    `bson:"_id" json:"id"`
	Message string    `bson:"message" json:"message"`
	Created time.Time `bson:"created" json:"created"`
	Checked bool      `bson:"checked" json:"checked"`
}

type TodoRequest struct {
	Message      string `json:"message"`
	Now          time.Time
	RandomString string
}

func NewTodo(req TodoRequest) (*Todo, error) {

	if req.Message == "" {
		return nil, errorenum.MessageMustNotEmpty
	}

	var obj Todo
	obj.Message = req.Message
	obj.Created = req.Now
	obj.Checked = false
	obj.ID = req.RandomString

	return &obj, nil
}

func (r *Todo) Check() error {

	if r.Checked {
		return errorenum.TodoAlreadyChecked
	}

	r.Checked = true

	return nil
}
