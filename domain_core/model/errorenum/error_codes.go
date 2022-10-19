package errorenum

import "demo3/shared/model/apperror"

const (
	SomethingError      apperror.ErrorType = "ER0000 something error"
	MessageMustNotEmpty apperror.ErrorType = "ER0001 message must not empty"
	TodoAlreadyChecked  apperror.ErrorType = "ER0002 todo already checked"
)
