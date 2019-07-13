package palindrome

import "testing"

func TestPalindrome(t *testing.T) {
	testIsPalindrome(t, isPalindrome, "cab", "abc", true)
	testIsPalindrome(t, isPalindrome, "ddd", "abc", false)
	testIsPalindrome(t, isPalindrome, "aab", "abc", false)
	testIsPalindrome(t, isPalindrome, "aabc", "abc", false)
}

func TestPalindrome2(t *testing.T) {
	testIsPalindrome(t, isPalindrome2, "cab", "abc", true)
	testIsPalindrome(t, isPalindrome2, "ddd", "abc", false)
	testIsPalindrome(t, isPalindrome2, "aab", "abc", false)
	testIsPalindrome(t, isPalindrome2, "aabc", "abc", false)
}

func TestPalindrome3(t *testing.T) {
	testIsPalindrome(t, isPalindrome3, "cab", "abc", true)
	testIsPalindrome(t, isPalindrome3, "ddd", "abc", false)
	testIsPalindrome(t, isPalindrome3, "aab", "abc", false)
	testIsPalindrome(t, isPalindrome3, "aabc", "abc", false)
}

func testIsPalindrome(t *testing.T, f func(string, string) bool, n string, h string, e bool) {
	if f(n, h) != e {
		if !e {
			t.Errorf("%q is not a palindrome of %q", n, h)
		} else {
			t.Errorf("%q is a palindrome of %q", n, h)
		}
	}
}
