module Q800.Q804
(
    uniqueMorseRepresentationsPlanA
) where

import Data.List as List

morse = [".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."]

uniqueMorseRepresentationsPlanA :: [String] -> Int
uniqueMorseRepresentationsPlanA ws = length . List.nub . map toCode $ ws
                                where
                                    charToMorse n = morse !! (fromEnum n - 97)
                                    toCode s = concat . map charToMorse $ s

