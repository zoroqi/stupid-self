module Tests.Q400.Q404
    (tests
    )
    where

import Test.HUnit
import Q400.Q404
import Struct

test1 :: Test
test1 =
    TestCase $ assertEqual "Should return 24" (24) (sumOfLeftLeaves (Node 3 (Node 9 Empty Empty) (Node 20 (Node 15 Empty Empty) (Node 7 Empty Empty))))

tests = TestList [test1]
