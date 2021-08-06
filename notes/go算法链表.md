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

## [234. 回文链表](https://leetcode-cn.com/problems/palindrome-linked-list/)

* 将链表赋值到数组中
* 通过双指针，从两端向中间移动，比较值是否相等

```go
func isPalindrome(head *ListNode) bool {
   var array []int

   for ; head != nil; head = head.Next {
      array = append(array, head.Val)
   }

   n := len(array)
   for i, v := range array[:n/2] {
      if v != array[n-i-1] {
         return false
      }
   }

   return true
}
```

# 迭代

## [206. 反转链表](https://leetcode-cn.com/problems/reverse-linked-list/)

假设链表为 1→2→3→∅，我们想要把它改成∅←1←2←3。

在遍历链表时，将当前节点的 next 指针改为指向前一个节点。由于节点没有引用其前一个节点，因此必须事先存储其前一个节点。在更改引用之前，还需要存储后一个节点。最后返回新的头引用。

```go
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode	// 定义节点，用于存储节点的前一节点的引用
    curr := head
    
    for curr != nil {
        next := curr.Next	//  存储下一个节点
        curr.Next = prev	// 指向前一个节点
        prev = curr 	// 赋值上一个节点，继续下一个节点
        curr = next	// 继续下一个节点
    }
}
```



## [21. 合并两个有序链表](https://leetcode-cn.com/problems/merge-two-sorted-lists)

* 链表1，2都不为空
* 判断那个链表的头节点值更小，将较小的值添加到结果，之后将链表的节点向后移一位
* 需要一个虚拟头节点，操作一次，向后移动一个节点

循环终止时，链表1，2至多有一个是非空，由于是有序的，可以直接将链表合并在上面处理的结果后面

```go
type ListNode struct {
    Val		int
    Next 	*ListNode
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
    preHead := &ListNode{}	// 创建一个虚拟头节点
    
    prev := preHead		// 以指针的形式操作
    for l1 !=nil && l2 != nil {		// 链表1，2都不为空
        if l1.Val <= l2.Val {
            prev.Next = l1		// 当其中一个链表的值较小时，指向此链表中的节点
            l1 = l1.Next		// 向后移动一个节点
        } else {
            prev.Next = l2
            l2 = l2.Next
        }
        prev = prev.Next		// 指针要随着链表的节点移动而移动
    }
    // 结束后，之多还有一个非空链表，将它拼接在后面
    if l1 != nil {
        prev.Next = l1
    } else {
        prev.Next = l2
    }
   
	return preHead.Next
}

```



## [83. 删除排序链表中的重复元素](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/)

```go
type ListNode struct {
    Val		int
    Next 	*ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	prev := head		// 建立指针

	for prev.Next != nil {
		if prev.Next.Val == prev.Val {
			prev.Next = prev.Next.Next
		} else {
			prev = prev.Next
		}
	}
	return head
}
```

# 递归

## [21. 合并两个有序链表](https://leetcode-cn.com/problems/merge-two-sorted-lists)

* 首先的找到终止条件（一开始为空或者当一个链表为空时）
* 判断哪个节点的值小，然后递归的判断指向的下一个节点哪个较小

```
type ListNode struct {
    Val		int
    Next 	*ListNode
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	if l1 == nil {			// 终止条件
		return l2
	}
	if l2 == nil {
		return l1
	}
	
	if l1.Val < l2.Val {		// 判断哪个节点的值小
		l1.Next = mergeTwoLists(l1.Next, l2)	// 递归的去寻找l1下一个节点，与l2的节点那个小
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)	// 递归的去寻找l2下一个节点，与l1的节点那个小
		return l2
	}
}
```



