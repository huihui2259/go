#include "c++/head.h"

class Solution
{
public:
    int sum = 0;
    int tmp = 0;
    int sumNumbers(TreeNode *root)
    {
        dfs(root);
        return sum;
    }
    void dfs(TreeNode *root)
    {
        if (!root->left && !root->right)
        {
            tmp = tmp * 10 + root->val;
            sum += tmp;
            return;
        }
        tmp = tmp * 10 + root->val;
        if (root->left)
        {
            dfs(root->left);
            tmp = (tmp - root->left->val) / 10;
        }

        if (root->right)
        {
            dfs(root->right);
            tmp = (tmp - root->right->val) / 10;
        }
    }
};