import qualified Data.ByteString.Lazy.Char8 as P
import Data.ByteString.Builder
import qualified Data.ByteString.Builder.Prim as Prim
import System.IO
import Control.Monad.Trans.State
import Data.Maybe

--import Data.ByteString.Builder.Extra (flush) -- for interactive problems
import Control.Monad
--import Data.List
--import Data.Array.Unboxed

--      scale = min p1 p2
--      s1 = p1 - scale
--      s2 = p2 - scale
--      y1 = x1 `quot` 10 ^ s2
--      y2 = x2 `quot` 10 ^ s1
--      z1 = toInteger x1 * 10 ^ s1
--      z2 = toInteger x2 * 10 ^ s2
--      ans = compare y1 y2 <> compare z1 z2

mainFun :: SP Builder
mainFun = do
  ~[t] <- getInts 1
  fmap mconcat $ replicateM t $ do
    ~[x1, p1] <- getInts 2
    ~[x2, p2] <- getInts 2
    let
      scale = min p1 p2
      s1 = min 7 $ p1 - scale
      s2 = min 7 $ p2 - scale
      y1 = x1 `quot` 10 ^ s2
      y2 = x2 `quot` 10 ^ s1
      z1 = toInteger x1 * 10 ^ s1
      z2 = toInteger x2 * 10 ^ s2
      ans = compare y1 y2 <> compare z1 z2
    pure $ string7 $ case ans of
      LT -> "<\n"
      EQ -> "=\n"
      GT -> ">\n"


type SP = State [P.ByteString]

getNext :: Int -> SP [P.ByteString]
getNext = state . splitAt

getInts :: Int -> SP [Int]
getInts k = map (fst . fromJust . P.readInt) <$> getNext k

putInts :: [Int] -> Builder
putInts vs = let
  sepPrim = (,) ' ' Prim.>$<
    Prim.liftFixedToBounded Prim.char7 Prim.>*< Prim.intDec
  in  case  vs  of
  [] -> char7 '\n'
  x : xs -> intDec x <> Prim.primMapListBounded sepPrim xs <> char7 '\n'

main :: IO ()
main = do
  hSetBuffering stdout NoBuffering
  inp <- P.getContents
  let  outp = evalState mainFun $ P.words inp
  P.putStr $ toLazyByteString outp
