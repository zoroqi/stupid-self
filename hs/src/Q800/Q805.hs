module Q800.Q805
(
    splitArraySameAveragePlanC,
    splitArraySameAveragePlanB
) where

import Data.Foldable(foldl')
import Data.List(subsequences)

splitArraySameAveragePlanC :: [Int] -> Bool
splitArraySameAveragePlanC xs = let
        s = sum xs
        l = length xs
        l2 = l `div` 2
        initlist = (take l2 (repeat []))++[[(0,0)]]
    in
        any (\(x,y) -> l*y == s*x) . concat $ init (foldl (\y x-> planc (+x) l2 y) initlist xs)
    where
        planc :: (Int -> Int) -> Int -> [[( Int,Int )]]  -> [[( Int, Int )]]
        planc addnum  0 x = x
        planc addnum  n (x:y:xs) = (x++[(n, ns) | ns <- map (addnum.snd) y]) : (planc addnum (n-1) (y:xs))

-- (x++map (\(_, s') -> (n, addnum s')) y) 这是最开始的最开始计算方式

-- init 是用来舍弃 [(0,0)] 这个元素用的

splitArraySameAveragePlanB :: [Int] -> Bool
splitArraySameAveragePlanB arr =
    let
        s = sum arr
        l = length arr
    in
        any (\(as,al) -> as*l == s*al) . filter (\(_,al) -> al /= 0 && al /= l) . map (\a -> (sum a,length a)) $  (subsequences arr)
