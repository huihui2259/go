#include "c++/head.h"

class Solution
{
public:
    bool isSymmetric(TreeNode *root)
    {
        if (!root)
            return true;
        return auxFunc(root->left, root->right);
    }

    bool auxFunc(TreeNode *left, TreeNode *right)
    {
        if (!left && !right)
            return true;
        if (left && !right)
            return false;
        if (!left && right)
            return false;
        if (left && right && left->val != right->val)
            return false;
        return auxFunc(left->right, right->left) && auxFunc(left->left, right->right);
    }
};