#include "c++/head.h"

class Solution
{
public:
    vector<int> productExceptSelf(vector<int> &nums)
    {
        int n = nums.size();
        vector<int> leftDp(n, 1);
        vector<int> rightDp(n, 1);
        vector<int> res(n, 1);
        for (int i = 1, j = n - 2; i < n; i++, j--)
        {
            leftDp[i] = leftDp[i - 1] * nums[i - 1];
            rightDp[j] = rightDp[j + 1] * nums[j + 1];
        }
        for (int i = 0; i < n; i++)
        {
            res[i] = leftDp[i] * rightDp[i];
        }
        return res;
    }
};