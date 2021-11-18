module Tests.Q400.Q406
    (tests
    )
    where

import Test.HUnit
import Q400.Q406
import Struct

test1 :: Test
test1 =
    TestCase $ assertEqual "Should return [[(5,0),(7,0),(5,2),(6,1),(4,4),(7,1)]]" ([[(5,0),(7,0),(5,2),(6,1),(4,4),(7,1)]]) (queuePlanA [(7,0),(4,4),(7,1),(5,0),(6,1),(5,2)])

tests = TestList [test1]
