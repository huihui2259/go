#include "c++/head.h"

class Solution
{
public:
    int maxProfit(vector<int> &prices)
    {
        int n = prices.size();
        vector<int> dp(n, 0);
        for (int i = 1; i < n; i++)
        {
            dp[i] = (dp[i - 1] + prices[i] - prices[i - 1]) > 0 ? (dp[i - 1] + prices[i] - prices[i - 1]) : 0;
        }
        return *max_element(dp.begin(), dp.end());
    }
};