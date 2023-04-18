#include "c++/head.h"

class Solution
{
public:
    vector<int> res;
    bool isValidBST(TreeNode *root)
    {
        dfs(root);
        for (int i = 1; i < res.size(); i++)
        {
            if (res[i] < res[i - 1])
                return false;
        }
        return true;
    }
    void dfs(TreeNode *root)
    {
        if (root == nullptr)
            return;
        dfs(root->left);
        res.push_back(root->val);
        dfs(root->right);
    }
};
