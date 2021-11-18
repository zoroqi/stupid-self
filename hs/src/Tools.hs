module Tools
    ( combination,
      permutation
    ) where

import Data.List

combination :: [a] -> [[a]]
combination [] = [[]]
combination (x:xs) = combination xs ++ map (x:) (combination xs)

permutation :: Eq a => [a] -> [[a]]
permutation [] = [[]]
permutation xs = [y:ys | y <- xs, ys <- permutation (delete y xs)]
