module Q800.Q841
(
    canVisitAllRoomsPlanB
) where

import qualified Data.HashSet as HashSet

canVisitAllRoomsPlanB :: [[Int]] -> Bool
canVisitAllRoomsPlanB rooms = HashSet.size s >= length rooms
                    where
                        s =  planb rooms (HashSet.fromList [0]) 0

planb :: [[Int]] -> HashSet.HashSet Int  -> Int -> HashSet.HashSet Int
planb rooms set n =
    let
        nset = HashSet.union set (HashSet.fromList (rooms!!n))
        df = planb rooms nset
        exist = filter (\x -> not (HashSet.member x set))
    in  HashSet.unions (nset : (map df . exist $ (rooms!!n)))
