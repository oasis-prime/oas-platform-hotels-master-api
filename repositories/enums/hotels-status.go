package enums

import (
	"database/sql/driver"
	"fmt"
)

type HotelStatus int

const (
	Null        HotelStatus = 0
	Draft       HotelStatus = 1
	Active      HotelStatus = 2
	InActive    HotelStatus = 3
	Maintenance HotelStatus = 4
)

//Value validate enum when set to database
func (t HotelStatus) Value() (driver.Value, error) {
	switch t {
	case Draft, Active, InActive, Maintenance: //valid case
		return t, nil
	}
	return nil, fmt.Errorf("invalid payment status type value") //else is invalid
}

//Scan validate enum on read from data base
func (t *HotelStatus) Scan(value interface{}) (err error) {
	var pt HotelStatus
	if value == nil {
		*t = 0
		return fmt.Errorf("value is not empty")
	}
	st, ok := value.(int8)
	if !ok {
		return fmt.Errorf("value is not cast string")
	}
	pt = HotelStatus(st)
	switch pt {
	case Draft, Active, InActive, Maintenance:
		*t = pt
		return nil
	}
	return fmt.Errorf("invalid payment status type value :%s", st)
}
