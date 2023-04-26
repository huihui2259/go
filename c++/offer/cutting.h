#include "head.h"

class Solution
{
public:
    int cuttingRope(int n)
    {
        vector<int> dp(n, 1);
        for (int i = 0; i < n; i++)
        {
            for (int j = 0; j < i; j++)
            {
                int a = max(dp[j], j);
                int b = max(dp[i - j], i - j);
                dp[i] = max(dp[i], dp[i - j] * dp[j]);
            }
        }
        return dp[n - 1];
    }
};