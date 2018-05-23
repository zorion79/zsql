package zsql

import (
	"testing"
	"time"
)

func TestZTime_Value(t *testing.T) {
	ztime := ZTime{
		Time:  time.Now(),
		Valid: true,
	}

	val, err := ztime.Value()
	if val == nil && err == nil {
		t.Errorf("zTime from string %v is: %v, "+
			"wanted: %v", ztime.Time, ztime.Valid, true)
	}

	ztime = ZTime{
		Time:  time.Now(),
		Valid: false,
	}

	val, err = ztime.Value()
	if val != nil && err == nil {
		t.Errorf("zTime from string %v is: %v, "+
			"wanted: %v", ztime.Time, ztime.Valid, false)
	}
}

func TestZTime_Scan(t *testing.T) {
	var zTime ZTime

	//var zTimeNil = struct{}{}

	zTime.Scan(nil)
	if zTime.Valid {
		t.Errorf("zTime from string %v is: %v, "+
			"wanted: %v", zTime.Time, zTime.Valid, true)
	}

	//zTime.Scan(zTimeNil)
	//if zTime.Valid {
	//	t.Errorf("zTimeNil is: %v, wanted: %v",
	//		zTime.Valid, false)
	//}

	zTime.Scan(time.Now())
	if !zTime.Valid {
		t.Errorf("zTime from time is: %v, wanted: %v",
			zTime.Valid, true)
	}

	zTimeBytes := []byte(time.Now().Format("2006-01-02T03:04:05-07:00"))
	zTime.Scan(zTimeBytes)
	if !zTime.Valid {
		t.Errorf("zTime from bytes is: %v, wanted: %v",
			zTime.Valid, true)
	}

	zTime.Scan("")
	if zTime.Valid {
		t.Errorf("zTime from string %v is: %v, wanted: %v", zTime.Time,
			zTime.Valid, false)
	}

	zTime.Scan("2007-12-03T00:00:00-07:00")
	//fmt.Println(zTime.Time.String())
	if !zTime.Valid {
		t.Errorf("zTime from string %v is: %v, "+
			"wanted: %v", zTime.Time, zTime.Valid, true)
	}

	zTime = ZTime{}
	zTime.Scan(int(10))
	if zTime.Valid {
		t.Errorf("zTime from string %v is: %v, "+
			"wanted: %v", zTime.Time, zTime.Valid, true)
	}
}
