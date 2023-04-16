#include "head.h"
class Node
{
public:
    int val;
    Node *left;
    Node *right;

    Node() {}

    Node(int _val)
    {
        val = _val;
        left = NULL;
        right = NULL;
    }

    Node(int _val, Node *_left, Node *_right)
    {
        val = _val;
        left = _left;
        right = _right;
    }
};

class Solution
{
public:
    Node *head = nullptr;
    Node *pre = nullptr;
    Node *last = nullptr;
    void dfs(Node *root)
    {
        if (root == nullptr)
        {
            return;
        }
        dfs(root->left);
        if (head == nullptr)
        {
            head = root;
        }
        if (pre != nullptr)
        {
            pre->right = root;
            root->left = pre;
        }
        pre = root;
        dfs(root->right);
    }
    Node *treeToDoublyList(Node *root)
    {
        if (!root)
            return root;
        dfs(root);
        pre->right = head;
        head->left = pre;
        return head;
    }
};