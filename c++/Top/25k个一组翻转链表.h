#include "c++/head.h"

class Solution
{
public:
    ListNode *reverseKGroup(ListNode *head, int k)
    {
        // 四个指针来定位
        auto dummy = new ListNode;
        auto preHead = dummy, Head = head, tail = head, proTail = head;
        preHead->next = Head;
        while (true)
        {
            int i = 0;
            while (i < k - 1 && tail)
            {
                tail = tail->next;
                i++;
            }
            if (tail == nullptr)
            {
                return dummy->next;
            }
            proTail = tail->next;
            tail->next = nullptr;
            tail = reverse(Head);
            print(tail);
            preHead->next = tail;
            head->next = proTail;

            preHead = head;
            Head = proTail;
            tail = proTail;
        }
        return dummy->next;
    }
    void print(ListNode *node)
    {
        cout << "链表：" << endl;
        while (node)
        {
            cout << " " << node->val;
        }
        cout << endl;
    }

    ListNode *reverse(ListNode *head)
    {
        if (!head || !head->next)
            return head;
        ListNode *pre = new ListNode;
        ListNode *cur = head;
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