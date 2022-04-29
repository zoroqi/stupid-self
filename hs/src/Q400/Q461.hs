module Q400.Q461
(hammingDistancePlanA
    ) where

hammingDistancePlanA ::  Int -> Int -> Int
hammingDistancePlanA a b = length . filter (\x -> fst x /= snd x) . zip (toReverseBinPlanB a 32) $ (toReverseBinPlanB b 32)

toReverseBinPlanA :: Int -> Int -> [Int]
toReverseBinPlanA n 0 = []
toReverseBinPlanA n 1 = [mod n 2]
toReverseBinPlanA n l = map (\x -> snd x) bin
        where
            bin = (div n 2, mod n 2) : zipWith (\b1 b2 -> (div (fst b1) 2, mod (fst b1) 2)) bin [1..l-1]

toReverseBinPlanB :: Int -> Int -> [Int]
toReverseBinPlanB n 0 = []
toReverseBinPlanB n 1 = [mod n 2]
toReverseBinPlanB n l = map (\x -> snd x) bin
        where
            bin = (divMod n 2) : zipWith (\b1 b2 -> divMod (fst b1) 2) bin [1..l-1]


-- 互联网找到一个比较接受的方案

helper :: Int -> [Int]
helper 0 = []
helper n = (n `mod` 2) : helper (n `div` 2)
