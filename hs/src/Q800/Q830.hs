module Q800.Q830 (
    largeGroupPositions
) where

import Data.List(group)

largeGroupPositions :: String -> [[Int]]
largeGroupPositions = let 
        plana  (tl, r) l = (tl+l, (tl,l):r)
    in 
    map (\(begin,l) -> [begin,begin+l-1]) . filter (\(_,l) -> l >=3 ) . snd . foldl plana (0,[]) . map length . group