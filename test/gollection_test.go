package test

import (
	"fmt"
	"github.com/ansuir/gollection"
	"testing"
)

func TestGollection(t *testing.T) {

	gollection.NewCollection([]int{1, 2, 3}).Map(func(i int) interface{} {
		return fmt.Sprintf("I is %d", i)
	}).(*gollection.Collection[any]).Each(func(s any) {
		fmt.Printf("S is '%s' \n", s.(string))
	})

}
