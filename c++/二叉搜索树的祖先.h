#include "head.h"

class Solution
{
public:
    TreeNode *lowestCommonAncestor(TreeNode *root, TreeNode *p, TreeNode *q)
    {
        int maxVal = max(p->val, q->val);
        int minVal = maxVal == p->val ? q->val : p->val;
        while (root)
        {
            if (root->val >= minVal && root->val <= maxVal)
            {
                return root;
            }
            if (root->val > maxVal)
            {
                root = root->left;
            }
            if (root->val < minVal)
            {
                root = root->right;
            }
        }
        return nullptr;
    }
};