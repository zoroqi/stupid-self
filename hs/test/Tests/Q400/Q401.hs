module Tests.Q400.Q401
    (tests
    )
    where

import Test.HUnit
import Q400.Q401

test1 :: Test
test1 =
    TestCase $ assertEqual "Should return [00:32, 00:16, 00:08, 00:04, 00:02, 00:01, 08:00, 04:00, 02:00, 01:00]" ([(0.0,32.0),(0.0,16.0),(0.0,8.0),(0.0,4.0),(0.0,2.0),(0.0,1.0),(8.0,0.0),(4.0,0.0),(2.0,0.0),(1.0,0.0)]) (readBinaryWatch (\x -> length x == 1) allTime)

test9 :: Test
test9 =
    TestCase $ assertEqual "Should return []" ([]) (readBinaryWatch (\x -> length x == 9) allTime)

tests = TestList [test1, test9]

