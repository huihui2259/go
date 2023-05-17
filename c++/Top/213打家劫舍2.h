#include "c++/head.h"

class Solution
{
public:
    int rob(vector<int> &nums)
    {
        if (nums.size() == 1)
            return nums[0];
        auto nums1 = sub(nums, 0, nums.size() - 2);
        auto nums2 = sub(nums, 1, nums.size() - 1);
        return max(rob1(nums1), rob1(nums));
    }

    vector<int> sub(vector<int> &nums, int i, int j)
    {
        vector<int> res;
        for (int pos = i; pos <= j; pos++)
        {
            res.push_back(nums[pos]);
        }
        return res;
    }

    int rob1(vector<int> &nums)
    {
        int n = nums.size();
        if (n == 1)
            return nums[0];
        if (n == 2)
            return max(nums[0], nums[1]);
        vector<int> dp(n, 0);
        dp[0] = nums[0];
        dp[1] = max(nums[0], nums[1]);
        for (int i = 2; i < n; i++)
        {
            dp[i] = max(dp[i - 1], dp[i - 2] + nums[i]);
        }
        return dp[n - 1];
    }
};