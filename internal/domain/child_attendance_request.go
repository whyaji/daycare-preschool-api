package domain

type CreateChildAttendanceRequest struct {
	ChildID uint   `json:"childId"`
	Date    string `json:"date"`
	Arrival string `json:"arrival"`
}
