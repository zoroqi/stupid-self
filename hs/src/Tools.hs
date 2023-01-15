module Tools
    ( combination,
      permutation,
      qsort,
      sublist
    ) where

import Data.List

combination :: [a] -> [[a]]
combination [] = [[]]
combination (x:xs) = combination xs ++ map (x:) (combination xs)

permutation :: Eq a => [a] -> [[a]]
permutation [] = [[]]
permutation xs = [y:ys | y <- xs, ys <- permutation (delete y xs)]

qsort :: Ord a => [a] -> [a]
qsort [] = []
qsort (x:xs) = smaller ++ [x] ++ bigger
        where
          smaller = qsort [a | a<-xs, a <= x]
          bigger = qsort [a | a <-xs, a > x]

sublist :: Int -> Int -> [a] -> [a]
sublist s e = drop s . take e
