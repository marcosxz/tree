package tree

import "testing"

func TestTrie(t *testing.T) {
	trie := NewTrie()
	trie.Insert("apple")
	trie.Insert("Apple")
	trie.Insert("orange")
	trie.Insert("Orange")
	trie.Insert("pineapple")
	trie.Insert("Pineapple")
	trie.Insert("banana")
	trie.Insert("Banana")
	t.Log("--------------")
	t.Log("has apple:", trie.Search("apple"))
	t.Log("has Apple:", trie.Search("Apple"))
	t.Log("has orange:", trie.Search("orange"))
	t.Log("has Orange:", trie.Search("Orange"))
	t.Log("has pineapple:", trie.Search("pineapple"))
	t.Log("has Pineapple:", trie.Search("Pineapple"))
	t.Log("has banana:", trie.Search("banana"))
	t.Log("has Banana:", trie.Search("Banana"))
	t.Log("--------------")
	t.Log("has appleX:", trie.Search("appleX"))
	t.Log("has AppleX:", trie.Search("AppleX"))
	t.Log("has orangeX:", trie.Search("orangeX"))
	t.Log("has OrangeX:", trie.Search("OrangeX"))
	t.Log("has pineappleX:", trie.Search("pineappleX"))
	t.Log("has PineappleX:", trie.Search("PineappleX"))
	t.Log("has bananaX:", trie.Search("bananaX"))
	t.Log("has BananaX:", trie.Search("BananaX"))
	t.Log("--------------")
	t.Log("has app:", trie.Search("app"))
	t.Log("has App:", trie.Search("App"))
	t.Log("has ora:", trie.Search("ora"))
	t.Log("has Ora:", trie.Search("Ora"))
	t.Log("has pine:", trie.Search("pine"))
	t.Log("has Pine:", trie.Search("Pine"))
	t.Log("has ban:", trie.Search("ban"))
	t.Log("has Ban:", trie.Search("Ban"))
	t.Log("--------------")
	t.Log("start with app:", trie.StartWith("app"))
	t.Log("start with App:", trie.StartWith("App"))
	t.Log("start with apple:", trie.StartWith("apple"))
	t.Log("start with Apple:", trie.StartWith("Apple"))
	t.Log("start with ora:", trie.StartWith("ora"))
	t.Log("start with Ora:", trie.StartWith("Ora"))
	t.Log("start with orange:", trie.StartWith("orange"))
	t.Log("start with Orange:", trie.StartWith("Orange"))
	t.Log("start with pine:", trie.StartWith("pine"))
	t.Log("start with Pine:", trie.StartWith("Pine"))
	t.Log("start with pineapple:", trie.StartWith("pineapple"))
	t.Log("start with Pineapple:", trie.StartWith("Pineapple"))
	t.Log("start with ban:", trie.StartWith("ban"))
	t.Log("start with Ban:", trie.StartWith("Ban"))
	t.Log("start with banana:", trie.StartWith("banana"))
	t.Log("start with Banana:", trie.StartWith("Banana"))
	t.Log("--------------")
	t.Log("start with appX:", trie.Search("appX"))
	t.Log("start with AppX:", trie.Search("AppX"))
	t.Log("start with oraX:", trie.Search("oraX"))
	t.Log("start with OraX:", trie.Search("OraX"))
	t.Log("start with pineX:", trie.Search("pineX"))
	t.Log("start with PineX:", trie.Search("PineX"))
	t.Log("start with banX:", trie.Search("banX"))
	t.Log("start with BanX:", trie.Search("BanX"))
}
