package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sut63/team01/ent"
	"github.com/sut63/team01/ent/annotation"
)

// AnnotationController defines the struct for the annotation controller
type AnnotationController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateAnnotation handles POST requests for adding annotation entities
// @Summary Create annotation
// @Description Create annotation
// @ID create-annotation
// @Accept   json
// @Produce  json
// @Param annotation body ent.Annotation true "Annotation entity"
// @Success 200 {object} ent.Annotation
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /annotations [post]
func (ctl *AnnotationController) CreateAnnotation(c *gin.Context) {
	obj := ent.Annotation{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "annotation binding failed",
		})
		return
	}

	a, err := ctl.client.Annotation.
		Create().
		SetMessages(obj.Messages).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving annotation failed",
		})
		return
	}

	c.JSON(200, a)
}

// GetAnnotation handles GET requests to retrieve a annotation entity
// @Summary Get a annotation entity by ID
// @Description get annotation by ID
// @ID get-annotation
// @Produce  json
// @Param id path int true "Annotation ID"
// @Success 200 {object} ent.Annotation
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /annotations/{id} [get]
func (ctl *AnnotationController) GetAnnotation(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	a, err := ctl.client.Annotation.
		Query().
		Where(annotation.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, a)
}

// ListAnnotation handles request to get a list of annotation entities
// @Summary List annotation entities
// @Description list annotation entities
// @ID list-annotation
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Annotation
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /annotations [get]
func (ctl *AnnotationController) ListAnnotation(c *gin.Context) {
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

	annotations, err := ctl.client.Annotation.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, annotations)
}

// DeleteAnnotation handles DELETE requests to delete a annotation entity
// @Summary Delete a annotation entity by ID
// @Description get annotation by ID
// @ID delete-annotation
// @Produce  json
// @Param id path int true "Annotation ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /annotations/{id} [delete]
func (ctl *AnnotationController) DeleteAnnotation(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Annotation.
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

// UpdateAnnotation handles PUT requests to update a annotation entity
// @Summary Update a annotation entity by ID
// @Description update annotation by ID
// @ID update-annotation
// @Accept   json
// @Produce  json
// @Param id path int true "Annotation ID"
// @Param annotation body ent.Annotation true "Annotation entity"
// @Success 200 {object} ent.Annotation
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /annotations/{id} [put]
func (ctl *AnnotationController) UpdateAnnotation(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Annotation{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "annotation binding failed",
		})
		return
	}

	a, err := ctl.client.Annotation.
		UpdateOneID(int(id)).
		SetMessages(obj.Messages).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "update annotation failed",
		})
		return
	}

	c.JSON(200, a)
}

// NewAnnotationController creates and registers handles for the AnnotationController
func NewAnnotationController(router gin.IRouter, client *ent.Client) *AnnotationController {
	ac := &AnnotationController{
		client: client,
		router: router,
	}
	ac.register()
	return ac
}

// InitAnnotationController registers routes to the main engine
func (ctl *AnnotationController) register() {
	annotations := ctl.router.Group("/annotations")

	annotations.GET("", ctl.ListAnnotation)

	// CRUD
	annotations.POST("", ctl.CreateAnnotation)
	annotations.GET(":id", ctl.GetAnnotation)
	annotations.PUT(":id", ctl.UpdateAnnotation)
	annotations.DELETE(":id", ctl.DeleteAnnotation)
}
