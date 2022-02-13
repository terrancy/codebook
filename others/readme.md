## 经典算法
#### 1、排序算法
- [ ] 快排的优化
#### 2、查找算法
- [x] 👏👏[二分查找框架](https://labuladong.gitee.io/algo/2/21/61/?_blank)
- [x] 👏👏[宫水三叶](https://github.com/SharingSource/LogicStack-LeetCode/tree/main/PDF?_blank)
- [x] 二分查找-II(左边界:1.mid=l+r>>2; 2.nums[mid]>=target r=mid;3.return r，右边界:1.mid=l+r+1>>2;2.nums[mid]<=target l=mid; 3.return l)
- [x] 数字在升序数组中出现的次数(通过二分查找计算数字的左右边界)
- [x] 旋转数组的最小数字(严格升序VS非降序数组:恢复二段性)
- [x] 😵‍💫在旋转过的有序数组中寻找目标值(严格升序VS非降序数组:恢复二段性)
- [x] 😵‍💫两个长度相等的排序数组中寻找中位数(上中位数(1.[二分法思路](https://www.bilibili.com/video/BV1BA411N7oe?_blank) 2.两数组中第N小的值))
- [ ] 寻找两个正序数组的中位数
#### 3、双指针与滑动窗口
- [x] 👏👏[双指针](https://labuladong.gitee.io/algo/2/21/59/?_blank)
- [x] 😵‍💫接雨水(1. leftMax, rightMax := 0, 0; 2. nums[left] <= nums[right])
- [x] 😵‍💫盛最多水的容器(LC11)
- [x] 👏👏[滑动窗口](https://labuladong.gitee.io/algo/1/11/?_blank)
- [ ] 最小覆盖子串
#### 4、回溯算法 
- [x] 👏👏[回溯算法框架](https://labuladong.gitee.io/algo/1/5/?_blank) (1、结束条件；2、剪枝-过滤父子兄弟重复节点以及改变起始位置；3、遍历、剪枝、选择、递归、撤销)
- [x] 😀集合的所有子集(1、排序 2、数组长度 3、遍历的起始值设置 4、回溯与记录)
- [x] 字符串的排列(下同)
- [x] 😵‍💫有重复项数字的全排列(剪枝：兄弟节点的过滤，i > 0 && !used[i-1] && num[i-1] == num[i])
- [x] 😵‍💫加起来和为目标值的组合(剪枝：) VS 树和为目标值的问题
- [ ] 😵‍💫通配符匹配
- [x] 😵‍💫括号生成
- [ ] ===== 图的遍历DFS ======
- [ ] 😵‍💫N皇后问题
- [ ] 😵‍💫数独
- [x] 机器人的运动范围(机器人在矩阵中行走,剪枝操作:1.越界 2.两数和 3.是否访问过)
- [x] 矩阵中的路径(1.矩阵的任意格子 2.directions列表)
#### 5、位运算与模拟
