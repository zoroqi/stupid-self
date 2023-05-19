module Q800.Q894 (
    allPossibleFBTPlanC
) where

import Struct

allPossibleFBTPlanC :: Num a => Int -> [BinTree a]
allPossibleFBTPlanC n = dfs n
    where 
        dfs 0 = []
        dfs 1 = (leaf 0):[]
        dfs num = concat [[Node 0 l r | l <- (dfs x), r <- (dfs (num-1-x))] | x <- [1..num-1] ]