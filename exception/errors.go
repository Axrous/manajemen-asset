package exception

import (
	"fmt"
	"net/http"
)

type Http struct {
    StatusCode  int    `json:"statusCode"`
    Description string `json:"description,omitempty"`
    // Metadata    string `json:"metadata,omitempty"`
}

func (e *Http) Error() string {
    return fmt.Sprintf("description: %s,  metadata: %s", e.Description)
}

func NewHttpError(description string, statusCode int) *Http {
    return &Http{
        StatusCode:  statusCode,
        Description: description,
    }
}

func BadRequestErr(description string) error  {
    return NewHttpError(description, http.StatusBadRequest)
}