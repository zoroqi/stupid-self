module Q400.Q496
(nextGreaterElementPlanA
    ) where

nextGreaterElementPlanA :: Integral a => [a] -> [a] -> [a]
nextGreaterElementPlanA [] _ = []
nextGreaterElementPlanA _ [] = []
nextGreaterElementPlanA (n:ns) n2 = (plana n (finda n n2)):(nextGreaterElementPlanA ns n2)

finda :: Integral a => a -> [a] -> [a]
finda _ [] = []
finda n (x:xs) = if x == n
                 then xs
                 else finda n xs

plana :: Integral a => a -> [a] -> a
plana _ [] = -1
plana n (x:xs) = if x > n
                then x
                else plana n xs

