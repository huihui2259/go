#include "head.h"

class Solution
{
public:
    string reverseWords(string s)
    {
        // 原地翻转(不太对，“   ”这种情况不对)
        reverse(s.begin(), s.end()); //"blue is sky the"
        int i = 0, j = 0;
        while (j < s.size())
        {
            int start = i;
            while (j < s.size() && s[j] == ' ')
            {
                j++;
            }
            while (j < s.size() && s[j] != ' ')
            {
                s[i++] = s[j++];
            }
            reverse(s.begin() + start, s.begin() + i);
            s[i++] = ' ';
        }
        if (i == s.size())
            return s.substr(0, i - 1);
        return s.substr(0, i - 2);
    }
};