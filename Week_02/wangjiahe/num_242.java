//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
//
// 示例 1:
//
// 输入: s = "anagram", t = "nagaram"
//输出: true
//
//
// 示例 2:
//
// 输入: s = "rat", t = "car"
//输出: false
//
// 说明:
//你可以假设字符串只包含小写字母。
//
// 进阶:
//如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
// Related Topics 排序 哈希表


import java.util.HashMap;
import java.util.Map;

//leetcode submit region begin(Prohibit modification and deletion)
class Solution {
    public boolean isAnagram(String s, String t) {
        if (s == null || t == null || s.length() != t.length()) {
            return false;
        }
        int[] array1 = new int[26]; // 字符串s的数组
        int[] array2 = new int[26]; // 字符串t的数组
        for (char char1 : s.toCharArray()) {
            array1[char1 - 'a']++;
        }
        for (char char2 : t.toCharArray()) {
            array2[char2 - 'a']++;
        }
        return Arrays.equals(array1, array2);
    }
}
//leetcode submit region end(Prohibit modification and deletion)
