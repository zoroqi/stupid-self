module Q400.Q492
(constructRectanglePlanB
    ) where

constructRectanglePlanB :: Integral a => a -> (a, a)
constructRectanglePlanB area = head . filter (\(l,m) -> l*m == area) . map (\x -> (div area x,x)) $ [m,m-1..1]
                        where m = floor . sqrt . fromIntegral $ area
