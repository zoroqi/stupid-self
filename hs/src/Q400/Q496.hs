module Q400.Q496
(nextGreaterElementPlanA,
    nextGreaterElementPlanC
    ) where

import qualified Data.HashSet as HashSet

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

nextGreaterElementPlanC :: Integral a => [a] -> [a] -> [a]
nextGreaterElementPlanC [] _ = []
nextGreaterElementPlanC _ [] = []
nextGreaterElementPlanC n1 n2 = map (\(_,x) -> x) . concat . f n1 . planc $ n2
    where f (n:ns) ns2 = filter (\(x,y) -> x == n) ns2:f ns ns2
          f [] ns2 = []

planc :: Integral a => [a] -> [(a,a)]
planc [] = []
planc ns = f ns [] []
    where
        f :: Integral a => [a] -> [a] -> [(a,a)] -> [(a,a)]
        f [] [] r= r
        f [] (s:ss) r = f [] ss ((s, -1):r)
        f (n:ns) [] r = f ns [n] r
        f (n:ns) s r
          | n > last s = f (n:ns) (init s) ((last s,n):r)
          | n < last s = f ns (s ++ [n]) r

