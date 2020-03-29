package liuzhan;

import java.util.HashMap;
import java.util.Map;

/**
 * message
 * Created by liuzhan
 * on 2020/3/29 9:33 下午
 */
public class TowSum {

    /**
     * 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
     *
     * 你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
     *
     * 示例:
     *
     * 给定 nums = [2, 7, 11, 15], target = 9
     *
     * 因为 nums[0] + nums[1] = 2 + 7 = 9
     * 所以返回 [0, 1]
     *
     * 来源：力扣（LeetCode）
     * 链接：https://leetcode-cn.com/problems/two-sum
     */

    public int[] twoSum(int[] nums, int target) {
        //利用 哈希表实现
        // 穷举数组元素: a = target - num[i]
        Map<Integer,Integer> numsMap = new HashMap<>(nums.length);
        for (int i = 0; i < nums.length ; i ++){
            numsMap.put(nums[i],i);
        }
        for (int i = 0 ; i < nums.length ; i ++){
            int targetNum = target - nums[i];
            if(numsMap.containsKey(targetNum) && numsMap.get(targetNum)!= i){
                return new int[]{i,numsMap.get(targetNum)};
            }
        }
        throw new IllegalArgumentException("没有找到符合两数字和的数组");
    }

}
