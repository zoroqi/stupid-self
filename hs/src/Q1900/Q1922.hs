module Q1900.Q1922 (
    countGoodNumbers
) where

countGoodNumbersMod :: Integer
countGoodNumbersMod = 1000000007

countGoodNumbers :: Integer -> Integer
countGoodNumbers n = if (n `mod` 2) == 0 then pow 20 (n `div` 2)
                     else pow 20 (n `div` 2) * 5 `mod` countGoodNumbersMod
                where
                    pow _ 0 = 1
                    pow x y =
                        let
                           num = pow x (y `div` 2) `mod` countGoodNumbersMod
                        in
                            if y `mod` 2 == 1 then (num * num * x)  `mod` countGoodNumbersMod
                              else (num * num) `mod` countGoodNumbersMod

--pow _ 0 r = r
--pow x y r = if y `mod` 2 == 1 then pow (x*x) (y `div` 2) (r * x)
--             else pow (x*x) (y `div` 2) r
