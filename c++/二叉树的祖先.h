#include "head.h"

class Solution
{
public:
    TreeNode *lowestCommonAncestor(TreeNode *root, TreeNode *p, TreeNode *q)
    {
        if (root == nullptr || root == p || root == q)
        {
            return root;
        }
        auto l = lowestCommonAncestor(root->left, p, q);
        auto r = lowestCommonAncestor(root->right, p, q);
        if (l && r)
        {
            return root;
        }
        else if (l && !r)
        {
            return l;
        }
        return r;
    }
};