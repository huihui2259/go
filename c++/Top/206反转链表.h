#include "c++/head.h"

class Solution
{
public:
    ListNode *reverseList(ListNode *head)
    {
        if (!head || !head->next)
            return head;
        auto pre = new ListNode;
        auto cur = head;
        pre->next = cur;
        auto pro = cur->next;
        while (pro)
        {
            cur->next = pre;
            pre = cur;
            cur = pro;
            pro = pro->next;
        }
        cur->next = pre;
        head->next = nullptr;
        return cur;
    }
};