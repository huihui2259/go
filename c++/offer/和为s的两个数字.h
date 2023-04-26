#include "head.h"

class Solution
{
public:
    vector<int> twoSum(vector<int> &nums, int target)
    {
        int i = 0, j = nums.size() - 1;
        while (i < j)
        {
            if (nums[i] + nums[j] < target)
                i++;
            else if (nums[i] + nums[j] > target)
                j--;
            else
            {
                return vector<int>{nums[i], nums[j]};
            }
        }
        return vector<int>();
    }
};