module Q800.Q884
(
    uncommonFromSentences
) where

import qualified Data.List as List

uncommonFromSentences :: String -> String -> [String]
uncommonFromSentences s1 s2 = concat . filter (\x -> length x == 1) . List.group . List.sort $ (ws1 ++ ws2)
    where
        ws1 = words s1
        ws2 = words s2
