#include "head.h"

/*
输入: [7,1,5,3,6,4]
输出: 5
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。
*/
class Solution
{
public:
    int maxProfit(vector<int> &prices)
    {
        int n = prices.size();
        // 当天出售股票得到的利润
        vector<int> dp(n, 0);
        int maxP = 0;
        for (int i = 1; i < n; i++)
        {
            dp[i] = dp[i - 1] + prices[i] - prices[i - 1];
            dp[i] = dp[i] > 0 ? dp[i] : 0;
            maxP = max(maxP, dp[i]);
        }
        return maxP;
    }
};