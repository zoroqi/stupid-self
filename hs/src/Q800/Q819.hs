module Q800.Q819
(
    mostCommonWord
) where

import qualified Data.Char as Char
import qualified Data.List as List
import qualified Data.HashSet as HashSet

mostCommonWord :: String -> [String] -> String
mostCommonWord paragraph banned = let
                                    ws = words . map Char.toLower . filter (\x -> not (x `elem` ",.;\'!?")) $ paragraph
                                    bans = HashSet.fromList banned
                                    counts = filter (\(x,_)-> not (HashSet.member x bans)) . map (\x -> (x!!0, length x)) . List.group . List.sort $ ws
                                  in
                                    comp counts 0 ""
                                  where
                                    comp :: [(String,Int)] -> Int -> String -> String
                                    comp [] _ ms = ms
                                    comp ((x,xn):xs) mc ms
                                                | xn > mc = comp xs xn x
                                                | otherwise = comp xs mc ms
