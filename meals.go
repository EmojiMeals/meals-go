package meals

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"
)

var (
	ErrNotFound = errors.New("recipe not found")
	recipes     map[string]string
)

// load recipes
func init() {
	recipes = make(map[string]string)
	data, err := ioutil.ReadFile("recipes.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(data, &recipes); err != nil {
		panic(err)
	}
}

func containsAll(str, sub string) bool {
	for _, ingredient := range strings.Split(sub, "") {
		if !strings.Contains(str, ingredient) {
			return false
		}
	}
	return true
}

func Mealify(ingredients string) (string, error) {
	for key, val := range recipes {
		if containsAll(key, ingredients) {
			return val, nil
		}
	}
	return "", ErrNotFound
}
