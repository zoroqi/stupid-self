module Main where

import Q400.Q401

main :: IO ()
main = putStrLn (show (readBinaryWatch (\x -> length x == 1) allTime))
