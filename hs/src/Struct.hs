{-# LANGUAGE DeriveFoldable #-}
{-# LANGUAGE DeriveFunctor #-}
module Struct
    ( BinTree(..),
        leaf,
        SStack,
        spush,
        spop,
        sempty,
        speek,
        isLeaf,
        listToTree,
        preOrder, inOrder, postOrder,
        bfsOrder
    ) where

import Control.Monad.State
    ( evalState, MonadState(put, get), State )

data BinTree a = Empty | Node a (BinTree a) (BinTree a) deriving(Eq,Show, Foldable, Functor)
--data BinTree a = Empty | Node a (BinTree a) (BinTree a) deriving(Eq,Show)

leaf :: a -> BinTree a
leaf a = Node a Empty Empty

isLeaf :: BinTree a -> Bool
isLeaf (Node _ Empty Empty) = True
isLeaf _ = False

-- 简单通过数组生成一棵二叉树, 第二个参数是用来表示 nil 值的.
-- 根节点不进行 nil 判断
listToTree :: Eq a => [a] -> a -> BinTree a
listToTree arr empty = buildTree 0 1 2
    where
        len = length arr
        buildTree n l r = case (((l >= len) || (arr!!l == empty)), ((r >= len) || (arr!!r == empty))) of
                (True, True) -> leaf (arr!!n)
                (True, False) -> Node (arr!!n) Empty (buildTree r (lNodeIndex r) (rNodeIndex r))
                (False, True) -> Node (arr!!n) (buildTree l (lNodeIndex l) (rNodeIndex l)) Empty
                (False, False) -> Node (arr!!n) (buildTree l (lNodeIndex l) (rNodeIndex l)) (buildTree r (lNodeIndex r) (rNodeIndex r))
        rNodeIndex num = (num + 1) * 2
        lNodeIndex num = (rNodeIndex num)- 1

-- 尝试自己实现, 最后优先使用 DeriveFoldable. 三个递归, 分别对应前, 中, 后的遍历
--instance Foldable BinTree where
--    foldr f z Empty = z
--    foldr f z (Node v l r) = f v (foldr f (foldr f z r) l)
--    foldr f z (Node v l r) = foldr f (f v (foldr f z r)) l
--    foldr f z (Node v l r) = foldr f (foldr f (f v z) r) l
--    foldMap _ Empty = mempty
--    foldMap f (Node v l r) = mappend (mappend (f v) (foldMap f (PROTree l))) (foldMap f (PROTree r))

preOrder :: BinTree a -> [a]
preOrder n = dfsOrder n (\v l r -> [v] ++ l ++ r)

inOrder :: BinTree a -> [a]
inOrder n = dfsOrder n (\v l r -> l ++[v]++ r)

postOrder :: BinTree a -> [a]
postOrder n = dfsOrder n (\v l r ->  l ++ r ++ [v])

dfsOrder :: BinTree a -> (a -> [a] -> [a] -> [a]) -> [a]
dfsOrder Empty _ = []
dfsOrder (Node v l r) merge = let
        ll = dfsOrder l merge
        rr = dfsOrder r merge
    in
        merge v ll rr

-- 就是尝试用 State 来实现
-- bfsOrder2 :: (Show a) => BinTree a -> [a]
-- bfsOrder2 Empty = []
-- bfsOrder2 root = do
--         evalState qo2 ([root], [])
--     where
--         qo2 :: State ([BinTree a], [a]) [a]
--         qo2 = do
--             (q,r) <- get
--             if null q then return r
--             else do
--                 case head q of 
--                     Empty -> put (tail q, r)
--                     (Node v left right) -> put (tail (q ++ [left,right]), v:r)
--                 qo2

bfsOrder :: BinTree a -> [a]
bfsOrder Empty = []
bfsOrder root = bfs [root] []
    where
        bfs [] r = r
        bfs (x:xs) r =
            case x of
                Empty -> bfs xs r
                (Node v left right) -> bfs (xs++[left,right]) (r++[v])

type SStack a = [a]

spush :: SStack a -> a -> SStack a
spush s n = s ++ [n]

spop :: SStack a -> SStack a
spop s = init s

sempty :: SStack a -> Bool
sempty s = null s

speek :: SStack a -> a
speek s = last s
