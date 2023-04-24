#include "c++/head.h"

class Solution
{
public:
    map<int, int> m;
    TreeNode *buildTree(vector<int> &preorder, vector<int> &inorder)
    {
        for (int i = 0; i < inorder.size(); i++)
        {
            m[inorder[i]] = i;
        }
        return dfs(preorder, 0, preorder.size() - 1, inorder, 0, inorder.size() - 1);
    }
    TreeNode *dfs(vector<int> &preorder, int pl, int pr, vector<int> &inorder, int il, int ir)
    {
        if (pl > pr)
            return nullptr;
        TreeNode *root = new TreeNode(preorder[pl]);
        int index = m[preorder[pl]];
        int leftLen = index - il, rightLen = ir - index;
        root->left = dfs(preorder, pl + 1, pl + leftLen, inorder, il, index - 1);
        root->right = dfs(preorder, pl + leftLen + 1, pr, inorder, index + 1, ir);
        return root;
    }
};