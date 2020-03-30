# 学习笔记
leetCode_2
class Solution {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode dummyHead = new ListNode(0);
        ListNode p = l1, q = l2 ,curr = dummyHead;
        int carry = 0;
        while (p != null || q != null){
            int x = (p != null) ? p.val : 0;
            int y = (q != null) ? q.val : 0;
            int sum = carry + x + y;
            carry = sum / 10;
            curr.next = new ListNode(sum % 10);
            curr = curr.next;
            if(p != null) p = p.next;
            if(q != null) q = q.next;
        }
        if (carry > 0){
            curr.next = new ListNode(carry);
        }
        return dummyHead.next;
    }
}


leetCode_13
class Solution {
    public int romanToInt(String s) {
        int sum = 0;
        int value = getValue(s.charAt(0));
        for (int i = 1;i< s.length();i++){
            int num = getValue(s.charAt(i));
            if (value < num){
                sum -= value;
            }else {
                sum += value;
            }
            value = num;
        }
        sum += value;
        return sum;
    }

    private int getValue(char str){
        switch (str){
            case 'I' : return 1;
            case 'V' : return 5;
            case 'X': return 10;
            case 'L': return 50;
            case 'C': return 100;
            case 'D': return 500;
            case 'M': return 1000;
            default: return 0;
        }
    }
}
