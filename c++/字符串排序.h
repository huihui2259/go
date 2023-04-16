#include "head.h"

class Solution
{
public:
    string tmp;
    vector<string> res;
    vector<bool> visit;
    vector<string> permutation(string s)
    {
        for (int i = 0; i < s.size(); i++)
        {
            visit.push_back(false);
        }
        sort(s.begin(), s.end());
        dfs(s);
        return res;
    }
    void dfs(string s)
    {
        if (tmp.size() == s.size())
        {
            res.push_back(tmp);
            return;
        }
        for (int i = 0; i < s.size(); i++)
        {
            if (visit[i])
                continue;
            // TODO:对重复字符处理
            visit[i] = true;
            tmp.push_back(s[i]);
            dfs(s);
            tmp.pop_back();
            visit[i] = false;
        }
    }
};