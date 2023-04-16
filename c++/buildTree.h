#include "head.h"

class Solution
{
public:
    vector<int> Preorder;
    map<int, int> m;
    TreeNode *buildTree(vector<int> &preorder, vector<int> &inorder)
    {
        if (preorder.size() == 0)
            return nullptr;
        for (int i = 0; i < inorder.size(); i++)
        {
            m[inorder[i]] = i;
        }
        Preorder = preorder;
        return build(0, preorder.size() - 1, 0, inorder.size() - 1);
    }
    TreeNode *build(int pl, int pr, int il, int ir)
    {
        if (pl > pr)
            return nullptr;
        TreeNode *root = new TreeNode(Preorder[pl]);
        int pos = m[Preorder[pl]];
        root->left = build(pl + 1, pl + pos - il, il, pos - 1);
        root->right = build(pl + pos - il + 1, pr, pos + 1, ir);
        return root;
    }
};