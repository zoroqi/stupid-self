module Q800.Q844
(
    backspaceComparePlanA
    )
where

backspaceComparePlanA :: String -> String -> Bool
backspaceComparePlanA s c =
                    let ss = m "" s
                        cs = m "" c
                    in ss == cs
                    where
                        m str [] = str
                        m str (x:xs) = if x == '#' then m (init str) xs
                                       else m (str++[x]) xs
