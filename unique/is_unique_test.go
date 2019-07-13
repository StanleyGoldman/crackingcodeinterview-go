package unique

import "testing"

func TestUnique(t *testing.T) {
	assertUnique(t, isUnique, "ab")
	assertNotUnique(t, isUnique, "aa")
}

func TestUnique2(t *testing.T) {
	assertUnique(t, isUnique2, "ab")
	assertNotUnique(t, isUnique2, "aa")
}

func assertUnique(t *testing.T, f func(string) bool, value string) {
	if f(value) != true {
		t.Errorf("%q is not Unique", value)
	}
}

func assertNotUnique(t *testing.T, f func(string) bool, value string) {
	if f(value) != false {
		t.Errorf("%q is Unique", value)
	}
}
