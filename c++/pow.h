#include "head.h"

class Solution
{
public:
    double myPow(double x, int n)
    {
        double res = 1.0;
        do
        {
            if (n % 2)
            {
                res *= x;
            }
            x *= x;
        } while (n / 2);
        return n > 0 ? res : 1 / res;
    }
};