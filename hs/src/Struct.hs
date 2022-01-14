module Struct
    ( BinTree(..),
        SStack(..)
    ) where

data BinTree a = Empty | Node a (BinTree a) (BinTree a)

type SStack a = [a]

spush :: SStack a -> a -> SStack a
spush s n = s ++ [n]

spop :: SStack a -> SStack a
spop s = init s

sempty :: SStack a -> Bool
sempty s = null s

speek :: SStack a -> a
speek s = last s

