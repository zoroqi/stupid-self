module Q400.Q406
    (queuePlanA
    ) where

import Tools(permutation)

queuePlanA :: (Ord a, Num b, Ord b) => [(a,b)] -> [[(a,b)]]
queuePlanA [] = []
queuePlanA xs = filter queueFilter (permutation xs)

-- 判断数量和
queueFilter :: (Ord a, Num b, Ord b) => [(a,b)] -> Bool
queueFilter [] = True
queueFilter xs = (checkLastPeople xs) && queueFilter (init xs)

checkLastPeople :: (Ord a, Num b, Ord b) => [(a,b)] -> Bool
checkLastPeople [] = False
checkLastPeople xs = f2 (init xs) (last xs) where
        f2 initx lastx = case initx of
            [] -> 0 == snd lastx
            xx -> count xx (fst lastx) 0 == snd lastx

-- 统计身高 >= 的数量
count :: (Ord a, Num c) => [(a,b)] -> a -> c -> c
count [] _  c = c
count (x:xs) n c = if (fst x) >= n
                     then count xs n c+1
                     else count xs n c
