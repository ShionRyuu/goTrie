package trie

import ()

type TrieNode struct {
	child map[rune]*TrieNode // subtree
	leaf  bool               // end of a word
}

/*
 * Create new TrieNode
 */
func NewTrieNode() *TrieNode {
	node := &TrieNode{
		child: make(map[rune]*TrieNode, 100),
		leaf:  false,
	}
	return node
}

/*
 * Find subtree with the node as the root
 */
func (node *TrieNode) get(e rune) (*TrieNode, bool) {
	if v, ok := node.child[e]; ok {
		return v, true
	}
	return nil, false
}

/*
 * Add child node
 */
func (node *TrieNode) add(e rune) *TrieNode {
	return node.addRune(e, false)
}

/*
 * Add leaf node
 */
func (node *TrieNode) addRune(e rune, leaf bool) *TrieNode {
	var childNode *TrieNode
	if v, ok := node.child[e]; ok {
		childNode = v
	} else {
		childNode = NewTrieNode()
	}
	childNode.leaf = childNode.leaf || leaf
	node.child[e] = childNode
	return childNode
}

func (node *TrieNode) setLeaf(leaf bool) {
	node.leaf = leaf
}

func (node *TrieNode) isLeaf() bool {
	return node.leaf
}
