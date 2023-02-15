module Q800.Q872 (
    leafSimilar
) where

import Struct(BinTree(..), isLeaf)

leafSimilar :: Eq a => BinTree a -> BinTree a -> Bool
leafSimilar r1 r2 = pre r1 == pre r2
    where
        pre Empty = []
        pre n@(Node v l r) = if isLeaf n
                             then
                                v:[]
                             else
                                pre(l)++pre(r)
