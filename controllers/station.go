package controllers

import (
	"net/http"
	"station-service/models"

	"github.com/gin-gonic/gin"
)

type StationController struct{}

type getStationRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type getStationListRequest struct {
	Offset int32 `form:"offset"`
	Limit  int32 `form:"limit" binding:"required,min=1,max=20"`
}

type createStationRequest struct {
	Name      string  `json:"name" db:"name"`
	Latitude  float32 `json:"lat" db:"lat"`
	Longitude float32 `json:"lng" db:"lng"`
	Provider  string  `json:"provider" db:"provider"`
}

type updateStationRequest struct {
	Name      string  `json:"name" db:"name"`
	Latitude  float32 `json:"lat" db:"lat"`
	Longitude float32 `json:"lng" db:"lng"`
	Provider  string  `json:"provider" db:"provider"`
}

var station = new(models.Station)

func (sc StationController) GetByID(ctx *gin.Context) {

	// Check if request has ID field in URI.
	var req getStationRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		ctx.Abort()
		return
	}

	// Execute query.
	result, err := station.GetByID(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (sc StationController) GetAll(ctx *gin.Context) {

	// Check if request has parameters offset and limit for pagination.
	var req getStationListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		ctx.Abort()
		return
	}

	arg := models.ListStationsParam{
		Offset: req.Offset,
		Limit:  req.Limit,
	}

	// Execute query.
	result, err := station.GetAll(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (sc StationController) Create(ctx *gin.Context) {

	// Check if request has all required fields in json body.
	var req createStationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		ctx.Abort()
		return
	}

	arg := models.CreateStationParam{
		Name:      req.Name,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
		Provider:  req.Provider,
	}

	// Execute query.
	result, err := station.Create(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, result)
}

func (sc StationController) Update(ctx *gin.Context) {

	// Check if request has ID field in URI.
	var reqID getStationRequest
	if err := ctx.ShouldBindUri(&reqID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		ctx.Abort()
		return
	}

	// Check if request has all required fields in json body.
	var req updateStationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		ctx.Abort()
		return
	}

	arg := models.UpdateStationParam{
		Name:      req.Name,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
		Provider:  req.Provider,
	}

	// Execute query.
	result, err := station.Update(ctx, arg, reqID.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, result)
}

func (sc StationController) Delete(ctx *gin.Context) {

	// Check if request has ID field in URI.
	var req getStationRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		ctx.Abort()
		return
	}

	// Execute query.
	if err := station.Delete(ctx, req.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
