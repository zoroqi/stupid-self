module Main where

import Test.HUnit
import Tests.Q800.Q824

main :: IO Counts
main = runTestTT $ tests
