#include "c++/head.h"

class Solution
{
public:
    vector<vector<int>> zigzagLevelOrder(TreeNode *root)
    {
        vector<vector<int>> res;
        if (!root)
            return res;
        queue<TreeNode *> q;
        q.push(root);
        int count = 0;
        while (!q.empty())
        {
            /* code */
            int n = q.size();

            vector<int> tmp;
            for (int i = 0; i < n; i++)
            {
                auto node = q.front();
                q.pop();
                tmp.push_back(node->val);
                if (node->left)
                    q.push(node->left);
                if (node->right)
                    q.push(node->right);
            }
            if (count % 2)
            {
                reverse(tmp.begin(), tmp.end());
            }
            res.push_back(tmp);
            count++;
        }
        return res;
    }
};