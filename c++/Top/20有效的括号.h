#include "c++/head.h"

class Solution
{
public:
    bool isValid(string s)
    {
        stack<char> st;
        for (auto ch : s)
        {
            if (ch == '{')
            {
                st.push('}');
            }
            else if (ch == '[')
            {
                st.push(']');
            }
            else if (ch == '(')
            {
                st.push(')');
            }
            else
            {
                if (!st.empty() && ch == st.top())
                {
                    st.pop();
                }
                else
                {
                    return false;
                }
            }
        }
        return st.empty();
    }
};