#include <iostream>
#include <vector>
#include <map>
#include <queue>
#include <algorithm>
#include <stack>
#include <cmath>
#include <string>
#include <list>
using namespace std;

struct TreeNode
{
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};
struct ListNode
{
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};
