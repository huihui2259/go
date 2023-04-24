#include "c++/head.h"

/*
输入: nums = [2,3,-2,4]
输出: 6
解释: 子数组 [2,3] 有最大乘积 6。

输入: nums = [-2,0,-1]
输出: 0
解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
*/

class Solution
{
public:
    int maxProduct(vector<int> &nums)
    {
        int n = nums.size();
        int res = nums[0];
        vector<int> maxDp(n, nums[0]);
        vector<int> minDp(n, nums[0]);
        for (int i = 1; i < n; i++)
        {
            maxDp[i] = max(nums[i], max(maxDp[i - 1] * nums[i], minDp[i - 1] * nums[i]));
            minDp[i] = min(nums[i], min(maxDp[i - 1] * nums[i], minDp[i - 1] * nums[i]));
            res = max(maxDp[i], res);
        }
        return res;
    }
};