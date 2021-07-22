[TOC]



`O(1)<O(logN)<O(N)<O(NlogN)<O(N2)<O(2N)<O(N!)`

# 前缀树

### `Trie`

`children[0] 对应小写字母 a,children[1] 对应小写字母 b;布尔字段isEnd，表示该节点是否为字符串的结尾`

```go
type Trie struct {
	children 	[26]*Trie
	isEnd 		bool
}
```

### 插入和查找

* 子节点存在。沿着指针移动到子节点，继续处理下一个字符。

* 子节点不存在。创建一个新的子节点，记录在 children 数组的对应位置上，然后沿着指针移动到子节点，继续搜索下一个字符。

  

* 子节点存在。沿着指针移动到子节点，继续搜索下一个字符。
* 子节点不存在。说明字典树中不包含该前缀，返回空指针。

```go
func (t *Trie) Insert(word string) {
    node := t
    
    for _, ch := range wod {
        ch -= 'a'	// 单词对应的索引
        if node.children[ch] = nil {	// 节点不存在
            node.children[ch] = &Trie{}	// 创建新的节点，记录在数组对应的位置上
        }
        node = node.children[ch]	// 继续判断下一个字符
    }
    node.isEnd = true	// 到达字符串尾，设置为true
}
```

```go
func (t *Trie) SearchPrefix(prefix string) *Trie {
    node := t
    
    for _, ch := range prefox {
        ch -= 'a'
        if node.children[ch] == nil {	// 节点不存在
            return nil 	// 返回空的
        }
        node = node.children[ch\] // 继续判断下一个字符
    }
    return node
}

// 查找
func (t *Trie) Search(word string) bool {
    node := t.SearchPrefix(word)
    return node != nil && node.isEnd  // 存在前缀节点并且isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
    return t.SearchPrefix(prefix) != nil	// 搜索到前缀末尾，说明存在该前缀
}
// 查找前缀是否存在
```

> 若搜索到了前缀的末尾，就说明字典树中存在该前缀。此外，若前缀末尾对应节点的`isEnd`为真，则说明字典树中存在该字符串。

# 双指针

## 160.相交链表

1. 当链表`headA`和`headB`都不为空时，才能相交
2. 创建两个指针，分别指向两个表得的头节点，依次遍历两个链表

* 每步操作需要同时更新指针 `pA` 和 `pB`
* 如果指针 `pA` 不为空，则将指针 `pA` 移动到下一个节点；如果指针 `pB` 不为空，则将指针 `pB` 移动到下一个节点。
* 如果指针 `pA` 为空，则将指针 `pA` 移到链表 `headB` 的头节点；如果指针 `pB` 为空，则将指针 `pB` 移到链表 `headA` 的头节点
* 当指针 `pA` 和 `pB` 指向同一个节点或都为空时，还会它们指向的节点或者 `null`

```go
type ListNode struct {
    Val		int
    Next	*ListNode 
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil {	// 两个链表都为空，不可能相交
        return nil 
    }
    
    pa, pb = headA, headB	// 创建两个指针，指向两个链表的头节点
    for pa != pb {			// 当指针相等时，返回他们相交的节点
        if pa == nil {		// 当指针遍历到尾节点时，指向链表b的头节点继续遍历
            pa = headB
        } else {
            pa = pa.Next	// 一直遍历到尾节点
        }
        
        if pb == nil {	   // 当指针遍历到尾节点时，指向链表a的头节点继续遍历
            pa == headA
        } else {
            pb = pb.Next	// 一直遍历到尾节点
        }
    }
    return pa
}
```



## [206. 反转链表](https://leetcode-cn.com/problems/reverse-linked-list/)

