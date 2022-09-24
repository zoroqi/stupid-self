module Q400.Q401
    (readBinaryWatch,
        allTime
    )
    where

import Tools (combination)

-- 分割
split :: Num a => (a -> Bool) -> [a] -> ([a], [a])
split _ [] = ([],[])
split f (x:xs) = if f x
              then (x:fst(split f xs),snd(split f xs))
              else (fst(split f xs),x:snd(split f xs))

-- 所有时间组合
allTime :: [[Double]]
allTime = combination [0.0 .. 9.0]

-- 计算返回
readBinaryWatch :: (Floating a, Ord a) => ([a] -> Bool) -> [[a]]-> [(a, a)]
readBinaryWatch _ [] = []
readBinaryWatch lengthFilter xs = filter (\x -> fst x < 12 && snd x < 60) .
             map (\x -> (
                    sum . map (2**) . fst $ x,
                    sum . map (\y->2**(y-4)) .snd $ x
                    )
                 ) . map (split (<4)) . filter lengthFilter $ xs

