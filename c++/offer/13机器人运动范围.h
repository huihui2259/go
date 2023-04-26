#include "head.h"

class Solution
{
public:
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
    int dfs(int i, int j, int k, int m, int n, vector<vector<bool>> &flag)
    {
        if (i == m || j == n || (digitsSum(i) + digitsSum(j)) > k || flag[i][j])
        {
            return 0;
        }
        flag[i][j] = true;
        return 1 + dfs(i + 1, j, k, m, n, flag) + dfs(i, j + 1, k, m, n, flag);
    }
    int movingCount(int m, int n, int k)
    {
        vector<vector<bool>> flag(m, vector<bool>(n));
        return dfs(0, 0, k, m, n, flag);
    }
};