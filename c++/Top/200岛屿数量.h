#include "c++/head.h"

class Solution
{
public:
    int numIslands(vector<vector<char>> &grid)
    {
        int count = 0;
        for (int i = 0; i < grid.size(); i++)
        {
            for (int j = 0; j < grid[0].size(); j++)
            {
                if (grid[i][j] == '1')
                {
                    count++;
                    dfs(i, j, grid);
                }
            }
        }
        return count;
    }
    void dfs(int i, int j, vector<vector<char>> &grid)
    {
        if (i > grid.size() || j > grid[0].size() || i <= 0 || j <= 0 || grid[i][j] == '0')
        {
            return;
        }
        grid[i][j] = '0';
        dfs(i, j + 1, grid);
        dfs(i + 1, j, grid);
        dfs(i - 1, j, grid);
        dfs(i, j - 1, grid);
    }
};