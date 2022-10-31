module Q800.Q876
(
    middleNode
) where

middleNode :: [a] -> [a]
middleNode arr = plana arr arr
    where
        plana [] s = s
        plana [_] s = s
        plana [_,_] (_:sx) = sx
        plana (_:_:fx) (_:sx) = plana fx sx
