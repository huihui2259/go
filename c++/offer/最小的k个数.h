#include "head.h"

class Solution
{
public:
    vector<int> getLeastNumbers(vector<int> &arr, int k)
    {
        priority_queue<int, vector<int>, greater<int>> q;
        vector<int> res;
        for (auto value : arr)
        {
            q.push(value);
        }
        for (int i = 0; i < k; i++)
        {
            res.emplace_back(q.top());
            q.pop();
        }
        return res;
    }
};