module Q800.Q824
(
    toGoatLatinPlanA
) where

toGoatLatinPlanA :: String -> String
toGoatLatinPlanA str = unwords . map goat $ (zipWith (,) (words str) [0..])
                where
                    goat (s,n) = case (s,n) of
                        (o@('a':_), cc) -> vowel o cc
                        (o@('e':_), cc) -> vowel o cc
                        (o@('i':_), cc) -> vowel o cc
                        (o@('o':_), cc) -> vowel o cc
                        (o@('u':_), cc) -> vowel o cc
                        ((x:xs), cc) -> consonant xs x cc
                        ([],_) -> ""
                    vowel s n = s++"ma"++(alist n)
                    consonant s t n = s++[t] ++"ma"++(alist n)
                    alist n = take n . repeat $ 'a'
