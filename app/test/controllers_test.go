package app_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lutfiharidha/haioo/app"
	mockmodel "github.com/lutfiharidha/haioo/app/test/mock"
	"github.com/lutfiharidha/haioo/helper"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestAddProduct(t *testing.T) {
	convey.Convey("TestAddProduct", t, func() {
		convey.Convey("Positive Scenario", func() {
			mockResponse := app.Cart{
				KodeProduk: "baju-2",
				NamaProduk: "Baju Nike",
				Kuantitas:  200,
			}
			expectedRes := helper.BuildSuccessResponse(mockResponse, true)
			jsonMockRes, _ := json.Marshal(expectedRes)

			mc := &mockmodel.CartModel{}
			ctrl := app.NewCartController()
			ctrl.Init(mc)
			mc.On("AddProduct", mock.Anything, mock.Anything).Return(mockResponse, nil)

			r := SetUpRouter()
			r.POST("/api/v1/cart/add", ctrl.AddProduct)

			request := app.AddProductRequest{
				KodeProduk: "baju-2",
				NamaProduk: "Baju Nike",
				Kuantitas:  200,
			}
			jsonValue, _ := json.Marshal(request)
			req, _ := http.NewRequest("POST", "/api/v1/cart/add", bytes.NewBuffer(jsonValue))

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			res, _ := ioutil.ReadAll(w.Body)
			convey.So(w.Code, convey.ShouldEqual, http.StatusOK)
			convey.So(string(res), convey.ShouldEqual, string(jsonMockRes))
		})

		convey.Convey("Negative Scenario", func() {
			convey.Convey("Error DB", func() {
				mockResponse := app.Cart{
					KodeProduk: "baju-2",
					NamaProduk: "Baju Nike",
					Kuantitas:  200,
				}
				expectedRes := helper.BuildSuccessResponse(mockResponse, true)
				jsonMockRes, _ := json.Marshal(expectedRes)

				mc := &mockmodel.CartModel{}
				ctrl := app.NewCartController()
				ctrl.Init(mc)
				mc.On("AddProduct", mock.Anything, mock.Anything).Return(nil, errors.New("Failed to process request"))

				r := SetUpRouter()
				r.POST("/api/v1/cart/add", ctrl.AddProduct)

				request := app.AddProductRequest{
					KodeProduk: "baju-2",
					NamaProduk: "Baju Nike",
					Kuantitas:  200,
				}
				jsonValue, _ := json.Marshal(request)
				req, _ := http.NewRequest("POST", "/api/v1/cart/add", bytes.NewBuffer(jsonValue))

				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				res, _ := ioutil.ReadAll(w.Body)
				convey.So(w.Code, convey.ShouldEqual, http.StatusInternalServerError)
				convey.So(string(res), convey.ShouldNotEqual, string(jsonMockRes))
			})

			convey.Convey("Incomplete request", func() {
				mockResponse := app.Cart{}
				jsonMockRes, _ := json.Marshal(mockResponse)

				mc := &mockmodel.CartModel{}
				ctrl := app.NewCartController()
				ctrl.Init(mc)
				mc.On("AddProduct", mock.Anything, mock.Anything).Return(mockResponse, nil)

				r := SetUpRouter()
				r.POST("/api/v1/cart/add", ctrl.AddProduct)

				request := app.AddProductRequest{
					NamaProduk: "Baju Adidas",
				}
				jsonValue, _ := json.Marshal(request)
				req, _ := http.NewRequest("POST", "/api/v1/cart/add", bytes.NewBuffer(jsonValue))

				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				res, _ := ioutil.ReadAll(w.Body)
				convey.So(w.Code, convey.ShouldEqual, http.StatusBadRequest)
				convey.So(string(res), convey.ShouldNotEqual, string(jsonMockRes))
			})

		})
	})
}

