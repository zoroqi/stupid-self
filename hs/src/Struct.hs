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
        preOrder, inOrder, postOrder
    ) where

data BinTree a = Empty | Node a (BinTree a) (BinTree a) deriving(Eq,Show)

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


preOrder :: Show a => BinTree a -> [a]
preOrder n = dfsOrder n (\v l r -> [v] ++ l ++ r)

inOrder :: Show a => BinTree a -> [a]
inOrder n = dfsOrder n (\v l r -> l ++[v]++ r)

postOrder :: Show a => BinTree a -> [a]
postOrder n = dfsOrder n (\v l r ->  l ++ r ++ [v])

dfsOrder :: BinTree a -> (a -> [a] -> [a] -> [a]) -> [a]
dfsOrder Empty _ = []
dfsOrder (Node v l r) merge = let
        ll = dfsOrder l merge
        rr = dfsOrder r merge
    in
        merge v ll rr


type SStack a = [a]

spush :: SStack a -> a -> SStack a
spush s n = s ++ [n]

spop :: SStack a -> SStack a
spop s = init s

sempty :: SStack a -> Bool
sempty s = null s

speek :: SStack a -> a
speek s = last s

