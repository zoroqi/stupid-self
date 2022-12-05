module Q800.Q829
(
    consecutiveNumbersSum
)
where

consecutiveNumbersSum :: Int -> Int
consecutiveNumbersSum n =
    let
        doubleN = 2 * n
        lMax = round . sqrt . fromIntegral $ doubleN
        f l = (doubleN-l*l+l) `mod` (2*l) == 0
    in
        length . filter f $ [1..lMax]
