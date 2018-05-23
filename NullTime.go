//Package zsql вот такой вот пакет :)
package zsql

import (
	"database/sql/driver"
	"fmt"
	"time"
)

//ZTime - структура для получения данных из sql
type ZTime struct {
	Time  time.Time
	Valid bool
}

//Scan - Стандартный метод для sql
func (nt *ZTime) Scan(value interface{}) (err error) {
	if value == nil {
		nt.Valid = false
		return
	}

	switch v := value.(type) {
	case time.Time:
		nt.Time, nt.Valid = v, true
		return
	case []byte:
		nt.Time, err = parseDateTime(string(v), time.UTC)
		nt.Valid = err == nil
		return
	case string:
		if v == "" {
			nt.Valid = false
			return
		}
		nt.Time, err = parseDateTime(v, time.UTC)
		nt.Valid = err == nil
		return
	}

	nt.Valid = false
	return fmt.Errorf("could not convert %T to time.ZTime", value)
}

// Value возвращает значение, если Valid true
func (nt ZTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}

func parseDateTime(str string, _ *time.Location) (t time.Time, err error) {
	t, err = time.Parse(
		"2006-01-02T03:04:05-07:00", str)
	return
}
