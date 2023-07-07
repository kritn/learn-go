package controllers

// มีหน้าที่ในการรับส่ง Context จาก HTTP Request หรือพูดง่ายๆก็คือ รับส่งข้อมูลหรือบริบทต่างๆที่ถูกยิงมากจาก API ที่ client ทำการยิงมา

import (
	"go_cleanarc/modules/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ItemHandler struct {
	itemUseCase entities.ItemsUseCase
}

func NewItemHandler(usecase entities.ItemsUseCase) *ItemHandler {
	return &ItemHandler{
		itemUseCase: usecase,
	}
}

func (t *ItemHandler) GetAllItems(c *gin.Context) {
	resp, err := t.itemUseCase.GetAllItems()

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

func (t *ItemHandler) CreateAItem(c *gin.Context) {
	req := new(entities.ItemReq)
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := t.itemUseCase.CreateAItem(req)// req=json

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, &req)
		return
	}
}

func (t *ItemHandler) GetAItem(c *gin.Context) {
	id := c.Params.ByName("id")
	var items entities.ItemRes

	err := t.itemUseCase.GetAItem(&items, id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, items)
	}
}

func (t *ItemHandler) UpdateAItem(c *gin.Context) {
	id := c.Params.ByName("id")
	var items entities.ItemReq
	if err := c.ShouldBindJSON(&items); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := t.itemUseCase.UpdateAItem(&items, id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, items)
	}
}

func (t *ItemHandler) DeleteAItem(c *gin.Context) {
	id := c.Params.ByName("id")
	var rows entities.RowsAffected
	err := t.itemUseCase.DeleteAItem(&rows, id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, rows)
	}
}
