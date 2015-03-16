package main

import (
	"fmt"
	"github.com/ShionRyuu/goTrie/trie"
)

func main() {
	fmt.Println("Hello goTrie!")
	tree := trie.NewTrieTree()
	tree.Add("sex")
	tree.Add("共产党")
	str := "Sexy girls are reading <<共产党宣言>>!"
	fmt.Printf("%v -> %v\n", str, tree.Filter(str))
}
