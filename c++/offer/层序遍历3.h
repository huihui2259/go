#include "head.h"

class Solution
{
public:
    vector<vector<int>> levelOrder(TreeNode *root)
    {
        vector<vector<int>> res;
        if (!root)
        {
            return res;
        }
        queue<TreeNode *> q;
        q.push(root);
        int nowLevel = 0;
        while (!q.empty())
        {
            auto n = q.size();
            vector<int> level;
            for (int i = 0; i < n; i++)
            {
                auto tmp = q.front();
                q.pop();
                level.push_back(tmp->val);
                if (tmp->left)
                    q.push(tmp->left);
                if (tmp->right)
                    q.push(tmp->right);
            }
            if (nowLevel % 2 != 0)
            {
                reverse(level.begin(), level.end());
            }
            res.push_back(level);
            nowLevel++;
        }
        return res;
    }
};