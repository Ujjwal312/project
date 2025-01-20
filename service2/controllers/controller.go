package controller

import (
	"net/http"

	"example.com/m/models"
	"example.com/m/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}
type ProductController struct {
	ProductService services.ProductService
}

func Newproduct(productService services.ProductService) ProductController {
	return ProductController{
		ProductService: productService,
	}
}
func Newuser(userService services.UserService) UserController {
	return UserController{
		UserService: userService,
	}
}
func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.UserService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	username := ctx.Param("name")
	user, err := uc.UserService.GetUser((&username))
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (uc *UserController) GetAll(ctx *gin.Context) {
	users, err := uc.UserService.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)

}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.UserService.UpdateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})

}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	username := ctx.Param("name")
	err := uc.UserService.DeleteUser(&username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}
func (uc *ProductController) CreateProduct(ctx *gin.Context) {
	var Product models.Product
	if err := ctx.ShouldBindJSON(&Product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.ProductService.CreateProduct(&Product)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *ProductController) GetProduct(ctx *gin.Context) {
	pname := ctx.Param("id")
	product, err := uc.ProductService.GetProduct((&pname))
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func (uc *ProductController) GetAllProduct(ctx *gin.Context) {
	product, err := uc.ProductService.GetAllProduct()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, product)

}

func (uc *ProductController) DeleteProduct(ctx *gin.Context) {
	pname := ctx.Param("name")
	err := uc.ProductService.DeleteProduct(&pname)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}
func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/user")
	userroute.POST("/create", uc.CreateUser)
	userroute.GET("/get/:name", uc.GetUser)
	userroute.GET("/getall", uc.GetAll)
	userroute.PATCH("/update", uc.UpdateUser)
	userroute.DELETE("/delete/:id", uc.DeleteUser)

}

func (uc *ProductController) RegisterProductRoutes(rg *gin.RouterGroup) {
	rg.GET("/getall", uc.GetAllProduct)
	rg.GET("/getproduct/:id", uc.GetProduct)
	rg.POST("/createproduct", uc.CreateProduct)
	rg.DELETE("/delete/:name", uc.DeleteProduct)

}
