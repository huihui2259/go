#include "c++/head.h"

class LRUCache
{
public:
    int cap;
    list<pair<int, int>> l;
    unordered_map<int, list<pair<int, int>>::iterator> m;
    LRUCache(int capacity)
    {
        cap = capacity;
    }

    int get(int key)
    {
        if (m.find(key) == m.end())
            return -1;
        int res = m[key]->second;
        l.erase(m[key]);
        l.push_front(pair<int, int>(key, res));
        m[key] = l.begin();
        return res;
    }

    void put(int key, int value)
    {
        if (m.find(key) != m.end())
        {
            l.erase(m[key]);
        }
        if (l.size() == cap)
        {
            m.erase(l.back().first);
            l.pop_back();
        }
        l.push_front(pair<int, int>(key, value));
        m[key] = l.begin();
    }
};