module Q200.Q233(
    countDigitOnePlanB
) where


countDigitOnePlanB :: (Integral a) => a -> a
countDigitOnePlanB num = dfs num 0
    where
        dfs 0 count = count
        dfs n count =
            let
                l = floor (logBase 10.0 (fromIntegral n))
                base = floor (10 ** (fromIntegral l))
                c09 = l * floor (10 ** (fromIntegral (l-1)))
            in
                if n `div` base == 1
                then dfs (n `mod` base) (count + c09 + (n `mod` base) + 1)
                else dfs (n `mod` base) (count + (n `div` base) * c09 + base)
