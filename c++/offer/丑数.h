#include "head.h"

class Solution
{
public:
    int nthUglyNumber(int n)
    {
        vector<int> dp(n, 1);
        int a = 1, b = 1, c = 1;
        for (int i = 1; i < n; i++)
        {
            int ugly = min(dp[c] * 5, min(dp[a] * 2, dp[b] * 3));
            if (ugly == dp[a] * 2)
                a++;
            if (ugly == dp[b] * 3)
                b++;
            if (ugly == dp[c] * 5)
                c++;
            dp[i] = ugly;
        }
        return dp.back();
    }
};