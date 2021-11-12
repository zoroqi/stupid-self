module Struct
    ( BinTree(..)
    ) where

data BinTree a = Empty | Node a (BinTree a) (BinTree a)
