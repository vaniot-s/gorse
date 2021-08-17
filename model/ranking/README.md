### ß4个排序模型
A，als：Alternative -Least-Squares
交替最小二乘法
它通过最小化误差的平方和寻找数据的最佳函数匹配。利用最小二乘法可以简便地求得未知的数据，并使得这些求得的数据与实际数据之间误差的平方和为最小。最小二乘法可用于曲线拟合。
B，bpr：Bayesian Personalized Ranking
基于贝叶斯后验优化的个性化排序算法
BPR算法将用户对物品的评分（显示反馈“1”，隐式反馈“0”）处理为一个pair对的集合<i,j>，其中i为评分为1的物品，j为评分为0的物品。假设某用户有M个“1”的评分，N个“0”的评分，则该用户共有M*N个pair对。
这样数据集就由三元组 <u,i,j>表示，该三元组的物理含义为：相对于物品“j”，用户“u”更喜欢物品“i”。
C，ccd：Cyclic Coordinate Descent (CCD)
是一个启发式的迭代搜索算法，
D，knn：k-Nearest Neighbor
K最近邻分类算法
如果一个样本在特征空间中的k个最相似(即特征空间中最邻近)的样本中的
大多数属于某一个类别，则该样本也属于这个类别