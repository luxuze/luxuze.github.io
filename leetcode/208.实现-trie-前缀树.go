package leetcode

// 15/15 cases passed (68 ms)
// Your runtime beats 70.09 % of golang submissions
// Your memory usage beats 30.67 % of golang submissions (17.9 MB)
/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 *
 * https://leetcode-cn.com/problems/implement-trie-prefix-tree/description/
 *
 * algorithms
 * Medium (69.20%)
 * Likes:    468
 * Dislikes: 0
 * Total Accepted:    62.5K
 * Total Submissions: 90.3K
 * Testcase Example:  '["Trie","insert","search","search","startsWith","insert","search"]\n' +
  '[[],["apple"],["apple"],["app"],["app"],["app"],["app"]]'
 *
 * 实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。
 *
 * 示例:
 *
 * Trie trie = new Trie();
 *
 * trie.insert("apple");
 * trie.search("apple");   // 返回 true
 * trie.search("app");     // 返回 false
 * trie.startsWith("app"); // 返回 true
 * trie.insert("app");
 * trie.search("app");     // 返回 true
 *
 * 说明:
 *
 *
 * 你可以假设所有的输入都是由小写字母 a-z 构成的。
 * 保证所有输入均为非空字符串。
 *
 *
*/

// @lc code=start
type Trie struct {
	Next  [26]*Trie
	IsEnd bool
}

/** Initialize your data structure here. */
func Trie_Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	cursor := this
	for i := range word {
		idx := word[i] - 'a'
		if cursor.Next[idx] == nil {
			cursor.Next[idx] = &Trie{}
		}
		cursor = cursor.Next[idx]
	}
	cursor.IsEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	cursor := this
	for i := range word {
		idx := word[i] - 'a'
		if cursor.Next[idx] == nil {
			return false
		}
		cursor = cursor.Next[idx]
	}
	if cursor.IsEnd {
		return true
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	cursor := this
	for i := range prefix {
		idx := prefix[i] - 'a'
		if cursor.Next[idx] == nil {
			return false
		}
		cursor = cursor.Next[idx]
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
// @lc code=end
