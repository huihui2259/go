#include "c++/head.h"

class Solution
{
public:
    string addStrings(string num1, string num2)
    {
        auto maxStr = num1.size() > num2.size() ? num1 : num2;
        auto minStr = num1 == maxStr ? num2 : num1;
        reverse(maxStr.begin(), maxStr.end());
        reverse(minStr.begin(), minStr.end());
        for (int i = minStr.size(); i < maxStr.size(); i++)
        {
            minStr.push_back('0');
        }
        int pos = 0;
        string res;
        for (int i = 0; i < maxStr.size(); i++)
        {
            int tmp = (maxStr[i] - '0') + (minStr[i] - '0') + pos;
            pos = tmp / 10;
            res.push_back(tmp % 10 + '0');
        }
        if (pos == 1)
        {
            res.push_back(pos + '0');
        }
        reverse(res.begin(), res.end());
        return res;
    }

    string addStrings2(string num1, string num2)
    {
        auto maxStr = num1.size() > num2.size() ? num1 : num2;
        auto minStr = num1 == maxStr ? num2 : num1;
        reverse(maxStr.begin(), maxStr.end());
        reverse(minStr.begin(), minStr.end());
        int pos = 0;
        string res;
        for (int i = 0; i < minStr.size(); i++)
        {
            int tmp = (maxStr[i] - '0') + (minStr[i] - '0') + pos;
            pos = tmp / 10;
            res.push_back(tmp % 10 + '0');
        }
        for (int i = minStr.size(); i < maxStr.size(); i++)
        {
            int tmp = (maxStr[i] - '0') + pos;
            pos = tmp / 10;
            res.push_back(tmp % 10 + '0');
        }

        if (pos == 1)
        {
            res.push_back(pos + '0');
        }
        reverse(res.begin(), res.end());
        return res;
    }
};