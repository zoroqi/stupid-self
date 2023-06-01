module Q800.Q856
(
    scoreOfParenthesesPlanB
) where


data Pairs = One
            | Double Pairs
            | Plus Pairs Pairs
            deriving(Show)

scoreOfParenthesesPlanB :: String -> Int
scoreOfParenthesesPlanB s = planb (read s)
    where
        planb One = 1
        planb (Double a) = 2 * (planb a)
        planb (Plus a b) = (planb a) + (planb b)

-- 翻译成数学公式打印
-- mathShow :: String -> String
-- mathShow s = pretty (read s)
--     where
--         pretty One = "1"
--         pretty (Double a) = "2 * (" ++ pretty a ++ ")"
--         pretty (Plus a b) = pretty a ++ " + " ++ pretty b

-- 可以使用 read 函数进行解析
instance Read Pairs where
    readsPrec _ ss = [(parse ss,"")]
        where
            parse "" = One
            parse "()" = One
            parse x =
                let
                    b = split x
                in
                    if snd b == ""
                    then Double (parse (remove(fst b)))
                    else Plus (parse (fst b)) (parse (snd b))
                where
                    remove s = init (tail s)

-- 分割成两个合法括号字符串,
split :: String -> (String, String)
split ss = spl (tail ss) "(" "("
    where
        spl [] (_:_) o = error ("Incomplete parentheses " ++ o)
        spl s [] o = (o, s)
        spl ('(':sx) sk o = spl sx ('(':sk) (o++"(")
        spl (')':sx) ('(':sk) o = spl sx sk (o++")")
