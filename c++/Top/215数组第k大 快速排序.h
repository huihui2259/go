#include "c++/head.h"

class Solution
{
public:
    int quickSort(int l, int r, vector<int> &nums, int target)
    {
        srand(time(0));
        int index = rand() % (r - l + 1) + l;
        swap(nums[l], nums[index]);
        int i = l, j = r, tmp = nums[l];
        while (i < j)
        {
            while (i < j && nums[j] < tmp)
            {
                j--;
            }
            if (i < j)
            {
                nums[i++] = nums[j];
            }

            while (i < j && nums[i] > tmp)
            {
                i++;
            }
            if (i < j)
            {
                nums[j--] = nums[i];
            }
        }
        nums[i] = tmp;
        if (i == target)
            return nums[i];
        else if (i < target)
            return quickSort(i + 1, r, nums, target);
        else
            return quickSort(l, i - 1, nums, target);
    }

    int findKthLargest(vector<int> &nums, int k)
    {
        return quickSort(0, nums.size() - 1, nums, k - 1);
    }
};