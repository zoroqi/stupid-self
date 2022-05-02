module Q400.Q456
(find132patternPlanA
    ) where

find132patternPlanA :: [Int] -> Bool
find132patternPlanA [] = False
find132patternPlanA [_] = False
find132patternPlanA [_,_] = False
find132patternPlanA nums = not . null . filter (\(x,y,z) -> x < z && z < y) . plana $ nums

plana :: [Int] -> [(Int,Int,Int)]
plana (x:y:[]) = []
plana nums = (ff nums) ++ (plana (tail nums))
                where
                    ff (a:b:[]) = []
                    ff (a:b:c) = [(i,j,k) | i <- [a], j <-[b],k <-c] ++ (ff (a:c))
