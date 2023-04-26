#include "head.h"

class Solution
{
public:
    bool exist(vector<vector<char>> &board, string word)
    {
        for (int i = 0; i < board.size(); i++)
        {
            for (int j = 0; j < board[0].size(); j++)
            {
                if (dfs(i, j, board, 0, word))
                {
                    return true;
                }
            }
        }
        return false;
    }

    bool dfs(int i, int j, vector<vector<char>> &board, int len, string word)
    {
        if (i < 0 || j < 0 || i >= board.size() || j >= board[0].size() || board[i][j] != word[len])
        {
            return false;
        }
        if (len == word.size() - 1)
        {
            return true;
        }
        auto tmp = board[i][j];
        board[i][j] = '0';
        if (dfs(i + 1, j, board, len + 1, word) || dfs(i, j + 1, board, len + 1, word) || dfs(i, j - 1, board, len + 1, word) || dfs(i - 1, j, board, len + 1, word))
            return true;
        board[i][j] = tmp;
        return false;
    }
};