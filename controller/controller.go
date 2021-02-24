package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"vue-admin-backend/models"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 显示首页
func SayHello(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func AddTodo(c *gin.Context) {
	// 1.从前端页面中把数据拿出来
	var todo models.Todo
	c.ShouldBindJSON(&todo)
	fmt.Println(todo)
	// 2.把数据存入数据库
	// 存入数据的操作是 DB.Create(&todo) 但是这里把存入数据和返回响应写在一起了
	// 3.返回一个响应
	if err := models.CreateATodo(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func FindAllTodos(c *gin.Context) {
	// 查询表中所有的数据
	var todoList []models.Todo
	if err := models.FindAllTodos(&todoList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func UpdateTodo(c *gin.Context) {
	var todo models.Todo
	c.ShouldBindBodyWith(&todo, binding.JSON)
	fmt.Println("---------------")
	fmt.Println(todo)
	inputID := strconv.Itoa(todo.ID)
	// 找到要修改的数据
	if err := models.FindTodoById(&todo, inputID); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
		return
	}
	// 保存更新之后的数据
	// 为什么就没有了呢？？？？
	// 采用 ShouldBindBodyWith 目的是为了多次绑定
	// 如果再用 BindWith 将不会有数据被绑定
	c.ShouldBindBodyWith(&todo, binding.JSON)
	fmt.Println("========")
	fmt.Println(todo)
	if err := models.SaveTodo(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	if err := models.DeleteOneTodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}

func FindUser(c *gin.Context) {
	var user models.User
	c.ShouldBindJSON(&user)

	if err := models.FindAUser(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "用户不存在或密码错误",
		})
	} else {
		c.JSON(http.StatusAccepted, gin.H{
			"message": "用户存在",
		})
	}
}

func FindItemById(c *gin.Context) {
	var todo models.Todo
	id, ok := c.Params.Get("id")
	fmt.Println(id)
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	if err := models.FindTodoById(&todo, id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		fmt.Println(todo)
		c.JSON(http.StatusOK, todo)
	}
}
