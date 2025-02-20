package domain

type CreateTeacherAttendanceRequest struct {
	UserID            string  `json:"userId"`
	Date              string  `json:"date"`
	ClockIn           string  `json:"clockIn"`
	ClockOut          string  `json:"clockOut"`
	WorkHour          string  `json:"workHour"`
	OvertimeRegular   string  `json:"overtimeRegular"`
	OvertimeMorning   string  `json:"overtimeMorning"`
	OvertimeEvening   string  `json:"overtimeEvening"`
	IsOvertimeMorning bool    `json:"isOvertimeMorning"`
	IsOvertimeEvening bool    `json:"isOvertimeEvening"`
	Latitude          float64 `json:"latitude"`
	Longitude         float64 `json:"longitude"`
}
