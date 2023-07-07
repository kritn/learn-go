package servers

// ทำหน้าที่ Maps Endpoint/Router เข้าด้วยกัน เพื่อให้ตัว Services API สามารถเข้าถึง Modules ต่างๆได้

import (
	auth_controller "go_cleanarc/modules/auth/controllers"
	auth_repositories "go_cleanarc/modules/auth/repositories"
	auth_usecases "go_cleanarc/modules/auth/usecases"
	items_controller "go_cleanarc/modules/items/controllers"
	items_repositories "go_cleanarc/modules/items/repositories"
	items_usecases "go_cleanarc/modules/items/usecases"
	"go_cleanarc/modules/users/controllers"
	"go_cleanarc/modules/users/repositories"
	"go_cleanarc/modules/users/usecases"
	"go_cleanarc/pkg/databases"
	"go_cleanarc/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	userRepo := repositories.NewUserRepository(databases.DB)
	userUseCase := usecases.NewUserUseCase(userRepo)
	userHandler := controllers.NewUserHandler(userUseCase)

	authRepo := auth_repositories.NewAuthRepository(databases.DB)
	authUseCase := auth_usecases.NewAuthUsecase(authRepo, userRepo)
	authHandler := auth_controller.NewAuthController(authUseCase)

	itemRepo := items_repositories.NewItemRepository(databases.DB)
	itemUseCase := items_usecases.NewItemUseCase(itemRepo)
	itemHandler := items_controller.NewItemHandler(itemUseCase)

	v0 := r.Group("/auth")
	{
		v0.POST("login", authHandler.Login)
		v0.POST("register", userHandler.Register)
	}

	v1 := r.Group("/v1", middlewares.JwtAuthentication())
	{
		v1.GET("users", userHandler.GetAllUsers)
	}

	v2 := r.Group("/v2")
	{
		v2.GET("items", itemHandler.GetAllItems)
		v2.POST("item", itemHandler.CreateAItem)
		v2.GET("item/:id", itemHandler.GetAItem)
		v2.PUT("item/:id", itemHandler.UpdateAItem)
		v2.DELETE("item/:id", itemHandler.DeleteAItem)
	}
	return r
}
