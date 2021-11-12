module Q400.Q404
    (sumOfLeftLeaves
    ) where

import Struct

sumOfLeftLeaves :: Num a => BinTree a -> a
sumOfLeftLeaves tree = dfs tree False

dfs :: Num a => BinTree a -> Bool -> a
dfs Empty _ = 0
dfs (Node k l r) b = val + (dfs l True) + (dfs r False) where
    val = case (l, r, b) of
        (Empty, Empty, True) -> k
        (_, _, _ ) -> 0
