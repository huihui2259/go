#include "c++/head.h"

class Solution
{
public:
    bool isPalindrome(ListNode *head)
    {
        ListNode *fast = head;
        ListNode *low = head;
        ListNode *pro = new ListNode(0);
        pro->next = low;
        while (!fast && !fast->next)
        {
            fast = fast->next->next;
            low = low->next;
            pro = pro->next;
        }
        pro->next = nullptr;
        auto newHead = reverse(head);
        // 节点数为偶数
        if (fast == nullptr)
        {
            return IsSame(newHead, low);
        }
        // 节点为奇数
        return IsSame(newHead, low->next);
    }
    // 反转链表
    ListNode *reverse(ListNode *head)
    {
        if (!head)
            return head;
        ListNode *pre = new ListNode(0);
        ListNode *cur = head;
        pre->next = cur;
        ListNode *pro = cur->next;
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
    bool IsSame(ListNode *l1, ListNode *l2)
    {
        while (l1 && l2)
        {
            if (l1->val != l2->val)
            {
                return false;
            }
            l1 = l1->next;
            l2 = l2->next;
        }
        return true;
    }
};