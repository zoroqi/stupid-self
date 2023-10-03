{-# LANGUAGE FunctionalDependencies, MultiParamTypeClasses #-}
{-# LANGUAGE GADTs, Rank2Types #-}
module Temp where

import System.IO (IOMode(..))

class Monad m => MonadHandle h m | m -> h where
        openFile :: FilePath -> IOMode -> m h
        hPutStr :: h -> String -> m ()
        hClose :: h -> m ()
        hGetContents :: h -> m String

        hPutStrLn :: h -> String -> m ()
        hPutStrLn h s = hPutStr h s >> hPutStr h "\n"


(+-) a b = a + b
(-+) a b = a / b
(-。-) a b = a + b

class MyC a where
    (++-) :: a -> a -> a

instance MyC Int where
    (++-) a b = a + b

infixl 5 +-
infix 6 -+

subsets :: [a] -> [[a]]
subsets [] = [[]]
subsets (x:xs) = subsets xs ++ map (x:) (subsets xs)

data Term a where
    Lit    :: Int -> Term Int
    Succ   :: Term Int -> Term Int
    IsZero :: Term Int -> Term Bool
    If     :: Term Bool -> Term a -> Term a -> Term a
    Pair   :: Term a -> Term b -> Term (a,b)

eval :: Term a -> a
eval (Lit i)      = i
eval (Succ t)     = 1 + eval t
eval (IsZero t)   = eval t == 0
eval (If b e1 e2) = if eval b then eval e1 else eval e2
eval (Pair e1 e2) = (eval e1, eval e2)

data Term2 a = TT {
    lit2 :: a,
    succ2 :: Term2 a -> Term2 a
}

ff :: Term2 Int -> Term2 Int
ff (TT n _) = TT {lit2 = n*2}
