module Q800.Q868
(
    binaryGapPlanA
) where 

binaryGapPlanA :: Int -> Int
binaryGapPlanA n = plana (divMod n 2) (-1) 0 0
    where
        plana (0, 0) _ _ m = m
        plana (b1,b2) before i m = plana (divMod b1 2) before' (i+1) m'
            where
                before' = if b2 == 1
                          then i
                          else before
                m' = if before > (-1) && b2 == 1
                     then max (i - before) m
                     else m
