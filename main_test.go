package main

import (
	"reflect"
	"testing"
)

func TestRightSideView(t *testing.T) {
	root := tn(1, tn(2, nil, tn(5, nil, nil)), tn(3, nil, tn(4, nil, nil)))
	if got := RightSideView(root); !reflect.DeepEqual(got, []int{1, 3, 4}) {
		t.Errorf("RightSideView = %v, want [1 3 4]", got)
	}

	// A deep level reachable only through the LEFT subtree still shows.
	r2 := tn(1, tn(2, tn(4, tn(8, nil, nil), nil), nil), tn(3, nil, nil))
	if got := RightSideView(r2); !reflect.DeepEqual(got, []int{1, 3, 4, 8}) {
		t.Errorf("RightSideView = %v, want [1 3 4 8]", got)
	}

	if got := RightSideView(nil); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("nil = %v", got)
	}
	if got := RightSideView(tn(9, nil, nil)); !reflect.DeepEqual(got, []int{9}) {
		t.Errorf("single = %v", got)
	}
}
