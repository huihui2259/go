#include "c++/head.h"

/*
给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。

完全平方数 是一个整数，其值等于另一个整数的平方；
换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。

示例 1：

输入：n = 12
输出：3 
解释：12 = 4 + 4 + 4
示例 2：

输入：n = 13
输出：2
解释：13 = 4 + 9
*/
// 感觉也像背包问题
// 其实是dp，但dp的位移不是每次+1，而是跳跃的

class Solution
{
public:
    int numSquares(int n)
    {
        vector<int> dp(n + 1, n);
        if (n == 1)
            return 1;
        dp[0] = 0;
        dp[1] = 1;
        for (int i = 2; i <= n; i++)
        {
            int minn = 0;
            for (int j = 1; i >= j * j; j++)
            {
                minn = dp[i - j * j] + 1;
                dp[i] = min(minn, dp[i]);
            }
        }
        return dp[n];
    }
};