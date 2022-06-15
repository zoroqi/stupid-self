module Main where

import Test.HUnit
import Tests.Q800.Q814

main :: IO Counts
main = runTestTT $ tests
