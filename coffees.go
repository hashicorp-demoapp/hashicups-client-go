package hashicups

import (
	"encoding/json"
	"fmt"
	"net/http"
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
