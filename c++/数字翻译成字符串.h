#include "head.h"

/*
给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，
11 翻译成 “l”，……，25 翻译成 “z”。一个数字可能有多个翻译。请编程实现一个函数，
用来计算一个数字有多少种不同的翻译方法。
输入: 12258
输出: 5
解释: 12258有5种不同的翻译，分别是"bccfi", "bwfi", "bczi", "mcfi"和"mzi"
*/

class Solution
{
public:
    int translateNum(int num)
    {
        auto s = to_string(num);
        auto a = s[0];
        vector<int> dp(s.size(), 1);
        for (int i = 1; i < s.size(); i++)
        {
            auto a = s[i - 1];
            auto b = s[i];
            if (i >= 2 && isCanTrans(s[i - 1], s[i]))
            {
                dp[i] = dp[i - 1] + dp[i - 2];
                continue;
            }
            if (i == 1 && isCanTrans(s[i - 1], s[i]))
            {
                dp[i] = 2;
                continue;
            }
            dp[i] = dp[i - 1];
        }
        return dp.back();
    }
    bool isCanTrans(char a, char b)
    {
        if (a == '1')
        {
            return true;
        }
        if (a == '2' && b <= '5')
        {
            return true;
        }
        return false;
    }
};