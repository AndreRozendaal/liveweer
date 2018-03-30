package liveweer_test

import (
	"testing"
	"strconv"
	"fmt"
	"liveweer"
)

func TestFoo(t *testing.T) {
	weer, err := liveweer.GetJSON("http://weerlive.nl/api/json-10min.php?locatie=Hilversum")
	if err != nil {
		t.Errorf("Error in GetJSON")
	} else {
		f, _ := strconv.ParseFloat(weer.Temp(),64)
		fmt.Printf("temperature: %v\n", f)   
	}
}


