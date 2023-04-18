#include "head.h"

class MaxQueue
{
public:
    deque<int> q;
    deque<int> simple;
    MaxQueue()
    {
    }

    int max_value()
    {
        if (q.empty())
            return -1;
        return q.front();
    }

    void push_back(int value)
    {
        while (!q.empty() && q.back() < value)
        {
            /* code */
            q.pop_back();
        }
        q.push_back(value);
        simple.push_back(value);
    }

    int pop_front()
    {
        if (simple.empty())
            return -1;
        if (!q.empty() && simple.front() == q.front())
        {
            q.pop_front();
        }
        int res = simple.front();
        simple.pop_front();
        return res;
    }
};