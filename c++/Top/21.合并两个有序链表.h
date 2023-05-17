#include "c++/head.h"

class Solution
{
public:
    ListNode *mergeTwoLists(ListNode *list1, ListNode *list2)
    {
        auto node = new ListNode;
        auto head = node;
        while (list1 && list2)
        {
            if (list1->val < list2->val)
            {
                node->next = list1;
                list1 = list1->next;
                node = node->next;
            }
            else
            {
                node->next = list2;
                list2 = list2->next;
                node = node->next;
            }
        }
        if (list1)
            node->next = list1;
        if (list2)
            node->next = list2;
        return head->next;
    }
};