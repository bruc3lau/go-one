class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


def binaryTreePaths(root: TreeNode) -> list[str]:
    if not root:
        return []

    stack = [(root, str(root.val))]
    result = []

    while stack:
        node, path = stack.pop()

        # 叶子节点
        if not node.left and not node.right:
            result.append(path)
        else:
            if node.right:
                stack.append((node.right, path + "->" + str(node.right.val)))
            if node.left:
                stack.append((node.left, path + "->" + str(node.left.val)))

    return result


# 测试1
root = TreeNode(1)
root.left = TreeNode(2, None, TreeNode(5))
root.right = TreeNode(3)
print(binaryTreePaths(root))  # ["1->2->5", "1->3"]

# 测试2：单节点
root = TreeNode(1)
print(binaryTreePaths(root))  # ["1"]
