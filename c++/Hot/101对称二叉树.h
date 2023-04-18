#include "c++/head.h"

class Solution
{
public:
    bool isSymmetric(TreeNode *root)
    {
        judge(root, root);
    }
    bool judge(TreeNode *left, TreeNode *right)
    {
        if (!left && !right)
            return true;
        if (left && !right)
            return false;
        if (!left && right)
            return false;
        if (left->val != right->val)
            return false;
        return judge(left->left, right->right) && judge(left->right, right->left);
    }
};