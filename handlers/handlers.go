package handlers

import (
	"net/http"

	"aktai/domain"
	"aktai/services"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Platform *services.Services
}

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func NewRouter(platform *services.Services) *gin.Engine {
	gin.SetMode(gin.DebugMode)
	router := gin.New()
	router.Use(gin.Logger())
	router.SetTrustedProxies(nil)

	h := Handlers{Platform: platform}

	college := router.Group("/college")
	college.GET("/", h.GetAllColleges)
	college.GET("/:id", h.GetCollege)
	college.POST("/", h.CreateCollege)
	college.PUT("/:id", h.UpdateCollege)
	college.DELETE("/:id", h.DeleteCollege)

	router.GET("/", h.HelloWorld)

	return router
}

func (h *Handlers) HelloWorld(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Hello World!")
}

func (h *Handlers) GetAllColleges(ctx *gin.Context) {
	var response Response

	colleges, err := h.Platform.GetAllColleges()
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response.Data = colleges
	response.Status = http.StatusOK

	ctx.JSON(http.StatusOK, response)
}

func (h *Handlers) GetCollege(ctx *gin.Context) {
	var response Response
	id := ctx.Param("id")

	college, err := h.Platform.GetCollege(id)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response.Data = college
	response.Status = http.StatusOK

	ctx.JSON(http.StatusOK, response)
}

func (h *Handlers) CreateCollege(ctx *gin.Context) {
	var response Response
	var request domain.College
	ctx.Header("Content-Type", "application/json")

	if err := ctx.BindJSON(&request); err != nil {
		response.Data = err.Error()
		response.Status = http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	college, code, err := h.Platform.CreateCollege(request)
	if err != nil {
		response.Data = err.Error()
		response.Status = code
		ctx.JSON(code, response)
		return
	}

	response.Data = college
	response.Status = http.StatusOK

	ctx.JSON(http.StatusOK, response)
}

func (h *Handlers) UpdateCollege(ctx *gin.Context) {
	var response Response
	var request domain.College
	ctx.Header("Content-Type", "application/json")
	id := ctx.Param("id")

	if err := ctx.BindJSON(&request); err != nil {
		response.Data = err.Error()
		response.Status = http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	request.ID = id

	college, err := h.Platform.UpdateCollege(request)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Data = college
	response.Status = http.StatusOK

	ctx.JSON(http.StatusOK, response)
}

func (h *Handlers) DeleteCollege(ctx *gin.Context) {
	var response Response
	id := ctx.Param("id")

	if err := h.Platform.DeleteCollege(id); err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Data = nil
	response.Status = http.StatusOK

	ctx.JSON(http.StatusOK, response)
}
