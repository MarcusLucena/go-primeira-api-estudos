package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Tasks struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

var taskList = []Tasks{
	{Id: 891, Title: "Estudar Go Lang"},
	{Id: 123, Title: "Seguir o Sujeito Programador no Youtube"},
}

func RouteTest(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Primeira API em Go",
	})
}

func TaskList(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, taskList)
}

func RegisterTask(ctx *gin.Context) {
	var newTask Tasks

	if err := ctx.BindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newTask.Id = len(taskList) + 1
	taskList = append(taskList, newTask)
	ctx.JSON(http.StatusOK, newTask)
}

func GetTaskById(ctx *gin.Context) {
	id := ctx.Param("id")

	for _, task := range taskList {
		if fmt.Sprintf("%d", task.Id) == id {
			ctx.JSON(http.StatusOK, task)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"message": "Id não encontrado",
	})
}

func DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")

	for index, task := range taskList {
		if fmt.Sprintf("%d", task.Id) == id {
			taskList = append(taskList[:index], taskList[index+1:]...)
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Tarefa deletada com sucesso!",
			})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"message": "Id não encontrado",
	})
}

func UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")

	var updateTask Tasks

	if err := ctx.BindJSON(&updateTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Formato dos dados invalido",
		})
		return
	}

	for index, task := range taskList {
		if fmt.Sprintf("%d", task.Id) == id {
			updateTask.Id = task.Id
			taskList[index] = updateTask
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Tarega atualizada com sucesso!",
				"tarefa":  task.Title,
			})
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"message": "Id não encontrado",
	})
}
