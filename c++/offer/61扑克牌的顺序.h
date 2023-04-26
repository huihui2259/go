#include "head.h"

class Solution
{
public:
    bool isStraight(vector<int> &nums)
    {
        set<int> s;
        int minVal = 14, maxVal = 0, count = 0;
        for (auto num : nums)
        {
            minVal = minVal > num && num != 0 ? num : minVal;
            maxVal = maxVal < num ? num : maxVal;
            s.emplace(num);
            if (num == 0)
                count++;
        }
        if (s.size() + count < 5)
        {
            return false;
        }
        return maxVal - minVal + 1 <= s.size() + count;
    }
};