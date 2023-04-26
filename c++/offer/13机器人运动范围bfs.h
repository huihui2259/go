#include "head.h"

class Solution
{
public:
    int movingCount(int m, int n, int k)
    {
        vector<vector<bool>> flag(m, vector<bool>(n));
        queue<vector<int>> q;
        q.push(vector<int>{0, 0});
        int count = 0;
        while (!q.empty())
        {
            auto x = q.front();
            q.pop();
            int i = x[0], j = x[1];
            if (i >= m || j >= n || digitsSum(i) + digitsSum(j) > k || flag[i][j])
            {
                continue;
            }
            flag[i][j] = true;
            count++;
            q.push(vector<int>{i + 1, j});
            q.push(vector<int>{i, j + 1});
        }
        return count;
    }
    int digitsSum(int num)
    {
        int sum = 0;
        while (num)
        {
            sum += num % 10;
            num /= 10;
        }
        return sum;
    }
};