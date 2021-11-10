module Main where

import Test.HUnit
import Tests.Q401

main :: IO Counts
main = runTestTT $ tests
