package adapter

import "testing"

var expect = "adaptee method"

func TestAdapter(t *testing.T) {
    adaptee := NewAdaptee()
    target := NewAdapter(adaptee)
    res := target.Request()
    t.Logf("adapter info %s",res)
    if res != expect {
        t.Fatalf("expectd: %s, actual: %s", expect, res)
    }
}
