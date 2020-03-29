//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
//
//
//
// 上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢 Mar
//cos 贡献此图。
//
// 示例:
//
// 输入: [0,1,0,2,1,0,1,3,2,1,2,1]
//输出: 6
// Related Topics 栈 数组 双指针


//leetcode submit region begin(Prohibit modification and deletion)
class Solution {

    public int trap(int[] height) {
        /**
         * 按列计算每列能装多少水
         */
        // 前一列与最后一列是装不下水的
        if (height.length <= 1) {
            return 0;
        }
        // 遍历每一列的左边与右边，如果有比当前列更高的列，则可以判断当前列是否可以装水，装多少水
        int sum = 0;
        for (int i = 1; i < height.length - 1; i++) {
            int now = height[i];
            // 遍历左边，找出最高的一段
            int leftMax = now;
            for (int j = i - 1; j >= 0; j--) {
                int left = height[j];
                if (left > leftMax) {
                    leftMax = left;
                }
            }
            // 遍历右边，找出最高的一段
            int rightMax = now;
            for (int b = i + 1; b < height.length; b++) {
                int right = height[b];
                if (right > rightMax) {
                    rightMax = right;
                }
            }
            // 当左右最长的部分都大于当前长度，才能存下水，存下的水的量为较短那根与当前的差
            if (rightMax > now && leftMax > now) {
                int min = Math.min(rightMax, leftMax);
                sum = sum + min - now;
            }
        }
        return sum;
    }
}
//leetcode submit region end(Prohibit modification and deletion)
