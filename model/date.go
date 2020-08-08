package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

type Date time.Time

const dateLayout = "2006-01-02"

// UnmarshalJSON Parses the json string in the date format
func (date *Date) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	nt, err := time.Parse(dateLayout, s)
	*date = Date(nt)
	return
}

// MarshalJSON writes a quoted string in the date format
func (date Date) MarshalJSON() ([]byte, error) {
	return []byte(date.String()), nil
}

// String returns the time in the date format
func (date *Date) String() string {
	t := time.Time(*date)
	return fmt.Sprintf("%q", t.Format(dateLayout))
}

func (date Date) Value() (driver.Value, error) {
	return json.Marshal(date)
}

func (date *Date) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, date)
}
