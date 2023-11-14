package sandsmad_parser

import (
	"errors"
	"time"
)

type DishDay string

const (
	Monday    DishDay = "Mandag"
	Tuesday   DishDay = "Tirsdag"
	Wednesday DishDay = "Onsdag"
	Thursday  DishDay = "Torsdag"
	Friday    DishDay = "Fredag"
)

var ErrWeekdayNotDishDay = errors.New("weekday is not a dish day")

// DishOfTheDay is a struct containing a dish of the day. A slice of lines describes the dish and the short description provides a summary of the main dish
type DishOfTheDay struct {
	NameOfDay        DishDay  `json:"name_of_day"`
	ShortDescription string   `json:"short_description"`
	Lines            []string `json:"lines"`
}

var WeekDayMap = map[time.Weekday]DishDay{
	time.Monday:    Monday,
	time.Tuesday:   Tuesday,
	time.Wednesday: Wednesday,
	time.Thursday:  Thursday,
	time.Friday:    Friday,
}

func (dish DishOfTheDay) IsDishOnTheMenuAt(t time.Time) (error, bool) {
	dishDay, ok := WeekDayMap[t.Weekday()]
	if !ok {
		return ErrWeekdayNotDishDay, false
	}

	return nil, dishDay == dish.NameOfDay
}

func (dish DishOfTheDay) IsDishOnTheMenuToday() (error, bool) {
	return dish.IsDishOnTheMenuAt(time.Now())
}
