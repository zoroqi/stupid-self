module Tests.Q800.Q806
    (tests
    )
    where

import Test.HUnit
import Q800.Q806
import Struct

test1 :: Test
test1 =
    TestCase $ assertEqual "Should return (2,4) " (2,4) (numberOfLinesPlanA [4,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10] "bbbcccdddaaa")

test2 :: Test
test2 =
    TestCase $ assertEqual "Should return (3,60) " (3,60) (numberOfLinesPlanB [10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10]  "abcdefghijklmnopqrstuvwxyz")

test3 :: Test
test3 =
    TestCase $ assertEqual "Should return (2, 54) " (2, 54) (numberOfLinesPlanA [2,3,4,5,6,7,8,9,10,2,3,4,5,6,7,8,9,10,2,3,4,5,6,7,8,9] "abcdefghijklmnopqrstuvwxyz")

test4 :: Test
test4 =
    TestCase $ assertEqual "Should return True " True ((numberOfLinesPlanB [2,3,4,5,6,7,8,9,10,2,3,4,5,6,7,8,9,10,2,3,4,5,6,7,8,9] "abcdefghijklmnopqrstuvwxyz") == (numberOfLinesPlanA [2,3,4,5,6,7,8,9,10,2,3,4,5,6,7,8,9,10,2,3,4,5,6,7,8,9] "abcdefghijklmnopqrstuvwxyz"))

tests = TestList [test1,test2,test3, test4]
