module Main where

import Test.HUnit
import Tests.Q400.Q404

main :: IO Counts
main = runTestTT $ tests
