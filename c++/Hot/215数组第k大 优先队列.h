#include "c++/head.h"

class Solution
{
public:
    int findKthLargest(vector<int> &nums, int k)
    {
        priority_queue<int, vector<int>, less<int>> q;
        for (auto num : nums)
        {
            q.push(num);
        }
        for (int i = 0; i < k - 1; i++)
        {
            q.pop();
        }
        return q.top();
    }
};