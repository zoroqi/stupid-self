module Q800.Q801
(
    minSwapPlanB
) where

minSwapPlanA :: [Int] -> [Int] -> ([Int],[Int])
minSwapPlanA x y = plana x y
        where
            plana [] [] = ([],[])
            plana [x] [y] = ([x],[y])
            plana (x1:xs@(x2:_)) (y1:ys@(y2:_))
                | x1 >= x2 || y1 >= y2 = (\(nx,ny) -> (x1:nx, y1:ny)) (plana ys xs)
                | otherwise = (\(nx,ny) -> (x1:nx, y1:ny)) (plana ys xs)

minSwapPlanB :: [Int] -> [Int] -> Int
minSwapPlanB x y = planb x y 0 1
        where
            planb [] [] dp0 dp1 = min dp0 dp1
            planb [_] [_] dp0 dp1 = min dp0 dp1
            planb (x1:xs@(x2:_)) (y1:ys@(y2:_)) dp0 dp1
                | x1 < x2 && y1 < y2 && x1 < y2 && y1 < x2 = planb xs ys (min dp0 dp1) ((min dp0 dp1) + 1)
                | x1 < y2 && y1 < x2 = planb xs ys dp1 (dp0+1)
                | otherwise = planb xs ys dp0 (dp1 + 1)
