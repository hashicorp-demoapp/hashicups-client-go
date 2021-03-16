package hashicups

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetCoffees - Returns list of coffees (no auth required)
func (c *Client) GetCoffees() ([]Coffee, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/coffees", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	coffees := []Coffee{}
	err = json.Unmarshal(body, &coffees)
	if err != nil {
		return nil, err
	}

	return coffees, nil
}

// GetCoffeeIngredients - Returns list of coffee ingredients (no auth required)
func (c *Client) GetCoffeeIngredients(coffeeID string) ([]Ingredient, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/coffees/%s/ingredients", c.HostURL, coffeeID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	ingredients := []Ingredient{}
	err = json.Unmarshal(body, &ingredients)
	if err != nil {
		return nil, err
	}

	return ingredients, nil
}

// CreateCoffee - Create new coffee
func (c *Client) CreateCoffee(coffee Coffee) (*Coffee, error) {
	rb, err := json.Marshal(coffee)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/coffees", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	newCoffee := Coffee{}
	err = json.Unmarshal(body, &newCoffee)
	if err != nil {
		return nil, err
	}

	return &newCoffee, nil
}

// CreateCoffeeIngredient - Create new coffee ingredient
func (c *Client) CreateCoffeeIngredient(coffee Coffee, ingredient Ingredient) (*Ingredient, error) {
	reqBody := struct {
		CoffeeID     int    `json:"coffee_id"`
		IngredientID int    `json:"ingredient_id"`
		Quantity     int    `json:"quantity"`
		Unit         string `json:"unit"`
	}{
		CoffeeID:     coffee.ID,
		IngredientID: ingredient.ID,
		Quantity:     ingredient.Quantity,
		Unit:         ingredient.Unit,
	}
	rb, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/coffees/%d/ingredients", c.HostURL, coffee.ID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	newIngredient := Ingredient{}
	err = json.Unmarshal(body, &newIngredient)
	if err != nil {
		return nil, err
	}

	return &newIngredient, nil
}
