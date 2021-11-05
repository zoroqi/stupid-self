import Data.List

-- 全组合
combination :: [a] -> [[a]]
combination [] = [[]]
combination (x:xs) = combination xs ++ map (x:) (combination xs)

-- 分割
split :: Num a => (a -> Bool) -> [a] -> ([a], [a])
split f [] = ([],[])
split f (x:xs) = if f x
              then (x:fst(split f xs),snd(split f xs))
              else (fst(split f xs),x:snd(split f xs))

-- 所有时间组合
allTime = combination [0.0 .. 9.0]

-- 计算返回
readBinaryWatch :: (Floating a, Ord a) => ([a] -> Bool) -> [[a]]-> [(a, a)]
readBinaryWatch lengthFilter [] = []
readBinaryWatch lengthFilter xs = filter (\x -> fst x < 12 && snd x < 60) . map (\x -> (sum .map (2**) . fst $x, sum .map (\x->2**(x-4)) .snd $x)) . map (split (<4)) . filter lengthFilter $ xs

