module Q800.Q896 (
    isMonotonic
) where

-- isMonotonic :: [Int] -> Bool
-- isMonotonic nums = monotonic nums up || monotonic nums low
--     where
--         monotonic [] _ = True
--         monotonic [_] _ = True
--         monotonic [x,y] mon = mon x y
--         monotonic (x:y:xs) mon = if mon x y
--                                  then monotonic (y:xs) mon
--                                  else False
--         up a b = a <= b
--         low a b = a >= b

isMonotonic :: [Int] -> Bool
isMonotonic nums = monotonic nums True True
    where
        monotonic [] u l = u || l
        monotonic [_] u l = u || l
        monotonic [x,y] u l = (u && up x y) || (l && low x y)
        monotonic (x:y:xs) u l = monotonic (y:xs) (u && up x y) (l &&low x y)
        up a b = a <= b
        low a b = a >= b
