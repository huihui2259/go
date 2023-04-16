#include "head.h"

class Solution
{
public:
    int missingNumber(vector<int> &nums)
    {
        int left = 0, right = nums.size() - 1;
        // 使用等号，以兼容一个数字的情况
        while (left <= right)
        {
            int mid = left + (right - left) / 2;
            if (nums[mid] == mid)
            {
                left = mid + 1;
            }
            else
            {
                right = right - 1;
            }
        }
        return left;
    }
};