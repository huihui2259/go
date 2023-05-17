#include "c++/head.h"

class Solution
{
public:
    void reorderList(ListNode *head)
    {
        if (!head || !head->next)
            return;
        auto head1 = head;
        auto low = head, fast = head;

        while (fast && fast->next)
        {
            low = low->next;
            fast = fast->next->next;
        }
        low->next = nullptr;
        cout << low->val;
        auto head2 = reverse(low);
        cout << head2->val;
        while (head2 != nullptr)
        {
            auto tmp1 = head1->next;
            auto tmp2 = head2->next;
            head1->next = head2;
            head2->next = tmp1;
            head1 = tmp1;
            head2 = tmp2;
        }
        return;
    }
    ListNode *reverse(ListNode *head)
    {
        if (!head || !head->next)
        {
            return head;
        }
        ListNode *pre = new ListNode;
        ListNode *cur = head;
        pre->next = cur;
        auto pro = cur->next;
        while (cur)
        {
            cur->next = pre;
            pre = cur;
            cur = pro;
            if (pro)
            {
                pro = pro->next;
            }
        }
        head->next = nullptr;
        return pre;
    }
};