package debug

import (
	"testing"
)
func TestBasic(t *testing.T) {

	type Author struct {
		Name string
		Age int
	}
	nimo := Author{
		Name: "nimo",
		Age: 27,
	}
	Debug(nimo)

}













