module Struct
    ( BinTree(..),
        leaf,
        SStack,
        spush,
        spop,
        sempty,
        speek,
        isLeaf
    ) where

data BinTree a = Empty | Node a (BinTree a) (BinTree a) deriving(Eq,Show)

leaf :: a -> BinTree a
leaf a = Node a Empty Empty

isLeaf :: BinTree a -> Bool
isLeaf (Node _ Empty Empty) = True
isLeaf _ = False

type SStack a = [a]

spush :: SStack a -> a -> SStack a
spush s n = s ++ [n]

spop :: SStack a -> SStack a
spop s = init s

sempty :: SStack a -> Bool
sempty s = null s

speek :: SStack a -> a
speek s = last s

