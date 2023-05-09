#include "c++/head.h"

/*
输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
*/

class Solution
{
public:
    vector<vector<int>> threeSum(vector<int> &nums)
    {
        sort(nums.begin(), nums.end());
        int n = nums.size();
        vector<vector<int>> res;
        for (int i = 0; i < n - 2; i++)
        {
            vector<int> tmp(3, 0);
            if (nums[i] > 0)
                break;
            if (i >= 1 && nums[i] == nums[i - 1])
                continue;
            int target = 0 - nums[i];
            int j = i + 1, k = n - 1;
            while (j < k)
            {

                if (nums[j] + nums[k] < target)
                    j++;
                else if (nums[j] + nums[k] > target)
                    k--;
                else
                {
                    tmp[0] = nums[i], tmp[1] = nums[j], tmp[2] = nums[k];
                    res.push_back(tmp);
                    while (j < k && nums[j] == nums[j + 1])
                    {
                        j++;
                    }
                    while (j < k && nums[k] == nums[k - 1])
                    {
                        k--;
                    }
                    j++;
                    k--;
                }
            }
        }
        return res;
    }
};