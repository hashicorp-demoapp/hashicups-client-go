package hashicups

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// GetOrder - Returns a specifc order
func (c *Client) GetOrder(orderID string) (*Order, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/orders/%s", c.HostURL, orderID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	order := Order{}
	err = json.Unmarshal(body, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

// CreateOrder - Create new order
func (c *Client) CreateOrder(orderItems []OrderItem) (*Order, error) {
	rb, err := json.Marshal(orderItems)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/orders", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	order := Order{}
	err = json.Unmarshal(body, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

// CreateCustomOrder - Create new custom order
func (c *Client) CreateCustomOrder(orderItems []OrderItem) (*Order, error) {
	coffees, err := c.GetCoffees()
	if err != nil {
		return nil, err
	}

	ois := []OrderItem{}
	for _, o := range orderItems {
		order := &o
		customIngredients := order.Coffee.Ingredient

		realIngredients, err := c.GetCoffeeIngredients(strconv.Itoa(order.Coffee.ID))
		if err != nil {
			return nil, err
		}

		// Update order with custom coffee
		for _, coffee := range coffees {
			if coffee.ID == order.Coffee.ID {
				coffee.Name = order.Coffee.Name
				newCoffe, err := c.CreateCoffee(coffee)
				if err != nil {
					return nil, err
				}
				coffee.ID = newCoffe.ID
				order.Coffee = coffee
				continue
			}
		}

		// Update custom coffee with ingredients
		for _, ingredient := range realIngredients {
			for _, customIngredient := range customIngredients {
				if ingredient.ID == customIngredient.ID {
					ingredient.Quantity = customIngredient.Quantity
					continue
				}
			}
			_, err := c.CreateCoffeeIngredient(order.Coffee, ingredient)
			if err != nil {
				return nil, err
			}
		}

		ois = append(ois, OrderItem{
			Coffee: Coffee{
				ID: order.Coffee.ID,
			},
			Quantity: order.Quantity,
		})
	}

	rb, err := json.Marshal(ois)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/orders", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	order := Order{}
	err = json.Unmarshal(body, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

// UpdateOrder - Updates an order
func (c *Client) UpdateOrder(orderID string, orderItems []OrderItem) (*Order, error) {
	rb, err := json.Marshal(orderItems)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/orders/%s", c.HostURL, orderID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	order := Order{}
	err = json.Unmarshal(body, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

// DeleteOrder - Deletes an order
func (c *Client) DeleteOrder(orderID string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/orders/%s", c.HostURL, orderID), nil)
	if err != nil {
		return err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return err
	}

	if string(body) != "Deleted order" {
		return errors.New(string(body))
	}

	return nil
}