func TestShowCart(t *testing.T) {
	convey.Convey("TestShowCart", t, func() {
		convey.Convey("Positive Scenario", func() {
			convey.Convey("Show All Product in Cart", func() {
				cart := []app.Cart{
					{
						KodeProduk: "baju-2",
						NamaProduk: "Baju Nike",
						Kuantitas:  200,
					},
					{
						KodeProduk: "baju-1",
						NamaProduk: "Baju Adidas",
						Kuantitas:  100,
					},
				}
				mc := &mockmodel.CartModel{}
				ctrl := app.NewCartController()
				ctrl.Init(mc)
				mc.On("ShowCart", mock.Anything, mock.Anything).Return(cart, nil)

				r := SetUpRouter()
				r.GET("/api/v1/cart/show", ctrl.ShowCart)

				request := app.ShowCartRequest{}
				jsonValue, _ := json.Marshal(request)
				req, _ := http.NewRequest("GET", "/api/v1/cart/show", bytes.NewBuffer(jsonValue))

				ex := []string{}
				for _, v := range cart {
					qty := strconv.Itoa(v.Kuantitas)
					ex = append(ex, v.KodeProduk+" - "+v.NamaProduk+" - "+qty)
				}
				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				res, _ := ioutil.ReadAll(w.Body)
				convey.So(w.Code, convey.ShouldEqual, http.StatusOK)
				convey.So(string(res), convey.ShouldEqual, ex)
			})

			convey.Convey("Show Product by Product Name in Cart", func() {
				cart := []app.Cart{
					{
						KodeProduk: "baju-2",
						NamaProduk: "Baju Nike",
						Kuantitas:  200,
					},
					{
						KodeProduk: "baju-1",
						NamaProduk: "Baju Adidas",
						Kuantitas:  100,
					},
				}
				expectedRes := helper.BuildSuccessResponse(cart, true)
				jsonMockRes, _ := json.Marshal(expectedRes)

				mc := &mockmodel.CartModel{}
				ctrl := app.NewCartController()
				ctrl.Init(mc)
				mc.On("ShowCart", mock.Anything, mock.Anything).Return(expectedRes, nil)

				r := SetUpRouter()
				r.GET("/api/v1/cart/show", ctrl.ShowCart)

				request := app.ShowCartRequest{
					NamaProduk: "Baju",
				}
				jsonValue, _ := json.Marshal(request)
				req, _ := http.NewRequest("GET", "/api/v1/cart/show", bytes.NewBuffer(jsonValue))

				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				res, _ := ioutil.ReadAll(w.Body)
				convey.So(w.Code, convey.ShouldEqual, http.StatusOK)
				convey.So(string(res), convey.ShouldEqual, string(jsonMockRes))
			})

			convey.Convey("Show Product Order by Quantity in Cart", func() {
				cart := []app.Cart{
					{
						KodeProduk: "baju-2",
						NamaProduk: "Baju Nike",
						Kuantitas:  200,
					},
					{
						KodeProduk: "baju-1",
						NamaProduk: "Baju Adidas",
						Kuantitas:  100,
					},
				}
				expectedRes := helper.BuildSuccessResponse(cart, true)
				jsonMockRes, _ := json.Marshal(expectedRes)

				mc := &mockmodel.CartModel{}
				ctrl := app.NewCartController()
				ctrl.Init(mc)
				mc.On("ShowCart", mock.Anything, mock.Anything).Return(expectedRes, nil)

				r := SetUpRouter()
				r.GET("/api/v1/cart/show", ctrl.ShowCart)

				request := app.ShowCartRequest{
					Kuantitas: "desc",
				}
				jsonValue, _ := json.Marshal(request)
				req, _ := http.NewRequest("GET", "/api/v1/cart/show", bytes.NewBuffer(jsonValue))

				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				res, _ := ioutil.ReadAll(w.Body)
				convey.So(w.Code, convey.ShouldEqual, http.StatusOK)
				convey.So(string(res), convey.ShouldEqual, string(jsonMockRes))
			})
		})

		convey.Convey("Negative Scenario", func() {
			convey.Convey("Error DB", func() {
				expectedRes := helper.BuildSuccessResponse("", false)
				jsonMockRes, _ := json.Marshal(expectedRes)

				mc := &mockmodel.CartModel{}
				ctrl := app.NewCartController()
				ctrl.Init(mc)
				mc.On("ShowCart", mock.Anything, mock.Anything).Return(nil, errors.New("Failed to process request"))

				r := SetUpRouter()
				r.GET("/api/v1/cart/show", ctrl.ShowCart)

				request := app.ShowCartRequest{}
				jsonValue, _ := json.Marshal(request)
				req, _ := http.NewRequest("GET", "/api/v1/cart/show", bytes.NewBuffer(jsonValue))

				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				res, _ := ioutil.ReadAll(w.Body)
				convey.So(w.Code, convey.ShouldEqual, http.StatusInternalServerError)
				convey.So(string(res), convey.ShouldNotEqual, string(jsonMockRes))
			})

		})
	})
}

