module Q800.Q867
(
    transpose,
    tt2
) where

transpose :: [[a]] -> [[a]]
transpose [] = []
transpose [xs] = [[x] | x <- xs]
transpose (x:xs) = zipWith (:) x (transpose xs)



-- 特殊版本
-- transpose :: [[a]] -> [[a]]
-- transpose = foldr (zipWith (:)) (repeat [])
