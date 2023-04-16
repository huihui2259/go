#include "head.h"

class Solution
{
public:
    int lengthOfLongestSubstring(string s)
    {
        deque<char> d;
        int res = 0;
        for (int i = 0; i < s.size(); i++)
        {
            if (find(d.begin(), d.end(), s[i]) == d.end())
            {
                d.push_back(s[i]);
            }
            else
            {
                while (d.front() != s[i])
                {
                    d.pop_front();
                }
                d.pop_front();
                d.push_back(s[i]);
            }
            res = res > d.size() ? res : d.size();
        }
        return res;
    }
};