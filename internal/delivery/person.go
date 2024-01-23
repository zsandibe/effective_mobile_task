package delivery

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zsandibe/effective_mobile_task/internal/domain"
)

const (
	defaultPage = 1
	defaultSize = 10
)

func (h *Handler) FindPersonsBySearchParam(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	// Check for errors or if the page is less than or equal to 0
	if err != nil || page <= 0 {
		page = defaultPage
	}
	size, err := strconv.Atoi(c.Query("size"))
	if err != nil || size <= 0 {
		size = defaultSize
	}
	searchParams := domain.SearchParams{
		Gender:      c.Query("gender"),
		Nationality: c.Query("nationality"),
		PageStart:   (page - 1) * size,
		PageLimit:   size,
	}
	persons, err := h.service.FindPersonsList(searchParams)
	if err != nil {
		if errors.Is(err, domain.NotFoundByParams) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, persons)
}

func (h *Handler) AddPerson(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer c.Request.Body.Close()

	var input domain.UserInput

	if err := json.Unmarshal(body, &input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if input.Name == "" || input.Surname == "" {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	person := domain.Person{
		Name:       input.Name,
		Surname:    input.Surname,
		Patronymic: input.Patronymic,
	}

	person, err = h.service.AddPerson(person)
	if err != nil {
		fmt.Println("OK")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, person)

}

func (h *Handler) FindPersonById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	person, err := h.service.FindPersonById(id)
	if err != nil {
		if errors.Is(err, domain.NotFoundById) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}

	c.JSON(http.StatusOK, person)

}

func (h *Handler) UpdatePersonById(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var person domain.Person

	if err := json.Unmarshal(body, &person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !UpdatedInputValidation(person) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.service.UpdatePersonById(person); err != nil {
		if errors.Is(err, domain.NothingWasUpdated) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)

}

func (h *Handler) DeletePersonById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.DeletePersonById(id); err != nil {
		if errors.Is(err, domain.NothingWasDeleted) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)

}

func UpdatedInputValidation(input domain.Person) bool {
	if input.Id <= 0 || input.Name == "" || input.Surname == "" || input.Age <= 0 || input.Gender == "" || input.Nationality == "" {
		return false
	}
	return true
}
