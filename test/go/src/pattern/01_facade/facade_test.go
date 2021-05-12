package facade

import "testing"

var expect = "A module runing\nB module runing"


// TestFAcadeAPI...
func TestFacadeAPI(t *testing.T) {
    api :=  NewAPI()
    ret :=  api.Test()

    t.Logf("ok %s", ret)

    if ret != expect {
        t.Fatalf("expect %s, return %s", expect, ret)
    }
}

