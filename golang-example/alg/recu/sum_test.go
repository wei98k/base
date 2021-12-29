package recu

import(
	"testing"
)

func TestSum(t *testing.T) {
	s := sum(100)
	
	var want = 5050

	if s != want {
		t.Errorf("result: %v", s)
	}
}