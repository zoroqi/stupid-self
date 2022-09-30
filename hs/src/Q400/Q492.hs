module Q400.Q492
(constructRectanglePlanB
    ) where

constructRectanglePlanB :: Integral a => a -> (a, a)
constructRectanglePlanB area = head . filter (\(l,m) -> l*m == area) . map (\x -> (div area x,x)) $ [mm,mm-1..1]
                    where mm = floor . (sqrt :: Double -> Double) . fromIntegral $ area
