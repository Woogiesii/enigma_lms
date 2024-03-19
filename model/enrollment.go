package model

import "time"

type Enrollment struct {
	Id                string             `json:"id"`
	Course            Course             `json:"course"`
	EnrollmentDetails []EnrollmentDetail `json:"enrollmentDetails"`
	Status            string             `json:"status"`
	CreatedAt         time.Time          `json:"createdAt"`
	UpdatedAt         time.Time          `json:"updatedAt"`
}

type EnrollmentDetail struct {
	Id           string    `json:"id"`
	EnrollmentId string    `json:"enrolmentId"`
	User         User      `json:"user"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
