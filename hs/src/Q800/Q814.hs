module Q800.Q814
(
    pruneTreePlanA
) where

import qualified Struct as Tree

pruneTreePlanA :: Integral a => Tree.BinTree a -> Tree.BinTree a
pruneTreePlanA Tree.Empty = Tree.Empty
pruneTreePlanA (Tree.Node a l r) =
                            case (nl, nr, a) of
                                (Tree.Empty, Tree.Empty, 0) -> Tree.Empty
                                (_, _, _) -> Tree.Node a nl nr
                            where
                                nl = pruneTreePlanA l
                                nr = pruneTreePlanA r
