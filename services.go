package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Tasks struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

func RouteTest(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Primeira API em Go",
	})
}

func TaskList(ctx *gin.Context) {
	rows, err := DB.Query("SELECT id, title FROM tasks")

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	defer rows.Close()

	var tasks []Tasks

	for rows.Next() {
		var task Tasks
		if err := rows.Scan(&task.Id, &task.Title); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		tasks = append(tasks, task)
	}

	ctx.JSON(http.StatusOK, tasks)
}

func RegisterTask(ctx *gin.Context) {
	var newTask Tasks

	if err := ctx.BindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := DB.Exec("INSERT INTO tasks (title) VALUES (?)", newTask.Title)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, err := result.LastInsertId()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newTask.Id = int(id)

	ctx.JSON(http.StatusCreated, newTask)

}

func GetTaskById(ctx *gin.Context) {
	id := ctx.Param("id")
	var task Tasks

	row := DB.QueryRow("SELECT id, title FROM tasks WHERE id = ?", id)

	if err := row.Scan(&task.Id, &task.Title); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, task)
}

func DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")

	_, err := DB.Exec("DELETE FROM tasks WHERE id = ?", id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Tarefa deletada com sucesso!"})
}

func UpdateTask(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var updateTask Tasks

	if err := ctx.BindJSON(&updateTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Formato dos dados invalido",
		})
		return
	}

	_, err := DB.Exec("UPDATE tasks SET title = ? WHERE id = ?", updateTask.Title, id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	updateTask.Id = id
	ctx.JSON(http.StatusOK, updateTask)
}
