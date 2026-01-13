
def solve(height):
    l = 0
    r = len(height) - 1
    # result=0
    max_area = 0
    result = ()
    while (l < r):
        area = min(height[l], height[r]) * (r - l)
        if area > max_area:
            max_area = area
            # result
            result = (l, r)

        # 移动
        if height[l] < height[r]:
            l += 1
        else:
            r -= 1
        pass
    pass
    return result


if __name__ == "__main__":
    print(solve([1, 8, 6, 2, 5, 4, 8, 3, 7]))
    pass

