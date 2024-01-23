package delivery

import "github.com/gin-gonic/gin"

func (h *Handler) Routes() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1")
	{
		person := api.Group("/person")
		{
			person.POST("/add", h.AddPerson)
			person.DELETE("/delete/:id", h.DeletePersonById)
			person.PUT("/update/:id", h.UpdatePersonById)
			person.GET("/:id", h.FindPersonById)
			person.GET("/", h.FindPersonsBySearchParam)
		}
	}
	return router
}
