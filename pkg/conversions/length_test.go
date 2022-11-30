package conversions_test

import (
	"fmt"
	"testing"

	"github.com/tommartensen/jaga/pkg/conversions"
)

func TestLength(t *testing.T) {
	var l conversions.Length = 456.45
	fmt.Println(l)
}
