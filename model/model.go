package model

//register//

type CreateRegisterRequest struct {
	Firstname      string `bson:"firstname" json:"firstname"`
	Lastname       string `bson:"lastname" json:"lastname"`
	Birthdate      string `bson:"birthdate" json:"birthdate"`
	Address        string `bson:"address" json:"address"`
	Email          string `bson:"email" json:"email" validate:"required"`
	Password       string `bson:"-" json:"password" validate:"required"`
	Phone          string `bson:"phone" json:"phone"`
	Education      string `bson:"education" json:"education"`
	WorkExperience string `bson:"workexperience" json:"workexperience"`
	FileID         string `bson:"fileid" json:"fileid"`
}

// type UpdateFileID struct {
// 	Email  string `bson:"email" json:"email" validate:"required"`
// 	FileID string `bson:"fileid" json:"fileid"`
// }

type RegisterResponseModel struct {
	Registers []map[string]interface{} `json:"registers"`
	Message   string                   `json:"message"`
}

//search job//
type JobResponseModel struct {
	Jobs    []map[string]interface{} `json:"jobs"`
	Message string                   `json:"message"`
}

//send file//
// type CreateFileRequest struct {
// 	Email    string                `bson:"email" json:"email" validate:"required"`
// 	Filename *multipart.FileHeader `form:"filename" binding:"required"`
// }

//get jobid//
type FileIDResponseModel struct {
	Fileids []map[string]interface{} `json:"Fileids"`
	Message string                   `json:"message"`
}
