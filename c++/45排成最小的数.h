#include "head.h"

class Solution
{
public:
    string minNumber(vector<int> &nums)
    {
        string s;
        sort(nums.begin(), nums.end(), [](int a, int b) {
            auto s1 = to_string(a);
            auto s2 = to_string(b);
            return s1 + s2 < s2 + s1;
        });
        for (auto num : nums)
        {
            s += to_string(num);
        }
        return s;
    }
};