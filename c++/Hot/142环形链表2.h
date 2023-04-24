#include "c++/head.h"

class Solution
{
public:
    ListNode *detectCycle(ListNode *head)
    {
        ListNode *fast = head, *low = head;
        int flag = 0;
        while (fast && fast->next)
        {
            low = low->next;
            fast = fast->next->next;
            if (fast == low)
            {
                flag = 1;
                break;
            }
        }
        if (!flag)
            return nullptr;
        // 已经确定有环
        fast = head;
        while (low != fast)
        {
            low = low->next;
            fast = fast->next;
        }
        return low;
    }
};

class Solution
{
public:
    bool hasCycle(ListNode *head)
    {
        ListNode *fast = head, *low = head;
        // 这段代码是错误的，原因如下：
        // 当fast到达最后一个节点时，if的条件不满足
        // 因此fast不会更新，而while的条件又满足，因此会出现：
        // low一直在更新，fast不更新，这样到最后就会low=fast
        // while (fast && low)
        // {
        //     if (low)
        //         low = low->next;
        //     if (fast && fast->next)
        //         fast = fast->next->next;
        //     if (low == fast)
        //         return true;
        // }
        while (fast && fast->next)
        {
            low = low->next;
            fast = fast->next->next;
            if (fast == low)
                return true;
        }

        return false;
    }
};