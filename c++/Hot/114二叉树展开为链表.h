#include "c++/head.h"

// [1,2,5,3,4,null,6]
class Solution
{
public:
    void flatten(TreeNode *root)
    {
        root = dfs(root);
    }

    // 传入一个节点, 将该节点下的子节点转为链表并返回该节点
    TreeNode *dfs(TreeNode *root)
    {
        if (root == nullptr)
            return;
        auto left = dfs(root->left);
        auto right = dfs(root->right);
        if (left)
        {
            auto last = Last(left);
            last->right = right;
            root->right = left;
            root->left = nullptr;
        }
        return root;
    }
    // 传入一个节点，返回该链表(转换后链表)的最后一个节点
    TreeNode *Last(TreeNode *root)
    {
        if (!root)
            return root;
        while (root->right)
        {
            root = root->right;
        }
        return root;
    }
};