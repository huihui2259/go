#include "c++/head.h"

class Solution
{
public:
    ListNode *mergeKLists(vector<ListNode *> &lists)
    {
        auto cmp = [](ListNode *a, ListNode *b) { return a->val > b->val; };
        priority_queue<ListNode *, vector<ListNode *>, decltype(cmp)> q(cmp);
        ListNode *dummy = new ListNode;
        auto cur = dummy;
        for (auto node : lists)
        {
            if (node)
            {
                q.push(node);
            }
        }
        while (!q.empty())
        {
            auto node = q.top();
            cur->next = node;
            if (node->next)
            {
                q.push(node->next);
            }
            q.pop();
            cur = cur->next;
        }
        cur->next = nullptr;
        return dummy->next;
    }
};