package controllers

// มีหน้าที่ในการรับส่ง Context จาก HTTP Request หรือพูดง่ายๆก็คือ รับส่งข้อมูลหรือบริบทต่างๆที่ถูกยิงมากจาก API ที่ client ทำการยิงมา

import (
	"go_cleanarc/modules/entities"
	"net/http"

	// "go_cleanarc/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

type AuthCon struct {
	AuthUse entities.AuthUsecase
}

func NewAuthController(authUse entities.AuthUsecase) *AuthCon {
	return &AuthCon{
		AuthUse: authUse,
	}
}

func (h *AuthCon) Login(c *gin.Context) {
	req := new(entities.UsersCredentials)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.AuthUse.Login(req)
	if err != nil {
		// c.AbortWithStatus(http.StatusNotFound)

		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}
