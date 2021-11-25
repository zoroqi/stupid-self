module Q400.Q448
(findDisappearedNumbersPlanA
    ) where

import qualified Data.HashSet as HashSet

findDisappearedNumbersPlanA :: [Integer] -> [Integer]
findDisappearedNumbersPlanA [] = []
findDisappearedNumbersPlanA n = filter (\x -> not (HashSet.member x fdn)) [1 .. len]
            where
                fdn = HashSet.fromList n
                len = toInteger . length $n
