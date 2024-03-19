package repository

import (
	"database/sql"
	"enigma-lms/model"
	"time"
)

type CourseRepository interface {
	Get(id string) (model.Course, error)
	Create(payload model.Course) (model.Course, error)
}

type courseRepository struct {
	db *sql.DB
}

func (cs *courseRepository) Get(id string) (model.Course, error) {
	var course model.Course
	err := cs.db.QueryRow(`SELECT id, course_full_name, course_short_name, description, course_start_date, course_end_date, course_image, created_at, updated_at FROM courses WHERE id = $1`, id).Scan(
		&course.Id,
		&course.CourseFullName,
		&course.CourseShortName,
		&course.Description,
		&course.CourseStartDate,
		&course.CourseEndDate,
		&course.CourseImage,
		&course.CreatedAt,
		&course.UpdatedAt,
	)

	if err != nil {
		return model.Course{}, err
	}

	return course, nil
}

func (cse *courseRepository) Create(payload model.Course) (model.Course, error) {
	var course model.Course
	err := cse.db.QueryRow(`INSERT INTO courses (course_full_name, course_short_name, description, course_start_date, course_end_date, course_image, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, course_full_name, course_short_name, description, course_start_date, course_end_date, course_image, created_at, updated_at`,
		payload.CourseFullName,
		payload.CourseShortName,
		payload.Description,
		payload.CourseStartDate,
		payload.CourseEndDate,
		payload.CourseImage,
		time.Now(),
	).Scan(
		&course.Id,
		&course.CourseFullName,
		&course.CourseShortName,
		&course.Description,
		&course.CourseStartDate,
		&course.CourseEndDate,
		&course.CourseImage,
		&course.CreatedAt,
		&course.UpdatedAt,
	)

	if err != nil {
		return model.Course{}, err
	}

	return course, nil
}

func NewCourseRepository(db *sql.DB) CourseRepository {
	return &courseRepository{db: db}
}
