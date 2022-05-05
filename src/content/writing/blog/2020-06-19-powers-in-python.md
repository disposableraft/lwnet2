---
#  Powers in Python
### 2020-06-19 {#date}
### python, self-help {#tags}
---

In Python the power operator can yield a surprising result if you expect it, as did I, to execute from left to right. The following expression evaluates to 🙀 `False`.

```python

>>> -2 ** 2 == 4
False

```

From the docs:

> The power operator binds more tightly than unary operators on its 
> left; it binds less tightly than unary operators on its right. The 
> syntax is:
>
>   power ::= (await_expr | primary) ["**" u_expr]
>
> Thus, in an unparenthesized sequence of power and unary operators, the
> operators are evaluated from right to left (this does not constrain
> the evaluation order for the operands): "-1**2" results in "-1".

For the expected result, parenthesize the left value (or use the `pow()` function).

```python

>>> (-2)**2 == 4
True

>>> math.pow(-2, 2) == 4
True

```



