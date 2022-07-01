module Tests.Q800.Q824
(
    tests
) where

import Test.HUnit
import Struct
import Q800.Q824

test1 :: Test
test1 =
    TestCase $ assertEqual "Should return 'Imaa peaksmaaa oatGmaaaa atinLmaaaaa'" "Imaa peaksmaaa oatGmaaaa atinLmaaaaa" (toGoatLatinPlanA "I speak Goat Latin")

tests = TestList [test1]
