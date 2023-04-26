#include "c++/head.h"

class mydeque
{
public:
    deque<int> q;
    mydeque() {}
    void push(int val)
    {
        while (!q.empty() && val > q.back())
        {
            q.pop_back();
        }
        q.push_back(val);
    }

    void pop(int val)
    {
        if (!q.empty() && q.front() == val)
        {
            q.pop_front();
        }
    }

    int top()
    {
        return q.front();
    }
};

class Solution
{
public:
    vector<int> maxSlidingWindow(vector<int> &nums, int k)
    {
        vector<int> res;
        mydeque d;
        for (int i = 0; i < k; i++)
        {
            d.push(nums[i]);
        }
        res.push_back(d.top());
        for (int i = k; i < nums.size(); i++)
        {
            d.pop(nums[i - k]);
            d.push(nums[i]);
            res.push_back(d.top());
        }
        return res;
    }
};