#include "c++/head.h"

class Solution
{
public:
    ListNode *reverseBetween(ListNode *head, int left, int right)
    {
        if (left < 0 || left >= right)
            return head;
        if (!head || !head->next)
            return head;
        int count = 0;
        auto dummy = new ListNode;
        auto pre = new ListNode;
        auto leftNode = head, rightNode = head, pro = rightNode->next, cur = head;
        pre->next = head;
        while (cur)
        {
            count++;
            if (left - 1 == count)
            {
                pre = cur;
            }
            if (left == count)
            {
                leftNode = cur;
            }
            if (right == count)
            {
                rightNode = cur;
                pro = rightNode->next;
            }
            cur = cur->next;
        }
        rightNode->next = nullptr;
        rightNode = reverse(leftNode);
        pre->next = rightNode;
        leftNode->next = pro;
        if (left == 1)
        {
            return pre->next;
        }
        return head;
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