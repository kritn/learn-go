package controllers

// มีหน้าที่ในการรับส่ง Context จาก HTTP Request หรือพูดง่ายๆก็คือ รับส่งข้อมูลหรือบริบทต่างๆที่ถูกยิงมากจาก API ที่ client ทำการยิงมา

import (
	"go_cleanarc/modules/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUseCase entities.UserUseCase
}

func NewUserHandler(usecase entities.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}

func (t *UserHandler) GetAllUsers(c *gin.Context) {

	resp, err := t.userUseCase.GetAllUsers()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

func (h *UserHandler) Register(c *gin.Context) {
	req := new(entities.UsersRegisterReq)
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.userUseCase.Register(req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
	return
}
