package main

import "fmt"

/**
 * 字母异位词
 * https://leetcode-cn.com/problems/valid-anagram/description/
 */

/**
 * 1.用一个HashMap
 * 2.编译每个字符，k:字符，v:数量，有重复字符，v+=1，
 * 3.遍历另一个字符串，遍历到有的字符，v-1
 * 4.遍历结束后看map是不是空的
 */
func isAnagram(s string, t string) bool {

	if len(t) != len(s) {
		return false
	}

	strMap := make(map[string]int)

	for _, v := range []rune(s) {
		if n, ok := strMap[string(v)]; ok {
			strMap[string(v)] = n + 1
		} else {
			strMap[string(v)] = 1
		}
	}

	for _, v := range []rune(t) {
		if n, ok := strMap[string(v)]; ok {
			if n > 1 {
				strMap[string(v)] = n - 1
			} else {
				delete(strMap, string(v))
			}
		}
	}

	if len(strMap) == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram(s, t))
}
