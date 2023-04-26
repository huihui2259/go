#include "head.h"
// 输入: nums = [-2,1,-3,4,-1,2,1,-5,4]
// 输出: 6
// 解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
class Solution
{
public:
    int maxSubArray(vector<int> &nums)
    {
        int n = nums.size();
        int res = nums[0];
        vector<int> dp(n, nums[0]);
        for (int i = 1; i < nums.size(); i++)
        {
            if (dp[i - 1] <= 0)
                dp[i] = nums[i];
            else
                dp[i] = dp[i - 1] + nums[i];
            res = max(res, dp[i]);
        }
        return res;
    }
};