#include "c++/head.h"

class Solution
{
public:
    ListNode *sortList(ListNode *head)
    {
        if (head == nullptr || head->next == nullptr)
            return head;
        ListNode *mid = middleNode(head);
        ListNode *left = sortList(head);
        ListNode *right = sortList(mid);
        return mergeTwoLists(left, right);
    }

    ListNode *middleNode(ListNode *head)
    {
        ListNode *fast = head, *slow = head;
        ListNode *prev = slow;
        while (fast && fast->next)
        {
            prev = slow;
            fast = fast->next->next;
            slow = slow->next;
        }
        prev->next = nullptr;
        return slow;
    }

    ListNode *mergeTwoLists(ListNode *l1, ListNode *l2)
    {
        ListNode *head = new ListNode(0);
        ListNode *tail = head;

        while (l1 && l2)
        {
            if (l1->val <= l2->val)
            {
                tail->next = l1;
                l1 = l1->next;
            }
            else
            {
                tail->next = l2;
                l2 = l2->next;
            }
            tail = tail->next;
        }
        tail->next = l1 ? l1 : l2;
        return head->next;
    }
};