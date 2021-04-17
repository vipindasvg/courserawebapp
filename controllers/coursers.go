package controllers

import (
	"net/http"
	"strconv"
	"io/ioutil"
	"encoding/json"
	"strings"

	"github.com/labstack/echo"
	"github.com/vipindasvg/courserawebapp/common"
    "github.com/vipindasvg/courserawebapp/models"
	"github.com/vipindasvg/courserawebapp/storage"
)

// request save courses /api/v1/getcourses?limit=2
func SaveCourses(c echo.Context) error {
	In := new(models.Input)
	if err := c.Bind(In); err != nil {
		common.Log.WithField("handler", "save-courses").WithField("issue", "request").Warnln("can not binds the request body into provided type:", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	querysplit := strings.Split(In.Link, "?")
	url := "https://api.coursera.org/api/courses.v1"
	res, err := http.Get(url + "?" + querysplit[1])

    if err != nil {
        common.Log.WithField("handler", "save-courses").WithField("issue", "request").Warnln("can not binds the request body into provided type:", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

	body, err := ioutil.ReadAll(res.Body)
	bodyString := string(body)

    if err != nil {
        common.Log.WithField("handler", "save-courses").WithField("issue", "request").Warnln("can not binds the request body into provided type:", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
	var data []models.Course
	split1 := strings.Split(bodyString, "[")
	split2 := strings.Split(split1[1], "]")
    final := "[" + split2[0] + "]"
	json.Unmarshal([]byte(final), &data)
	
	curs := storage.GetCursor()
	for _, cs := range data {
		_, err := curs.CreateCourses(&cs)
		if err != nil {
			common.Log.WithField("handler", "save-courses").WithField("issue", "cursor").Errorln("can not creates course record in the database:", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}	
	return c.JSON(http.StatusCreated, data)
}

// request get courses /api/v1/getcourses?limit=2
func GetCourses(c echo.Context) error {
	curs := storage.GetCursor()
	limit := c.QueryParam("limit")
	intlimit, err := strconv.Atoi(limit)
	if err != nil {
		common.Log.WithField("handler", "get-courses").WithField("issue", "request").Warnln("Can not getting the limit:", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	cs, err := curs.GetCourses(intlimit)
	if err != nil {
		common.Log.WithField("handler", "get-courses").WithField("issue", "cursor").Errorln("can not get:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, cs)
}