package api_client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	client *http.Client
}

func NewClient(timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("time out can't be zero")
	}

	return &Client{
		client: &http.Client{
			Timeout: timeout,
		},
	}, nil
}

func (c Client) GetProducts() ([]Product, error) {
	resp, err := c.client.Get("https://fakestoreapi.com/products")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, resp.Body)
	if err != nil {
		return nil, err
	}

	r := make([]Product, 10)

	if err = json.Unmarshal(buf.Bytes(), &r); err != nil {
		return nil, err
	}

	return r, nil
}

func (c Client) GetProduct(id int) (Product, error) {
	resp, err := c.client.Get(fmt.Sprint("https://fakestoreapi.com/products/", id))

	if err != nil {
		return Product{}, err
	}

	defer resp.Body.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, resp.Body)
	if err != nil {
		return Product{}, err
	}

	var r Product

	if err = json.Unmarshal(buf.Bytes(), &r); err != nil {
		return Product{}, err
	}

	return r, nil
}

func (c Client) GetCategories() ([]string, error) {
	resp, err := c.client.Get("https://fakestoreapi.com/products/categories")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, resp.Body)
	if err != nil {
		return nil, err
	}

	var r []string

	if err = json.Unmarshal(buf.Bytes(), &r); err != nil {
		return nil, err
	}

	return r, nil
}

func (c Client) GetProductsFromCategory(category string) ([]Product, error) {
	resp, err := c.client.Get(fmt.Sprint("https://fakestoreapi.com/products/category/", category))

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, resp.Body)
	if err != nil {
		return nil, err
	}

	r := make([]Product, 10)

	if err = json.Unmarshal(buf.Bytes(), &r); err != nil {
		return nil, err
	}

	return r, nil
}

func (c Client) GetCarts() ([]Cart, error) {
	resp, err := c.client.Get("https://fakestoreapi.com/carts")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, resp.Body)
	if err != nil {
		return nil, err
	}

	r := make([]Cart, 1)

	if err = json.Unmarshal(buf.Bytes(), &r); err != nil {
		return nil, err
	}

	return r, nil
}
