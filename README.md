# Лабораторная работа по Теории Автоматов
Прототип LL(1) анализатора

## Грамматики

Исходная грамматика
```
<PROG>   -> PROG id <VAR> begin < LISTST> end
<VAR>    -> VAR <IDLIST> : <TYPE>
<IDLIST> -> id | <IDLIST>, id
<LISTST> -> <ST> | <LISTST> <ST>
<TYPE>   -> int|float|bool|string
<ST>     -> <READ>|<WRITE>| <ASSIGN>
<ASSIGN> -> id := <EXP>;
<EXP>    -> <EXP> + <T> | <T>
<T>      -> <T>*<F> | <F>
<F>      -> -<F> | (<EXP>) | id | num
<READ>   -> READ (<IDLIST>);
<WRITE>  -> WRITE (<IDLIST>);
```

Приведенная грамматика
```
<PROG>    -> PROG id <VAR> begin <LISTST> end
<VAR>     -> VAR <IDLIST> : <TYPE>
<IDLIST>  -> id <IDLISTA>
<IDLISTA> -> ε | , id <IDLISTA>
<LISTST>  -> <ST> <LISTSTA>
<LISTSTA> -> <ST> <LISTSTA> | ε
<TYPE>    -> int|float|bool|string
<ST>      -> <READ> | <WRITE> | <ASSIGN>
<ASSIGN>  -> id := <EXP>;
<EXP>     -> <T><EXPA>
<EXPA>    -> ε | + <T><EXPA>
<T>       -> <F> | <TA>
<TA>      -> ε | * <F><TA>
<F>       -> -<F> | (<EXP>) | id | num
<READ>    -> READ(<IDLIST>);
<WRITE>   -> WRITE(<IDLIST>);
```

