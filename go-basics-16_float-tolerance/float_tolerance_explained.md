# 浮点数容差详解

浮点数（float32、float64）的运算和存储有误差。不能直接比较是否相等。

## 容差法

### 绝对容差比较

```go
func absIsClose(a, b, tol float64) bool {
    diff := math.Abs(a - b)
    return diff <= tol
}
```

> [!CAUTION]
>
> 问题：当数值较大时，绝对容差比较会失效。

### 相对容差比较

```go
func relIsClose(a, b, tol float64) bool {
    diff := math.Abs(a - b)
    return diff <=
        tol*max(math.Abs(a), math.Abs(b))
}
```

> [!CAUTION]
>
> 问题：当数值较小（趋近于0）时，相对容差比较会失效。

### 混合容差比较

```go
func isClose(a, b, absTol float64, relTol float64) bool {
    diff := math.Abs(a - b)
    return diff <= absTol ||
        diff <= relTol*max(math.Abs(a), math.Abs(b))
}
```

> [!NOTE]
>
> 扬长避短，适用于各种数值范围。

## 思考：相对容差计算的科学性

`tol*max(math.Abs(a), math.Abs(b))` 科学吗？

### 浮点数（float64）的二进制表示

例如：−26.375

```go
b := math.Float64bits(-26.375)
fmt.Printf("%0x\n%064b\n", b, b)
// 输出：
// c03a600000000000
// 1_10000000011_10100110000……（省略）
```

| S（符号位）1位 | E（指数）11位 | M（尾数）52位 |
| --- | --- | --- |
| 1 | 10000000011 | 1010011000…… |
| 负数 | 1027 | 0xa600000000000 |

$Value = (-1)^S \times (1 + M \times 2^{-52}) \times 2^{E-1023}$

$= (-1)^1 \times (1 + 0xa600000000000 \times 2^{-52}) \times 2^{4}$

$= -1 \times 1.6484375 \times 16$

$= -26.375$

用代码验证：

```go
f := math.Float64frombits(0xc03a600000000000)
fmt.Println(f)
// 输出：-26.375
```

### 尾数最低位的含义

$Value = (-1)^S \times (1 + M \times 2^{-52}) \times 2^e$

$|Value| = 2^e + M \times 2^{e-52}$

M 的最低位（第52位）表示 $2^{e-52}$。

Value 和下一个可表示的数直接的距离（ULP，Unit in the Last Place）为 $2^{e-52}$，即

$ULP = 2^{e-52}$

代码验证：

```go
 a := -26.375
 b := math.Nextafter(a, 0)
 diff := b - a
 fmt.Printf("a-b: %g\n", diff)

 ulp := math.Pow(2, 4-52)
 fmt.Printf("ULP: %g\n", ulp)

// 输出：
// a-b: 3.552713678800501e-15
// ULP: 3.552713678800501e-15
```

ULP 取决于指数e，数值越大，ULP越大；数值越小，ULP越小。

* 数轴：越靠近0，刻度越密集；越远离0，刻度越稀疏。

* 精度：越靠近0，精度越小，数值越精确；越远离0，精度越大，数值越不精确。

* 最小ULP： $\epsilon = 2^{-52}$

### $\epsilon \times |Value|$ 的含义

$\epsilon \times |Value|$

$= 2^{-52} \times |Value|$

$= 2^{-52} \times (1.f \times 2^e)$

$= 2^{-52} \times (1.f \times 2^e) \times (2^{-52} \times 2^{52})$

$= 2^{-52} \times 1.f \times (2^e \times 2^{-52}) \times 2^{52}$

$= 1.f \times 2^{e-52}$

$= 1.f \times ULP$

由于 $1.f \in [1, 2)$，因此

$\epsilon \times |Value| \in [ULP, \quad 2 \times ULP)$

如果 $tol = N \times \epsilon$，那么

$tol \times |Value| \in [N \times ULP, \quad 2N \times ULP)$

* 容差值 $tol × \max (|a|, |b|)$ 将位于 *最大ULP的N倍和2N倍之间*。

* 容差值 $tol × \max (|a|, |b|)$ 可以根据数值的大小*自动调整*，适用于各种数值范围。

```go
func isClose(a, b, absTol float64, relTol float64) bool {
    diff := math.Abs(a - b)
    return diff <= absTol ||
        diff <= relTol*max(math.Abs(a), math.Abs(b))
}
```

### 相对容差的缺陷

#### 当b趋近于0时

$|a-0| < tol × \max (|a|, |0|) \rightarrow$

$|a| < tol × |a| \rightarrow$

$tol > 1$ 是否成立。

这显然是不合适的。此时相对容差比较法就失效了。

#### 次正规数

$Value = (-1)^S × M×2^{-52} × 2^{-1022}$

* 指数e固定为-1022
* ULP为固定值 $2^{-1074}$，非常小

只适合绝对容差比较法。
