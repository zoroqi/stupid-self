module Q800.Q889 (
    constructFromPrePost
) where

import Struct(BinTree(..), leaf)
import Tools(sublist)

-- [1, 2, 4, 5, 3, 6, 7]
-- [4, 5, 2, 6, 7, 3, 1]
constructFromPrePost :: [Int] -> [Int] -> BinTree Int
constructFromPrePost pre post =
    let
        preS = 1
        postE = (length pre) - 2
        preE = findIndex (post!!postE) pre
        postS = findIndex (pre!!preS) post
    in
        if length pre == 1
        then leaf (pre!!0)
        else
            if (pre!!preS) == (post!!postE)
            then Node (pre!!0) (constructFromPrePost (tail pre) (init post)) Empty
            else
                Node (pre!!0) (constructFromPrePost (sublist preS preE pre) (take postS post)) (constructFromPrePost (drop preE pre) (sublist (postS+1) (postE+1) post))


findIndex :: Int -> [Int] -> Int
findIndex target arr = find target arr 0
    where
        -- 在这个算法内一定存在
        find t (x:xs) n= if t == x
            then n
            else find t xs n+1
        find _ [] n = n - 1
