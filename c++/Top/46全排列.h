#include "c++/head.h"

class Solution
{
public:
    vector<vector<int>> res;
    vector<int> tmp;
    vector<bool> pos;
    vector<vector<int>> permute(vector<int> &nums)
    {
        for (int i = 0; i < nums.size(); i++)
        {
            pos.push_back(false);
        }
        dfs(nums);
        return res;
    }
    void dfs(vector<int> &nums)
    {
        if (tmp.size() == nums.size())
        {
            res.push_back(tmp);
            return;
        }
        for (int i = 0; i < nums.size(); i++)
        {
            if (pos[i])
                continue;
            pos[i] = true;
            tmp.push_back(nums[i]);
            dfs(nums);
            tmp.pop_back();
            pos[i] = false;
        }
    }
};