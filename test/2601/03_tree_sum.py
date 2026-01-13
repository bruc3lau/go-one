

# 3. 示例:
#     1. 输入：nums = [-1,0,1,2,-1,-4]
# 输出：[[-1,-1,2],[-1,0,1]]
#     2. 输入：nums = [0,1,1]
# 输出：[]
#     3. 输入：nums = [0,0,0]
# 输出：[[0,0,0]]

def solve(nums):
    nums.sort()
    res = []
    n = len(nums)

    for i in range(n - 2):
        # todo
        #  去重
        if i > 0 and nums[i] == nums[i - 1]:
            continue
        l, r = i + 1, n - 1
        while l < r:
            s = nums[i] + nums[l] + nums[r]
            if s == 0:
                res.append([nums[i], nums[l], nums[r]])

                #todo
                # 去重
                while l < r and nums[l] == nums[l + 1]:
                    l += 1
                while l < r and nums[r] == nums[r - 1]:
                    r -= 1
                l += 1
                r -= 1
            elif s < 0:
                l += 1
            else:
                r -= 1
    return res
    pass


if __name__ == "__main__":
    print(solve([-1, 0, 1, 2, -1, -4]))
    print(solve([0,1,1]))
    print(solve([0,0,0]))
    pass
