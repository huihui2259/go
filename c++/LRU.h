#include "head.h"

class LRUCache
{
public:
    int cap;
    list<pair<int, int>> l;
    map<int, list<pair<int, int>>::iterator> m;
    LRUCache(int capacity)
    {
        cap = capacity;
    }

    int get(int key)
    {
        if (m.find(key) == m.end())
            return -1;
        int res = m[key]->second;
        // auto tmp = m[key];
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
        else
        {
            if (m.size() == cap)
            {
                m.erase(l.back().first);
                l.pop_back();
            }
        }
        l.push_back(pair<int, int>(key, value));
        m[key] = l.begin();
    }
};

/**
 * Your LRUCache object will be instantiated and called as such:
 * LRUCache* obj = new LRUCache(capacity);
 * int param_1 = obj->get(key);
 * obj->put(key,value);
 */