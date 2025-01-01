package main

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/", RouteTest)

	// Lista todas tarefas
	router.GET("/tarefas", TaskList)

	// Cadastrar tarefa
	router.POST("/tarefas", RegisterTask)

	// Buscar tarefa pelo ID
	router.GET("/tarefas/:id", GetTaskById)

	// Deletar tarefa pelo Id
	router.DELETE("/tarefas/:id", DeleteTask)

	// Atualizar tarefa pelo Id
	router.PUT("/tarefas/:id", UpdateTask)
}
