package _0208

type Trie struct {
	val  byte
	sons [26]*Trie
	end  bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this
	for i := 0; i < len(word); i++ {
		j := word[i] - 'a'
		if node.sons[j] == nil {
			node.sons[j] = &Trie{val: word[i]}
		}
		node = node.sons[j]
	}
	node.end = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this
	for i := 0; i < len(word); i++ {
		j := word[i] - 'a'
		if node.sons[j] == nil {
			return false
		}
		node = node.sons[j]
	}
	if !node.end {
		return false
	}
	return true
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for i := 0; i < len(prefix); i++ {
		j := prefix[i] - 'a'
		if node.sons[j] == nil {
			return false
		}
		node = node.sons[j]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
