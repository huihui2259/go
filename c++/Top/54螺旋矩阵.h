#include "c++/head.h"

class Solution
{
public:
    vector<int> spiralOrder(vector<vector<int>> &matrix)
    {
        int m = matrix.size(), n = matrix[0].size();
        int l = 0, r = n - 1, up = 0, down = m - 1;
        vector<int> res;
        while (true)
        {
            for (int i = l; i <= r; i++)
            {
                res.push_back(matrix[up][i]);
            }
            if (++up > down)
            {
                break;
            }
            for (int i = up; i <= down; i++)
            {
                res.push_back(matrix[i][r]);
            }
            if (--r < l)
            {
                break;
            }
            for (int i = r; i >= l; i--)
            {
                res.push_back(matrix[down][i]);
            }
            if (--down < up)
            {
                break;
            }
            for (int i = down; i >= up; i--)
            {
                res.push_back(matrix[i][l]);
            }
            if (++l > r)
            {
                break;
            }
        }
        return res;
    }
};