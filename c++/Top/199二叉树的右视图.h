#include "c++/head.h"

class Solution
{
public:
    vector<int> rightSideView(TreeNode *root)
    {
        vector<int> res;
        if (!root)
            return res;
        queue<TreeNode *> q;
        q.push(root);
        while (!q.empty())
        {
            int n = q.size();
            for (int i = 0; i < n; i++)
            {
                auto top = q.front();
                if (i == n - 1)
                {
                    res.push_back(top->val);
                }
                q.pop();
                if (top->left)
                    q.push(top->left);
                if (top->right)
                    q.push(top->right);
            }
        }
        return res;
    }
};