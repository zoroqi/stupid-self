module Q800.Q806
(
    numberOfLinesPlanA,
    numberOfLinesPlanB
) where

numberOfLinesPlanA :: [Int] -> String ->(Int,Int)
numberOfLinesPlanA widths s = count s 1 0
                where
                    len n = widths !! (fromEnum n - 97)
                    count :: [Char] -> Int -> Int -> (Int,Int)
                    count [] a b = (a,b)
                    count (x:xs) a b = if b + l > 100
                                       then count xs (a+1) l
                                       else count xs a $ b+l
                        where l = len x

numberOfLinesPlanB :: [Int] -> String -> (Int,Int)
numberOfLinesPlanB widths s = last nums
                    where
                    len n = widths !! (fromEnum n - 97)
                    nums = (1,0):zipWith count nums s
                        where
                            count (a, b) x = if b + l > 100
                                               then (a+1, l)
                                               else (a, b+l)
                               where l = len x
