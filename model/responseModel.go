package model

type JobsResponseModel struct {
	Position  []map[string]interface{} `json:"position"`
	Jobsource string                   `json:"jobsource"`
	Company   string                   `json:"company"`
	Salary    string                   `json:"salary"`
	Location  string                   `json:"location"`
	Urllink   string                   `json:"urllink"`
	Email     string                   `json:"email"`
	Status    string                   `json:"status"`
	Message   string                   `json:"message"`
}

//search job//
type JobResponseModelscape struct {
	Jobs    []map[string]interface{} `json:"jobs"`
	Message string                   `json:"message"`
}

//แสดงข้อมูลการสมัครงานของแต่ละ User//
type UserResponseJobModel struct {
	Userjobs []map[string]interface{} `json:"email"`
	Message  string                   `json:"message"`
}

type CreateJobRequest struct {
	Position string `bson:"position" json:"position"`
	Company  string `bson:"company" json:"company"`
	Urllink  string `bson:"urllink" json:"urllink"`
	Email    string `bson:"email" json:"email"`
	Status   string `bson:"status" json:"status"`
}
