module Q600.Q606 (
    tree2str
) where

import Struct

tree2str :: (Num a, Show a) => BinTree a -> String
tree2str Empty = ""
tree2str (Node v Empty Empty) = show v
tree2str (Node v l Empty) = (show v) ++ "(" ++ (tree2str l) ++ ")"
tree2str (Node v Empty r ) = (show v) ++ "()("++(tree2str r)++")"
tree2str (Node v l r ) = (show v) ++ "("++(tree2str l)++")(" ++ (tree2str r) ++")"
