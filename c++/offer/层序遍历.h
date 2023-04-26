#include "head.h"

class Solution
{
public:
    vector<int> levelOrder(TreeNode *root)
    {
        vector<int> res;
        if (!root)
        {
            return res;
        }
        queue<TreeNode *> q;
        q.push(root);
        while (!q.empty())
        {
            auto tmp = q.front();
            q.pop();
            res.push_back(tmp->val);
            if (tmp->left)
                q.push(tmp->left);
            if (tmp->right)
                q.push(tmp->right);
        }
        return res;
    }
};