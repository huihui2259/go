#include "head.h"
/*
输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
输出: [3,3,5,5,6,7] 
解释: 

  滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
*/

class Solution
{
public:
    class mydeque
    {
    public:
        deque<int> d;
        mydeque()
        {
        }
        void push(int val)
        {
            while (!d.empty() && d.back() < val)
            {
                d.pop_back();
            }
            d.push_back(val);
        }
        void pop(int val)
        {
            if (!d.empty() && d.front() == val)
            {
                d.pop_front();
            }
        }
        int front()
        {
            return d.front();
        }
    };
    vector<int> maxSlidingWindow(vector<int> &nums, int k)
    {
        vector<int> res;
        mydeque deq;
        for (int i = 0; i < k; i++)
        {
            deq.push(nums[i]);
        }
        res.push_back(deq.front());
        for (int i = k; i < nums.size(); i++)
        {
            // 注意，这里push和pop可以调换顺序
            deq.push(nums[i]);
            deq.pop(nums[i - k]);
            res.push_back(deq.front());
        }
        return res;
    }
};
