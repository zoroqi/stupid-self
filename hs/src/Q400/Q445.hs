module Q400.Q445
    (addTwoNumbersPlanA
    ) where

addTwoNumbersPlanA :: Integral a => [a] -> [a] -> [a]
addTwoNumbersPlanA [] [] = []
addTwoNumbersPlanA [] x = x++[]
addTwoNumbersPlanA x [] = x++[]
addTwoNumbersPlanA x1 x2 = sumPlanA (reverse x1) (reverse x2) 0 []

sumPlanA :: Integral a =>  [a] -> [a] -> a -> [a] -> [a]
sumPlanA [] [] 0 xs = xs
sumPlanA [] [] x xs = x:xs
sumPlanA [] (r:rs) x xs = sumPlanA [] rs ((x+r) `div` 10) (((x+r) `mod` 10):xs)
sumPlanA (l:ls) [] x xs = sumPlanA ls [] ((x+l) `div` 10) (((x+l) `mod` 10):xs)
sumPlanA (l:ls) (r:rs) x xs = sumPlanA ls rs ((x+l+r) `div` 10 ) (((x+l+r) `mod` 10):xs)

numm :: (Int, String) b => b -> b
numm b -> b
