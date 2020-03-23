package liuzhan;

import java.util.Arrays;
import java.util.NoSuchElementException;

/**
 * 算法训练营 第一周
 * commitBy: liuzhan
 * Date: 2020-03-21
 */

public class Leetcode_week_1 {

    /**
     删除排序数组中的重复项
     给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
     不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
     来源：力扣（LeetCode）
     链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array
     */
    public static int removeDuplicates(int[] nums) {
        if(nums.length == 0){
            return 0;
        }
        //双指针i = 0 , j = i + 1
        int i = 0;
        for(int j = 0 ; j< nums.length;j++){
            //比较连续两个位置是否相等；将不重复的数字移到数组左侧；
            if(nums[i] != nums[j]) {
                i ++;
                nums[i] = nums[j];
            }
        }
        return i+1;
    }

    /**
     * 给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数
     * 输入: [1,2,3,4,5,6,7] 和 k = 3
     * 输出: [5,6,7,1,2,3,4]
     * 解释:
     * 向右旋转 1 步: [7,1,2,3,4,5,6]
     * 向右旋转 2 步: [6,7,1,2,3,4,5]
     * 向右旋转 3 步: [5,6,7,1,2,3,4]
     *
     * 来源：力扣（LeetCode）
     * 链接：https://leetcode-cn.com/problems/rotate-array
     */
    public void rotate(int[] nums, int k) {
        for(int i = 0 ; i < k ; i++){
            //把数组翻转K次，每次把最后一个元素，放在一个位置
            int last = nums[nums.length - 1];
            for(int j = 0 ; j< nums.length ; j ++){
                int temp  = nums[j];
                nums[j] = last;
                last = temp;
            }
        }
    }

    /**
     * 将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 
     *
     * 示例:
     * 输入：1->2->4, 1->3->4
     * 输出：1->1->2->3->4->4
     *
     * 来源：力扣（LeetCode）
     * 链接：https://leetcode-cn.com/problems/merge-two-sorted-lists
     */

    public class ListNode {
          int val;
          ListNode next;
          ListNode(int x) {
              val = x;
          }
    }
    public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
        //递归的第一步：终止条件，l1 == null,则返回l2. l2 == null,则返回l1
        if (l1 == null) {
            return l2;
        }
        if (l2 == null) {
            return l1;
        }
        if (l1.val < l2.val) {
            //如果 l1 的 val 值更小，则将 l1.next 与排序好的链表头相接
            l1.next = mergeTwoLists(l1.next, l2);
            return l1;
        } else {
            //如果 l2 的 val 值更小，则将 l2.next 与排序好的链表头相接
            l2.next = mergeTwoLists(l1, l2.next);
            return l2;
        }
    }

    /**
     * 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
     *
     * 你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
     *
     * 示例:
     *
     * 给定 nums = [2, 7, 11, 15], target = 9
     * 因为 nums[0] + nums[1] = 2 + 7 = 9
     * 所以返回 [0, 1]
     *
     * 来源：力扣（LeetCode）
     * 链接：https://leetcode-cn.com/problems/two-sum
     *
     * @param nums
     * @param target
     * @return
     */

    public int[] twoSum(int[] nums, int target) {
        //暴力破解
//        for(int i = 0; i< nums.length;i++){
//            for (int j = i+1; j< nums.length;j++){
//                if((nums[i] +nums[j]) == target){
//                    return new int[]{i,j};
//                }
//            }
//        }

        throw new IllegalArgumentException("No two sum solution");
    }


    /**
     * 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
     *
     * 示例:
     *
     * 输入: [0,1,0,3,12]
     * 输出: [1,3,12,0,0]
     * 说明:
     *
     * 必须在原数组上操作，不能拷贝额外的数组。
     * 尽量减少操作次数。
     *
     * 来源：力扣（LeetCode）
     * 链接：https://leetcode-cn.com/problems/move-zeroes
     */
    public static void moveZeroes(int[] nums) {
        int zeroIndex = 0;
        for (int i = 0 ; i < nums.length;i++){
            if(nums[i] != 0){
                nums[zeroIndex] = nums[i];
                zeroIndex ++;
            }
        }
        for (int i = zeroIndex;i< nums.length;i++){
            nums[i] = 0;
        }
//        for (int i = 0 ; i < nums.length;i++){
//            System.out.println(nums[i]);
//        }
    }

    /**
     * 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
     *
     * 最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
     *
     * 你可以假设除了整数 0 之外，这个整数不会以零开头。
     *
     * 示例 1:
     *
     * 输入: [1,2,3]
     * 输出: [1,2,4]
     * 解释: 输入数组表示数字 123。
     *
     * 来源：力扣（LeetCode）
     * 链接：https://leetcode-cn.com/problems/plus-one
     */

    public static int[] plusOne(int[] digits) {
        //从末尾往前开始 循环加1；遇9 则开始进位；
        for (int i = digits.length-1 ; i >= 0; i--){
            if(digits[i] == 9){
                digits[i] = 0;
                continue;
            }
            if(digits[i] != 9){
                //直到不需要进位 循环结束，在该位加1；
                digits[i]++;
                return digits;
            }
        }
        //考虑特殊情况，9，99，999 需要在首位加1
        int[] nums2 = new int[digits.length+1];
        nums2[0]=1;
        return nums2;
    }

    public static void main(String args[]){
        int[] nums = new int[]{1,1,0,2,3};
//        System.out.println(removeDuplicates(nums));
//        moveZeroes(nums);
        plusOne(nums);
    }
}