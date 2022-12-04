{-# LANGUAGE FunctionalDependencies, MultiParamTypeClasses #-}
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
