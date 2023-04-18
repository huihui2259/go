#include "c++/head.h"

class Solution
{
public:
    vector<vector<int>> levelOrder(TreeNode *root)
    {
        vector<vector<int>> res;
        if (!root)
            return res;
        queue<TreeNode *> q;
        q.push(root);
        while (!q.empty())
        {
            int len = q.size();
            vector<int> level;
            for (int i = 0; i < len; i++)
            {
                auto front = q.front();
                q.pop();
                level.push_back(front->val);
                if (front->left)
                {
                    q.push(front->left);
                }
                if (front->right)
                {
                    q.push(front->right);
                }
            }
            res.push_back(level);
        }
        return res;
    }
};