package error

import "fmt"

// ErrorID ... アプリケーションエラーの識別子
type ErrorID string

// ApplicationError ... アプリケーション固有のエラーを表す。
type ApplicationError interface {
	Error() string
}

type applicationError struct {
	errorID     ErrorID
	errorString string
	cause       error
}

func (e *applicationError) Error() string {
	if e == nil {
		return ""
	}
	if e.cause == nil {
		return fmt.Sprintf("[ERROR_ID:%s) %s", e.errorID, e.errorString)
	}
	return fmt.Sprintf("[ERROR_ID:%s) %s [CAUSE:%s]", e.errorID, e.errorString, e.cause)
}

// CreateValidationError ... バリデーション失敗を表す。
func CreateValidationError(errorID ErrorID, cause error) ApplicationError {
	return &applicationError{
		errorID:     errorID,
		errorString: "Validation failed",
		cause:       cause,
	}
}
