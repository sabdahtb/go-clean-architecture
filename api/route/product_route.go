package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"sabda.go/learn/api/controller"
	"sabda.go/learn/bootstrap"
	"sabda.go/learn/repository"
	"sabda.go/learn/usecase"
)

func NewProductRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, route *gin.RouterGroup) {
	pr := repository.NewProductRepository(db)
	pc := &controller.ProductController{
		ProductUsecase: usecase.NewProductUsecase(pr, timeout),
		Env:            env,
		DB:             db,
	}

	route.POST("/create", pc.CreateProduct)
	route.GET("", pc.GetAllProduct)
	route.GET("/detail/:id", pc.GetProductById)
	route.DELETE("/delete/:id", pc.DeleteProductById)
	route.PATCH("/update/:id", pc.UpdateProduct)
}
