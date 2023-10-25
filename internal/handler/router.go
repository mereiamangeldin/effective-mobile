package handler

import "github.com/gin-gonic/gin"

func (h *Handler) InitRouter() *gin.Engine {
	h.logger.Info("initialization of router")
	router := gin.Default()

	apiv1 := router.Group("/api/v1")

	people := apiv1.Group("/people")
	people.POST("", h.addPerson)
	people.GET("", h.getPeople)
	people.GET("/:person_id", h.getPerson)
	people.DELETE("/:person_id", h.deletePerson)
	people.PUT("/:person_id", h.updatePerson)
	return router
}
