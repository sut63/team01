package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sut63/team01/ent"
	"github.com/sut63/team01/ent/levelofdangerous"
)

// LevelOfDangerousController defines the struct for the LevelOfDangerous controller
type LevelOfDangerousController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateLevelOfDangerous handles POST requests for adding LevelOfDangerous entities
// @Summary Create LevelOfDangerous
// @Description Create LevelOfDangerous
// @ID create-LevelOfDangerous
// @Accept   json
// @Produce  json
// @Param LevelOfDangerous body ent.LevelOfDangerous true "LevelOfDangerous entity"
// @Success 200 {object} ent.LevelOfDangerous
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /LevelOfDangerous [post]
func (ctl *LevelOfDangerousController) CreateLevelOfDangerous(c *gin.Context) {
	obj := ent.LevelOfDangerous{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "LevelOfDangerous binding failed",
		})
		return
	}
	u, err := ctl.client.LevelOfDangerous.
		Create().
		SetName(obj.Name).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
	c.JSON(200, u)
}

// GetLevelOfDangerous handles GET requests to retrieve a LevelOfDangerous entity
// @Summary Get a LevelOfDangerous entity by ID
// @Description get LevelOfDangerous by ID
// @ID get-LevelOfDangerous
// @Produce  json
// @Param id path int true "LevelOfDangerous ID"
// @Success 200 {object} ent.LevelOfDangerous
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /LevelOfDangerous/{id} [get]
func (ctl *LevelOfDangerousController) GetLevelOfDangerous(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	u, err := ctl.client.LevelOfDangerous.
		Query().
		Where(levelofdangerous.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, u)
}

// ListLevelOfDangerous handles request to get a list of LevelOfDangerous entities
// @Summary List LevelOfDangerous entities
// @Description list LevelOfDangerous entities
// @ID list-LevelOfDangerous
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.LevelOfDangerous
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /LevelOfDangerous [get]
func (ctl *LevelOfDangerousController) ListLevelOfDangerous(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}
	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}
	LevelOfDangerous, err := ctl.client.LevelOfDangerous.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, LevelOfDangerous)
}

// DeleteLevelOfDangerous handles DELETE requests to delete a LevelOfDangerous entity
// @Summary Delete a LevelOfDangerous entity by ID
// @Description get LevelOfDangerous by ID
// @ID delete-LevelOfDangerous
// @Produce  json
// @Param id path int true "LevelOfDangerous ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /LevelOfDangerous/{id} [delete]
func (ctl *LevelOfDangerousController) DeleteLevelOfDangerous(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = ctl.client.LevelOfDangerous.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// UpdateLevelOfDangerous handles PUT requests to update a LevelOfDangerous entity
// @Summary Update a LevelOfDangerous entity by ID
// @Description update LevelOfDangerous by ID
// @ID update-LevelOfDangerous
// @Accept   json
// @Produce  json
// @Param id path int true "LevelOfDangerous ID"
// @Param LevelOfDangerous body ent.LevelOfDangerous true "LevelOfDangerous entity"
// @Success 200 {object} ent.LevelOfDangerous
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /LevelOfDangerous/{id} [put]
func (ctl *LevelOfDangerousController) UpdateLevelOfDangerous(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	obj := ent.LevelOfDangerous{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "LevelOfDangerous binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.LevelOfDangerous.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}
	c.JSON(200, u)
}

// NewLevelOfDangerousController creates and registers handles for the LevelOfDangerous controller
func NewLevelOfDangerousController(router gin.IRouter, client *ent.Client) *LevelOfDangerousController {
	uc := &LevelOfDangerousController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitLevelOfDangerousController registers routes to the main engine
func (ctl *LevelOfDangerousController) register() {
	LevelOfDangerous := ctl.router.Group("/LevelOfDangerous")

	LevelOfDangerous.GET("", ctl.ListLevelOfDangerous)

	// CRUD
	LevelOfDangerous.POST("", ctl.CreateLevelOfDangerous)
	LevelOfDangerous.GET(":id", ctl.GetLevelOfDangerous)
	LevelOfDangerous.PUT(":id", ctl.UpdateLevelOfDangerous)
	LevelOfDangerous.DELETE(":id", ctl.DeleteLevelOfDangerous)
}
