#include "head.h"

class Solution
{
public:
    vector<int> singleNumbers(vector<int> &nums)
    {
        int res = 0;
        for (auto num : nums)
        {
            res ^= num;
        }
        int mask = 1;
        while ((res & mask) == 0)
        {
            mask <<= 1;
        }
        int a = 0, b = 0;
        for (auto num : nums)
        {
            if ((num & mask) == 0)
            {
                a ^= num;
                continue;
            }
            b ^= num;
        }
        return vector<int>{a, b};
    }
};