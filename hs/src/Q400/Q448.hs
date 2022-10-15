module Q400.Q448
(findDisappearedNumbersPlanA,
    findDisappearedNumbersPlanB
    ) where

import qualified Data.HashSet as HashSet
import qualified Data.List as List

findDisappearedNumbersPlanA :: [Integer] -> [Integer]
findDisappearedNumbersPlanA [] = []
findDisappearedNumbersPlanA n = filter (\x -> not (HashSet.member x fdn)) [1 .. len]
            where
                fdn = HashSet.fromList n
                len = toInteger . length $n


findDisappearedNumbersPlanB :: [Int] -> [Int]
findDisappearedNumbersPlanB [] = []
findDisappearedNumbersPlanB n = planb nl [1..(length n)]
            where
                nl = List.nub . List.sort $ n


planb :: [Int] -> [Int] -> [Int]
planb _ [] = []
planb [] y = y
planb (x@(xs:xss)) (y:ys)
             | xs == y = planb xss ys
             | xs > y  = y : planb x ys
             | otherwise = ys

-- [_] [_]
-- [_] (_:_:_)
-- (_:_:_) [_]
-- (_:_:_) (_:_:_)

-- n := len(nums)
-- for _, v := range nums {
--     v = (v - 1) % n
--     nums[v] += n
-- }
-- for i, v := range nums {
--     if v <= n {
--         ans = append(ans, i+1)
--     }
-- }
-- return
-- findDisappearedNumbersPlanC :: [Int] -> [Int]
-- findDisappearedNumbersPlanC [] = []
-- findDisappearedNumbersPlanC n = planc n (length n)
--
-- planc :: [Int] -> Int -> [Int]



