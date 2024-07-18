package soma

import (
	"testing"
)

func TestShouldSum(t *testing.T){
	
	val := Soma(1,2)

	if val != 3 {
		t.Error("expecting:", 3, " got: ", val)
	}
}

func TestSomaN(t *testing.T) {
	val := SomaNums(2,3,4)

	if val != 9 {
		t.Error("expecting:", 9, " got: ", val)
	}
}