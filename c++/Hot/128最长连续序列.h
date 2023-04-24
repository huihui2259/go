#include "c++/head.h"
/*
输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
[9,1,4,7,3,-1,0,5,8,-1,6]
7
*/
class Solution
{
public:
    int longestConsecutive(vector<int> &nums)
    {
        unordered_set<int> s;
        for (auto num : nums)
        {
            s.insert(num);
        }
        int max = 0;
        for (auto val : s)
        {
            if (s.count(val - 1) == 0)
            {
                int count = 1;
                val++;
                while (s.count(val))
                {
                    val++;
                    count++;
                }
                max = max > count ? max : count;
            }
        }
        return max;
    }
};
