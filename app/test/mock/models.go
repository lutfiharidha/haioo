package mock

import (
	"github.com/lutfiharidha/haioo/app"
	"github.com/stretchr/testify/mock"
)

type CartModel struct {
	mock.Mock
}

func (c *CartModel) Init() {}

func (c *CartModel) AddProduct(req app.AddProductRequest) (res app.Cart, err error) {
	res = app.Cart{}

	call := c.Called()
	re := call.Get(0)
	er := call.Get(1)

	if er != nil {
		ree, ok := er.(error)
		if ok {
			return res, ree
		}
	}

	ree, ok := re.(app.Cart)
	if ok {
		return ree, nil
	}

	return res, err
}

func (c *CartModel) DeleteProduct(req app.DeleteProductRequest) (res bool, err error) {

	call := c.Called()
	re := call.Get(0)
	er := call.Get(1)

	if er != nil {
		ree, ok := er.(error)
		if ok {
			return res, ree
		}
	}

	ree, ok := re.(bool)
	if ok {
		return ree, nil
	}

	return res, err
}

func (c *CartModel) ShowCart(req app.ShowCartRequest) (res []app.Cart, err error) {

	res = []app.Cart{}

	call := c.Called()
	re := call.Get(0)
	er := call.Get(1)

	if er != nil {
		ree, ok := er.(error)
		if ok {
			return res, ree
		}
	}

	ree, ok := re.([]app.Cart)
	if ok {
		return ree, nil
	}

	return res, err
}
