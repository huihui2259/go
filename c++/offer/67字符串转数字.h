#include "head.h"

class Solution
{
public:
    int strToInt(string str)
    {
        // 第一个非空字符必须是+、-或数字
        int pos = 0, n = str.size();
        while (pos < n && str[pos] == ' ')
        {
            pos++;
        }
        // str中全是空格
        if (pos == n)
        {
            return 0;
        }
        // 第一个非空字符必须合法
        if (str[pos] != '-' && str[pos] != '+' && (str[pos] < '0' || str[pos] > '9'))
        {
            return 0;
        }
        // 以下是合法情况，另开辟一个函数
        char sign = ' ';
        if (str[pos] == '-' || str[pos] == '+')
        {
            sign = str[pos++];
        }
        string s;
        while (pos < n && (str[pos] >= '0' && str[pos] <= '9'))
        {
            s += str[pos++];
        }
        double res = toInt(s);
        if (res > INT32_MAX && sign != '-')
        {
            return INT32_MAX;
        }
        if (res > INT32_MAX && sign == '-')
        {
            return INT32_MIN;
        }
        if (sign == '-')
        {
            return 0 - res;
        }
        return res;
    }
    double toInt(string s)
    {
        double res = 0;
        for (int i = s.size() - 1; i >= 0; i--)
        {
            res += (s[i] - '0') * pow(10, s.size() - i - 1);
        }
        return res;
    }
};