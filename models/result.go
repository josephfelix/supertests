package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
)

// Result is used by pop to map your results database table to your go code.
type Result struct {
	ID        string    `json:"id" db:"id"`
	UserID    string    `json:"user_id" db:"userid"`
	Guid      string    `json:"guid" db:"guid"`
	Result    string    `json:"result" db:"result"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (r Result) String() string {
	jr, _ := json.Marshal(r)
	return string(jr)
}

// Results is not required by pop and may be deleted
type Results []Result

// String is not required by pop and may be deleted
func (r Results) String() string {
	jr, _ := json.Marshal(r)
	return string(jr)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (r *Result) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (r *Result) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (r *Result) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
