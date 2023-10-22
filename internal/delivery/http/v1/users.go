package v1

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/danluki/effective-mobile-test/internal/service"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initUsersRoutes(api *gin.RouterGroup) {
	users := api.Group("/users")
	{
		users.POST("/", h.userCreate)
		users.GET("/", h.userGetMany)
		users.PATCH("/", h.userPatch)
		users.DELETE("/:id", h.userDelete)
	}
}

type userCreateInput struct {
	Name       string `json:"name"       binding:"requied"`
	Surname    string `json:"surname"    binding:"required"`
	Patronymic string `json:"patronymic" binding:"required"`
}

func (h *Handler) userCreate(c *gin.Context) {
	var input userCreateInput
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.services.Users.Create(c.Request.Context(), service.CreateUserInput{
		Name:       input.Name,
		Surname:    input.Surname,
		Patronymic: input.Patronymic,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, user)
}

type userGetManyInput struct {
	Gender   *string `json:"gender"    form:"gender"`
	MinAge   *int    `json:"min_age"   form:"min_age"   binding:"min=0"`
	MaxAge   *int    `json:"max_age"   form:"max_age"`
	Country  *string `json:"country"   form:"country"`
	Page     int     `json:"page"      form:"page"      binding:"required"`
	PageSize int     `json:"page_size" form:"page_size" binding:"required"`
}

func (h *Handler) userGetMany(c *gin.Context) {
	var input userGetManyInput
	if err := c.BindQuery(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}

	users, err := h.services.Users.List(c.Request.Context(), service.ListUsersInput{
		Gender:   input.Gender,
		MinAge:   input.MinAge,
		MaxAge:   input.MaxAge,
		Country:  input.Country,
		Page:     input.Page,
		PageSize: input.PageSize,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, users)
}

type userPatchInput struct {
	Name       *string `json:"name"`
	Surname    *string `json:"surname"`
	Patronymic *string `json:"patronymic"`
	Age        *uint   `json:"age"`
	Gender     *string `json:"gender"`
	Country    *string `json:"country"`
}

func (h *Handler) userPatch(c *gin.Context) {
	var input userPatchInput
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}

	id := c.Param("id")

	convertedId, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}

	user, err := h.services.Users.Update(c.Request.Context(), service.UpdateUserInput{
		ID:         convertedId,
		Name:       input.Name,
		Surname:    input.Surname,
		Patronymic: input.Patronymic,
		Age:        input.Age,
		Gender:     input.Gender,
		Country:    input.Country,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) userDelete(c *gin.Context) {
	id := c.Param("id")

	convertedId, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Users.Delete(c.Request.Context(), int32(convertedId))
	if err != nil {
		if errors.Is(err, errors.New("not found")) {
			c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, "Internal Server Error")
	}

	c.JSON(http.StatusOK, nil)
}
