module Main where

import Test.HUnit
import Tests.Q400.Q406

main :: IO Counts
main = runTestTT $ tests
