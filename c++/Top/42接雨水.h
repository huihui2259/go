#include "c++/head.h"

class Solution
{
public:
    int trap(vector<int> &height)
    {
        int n = height.size();
        int res = 0;
        vector<int> Left(height);
        vector<int> Right(height);
        for (int i = 1, j = n - 2; i < n, j >= 0; i++, j--)
        {
            Left[i] = max(Left[i - 1], height[i]);
            Right[j] = max(Right[j + 1], height[j]);
        }
        for (int i = 0; i < n; i++)
        {
            res += min(Left[i], Right[i]) - height[i];
        }
        return res;
    }
};