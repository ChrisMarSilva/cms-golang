package utils

import (
	"fmt"
	"time"
)

type Date struct {
	time.Time
}

func (t *Date) UnmarshalJSON(b []byte) (err error) {
	date, err := time.Parse(`"2006-01-02"`, string(b))
	if err != nil {
		return fmt.Errorf("error on converting date: %w", err)
	}

	t.Time = date
	return
}

func (t Date) MarshalJSON() ([]byte, error) {
	return []byte(t.Time.Format(`"2006-01-02"`)), nil
}
