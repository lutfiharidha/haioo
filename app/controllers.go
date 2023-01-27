package app

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lutfiharidha/haioo/helper"
)

type CartController struct {
	model CartModel
}

func NewCartController() *CartController {
	return &CartController{
		model: NewCartModels(),
	}
}

func (c *CartController) Init(m CartModel) {
	c.model = m //includes a model for creating transaction data in the database
}

func (c *CartController) AddProduct(context *gin.Context) {
	var req AddProductRequest
	err := context.ShouldBindJSON(&req)
	if err != nil {
		res := helper.BuildErrorResponse(err.Error())
		context.JSON(http.StatusBadRequest, res)
		return
	}
	res, err := c.model.AddProduct(req)
	if err != nil {
		res := helper.BuildErrorResponse(err.Error())
		context.JSON(http.StatusInternalServerError, res)
		return
	}
	success := helper.BuildSuccessResponse(res, true)
	context.JSON(http.StatusOK, success)
}

func (c *CartController) DeleteProduct(context *gin.Context) {
	var req DeleteProductRequest
	err := context.ShouldBindJSON(&req)
	if err != nil {
		res := helper.BuildErrorResponse(err.Error())
		context.JSON(http.StatusBadRequest, res)
		return
	}

	res, err := c.model.DeleteProduct(req)
	if err != nil {
		res := helper.BuildErrorResponse(err.Error())
		context.JSON(http.StatusInternalServerError, res)
		return
	}
	success := helper.BuildSuccessResponse("", res)
	context.JSON(http.StatusOK, success)
}

func (c *CartController) ShowCart(context *gin.Context) {
	var req ShowCartRequest

	err := context.BindQuery(&req)

	if err != nil {
		res := helper.BuildErrorResponse(err.Error())
		context.JSON(http.StatusBadRequest, res)
		return
	}
	res, err := c.model.ShowCart(req)
	if err != nil {
		res := helper.BuildErrorResponse(err.Error())
		context.JSON(http.StatusInternalServerError, res)
		return
	}
	ress := []string{}
	for _, v := range res {
		qty := strconv.Itoa(v.Kuantitas)
		ress = append(ress, v.KodeProduk+" - "+v.NamaProduk+" - "+qty)
	}
	context.JSON(http.StatusOK, ress)
}
