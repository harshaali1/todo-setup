package main

import (
	"tms-backend/src/database"
	"tms-backend/src/handlers"
	"tms-backend/src/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	// Initialize database
	database.Connect()
	database.Migrate()

	router := gin.Default()

	// Public routes
	router.POST("/register", handlers.Register)

	// JWT middleware
	authMiddleware := middleware.AuthMiddleware()
	router.POST("/login", authMiddleware.LoginHandler)

	// Protected routes
	auth := router.Group("/")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/todos", handlers.GetTodos)
		auth.POST("/todos", handlers.CreateTodo)
	}

	router.Run(":8000")
}

/*package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Todo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	UserID      string `json:"user_id"`
}

var todos = []Todo{}

func main1() {
	r := gin.Default()

	// Auth routes
	r.POST("/login", func(c *gin.Context) {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": "123",
			"exp":     time.Now().Add(24 * time.Hour).Unix(),
		})
		tokenString, _ := token.SignedString([]byte("secret"))
		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	})

	// Protected todo routes
	todoRoutes := r.Group("/todos")
	todoRoutes.Use(AuthMiddleware())
	{
		todoRoutes.GET("/", getTodos)
		todoRoutes.POST("/", createTodo)
		todoRoutes.PUT("/:id", updateTodo)
		todoRoutes.DELETE("/:id", deleteTodo) // New DELETE endpoint
	}

	r.Run(":8080")
}

// Delete todo handler
func deleteTodo(c *gin.Context) {
	userID := c.MustGet("user_id").(string)
	todoID := c.Param("id")

	for index, todo := range todos {
		if todo.ID == todoID && todo.UserID == userID {
			// Remove from slice
			todos = append(todos[:index], todos[index+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

// Existing handlers (GET, POST, PUT)
func getTodos(c *gin.Context) {
	userID := c.MustGet("user_id").(string)
	var userTodos []Todo
	for _, todo := range todos {
		if todo.UserID == userID {
			userTodos = append(userTodos, todo)
		}
	}
	c.JSON(http.StatusOK, userTodos)
}

func createTodo(c *gin.Context) {
	userID := c.MustGet("user_id").(string)
	var todo Todo
	c.ShouldBindJSON(&todo)
	todo.ID = "id_" + strconv.Itoa(len(todos)+1)
	todo.UserID = userID
	todos = append(todos, todo)
	c.JSON(http.StatusCreated, todo)
}

func updateTodo(c *gin.Context) {
	userID := c.MustGet("user_id").(string)
	todoID := c.Param("id")

	var updateData struct {
		Title       *string `json:"title"`
		Description *string `json:"description"`
		Completed   *bool   `json:"completed"`
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	for i, todo := range todos {
		if todo.ID == todoID && todo.UserID == userID {
			if updateData.Title != nil {
				todos[i].Title = *updateData.Title
			}
			if updateData.Description != nil {
				todos[i].Description = *updateData.Description
			}
			if updateData.Completed != nil {
				todos[i].Completed = *updateData.Completed
			}
			c.JSON(http.StatusOK, todos[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

// Auth middleware (unchanged)
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Bearer token not found"})
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("secret"), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token", "details": err.Error()})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("user_id", claims["user_id"])
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		}
	}
}
*/
