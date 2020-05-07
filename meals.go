package meals

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"
)

var ErrNotFound = errors.New("recipe not found")

type Cookbook struct {
	recipes map[string]string
}

func NewCookbook(path string) *Cookbook {
	recipes := make(map[string]string)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(data, &recipes); err != nil {
		panic(err)
	}
	return &Cookbook{recipes}
}

func containsAll(str string, sub ...string) bool {
	for _, ingredient := range sub {
		if !strings.Contains(str, ingredient) {
			return false
		}
	}
	return true
}

func (cb *Cookbook) Mealify(ingredients ...string) (string, error) {
	for key, val := range cb.recipes {
		if containsAll(key, ingredients...) {
			return val, nil
		}
	}
	return "", ErrNotFound
}
