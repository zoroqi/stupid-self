module Q800.Q852 (
    peakIndexInMountainArrayPlanB,
    peakIndexInMountainArrayPlanC
) where

peakIndexInMountainArrayPlanB :: [Int] -> Int
peakIndexInMountainArrayPlanB arr = top 0 (length arr -1)
    where
        top l r = if l < r
                  then let
                          m = (l+r) `div` 2
                       in
                          if (arr!!m) > (arr!!(m+1)) then top l m
                          else top (m+1) r
                  else l

peakIndexInMountainArrayPlanC :: [Int] -> Int
peakIndexInMountainArrayPlanC = third . foldl f (-1, -1, 0)
    where
        f (index,m,i) x = if m > x then (index,m,i+1)
                    else (i, x, i+1)
        third (a,_,_) = a
