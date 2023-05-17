#include "c++/head.h"

/*
输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

*/
class Solution
{
public:
    vector<vector<int>> merge(vector<vector<int>> &intervals)
    {
        vector<vector<int>> res;
        if (intervals.size() == 0)
            return res;
        sort(intervals.begin(), intervals.end());
        res.push_back(intervals[0]);
        for (int i = 1; i < intervals.size(); i++)
        {
            auto &tmp = res.back();
            // 区间重叠，需要合并
            if (tmp[1] > intervals[i][0])
            {
                tmp[1] = max(tmp[1], intervals[i][1]);
            }
            else
            {
                res.push_back(intervals[i]);
            }
        }
        return res;
    }
};