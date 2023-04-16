#include "head.h"

class Solution
{
public:
    bool isBalanced(TreeNode *root)
    {
        if (root == nullptr)
            return true;
        if (abs(depth(root->left) - depth(root->right)) > 1)
        {
            return false;
        }
        return isBalanced(root->left) && isBalanced(root->right);
    }
    int depth(TreeNode *root)
    {
        if (root == nullptr)
            return 0;
        return max(depth(root->left), depth(root->right)) + 1;
    }
};