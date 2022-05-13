// Package error defines the basic and common error response in srs-api
// The detail describes in the following page
// https://17media.atlassian.net/wiki/spaces/TEAMZ/pages/662176943/Error
package srserr

import "net/http"

// Error wraps lower level errors with status, code, and message.
type Error interface {
	error
	Status() int
	Code() string
	Message() string
	OrigErr() error
	WithMessage(string) Error
	WithError(error) Error
}

func New(status int, code, message string, OrigErr error) Error {
	return newBaseError(status, code, message, OrigErr)
}

const (
	CodeInvalidRequest      = "invalid_request"
	CodeInsufficientRequest = "insufficient_request"
	CodeInvalidGrant        = "invalid_grant"
	CodeInsufficientPerms   = "insufficient_perms"
	CodeResourceNotFound    = "reaource_not_found"
	CodeConflictRequest     = "conflict_request"
	CodeResourceExhausted   = "resource_exhausted"
	CodeConnectionRefused   = "connection_refused"
	CodeUpstreamRefused     = "upstream_refused"
)

var (
	ErrInvalidRequest      = New(http.StatusBadRequest, CodeInvalidRequest, CodeInvalidRequest, nil)
	ErrInsufficientRequest = New(http.StatusBadRequest, CodeInsufficientRequest, CodeInsufficientRequest, nil)
	ErrInvalidGrant        = New(http.StatusUnauthorized, CodeInvalidGrant, CodeInvalidGrant, nil)
	ErrInsufficientPerms   = New(http.StatusForbidden, CodeInsufficientPerms, CodeInsufficientPerms, nil)
	ErrResourceNotFound    = New(http.StatusNotFound, CodeResourceNotFound, CodeResourceNotFound, nil)
	ErrConflictRequest     = New(http.StatusConflict, CodeConflictRequest, CodeConflictRequest, nil)
)
