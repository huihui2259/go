#include "c++/head.h"

class LRUCache
{
public:
    list<pair<int, int>> li;
    unordered_map<int, list<pair<int, int>>::iterator> m;
    int cap = 0;
    LRUCache(int capacity)
    {
        cap = capacity;
    }

    int get(int key)
    {
        // 没找到直接返回-1
        if (m.find(key) == m.end())
            return -1;
        int res = m[key]->second;
        li.erase(m[key]);
        li.push_front(pair<int, int>(key, res));
        m[key] = li.begin();
        return res;
    }

    void put(int key, int value)
    {
        // 下述的代码有问题，比如有容量为2，两个元素
        // (1,3)(2,3),这时再加入(1,5),此时会先删除(2,3),这样就会变成最终元素为(1,5)
        // 其实应该是(1,5)(2,3)
        // // 容量满了，删除最不常用
        // if (m.size() == cap)
        // {
        //     m.erase(li.back().first);
        //     li.pop_back();
        //     // li.erase(li.end());
        // }
        // // 容量未满，且已存在，删除
        // if (m.find(key) != m.end())
        // {
        //     li.erase(m[key]);
        // }
        // 已存在，删除
        if (m.find(key) != m.end())
        {
            li.erase(m[key]);
        }
        // 容量满了，删除最不常用
        if (li.size() == cap)
        {
            m.erase(li.back().first);
            li.pop_back();
            // li.erase(li.end());
        }
        // 插入新的
        li.push_front(pair<int, int>(key, value));
        m[key] = li.begin();
    }
};