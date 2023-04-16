#include <iostream>
#include "head.h"
using namespace std;

void test(int &&val)
{
    cout << "val: " << val;
}
string reverseWords(string s)
{
    int i = 0;
    vector<string> res;
    while (i < s.size())
    {
        string tmp;
        while (i < s.size() && s[i] == ' ')
        {
            i++;
        }
        while (i < s.size() && s[i] != ' ')
        {
            tmp += s[i];
            i++;
        }
        if (!tmp.empty())
            res.push_back(tmp);
    }
    string str;
    for (i = res.size() - 1; i >= 0; i--)
    {
        str += res[i];
        if (i != 0)
        {
            str += " ";
        }
    }
    return str;
}

int main()
{
    string s = "the sky is blue";
    reverseWords(s);
}