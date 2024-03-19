package repository

import (
	"database/sql"
	"enigma-lms/model"
	"time"
)

type EnrollmentRepository interface {
	Create(payload model.Enrollment) (model.Enrollment, error)
}

type enrollmentRepository struct {
	db *sql.DB
}

func (e *enrollmentRepository) Create(payload model.Enrollment) (model.Enrollment, error) {
	tx, err := e.db.Begin()
	if err != nil {
		return model.Enrollment{}, err
	}

	var enrollment model.Enrollment
	var enrollmentDetails []model.EnrollmentDetail

	err = tx.QueryRow(`
	INSERT INTO enrollments (course_id, status, updated_at) VALUES ($1, $2, $3)
	RETURNING id, status, created_at, updated_at`,
		payload.Course.Id,
		"active",
		time.Now(),
	).Scan(
		&enrollment.Id,
		&enrollment.Status,
		&enrollment.CreatedAt,
		&enrollment.UpdatedAt,
	)
	if err != nil {
		return model.Enrollment{}, tx.Rollback()
	}

	for _, v := range payload.EnrollmentDetails {
		var enrollmentDetail model.EnrollmentDetail
		err = tx.QueryRow(`INSER INTO enrollment_detais (enrollment_id, user_id, updatedat) VALUES ($1, $2, $3)
		RETURNING id, enrollment_id, created_at, updated_at`,
			enrollment.Id,
			v.User.Id,
			time.Now(),
		).Scan(
			&enrollmentDetail.Id,
			&enrollmentDetail.EnrollmentId,
			&enrollmentDetail.CreatedAt,
			&enrollmentDetail.UpdatedAt,
		)
		if err != nil {
			return model.Enrollment{}, tx.Rollback()
		}
		enrollmentDetail.User = v.User
		enrollmentDetails = append(enrollmentDetails, enrollmentDetail)
	}
	if err := tx.Commit(); err != nil {
		return model.Enrollment{}, err
	}

	enrollment.Course = payload.Course
	enrollment.EnrollmentDetails = enrollmentDetails
	return enrollment, nil
}

func NewEnrollmentRepository(db *sql.DB) EnrollmentRepository {
	return &enrollmentRepository{db: db}
}
