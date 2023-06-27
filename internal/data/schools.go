// Filename: internal/data/schools.go
package data

import "time"

type School struct {
	ID int64 `json:"id"`
	// If we do not want the user to see some field, we can set the json to "-", For Eg: `json:"-"`
	CreatedAt time.Time `json:"-"`
	Name      string    `json:"name"`
	Level     string    `json:"level"`
	Contact   string    `json:"contact"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email,omitempty"`
	Website   string    `json:"website,omitempty"`
	Address   string    `json:"address"`
	Mode      []string  `json:"mode"`
	Version   int32     `json:"version"`
}
