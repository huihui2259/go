#include "c++/head.h"
/*
输入：[7,1,5,3,6,4]
输出：5
解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
*/
class Solution
{
public:
    int maxProfit(vector<int> &prices)
    {
        int n = prices.size();
        int max = 0;
        vector<int> dp(n, 0);
        for (int i = 1; i < prices.size(); i++)
        {
            dp[i] = (dp[i - 1] + prices[i] - prices[i - 1]) > 0 ? (dp[i - 1] + prices[i] - prices[i - 1]) : 0;
            max = max > dp[i] ? max : dp[i];
        }
        return max;
    }
};