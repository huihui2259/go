#include "c++/head.h"

/*
输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
*/
class Solution
{
public:
    int maxSubArray(vector<int> &nums)
    {
        vector<int> dp(nums);
        for (int i = 1; i < nums.size(); i++)
        {
            dp[i] = dp[i - 1] < 0 ? nums[i] : (dp[i - 1] + nums[i]);
        }
        return *max_element(dp.begin(), dp.end());
    }
};