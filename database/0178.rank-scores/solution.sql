-- 解题思路：
-- dense_rank() 函数

SELECT score, dense_rank() over (ORDER BY score DESC) AS 'rank'
FROM Scores;