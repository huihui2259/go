#include "head.h"

class MinStack
{
public:
    stack<int> st;
    stack<int> minSt;
    /** initialize your data structure here. */
    MinStack()
    {
    }

    void push(int x)
    {
        st.push(x);
        while (!minSt.empty() && x < minSt.top())
        {
            minSt.pop();
        }
        minSt.push(x);
    }

    void pop()
    {
        int top = st.top();
        st.pop();
        if (minSt.top() == top)
        {
            minSt.pop();
        }
    }

    int top()
    {
        return st.top();
    }

    int min()
    {
        return minSt.top();
    }
};