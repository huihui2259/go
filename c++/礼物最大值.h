#include "head.h"

class Solution
{
public:
    int maxValue(vector<vector<int>> &grid)
    {
        vector<vector<int>> res(grid);
        for (int i = 1; i < res.size(); i++)
        {
            res[i][0] = res[i - 1][0] + res[i][0];
        }
        for (int i = 1; i < res[0].size(); i++)
        {
            res[0][i] = res[0][i - 1] + res[0][i];
        }
        for (int i = 1; i < res.size(); i++)
        {
            for (int j = 1; j < res[0].size(); j++)
            {
                res[i][j] = max(res[i - 1][j], res[i][j - 1]) + res[i][j];
            }
        }
        return res[res.size() - 1][res[0].size() - 1];
    }
};
