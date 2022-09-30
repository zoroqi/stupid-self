module Q400.Q441 
    (arrangeCoins,
     arrangeCoinsPlanA
    ) where

arrangeCoins :: Integer -> Integer
arrangeCoins 0 = 0
arrangeCoins n = floor $ ((sqrt . fromIntegral $ (1+8*n)::Double) - 1) / 2


arrangeCoinsPlanA :: Integer -> Integer
arrangeCoinsPlanA 0 = 0
arrangeCoinsPlanA n = pa 1 n
        where
            pa num total
                | total >= num = pa (num+1) (total-num)
                | otherwise = num -1
