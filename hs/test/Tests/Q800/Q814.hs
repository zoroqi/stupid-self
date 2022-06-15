module Tests.Q800.Q814
(
    tests
) where

import Test.HUnit
import Struct
import Q800.Q814

test1 :: Test
test1 =
    TestCase $ assertEqual "Should return True " (Node 1 Empty (Node 1 Empty (leaf 1)))  (pruneTreePlanA (Node 1 (Node 0 (leaf 0) (leaf 0)) (Node 1 (leaf 0) (leaf 1))))

tests = TestList [test1]
