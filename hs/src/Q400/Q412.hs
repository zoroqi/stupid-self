module Q400.Q412 
    (fizzBuzz,
        fizzBuzzPlanB
    ) where

fizzBuzz :: Integer -> [String]
fizzBuzz 0 = []
fizzBuzz n = map fb [1..n]

fb :: Integer -> String
fb x | x `mod` 3 == 0 && x `mod` 5 == 0 = "FizzBuzz"
     | x `mod` 3 == 0 = "Fizz"
     | x `mod` 5 == 0 = "Buzz"
     | otherwise = show x

fizzBuzzPlanB :: Integer -> [String]
fizzBuzzPlanB 0 = []
fizzBuzzPlanB n = map fb1 [1..n]
         where
            fb1 x
                | x `mod` 3 == 0 && x `mod` 5 == 0 = "FizzBuzz"
                | x `mod` 3 == 0 = "Fizz"
                | x `mod` 5 == 0 = "Buzz"
                | otherwise = show x

