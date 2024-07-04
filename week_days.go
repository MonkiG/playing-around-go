package main

type WeekDay int

const (
	Sunday WeekDay = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

var weekDayName = map[WeekDay]string{
	Sunday:    "sunday",
	Monday:    "monday",
	Tuesday:   "tuesday",
	Wednesday: "wednesday",
	Thursday:  "thursday",
	Friday:    "friday",
	Saturday:  "saturday",
}

func (w WeekDay) String() string {
	return weekDayName[w]
}
