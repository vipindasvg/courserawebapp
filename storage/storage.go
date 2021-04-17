package storage

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/vipindasvg/courserawebapp/common"
	"github.com/vipindasvg/courserawebapp/models"
)

const errnorec = "record not found"

type cursor struct {
	Db *gorm.DB
}

// cursor is used for interaction with the database
func GetCursor() *cursor {
	c := new(cursor)
	c.Db = common.Db
	return c
}

func (c *cursor) CreateCourses(cs *models.Course) (*models.Course, error) {
	result := c.Db.Create(cs)
	if err := result.Error; err != nil {
		return nil, err
	}
	record := result.Value.(*models.Course)
	return record, nil
}

func (c *cursor) GetCourses(limit int) ([]models.Course, error) {
	var cs []models.Course
	t := "select * from courses ORDER BY id DESC LIMIT '%d'"
	queryString := fmt.Sprintf(t, limit)
	if err := c.Db.Raw(queryString).Scan(&cs).Error; err != nil {
			return nil, err
	}
	return cs, nil
}