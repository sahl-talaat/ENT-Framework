package routes

import (
	controller "entdemo/Controller"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	v1 := r.Group("api/v1")
	{
		v1.GET("employees/:id", controller.GetEmployee)
		v1.POST("employees", controller.CreateEmployee)
		v1.PUT("employees/:id", controller.UpdateEmployee)
		v1.DELETE("employees/:id", controller.DeleteEmployee)

		v1.GET("departments/:id", controller.GetDepartment)
		v1.POST("departments", controller.CreateDepartment)
		v1.PUT("departments/:id", controller.UpdateDepartment)
		v1.DELETE("departments/:id", controller.DeleteDepartment)

		v1.GET("branches/:id", controller.GetBranch)
		v1.POST("branches", controller.CreateBranch)
		v1.PUT("branches/:id", controller.UpdateBranch)
		v1.DELETE("branches/:id", controller.DeleteBranch)

		v1.GET("products/:id", controller.GetProduct)
		v1.POST("products", controller.CreateProduct)
		v1.PUT("products/:id", controller.UpdateProduct)
		v1.DELETE("products/:id", controller.DeleteProduct)
	}
}
