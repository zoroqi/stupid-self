module Q800.Q805
(
    splitArraySameAveragePlanC
) where

import Data.Foldable(foldl')

splitArraySameAveragePlanC :: [Int] -> Bool
splitArraySameAveragePlanC xs = let
        s = sum xs
        l = length xs
        l2 = l `div` 2
        initlist = (take l2 (repeat []))++[[(0,0)]]
    in
        any (\(x,y) -> l*y == s*x) . concat $ init (foldl' (\y x-> planc (+x) l2 y) initlist xs)
    where
        planc :: (Int -> Int) -> Int -> [[(Int,Int)]]  -> [[(Int,Int)]]
        planc addnum  0 x = x
        planc addnum  n (x:y:xs) = (x++[(n, ns) | ns <- map (addnum.snd) y]) : (planc addnum (n-1) (y:xs))

-- (x++map (\(_, s') -> (n, addnum s')) y) 这是最开始的最开始计算方式

-- init 是用来舍弃 [(0,0)] 这个元素用的
