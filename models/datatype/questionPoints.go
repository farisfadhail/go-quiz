package datatype

import (
	"database/sql/driver"
	"errors"
	"strings"
)

type Points []string

func (p *Points) Scan(src any) error {
	bytes, ok := src.([]byte)
	if !ok {
		return errors.New("src value can't cast to []byte")
	}

	*p = strings.Split(string(bytes), ",")

	return nil
}

func (p Points) Value() (driver.Value, error) {
	if len(p) == 0 {
		return nil, nil
	}

	return strings.Join(p, ","), nil
}
