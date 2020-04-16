# 学习笔记

leetcode_20

import java.io.*;
import java.util.*;

//leetcode submit region begin(Prohibit modification and deletion)
class Solution {
    private HashMap<Character,Character> hashMap;
    public Solution(){
        this.hashMap = new HashMap<>();
        hashMap.put(')','(');
        hashMap.put(']','[');
        hashMap.put('}','{');
    }
    public boolean isValid(String s) {
        Stack<Character> stack = new Stack<>();
        for (int i=0;i<s.length();i++){
            char c = s.charAt(i);
            if (this.hashMap.containsKey(c)){
                char topElement = stack.empty() ? '#' : stack.pop();
                if (topElement != this.hashMap.get(c)){
                    return false;
                }
            }else {
                stack.push(c);
            }
        }
        return stack.isEmpty();
    }
}



leetcode_14
import org.apache.logging.log4j.util.*;

//leetcode submit region begin(Prohibit modification and deletion)
class Solution {
    public String longestCommonPrefix(String[] strs) {
        if (strs == null || strs.length ==0) return "";
        String one = strs[0];
        for (int j = 1;j<strs.length;j++){
            while (strs[j].indexOf(one) != 0){
                one = one.substring(0,one.length()-1);
                if (one.isEmpty()) return "";
            }
        }
        return  one;
    }
}
