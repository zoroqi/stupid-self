module Q400.Q481
(magicalStringPlanA,
    magicalStringPlanB,
    ) where

magicalStringPlanA :: Int -> Int
magicalStringPlanA a = foldl (+) 0 . filter (\x -> x == 1) . plana $ a

plana :: Int -> [Int]
plana 1 = [1]
plana 2 = [1,2]
plana n = next n [2,2]
          where
            next 0 x = x
            next m x = next (m-1) (level (tail x) 1 [1])

level :: [Int] -> Int -> [Int] -> [Int]
level [] _ x = x
level (x:xs) before nums = level xs (last newn) (nums ++ newn)
                 where
                    newn = expand before x

expand :: Int -> Int -> [Int]
expand 1 2 = [2, 2]
expand 2 2 = [1, 1]
expand 1 1 = [2]
expand 2 1 = [1]
expand _ _ = []

magicalStringPlanB :: Int -> Int
magicalStringPlanB a = foldl (+) 0 . filter (\x -> x == 1) $ (take a ([1,2] ++ (concat planb)))

planb :: [[Int]]
planb = [[2]]++(zipWith expand (cycle [2,1]) (concat planb))
