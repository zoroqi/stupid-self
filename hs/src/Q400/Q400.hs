module Q400.Q400
(
    findNthDigitPlanA,
    findNthDigitPlanB,
    splitNum
) where

findNthDigitPlanA :: Int -> Integer
findNthDigitPlanA 0 = 0
findNthDigitPlanA n = last . take n . concat . map (reverse . splitNum) $ [1..]

splitNum :: Integer -> [Integer]
splitNum 0 = []
splitNum x = (x `mod` 10): splitNum (x `div` 10);

findNthDigitPlanB :: Int -> Integer
findNthDigitPlanB 0 = 0
findNthDigitPlanB n = last . take n . concat . map (reverse . splitNum) $ [1..]
