#include "c++/head.h"

class Solution
{
public:
    string longestPalindrome(string s)
    {
        int n = s.size();
        int len = 0;
        string res;
        vector<vector<bool>> dp(n, vector<bool>(n, false));
        for (int i = n - 1; i >= 0; i--)
        {
            for (int j = i; j < n; j++)
            {
                if (i == j)
                    dp[i][j] = true;
                else if (j - i == 1 && s[i] == s[j])
                    dp[i][j] = true;
                else if (j - i > 1 && dp[i + 1][j - 1] && s[i] == s[j])
                    dp[i][j] = true;
                else
                    dp[i][j] = false;
                if (dp[i][j])
                {
                    if (j - i + 1 > len)
                    {
                        len = j - i + 1;
                        res = s.substr(i, len);
                    }
                }
            }
        }
        return res;
    }
};