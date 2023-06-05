module Q600.Q668 (
    findKthNumberPlanA,
    findKthNumberPlanD
) where

import Data.List(sort)

-- 不做任何检查, 简单暴力方案
findKthNumberPlanA :: Int -> Int -> Int -> Int
findKthNumberPlanA m n k = last . take k $ sort [x*y | x <- [1..m], y<-[1..n]]

findKthNumberPlanD :: Int -> Int -> Int -> Int
findKthNumberPlanD m n k = binsearch 0 (m*n) (\x -> (count x) >= k)
    where
        binsearch l r f =
            let
               h = (l+r) `div` 2
            in
                if l < r then
                    if f h then binsearch l h f
                    else binsearch (h+1) r f
                else l
        count x =
            let
                c = (x `div` n) * n
            in
                c + (foldr (+) 0 [ x `div` a | a <-[x`div`n+1..m]])

-- [10..4] 返回的是空集合, 而 [10,9..4] 返回的是 10 到 4 部分, 很智能啊.
