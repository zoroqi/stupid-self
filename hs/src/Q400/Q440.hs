module Q400.Q440
    (findKthNumberPlanA
    ) where

import qualified Data.List as List

findKthNumberPlanA :: Int -> Int -> String
findKthNumberPlanA n k = last . take k . List.sort . map show $ [1..n]
