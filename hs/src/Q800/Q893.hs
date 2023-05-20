module Q800.Q893
(
    numSpecialEquivGroupsPlanA
) where

import Data.List(sort,group)

numSpecialEquivGroupsPlanA :: [String] -> Int
numSpecialEquivGroupsPlanA  = length . group . sort . map (\(e,o) -> e++o) . map (\(e,o) -> (sort e,sort o)) . map split
    where
        split = \list ->
            case list of
                [] -> ([], [])
                x:xs ->
                    let (evens, odds) = split xs
                    in (x:odds, evens)
