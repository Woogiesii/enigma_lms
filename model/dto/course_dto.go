package dto

type CourseRequestDto struct {
	Id              string `json:"id"`
	CourseFullName  string `json:"courseFullName"`
	CourseShortName string `json:"courseShortName"`
	Description     string `json:"description"`
	CourseStartDate string `json:"courseStartDate"`
	CourseEndDate   string `json:"courseEndDate"`
	CourseImage     string `json:"courseImage"`
}
