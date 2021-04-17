package models

type Input struct {
	Link   string      `json:"link"`
}

type Course struct {
	CourseType   string      `json:"courseType"`
	Id			 string 	 `json:"id" gorm:"column:cid"`
	Slug 		 string 	 `json:"slug"`
	Name 		 string 	 `json:"Name"`
}