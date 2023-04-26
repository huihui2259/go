#include "head.h"

// 输入: [1,2,3,4,5]
// 输出: [120,60,40,30,24]

class Solution
{
public:
    vector<int> constructArr(vector<int> &a)
    {
        int n = a.size();
        vector<int> left(n, 1);
        vector<int> right(n, 1);
        vector<int> res(n);
        for (int i = 1, j = n - 2; i < n; i++, j--)
        {
            left[i] = left[i - 1] * a[i - 1];
            right[j] = right[j + 1] * a[j + 1];
        }
        for (int i = 0; i < n; i++)
        {
            res[i] = left[i] * right[i];
        }
        return res;
    }
};