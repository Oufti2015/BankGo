package model

import (
	"strings"
	"time"
)

type CivilTime time.Time

type Period struct {
	StartDate CivilTime `json:"startDate"`
	EndDate   CivilTime `json:"endDate"`
}

func (c *CivilTime) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`) //get rid of "
	if value == "" || value == "null" {
		return nil
	}

	t, err := time.Parse("2006-01-02", value) //parse time
	if err != nil {
		return err
	}
	*c = CivilTime(t) //set result using the pointer
	return nil
}

func (c CivilTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(c).Format("2006-01-02") + `"`), nil
}

func (p Period) StartTime() time.Time {
	return time.Time(p.StartDate)
}

func (p Period) EndTime() time.Time {
	return time.Time(p.EndDate)
}
