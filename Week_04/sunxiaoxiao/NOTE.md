# 学习笔记

leetcode_21

class Solution {
//    public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
//        if (l1 == null){
//            return l2;
//        }else if (l2 == null){
//            return l1;
//        }else if (l1.val < l2.val){
//            l1.next = mergeTwoLists(l1.next,l2);
//            return l1;
//        }else {
//            l2.next = mergeTwoLists(l2.next,l1);
//            return l2;
//        }
//    }

    public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
        ListNode prehead = new ListNode(-1);
        ListNode listNode = prehead;
        while (l1 != null && l2 != null) {
            if (l1.val <= l2.val){
                listNode.next = l1;
                l1 = l1.next;
            }else {
                listNode.next = l2;
                l2 = l2.next;
            }
            listNode = listNode.next;
        }
        listNode.next = l1 == null ? l2 : l1;
        return prehead.next;
    }
}



leetcode_26

class Solution {
    public int removeDuplicates(int[] nums) {
        if (nums.length == 0) return 0;
        int j = 0;
        for (int i=1;i<nums.length; i++){
            if (nums[j] != nums[i]){
                j++;
                nums[j] = nums[i];
            }
        }
        return j+1;
    }
}
