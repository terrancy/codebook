package others

import (
	"math/rand"
)

// SkipList
// @Title: LC1206.设计跳表
// @Description: 设计一个跳表，支持 add/search/erase，时间复杂度 O(log n)
// @Link: https://leetcode.cn/problems/design-skiplist/
//
//	跳表示意 (MAX_LEVEL=4):
//	L4:  head ──────────────────────────> nil
//	L3:  head ────────────── 9 ────────> nil
//	L2:  head ──── 3 ────── 9 ── 12 ───> nil
//	L1:  head ── 1 ─ 3 ── 5 ─ 9 ─ 12 ──> nil
//	L0:  head ─ 1 ─ 2 ─ 3 ─ 5 ─ 7 ─ 9 ─ 12 ─> nil

const (
	skipListMaxLevel    = 32
	skipListProbability = 0.25
)

type skipListNode struct {
	val int
	next []*skipListNode
}

type SkipList struct {
	head  *skipListNode
	level int
}

func SkipListConstructor() SkipList {
	return SkipList{
		head: &skipListNode{
			val:  -1,
			next: make([]*skipListNode, skipListMaxLevel),
		},
		level: 1,
	}
}

func (sl *SkipList) randomLevel() int {
	lv := 1
	for lv < skipListMaxLevel && rand.Float64() < skipListProbability {
		lv++
	}
	return lv
}

func (sl *SkipList) Search(target int) bool {
	cur := sl.head
	for i := sl.level - 1; i >= 0; i-- {
		for cur.next[i] != nil && cur.next[i].val < target {
			cur = cur.next[i]
		}
	}
	cur = cur.next[0]
	return cur != nil && cur.val == target
}

func (sl *SkipList) Add(num int) {
	update := make([]*skipListNode, skipListMaxLevel)
	cur := sl.head
	for i := sl.level - 1; i >= 0; i-- {
		for cur.next[i] != nil && cur.next[i].val < num {
			cur = cur.next[i]
		}
		update[i] = cur
	}

	newLv := sl.randomLevel()
	if newLv > sl.level {
		for i := sl.level; i < newLv; i++ {
			update[i] = sl.head
		}
		sl.level = newLv
	}

	newNode := &skipListNode{
		val:  num,
		next: make([]*skipListNode, newLv),
	}
	for i := 0; i < newLv; i++ {
		newNode.next[i] = update[i].next[i]
		update[i].next[i] = newNode
	}
}

func (sl *SkipList) Erase(num int) bool {
	update := make([]*skipListNode, skipListMaxLevel)
	cur := sl.head
	for i := sl.level - 1; i >= 0; i-- {
		for cur.next[i] != nil && cur.next[i].val < num {
			cur = cur.next[i]
		}
		update[i] = cur
	}

	cur = cur.next[0]
	if cur == nil || cur.val != num {
		return false
	}

	for i := 0; i < sl.level; i++ {
		if update[i].next[i] != cur {
			break
		}
		update[i].next[i] = cur.next[i]
	}

	for sl.level > 1 && sl.head.next[sl.level-1] == nil {
		sl.level--
	}
	return true
}