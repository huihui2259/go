#include "c++/head.h"

class Solution
{
public:
    vector<int> twoSum(vector<int> &nums, int target)
    {
        map<int, int> mp;
        for (int i = 0; i < nums.size(); i++)
        {
            if (mp.find(target - nums[i]) == mp.end())
            {
                mp[nums[i]] = i;
            }
            else
            {
                return vector<int>{i, mp[target - nums[i]]};
            }
        }
        return {0, 0};
    }
};