func TestDeleteProduct(t *testing.T) {
	convey.Convey("TestDeleteProduct", t, func() {
		convey.Convey("Positive Scenario", func() {
			expectedRes := helper.BuildSuccessResponse("", true)
			jsonMockRes, _ := json.Marshal(expectedRes)

			mc := &mockmodel.CartModel{}
			ctrl := app.NewCartController()
			ctrl.Init(mc)
			mc.On("DeleteProduct", mock.Anything, mock.Anything).Return(expectedRes, nil)

			r := SetUpRouter()
			r.DELETE("/api/v1/cart/delete", ctrl.DeleteProduct)

			request := app.DeleteProductRequest{
				KodeProduk: "baju-2",
			}
			jsonValue, _ := json.Marshal(request)
			req, _ := http.NewRequest("DELETE", "/api/v1/cart/delete", bytes.NewBuffer(jsonValue))

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			res, _ := ioutil.ReadAll(w.Body)
			convey.So(w.Code, convey.ShouldEqual, http.StatusOK)
			convey.So(string(res), convey.ShouldEqual, string(jsonMockRes))
		})

		convey.Convey("Negative Scenario", func() {
			convey.Convey("Error DB", func() {
				expectedRes := helper.BuildSuccessResponse("", false)
				jsonMockRes, _ := json.Marshal(expectedRes)

				mc := &mockmodel.CartModel{}
				ctrl := app.NewCartController()
				ctrl.Init(mc)
				mc.On("DeleteProduct", mock.Anything, mock.Anything).Return(nil, errors.New("Failed to process request"))

				r := SetUpRouter()
				r.DELETE("/api/v1/cart/delete", ctrl.DeleteProduct)

				request := app.DeleteProductRequest{
					KodeProduk: "baju-2",
				}
				jsonValue, _ := json.Marshal(request)
				req, _ := http.NewRequest("DELETE", "/api/v1/cart/delete", bytes.NewBuffer(jsonValue))

				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				res, _ := ioutil.ReadAll(w.Body)
				convey.So(w.Code, convey.ShouldEqual, http.StatusInternalServerError)
				convey.So(string(res), convey.ShouldNotEqual, string(jsonMockRes))
			})

			convey.Convey("Incomplete request", func() {
				mockResponse := app.Cart{}
				jsonMockRes, _ := json.Marshal(mockResponse)

				mc := &mockmodel.CartModel{}
				ctrl := app.NewCartController()
				ctrl.Init(mc)
				mc.On("DeleteProduct", mock.Anything, mock.Anything).Return(mockResponse, nil)

				r := SetUpRouter()
				r.DELETE("/api/v1/cart/delete", ctrl.DeleteProduct)

				request := app.DeleteProductRequest{}
				jsonValue, _ := json.Marshal(request)
				req, _ := http.NewRequest("DELETE", "/api/v1/cart/delete", bytes.NewBuffer(jsonValue))

				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				res, _ := ioutil.ReadAll(w.Body)
				convey.So(w.Code, convey.ShouldEqual, http.StatusBadRequest)
				convey.So(string(res), convey.ShouldNotEqual, string(jsonMockRes))
			})

		})
	})
}
