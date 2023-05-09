#include "c++/head.h"

class Solution
{
public:
    vector<int> sortArray(vector<int> &nums)
    {
        quickSort(0, nums.size() - 1, nums);
        return nums;
    }

    void quickSort(int start, int end, vector<int> &nums)
    {
        if (start >= end)
            return;
        srand(time(0));
        int index = rand() % (end - start + 1) + start;
        swap(nums[start], nums[index]);
        int i = start, j = end, tmp = nums[start];
        while (i < j)
        {
            /* code */
            while (i < j && nums[j] >= tmp)
            {
                j--;
            }
            if (i < j)
            {
                nums[i++] = nums[j];
            }

            while (i < j && nums[i] <= tmp)
            {
                i++;
            }
            if (i < j)
            {
                nums[j--] = nums[i];
            }
        }
        nums[i] = tmp;
        quickSort(start, i - 1, nums);
        quickSort(i + 1, end, nums);
    }
};