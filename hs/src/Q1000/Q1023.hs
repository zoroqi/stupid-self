module Q1000.Q1023(
    camelMatchPlanA
) where

import Text.Regex.Posix

camelMatchPlanA:: [String] -> String -> [Bool]
camelMatchPlanA queries pattern =
    let
       reg = ("[a-z]*"++concat [v:"[a-z]*"|v<-pattern])
    in
        map (=~ reg) queries
