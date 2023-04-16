#include "head.h"

// 输入: s = "abcdefg", k = 2
// 输出: "cdefgab"
class Solution
{
public:
    string reverseLeftWords(string s, int n)
    {
        int len = s.size();
        s.resize(s.size() + n);
        int j = 0;
        for (int i = len; i < s.size(); i++, j++)
        {
            s[i] = s[j];
        }
        return s.substr(j, s.size());
    }
};