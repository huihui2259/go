#include "c++/head.h"

class Solution
{
public:
    int maximalSquare(vector<vector<char>> &matrix)
    {
        int m = matrix.size(), n = matrix[0].size();
        int res = 0;
        vector<vector<int>> dp(m, vector<int>(0, 1));
        for (int i = 0; i < m; i++)
        {
            dp[i][0] = matrix[i][0] - '0';
        }
        for (int i = 0; i < n; i++)
        {
            dp[0][i] = matrix[0][i] - '0';
        }
        for (int i = 1; i < m; i++)
        {
            for (int j = 1; j < n; j++)
            {
                if (dp[i][j] == '1')
                {
                    dp[i][j] = min(dp[i - 1][j - 1], min(dp[i - 1][j], dp[i][j - 1])) + 1;
                }
                res = max(dp[i][j], res);
            }
        }
        return res * res;
    }
};