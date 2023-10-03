module Main where

import Game.TicTacToe2

main = do
    g <- getStdGen
    runGame  (playerAI g) playerHuman
