// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: v1/person.proto

package person

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on Person with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Person) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Person with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in PersonMultiError, or nil if none found.
func (m *Person) ValidateAll() error {
	return m.validate(true)
}

func (m *Person) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Age

	if m.Password != nil {
		// no validation rules for Password
	}

	if len(errors) > 0 {
		return PersonMultiError(errors)
	}

	return nil
}

// PersonMultiError is an error wrapping multiple validation errors returned by
// Person.ValidateAll() if the designated constraints aren't met.
type PersonMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PersonMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PersonMultiError) AllErrors() []error { return m }

// PersonValidationError is the validation error returned by Person.Validate if
// the designated constraints aren't met.
type PersonValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PersonValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PersonValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PersonValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PersonValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PersonValidationError) ErrorName() string { return "PersonValidationError" }

// Error satisfies the builtin error interface
func (e PersonValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPerson.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PersonValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PersonValidationError{}

// Validate checks the field values on GetPersonRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetPersonRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetPersonRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetPersonRequestMultiError, or nil if none found.
func (m *GetPersonRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetPersonRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for PersonName

	if len(errors) > 0 {
		return GetPersonRequestMultiError(errors)
	}

	return nil
}

// GetPersonRequestMultiError is an error wrapping multiple validation errors
// returned by GetPersonRequest.ValidateAll() if the designated constraints
// aren't met.
type GetPersonRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetPersonRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetPersonRequestMultiError) AllErrors() []error { return m }

// GetPersonRequestValidationError is the validation error returned by
// GetPersonRequest.Validate if the designated constraints aren't met.
type GetPersonRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPersonRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPersonRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPersonRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPersonRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPersonRequestValidationError) ErrorName() string { return "GetPersonRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetPersonRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPersonRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPersonRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPersonRequestValidationError{}

// Validate checks the field values on GetPersonResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetPersonResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetPersonResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetPersonResponseMultiError, or nil if none found.
func (m *GetPersonResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetPersonResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetPerson()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetPersonResponseValidationError{
					field:  "Person",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetPersonResponseValidationError{
					field:  "Person",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPerson()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetPersonResponseValidationError{
				field:  "Person",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Message

	// no validation rules for Retcode

	if len(errors) > 0 {
		return GetPersonResponseMultiError(errors)
	}

	return nil
}

// GetPersonResponseMultiError is an error wrapping multiple validation errors
// returned by GetPersonResponse.ValidateAll() if the designated constraints
// aren't met.
type GetPersonResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetPersonResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetPersonResponseMultiError) AllErrors() []error { return m }

// GetPersonResponseValidationError is the validation error returned by
// GetPersonResponse.Validate if the designated constraints aren't met.
type GetPersonResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPersonResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPersonResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPersonResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPersonResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPersonResponseValidationError) ErrorName() string {
	return "GetPersonResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetPersonResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPersonResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPersonResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPersonResponseValidationError{}
