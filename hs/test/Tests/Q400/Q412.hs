module Tests.Q400.Q412
    (tests
    )where

import Test.HUnit
import Q400.Q412

test1 :: Test
test1 =
    TestCase $ assertEqual "Should return [\"1\",\"2\",\"Fizz\",\"4\",\"Buzz\",\"Fizz\",\"7\",\"8\",\"Fizz\",\"Buzz\",\"11\",\"Fizz\",\"13\",\"14\",\"FizzBuzz\"]" (["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"]) (fizzBuzz 15)
test2 :: Test
test2 =
    TestCase $ assertEqual "Should return [\"1\",\"2\",\"Fizz\",\"4\",\"Buzz\",\"Fizz\",\"7\",\"8\",\"Fizz\",\"Buzz\",\"11\",\"Fizz\",\"13\",\"14\",\"FizzBuzz\"]" (["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"]) (fizzBuzzPlanB 15)


tests = TestList [test1, test2]
