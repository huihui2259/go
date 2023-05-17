#include "c++/head.h"

class Solution
{
public:
    int mySqrt(int x)
    {
        if (x == 0)
            return x;
        int i = 1, j = x;
        while (i <= j)
        {
            int mid = (j - i) / 2 + i;
            if (mid <= x - 1 && x / mid >= mid && x / (mid + 1) < (mid + 1))
            {
                return mid;
            }
            else if (x / mid > mid)
            {
                i = mid + 1;
            }
            else
            {
                j = mid - 1;
            }
        }
        return i;
    }

    int mySqrt(int x)
    {
        if (x == 0)
            return x;
        int i = 1, j = x;
        int ans = 0;
        while (i <= j)
        {
            int mid = (j - i) / 2 + i;
            if (x / mid >= mid)
            {
                ans = mid;
                i = mid + 1;
            }
            else
            {
                j = mid - 1;
            }
        }
        return ans;
    }
};