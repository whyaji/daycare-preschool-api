package domain

import (
	"time"

	"gorm.io/gorm"
)

// Registered user Email list model
type RegisteredEmail struct {
	ID           uint       `gorm:"primaryKey"`
	Email        string     `gorm:"size:255;unique;not null"`
	Roles        []Role     `gorm:"many2many:registered_email_roles;"`
	RegisteredAt *time.Time `gorm:"default:null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

// User model (for both Parents and Bunda Workers) with multi role
type User struct {
	ID                uint    `gorm:"primaryKey"`
	Name              string  `gorm:"size:255;not null"`
	Email             string  `gorm:"size:255;unique;not null"`
	Password          string  `gorm:"size:255;not null"`
	Gender            string  `gorm:"type:enum('male','female');not null"`
	Phone             string  `gorm:"size:255;not null"`
	Address           string  `gorm:"type:text;not null"`
	JobTitle          string  `gorm:"size:255;default:null"`
	JobPlace          string  `gorm:"size:255;default:null"`
	Roles             []Role  `gorm:"many2many:user_roles;"`
	ChildrenAsParent  []Child `gorm:"many2many:child_parents;"`
	ChildrenAsTeacher []Child `gorm:"many2many:child_teachers;"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}

// Role model (for User)
type Role struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:255;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// Attendance for Bunda (Workers)
type TeacherAttendance struct {
	ID              uint      `gorm:"primaryKey"`
	UserID          uint      `gorm:"not null"`
	Date            time.Time `gorm:"not null"`
	ClockIn         *time.Time
	ClockOut        *time.Time
	WorkHour        float32 `gorm:"default:0"`
	OvertimeRegular int     `gorm:"default:0"`
	OvertimeMorning int     `gorm:"default:0"`
	OvertimeEvening int     `gorm:"default:0"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

// Leave Requests (For Bunda Workers)
type LeaveRequest struct {
	ID         uint      `gorm:"primaryKey"`
	UserID     uint      `gorm:"not null"`
	LeaveDate  time.Time `gorm:"not null"`
	Reason     string    `gorm:"type:text;not null"`
	Status     string    `gorm:"type:enum('pending','approved','rejected');default:'pending'"`
	ApprovedBy *uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

// Child model (supports multiple parents & teachers)
type Child struct {
	ID               uint      `gorm:"primaryKey"`
	Name             string    `gorm:"size:255;not null"`
	Nickname         string    `gorm:"size:255;not null"`
	BirthPlace       string    `gorm:"size:255;not null"`
	BirthDate        time.Time `gorm:"not null"`
	Gender           string    `gorm:"type:enum('male','female');not null"`
	AlergyInfo       string    `gorm:"type:text"`
	Notes            string    `gorm:"type:text"`
	NumberOfSiblings int       `gorm:"default:0"`
	LivingWith       string    `gorm:"size:255;not null"`
	RegisteredDate   time.Time `gorm:"not null"`
	Parents          []User    `gorm:"many2many:child_parents;"`
	Teachers         []User    `gorm:"many2many:child_teachers;"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}

// Attendance for Child
type ChildAttendance struct {
	ID              uint      `gorm:"primaryKey"`
	ChildID         uint      `gorm:"not null"`
	Date            time.Time `gorm:"not null"`
	Arrival         time.Time `gorm:"not null"`
	Departure       *time.Time
	OvertimeMorning int `gorm:"default:0"`
	OvertimeEvening int `gorm:"default:0"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

// Child Diary (Daily Report for Each Child)
type ChildDiary struct {
	ID                uint      `gorm:"primaryKey"`
	ChildID           uint      `gorm:"not null"`
	Date              time.Time `gorm:"not null"`
	DeliveredBy       string    `gorm:"size:255;not null"`
	HealthCondition   string    `gorm:"type:text"`
	ActivityCondition string    `gorm:"type:text"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}

// Child Meals (Linked to Child Diary)
type ChildMeal struct {
	ID        uint      `gorm:"primaryKey"`
	DiaryID   uint      `gorm:"not null"`
	MealTime  time.Time `gorm:"not null"`
	MealName  string    `gorm:"size:255;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// Child Sleep (Linked to Child Diary)
type ChildSleep struct {
	ID         uint      `gorm:"primaryKey"`
	DiaryID    uint      `gorm:"not null"`
	SleepStart time.Time `gorm:"not null"`
	SleepEnd   time.Time `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

// Child Toilet (Linked to Child Diary)
type ChildToilet struct {
	ID        uint `gorm:"primaryKey"`
	DiaryID   uint `gorm:"not null"`
	PeeCount  int  `gorm:"default:0"`
	PoopCount int  `gorm:"default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// Pre-School Condition (Filled by Parents Before School)
type ChildCondition struct {
	ID             uint      `gorm:"primaryKey"`
	ChildID        uint      `gorm:"not null"`
	Date           time.Time `gorm:"not null"`
	ConditionNotes string    `gorm:"type:text"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

type WorkLocation struct {
	ID        uint    `gorm:"primaryKey"`
	Name      string  `gorm:"size:255;not null"`
	Address   string  `gorm:"type:text;not null"`
	Latitude  float64 `gorm:"not null"`
	Longitude float64 `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
