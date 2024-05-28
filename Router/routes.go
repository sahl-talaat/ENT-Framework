package routes

import (
	controllers "myapp/Controller"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	v1 := r.Group("api/v1")
	{
		v1.GET("employees", controllers.GetEmployees)
		v1.POST("employees", controllers.CreateEmployee)
		v1.GET("employees/:id", controllers.GetEmployee)
		v1.PUT("employees/:id", controllers.UpdateEmployee)
		v1.DELETE("employees/:id", controllers.DeleteEmployee)

		v1.GET("departments", controllers.GetDepartments)
		v1.POST("departments", controllers.CreateDepartment)
		v1.GET("departments/:id", controllers.GetDepartment)
		v1.PUT("departments/:id", controllers.UpdateDepartment)
		v1.DELETE("departments/:id", controllers.DeleteDepartment)

		v1.GET("branches", controllers.GetBranches)
		v1.POST("branches", controllers.CreateBranch)
		v1.GET("branches/:id", controllers.GetBranch)
		v1.PUT("branches/:id", controllers.UpdateBranch)
		v1.DELETE("branches/:id", controllers.DeleteBranch)

		v1.GET("products", controllers.GetBranches)
		v1.POST("products", controllers.CreateProduct)
		v1.GET("products/:id", controllers.GetProduct)
		v1.PUT("products/:id", controllers.UpdateProduct)
		v1.DELETE("products/:id", controllers.DeleteProduct)
	}
}
