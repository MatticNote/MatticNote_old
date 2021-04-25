package apiV1

import (
	"database/sql"
	"encoding/json"
)

type (
	NullString struct {
		sql.NullString
	}
	NullTime struct {
		sql.NullTime
	}
	MNAPIV1User struct {
		Uuid           string     `json:"uuid"`
		Username       string     `json:"username"`
		Host           NullString `json:"host"`
		DisplayName    NullString `json:"display_name"`
		Summary        NullString `json:"summary"`
		IsBot          bool       `json:"is_bot"`
		AcceptManually bool       `json:"accept_manually"`
		CreatedAt      NullTime   `json:"created_at"`
		UpdatedAt      NullTime   `json:"updated_at"`
	}
)

// NullString

func (ns *NullString) UnmarshalJSON(value []byte) error {
	err := json.Unmarshal(value, &ns.String)
	ns.Valid = err == nil
	return err
}

func (ns NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(ns.String)
}

// NullTime

func (ns *NullTime) UnmarshalJSON(value []byte) error {
	err := json.Unmarshal(value, &ns.Time)
	ns.Valid = err == nil
	return err
}

func (ns NullTime) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(ns.Time)
}
