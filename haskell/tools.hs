import Data.List

combination :: [a] -> [[a]]
combination [] = [[]]
combination (x:xs) = combination xs ++ map (x:) (combination xs)
