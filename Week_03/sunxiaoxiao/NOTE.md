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
