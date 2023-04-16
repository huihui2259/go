#include "head.h"

class Solution
{
public:
    vector<vector<int>> findContinuousSequence(int target)
    {
        int i = 0, j = 0, sum = 0;
        vector<vector<int>> res;
        while (i <= target / 2)
        {
            while (sum < target)
            {
                sum += j;
                j++;
            }
            while (sum > target)
            {
                sum -= i;
                i++;
            }
            if (sum == target)
            {
                vector<int> tmp;
                for (int k = i; k < j; k++)
                {
                    tmp.push_back(k);
                }
                res.push_back(tmp);
                sum -= i;
                i++;
            }
        }
        return res;
    }
};