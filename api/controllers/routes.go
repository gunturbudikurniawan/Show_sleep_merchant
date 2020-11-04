package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gunturbudikurniawan/Show_sleep_merchant/api/middlewares"
)

// func hello(w http.ResponseWriter, r *http.Request) {
// 	io.WriteString(w, "Hello")
// }
func (s *Server) initializeRoutes() {

	v1 := s.Router.Group("/api/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "home.tmpl", nil)
		})
		// Login Route
		v1.POST("/login", s.Login)

		// Reset password:
		v1.POST("/password/forgot", s.ForgotPassword)
		v1.POST("/password/reset", s.ResetPassword)

		// merchants
		v1.POST("/merchants", s.CreateMerchants)
		v1.POST("/merchants/login", s.LoginMerchant)
		v1.POST("/merchants/savedorder", middlewares.TokenAuthMiddleware(), s.CreateSavedOrder)
		v1.POST("/merchants/onlinesales", middlewares.TokenAuthMiddleware(), s.CreateOnlineSales)
		v1.POST("/merchants/sales", middlewares.TokenAuthMiddleware(), s.CreateSales)

		//Users routes
		v1.POST("/users", s.CreateUser)
		v1.GET("/users", s.GetUsers)
		v1.GET("/users/:id", s.GetUser)
		v1.PUT("/users/:id", middlewares.TokenAuthMiddleware(), s.UpdateUser)
		v1.PUT("/avatar/users/:id", middlewares.TokenAuthMiddleware(), s.UpdateAvatar)
		v1.DELETE("/users/:id", middlewares.TokenAuthMiddleware(), s.DeleteUser)

		//Posts routes
		v1.POST("/posts", middlewares.TokenAuthMiddleware(), s.CreatePost)
		v1.GET("/posts", s.GetPosts)
		v1.GET("/posts/:id", s.GetPost)
		v1.PUT("/posts/:id", middlewares.TokenAuthMiddleware(), s.UpdatePost)
		v1.DELETE("/posts/:id", middlewares.TokenAuthMiddleware(), s.DeletePost)
		v1.GET("/user_posts/:id", s.GetUserPosts)

		//Like route
		v1.GET("/likes/:id", s.GetLikes)
		v1.POST("/likes/:id", middlewares.TokenAuthMiddleware(), s.LikePost)
		v1.DELETE("/likes/:id", middlewares.TokenAuthMiddleware(), s.UnLikePost)

		//Comment routes
		v1.POST("/comments/:id", middlewares.TokenAuthMiddleware(), s.CreateComment)
		v1.GET("/comments/:id", s.GetComments)
		v1.PUT("/comments/:id", middlewares.TokenAuthMiddleware(), s.UpdateComment)
		v1.DELETE("/comments/:id", middlewares.TokenAuthMiddleware(), s.DeleteComment)

		//wishlist
		v1.POST("/wishlist/:id", middlewares.TokenAuthMiddleware(), s.AddToWishlist)
		v1.DELETE("/wishlist/:id", middlewares.TokenAuthMiddleware(), s.DeleteFromWishlist)
	}
}
