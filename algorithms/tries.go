package algorithms

type TrieNode struct {
	Value    rune
	Children []TrieNode
}

func (t *TrieNode) SearchTrie(search_term rune) *TrieNode {
	// Given a pointer to a TrieNode and a rune to search for, find the TrieNode that contains the rune in its val
	if t.Value == search_term {
		return t
	}
}

func (t *TrieNode) GenerateTrie(words []string) *TrieNode {
	// Given an array of words, fit it into a trie
	root = t
	for _, word := range words {
		for i := 0; i < len(word); i++ {
			new_trie_node := TrieNode{word[i], make([]TrieNode)}
		}
	}
}

func SearchWords(words []string, target string) {
	// Given an array of words, search it efficiently
	root := GenerateTrie(words)
}
