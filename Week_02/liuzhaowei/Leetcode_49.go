package main

import (
	"fmt"
	"sort"
)

/**
 * 字母异位词分组
 * https://leetcode-cn.com/problems/group-anagrams/
 */
/**
 * 1. 对每个字符串进行排序
 * 2. 分组，记录下标
 */
func groupAnagrams(strs []string) [][]string {
	strMap := make(map[string][]string)
	for _, s := range strs {
		b := []byte(s)
		key := string(sortByte(b))

		if v, ok := strMap[key]; ok {
			v = append(v, s)
			strMap[key] = v
		} else {
			vs := make([]string, 0)
			vs = append(vs, s)
			strMap[key] = vs
		}
	}

	back := make([][]string, 0)

	for _, v := range strMap {
		back = append(back, v)
	}
	return back
}

func sortByte(b []byte) []byte {
	a := make([]byte, len(b), cap(b))

	it := make([]int, len(b), cap(b))
	for i, bt := range b {
		it[i] = int(bt)
	}

	sort.Ints(it)

	for i, v := range it {
		a[i] = byte(v)
	}

	return a
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}
