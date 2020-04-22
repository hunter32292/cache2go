package cache

import (
	"fmt"
)

var currentId int

var values Values

func init() {
	RepoCreateValue(Value{Value: "Example A"})
	RepoCreateValue(Value{Value: "Example B"})
}

func RepoFindValue(key string) Value {
	for _, t := range values {
		if t.Value == key {
			return t
		}
	}
	// return empty Value if not found
	return Value{}
}

func RepoCreateValue(t Value) Value {
	currentId += 1
	t.Id = currentId
	values = append(values, t)
	return t
}

func RepoDestroyValue(value string) error {
	for i, t := range values {
		if t.Value == value {
			values = append(values[:i], values[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Value with id of %q to delete", value)
}
