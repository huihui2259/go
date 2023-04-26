#include "c++/head.h"

/*
输入: nums = [0,1,0,3,12]
输出: [1,3,12,0,0]
*/

class Solution
{
public:
    void moveZeroes(vector<int> &nums)
    {
        int n = nums.size(), left = 0, right = 0;
        while (right < n)
        {
            if (nums[right])
            {
                swap(nums[left], nums[right]);
                left++;
            }
            right++;
        }
    }
};
