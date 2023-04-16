#include "head.h"

class Solution
{
public:
    int res = 0;
    int c = 0;
    int kthLargest(TreeNode *root, int k)
    {
        c = k;
        dfs(root);
        return res;
    }
    void dfs(TreeNode *root)
    {
        if (root == nullptr)
        {
            return;
        }

        dfs(root->left);
        c--;
        if (c == 0)
        {
            res = root->val;
            return;
        }
        dfs(root->right);
    }
};