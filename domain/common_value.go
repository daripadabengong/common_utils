package domain

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/google/uuid"
)

type EntityID struct {
	value uuid.UUID
}

type PhoneNumber struct {
	value string
}

type EmailAddress struct {
	value string
}

type NullableString struct {
	value string
}

type RequiredString struct {
	value string
}

func (v EntityID) GetValue() uuid.UUID    { return v.value }
func (v PhoneNumber) GetValue() string    { return v.value }
func (v EmailAddress) GetValue() string   { return v.value }
func (v NullableString) GetValue() string { return v.value }
func (v RequiredString) GetValue() string { return v.value }

func NewEntityID(value uuid.UUID) (EntityID, error) {
	if value == uuid.Nil {
		value = uuid.New()
	}
	return EntityID{value: value}, nil
}

func NewPhoneNumber(value string) (PhoneNumber, error) {
	if value == "" {
		return PhoneNumber{}, errors.New("phone number can't be empty")
	}
	re := regexp.MustCompile(`^[0-9]+$`)
	if !re.MatchString(value) {
		return PhoneNumber{}, errors.New("username can only contains numbers")
	}
	return PhoneNumber{value: value}, nil
}

func NewEmailAddress(value string) (EmailAddress, error) {
	if value == "" {
		return EmailAddress{}, errors.New("email address can't be empty")
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(value) {
		return EmailAddress{}, errors.New("invalid email address")
	}
	return EmailAddress{value: value}, nil
}

func NewNullableString(value string) (NullableString, error) {
	return NullableString{value: value}, nil
}

func NewRequiredString(fieldName, value string) (RequiredString, error) {
	if value == "" {
		return RequiredString{}, fmt.Errorf("%s can't be empty", fieldName)
	}
	return RequiredString{value: value}, nil
}
