package trie

import (
	"strings"
)

const (
	Split = " "
	Repl  = "*"
)

/*
 *
 */
type TrieTree struct {
	root *TrieNode
}

func NewTrieTree() *TrieTree {
	tree := &TrieTree{
		root: NewTrieNode(),
	}
	return tree
}

/*
 *
 */
func (tree *TrieTree) Add(str string) {
	var node *TrieNode = tree.root
	str = strings.Trim(strings.ToLower(str), Split)
	for _, c := range str {
		node = node.add(c)
	}
	node.setLeaf(true)
}

/*
 *
 */
func (tree *TrieTree) Filter(str string) string {
	str = strings.Trim(strings.ToLower(str), Split)
	words := []rune(str)
	length := len(words)
	filter := make([]rune, length)
	for i := 0; i < length; {
		var node *TrieNode = tree.root

		for j := i; j < length; j++ {
			e := words[j]
			if v, ok := node.get(e); ok {
				// leaf, sensitive worlds
				if v.isLeaf() {
					// append
					filter = append(filter, []rune(strings.Repeat(Repl, j-i+1))...)
					i = j + 1
					break
				} else {
					node = v
				}
			} else {
				break
			}
		}

		// not sensitive head char
		filter = append(filter, words[i])
		i++
	}

	return string(filter)
}
