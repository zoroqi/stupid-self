module Main where

import Test.HUnit
import Tests.Q800.Q806

main :: IO Counts
main = runTestTT $ tests
