package controller

import (
	"fmt"
	u "member-db-api/util"
	"net/http"
	"strconv"
	"strings"

	m "member-db-api/model"

	"github.com/gin-gonic/gin"
)

// GetAllMember Controller will return all members
var GetAllMember = func(w http.ResponseWriter, r *http.Request) {
	rawCookieToken := r.Context().Value("ckiftk")
	cookieTOken, _ := rawCookieToken.(map[string]interface{})
	fmt.Println("FROM HERE I GET", cookieTOken["message"])
	rawHeaderToken := r.Context().Value("authHeader")
	fmt.Println("FROM HERE I GET", rawHeaderToken)
	data := memberContext().GetMembers()
	if data != nil {
		message := u.BaseSuccessResponse
		message["data"] = data
		u.Respond(w, message)
	} else {
		w.WriteHeader(http.StatusNotFound)
		u.Respond(w, u.BaseFailedResponse)
	}
}

var GetMemberById = func(c *gin.Context) {
	id := c.Param("id")
	rawCookieToken := c.Request.Context().Value("ckiftk")
	cookieTOken, _ := rawCookieToken.(map[string]interface{})
	fmt.Println("FROM HERE I GET", cookieTOken["message"])
	rawHeaderToken := c.Request.Context().Value("authHeader")
	fmt.Println("FROM HERE I GET", rawHeaderToken)
	if intVal, err := strconv.Atoi(id); err == nil {
		data := memberContext().GetMemberByID(intVal, "")
		if data != nil {
			message := u.BaseSuccessResponse
			message["data"] = data
			c.JSON(http.StatusFound, message)
		} else {
			c.AbortWithStatusJSON(http.StatusNotFound, u.BaseFailedResponse)
		}
	} else {
		c.AbortWithStatusJSON(http.StatusBadRequest, u.BaseFailedResponse)
	}
}

var GetMemberByIdType = func(c *gin.Context) {
	id := c.Param("id")
	t := strings.TrimPrefix(c.Param("type"), "/")
	rawCookieToken := c.Request.Context().Value("ckiftk")
	cookieTOken, _ := rawCookieToken.(map[string]interface{})
	fmt.Println("FROM HERE I GET", cookieTOken["message"])
	rawHeaderToken := c.Request.Context().Value("authHeader")
	fmt.Println("FROM HERE I GET", rawHeaderToken)
	fmt.Println("PARAM", id, t)
	if intVal, err := strconv.Atoi(id); err == nil {
		data := memberContext().GetMemberByID(intVal, t)
		if data != nil {
			message := u.BaseSuccessResponse
			message["data"] = data
			c.JSON(http.StatusOK, message)
		} else {
			c.AbortWithStatusJSON(http.StatusNotFound, u.BaseFailedResponse)
		}
	} else {
		c.AbortWithStatusJSON(http.StatusBadRequest, u.BaseFailedResponse)
	}
}

var PostMemberByIdType = func(c *gin.Context) {
	var m m.MemberInput
	c.BindJSON(&m)
	rawCookieToken := c.Request.Context().Value("ckiftk")
	cookieTOken, _ := rawCookieToken.(map[string]interface{})
	fmt.Println("FROM HERE I GET", cookieTOken["message"])
	rawHeaderToken := c.Request.Context().Value("authHeader")
	fmt.Println("FROM HERE I GET", rawHeaderToken)
	fmt.Println("PARAM", m.Id, m.T)
	if m.Id != 0 {
		data := memberContext().GetMemberByID(m.Id, m.T)
		fmt.Println(data)
		if data != nil {
			message := u.BaseSuccessResponse
			message["data"] = *data
			c.JSON(http.StatusOK, message)
		} else {
			c.AbortWithStatusJSON(http.StatusNotFound, u.BaseFailedResponse)
		}
	} else {
		c.AbortWithStatusJSON(http.StatusBadRequest, u.BaseFailedResponse)
	}
}
