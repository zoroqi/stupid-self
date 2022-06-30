module Q800.Q831
(
    maskPII
) where

import Text.Regex.Posix as Regex
import Data.Char

maskPII :: String -> String
maskPII s = if elem '@' s
            then mailPlanA s
            else phonePlanA s

mailPlanA :: String -> String
mailPlanA s = map toLower ((ss!!1)++"*****"++(ss!!3)++"@"++(ss!!4))
            where
                ss = concat ( s =~ "(.)(.*)(.)@(.+)" :: [[String]])

phonePlanA :: String -> String
phonePlanA s = case length phone of
                10 -> "***-***-"++(drop 6 phone)
                11 -> "+*-***-***-"++(drop 7 phone)
                12 -> "+**-***-***-"++(drop 8 phone)
                13 -> "+***-***-***-"++(drop 9 phone)
                _ -> phone
    where
        phone = filter (\x -> fromEnum x >= 48 && fromEnum x <= 57) s

