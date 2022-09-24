module Q400.Q400
(
    findNthDigitPlanA,
    findNthDigitPlanB,
) where

findNthDigitPlanA :: Int -> Int
findNthDigitPlanA n | n <= 0 = 0
findNthDigitPlanA n = last . take n . concat . map (reverse . splitNum) $ [1..]

splitNum :: Int -> [Int]
splitNum 0 = []
splitNum x = (x `mod` 10): splitNum (x `div` 10);

findNthDigitPlanB :: Int -> Int
findNthDigitPlanB n | n <= 0 = 0
findNthDigitPlanB n = planB m c where
    c = findRange n
    m = n - (sum [ 9 * x * (pow 10 (x-1)) | x<-[0..(c-1)]])

findRange :: Int -> Int
findRange n = local n 1 where
        local x1 x2
            | x1 <= 0 = x2 -1
            | otherwise = local (x1 - x2 * 9 * (pow 10 (x2-1))) (x2+1)

planB :: Int -> Int -> Int
planB m c = (splitNum num)!!index where
        num = ((m-1) `div` c) + (pow 10 (c-1))
        index = c - ((m-1) `mod` c) - 1

pow :: Int -> Int -> Int
pow _ 0 = 1
pow x1 x2 = foldl (*) 1 . take x2 . repeat $ x1
