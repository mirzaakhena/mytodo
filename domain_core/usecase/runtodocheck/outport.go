package runtodocheck

import "demo3/domain_core/model/repository"

// Outport of usecase
type Outport interface {
	repository.FindOneTodoRepo
	repository.SaveTodoRepo
}
