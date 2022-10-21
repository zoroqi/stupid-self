module Q100.Q152
(
    maxProduct
) where

maxProduct :: [Int] -> Int
maxProduct [] = 0
maxProduct (x:xs) = plana (x,x,x) xs

plana :: (Int,Int,Int) -> [Int] -> Int
plana (m, _, _) [] = m
plana (m, imax, imin) (x:xs)
    | x < 0 =
        let
            imax' = max (imin*x) x
            imin' = min (imax*x) x
            m' = max m imax'
        in
            plana (m', imax', imin') xs
    | otherwise =
        let
            imax' = max (imax*x) x
            imin' = min (imin*x) x
            m' = max imax' m
        in
            plana (m', imax', imin') xs
