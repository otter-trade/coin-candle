# 持仓管理

持仓的计算比较复杂，平仓，计算收益等有大量的公共方法和函数。
最好以面向对象的方式去进行管理：

写一下伪代码

```go
// 建立一个 action
action,err := NewPositionAction({
  StrategyID string
  MockName string
})
// action 包含交易需要 的各种信息，包括手续费率，余额等。

action.SetNewPositionList
// SetNewPositionList 仓位相关的信息，需要判断参数，需要根据币种状态进行过滤，需要计算单项手续费。

action.UpdatePosition
// 将 NewPositionList 更新到 PositionAction 里面去。

action.ReadPosition
// 计算和读取 Position 的相关状态

action.ClosePosition
// 平仓计算，落地收益

action.SavePosition
// 将 Position 保存到本地

action.ReadPositionHistory
// 读取 持仓历史
```

## 收益和收益率计算

```txt

计算收益率的公式：

收益率 = (平仓时的价格 - 下单时的价格) / 下单时候的价格

平仓后收益 = 余额 * 收益率 - 手续费


简化概念：

在同一条K线上 若 A时刻下单，B时刻平仓； A点 到B点 时，B点 的收益率为

(B点收盘价 - A点收盘价)/A点收盘价 * 100%


例子：

昨天股价 4 元，今天股价 5 元 ，若昨天买了股票，今天的收益率为 ：
(5-4)/4 * 100% ≈ 25%
若昨天买了 1000 元的股票，今天的收益额为  1000 * 0.33 =  250 元
则总余额 =  1000 * (1+0.25) =  1250 元

换个计算逻辑来验证：
1000 元昨天可以购买  1000 / 4  = 250 股
今天每股 5 元，全部卖出 可得    250 * 5 = 1250 元




```

## 手续费收取规则

参照 https://www.okx.com/zh-hans/fees

```txt

手续费计算公式：
现货交易手续费 = 手续费率 × 成交时买入币种的数量。

以 USDT/USD 现货为例，假设 USDT 目前价格为 10 USD；交易者 A (挂单手续费为 0.04%，吃单手续费为 0.1%) 以市价单买入 1 USDT，成交时为吃单方，需要支付的手续费 = 0.1% × 1 = 0.001 USDT，成交后将获得 0.999 USDT；

若交易者 A 以限价单卖出 1 USDT，即买入 10 USD，则成交时为挂单方，需要支付的手续费 = 0.04% × 10 = 0.004 USD，成交后将获得 9.996 USD。

```

> 买入卖出 均要支付手续费

## 手续费计算公式

```txt

若 手续费率为  0.1%

当前持有  1000 USTD 的 BTC 想要卖掉，则代表买入  1000 USDT
手续费 = 1000USDT *  0.1%   =  1USDT
成交后得到  1000USDT - 1USDT =   999 USDT


当你持有 1000USDT 想要购买 BTC,  相当于你要买入 x 个BTC
手续费 =  x BTC *  0.1%   =  a BTC
成交后得到   x BTC - a BTC =   x - a BTC

已知 xBTC = 1000USDT  则上述公式可等效为
手续费 = 1000USDT *  0.1%   =  1USDT
成交后得到  1000USDT - 1USDT =   999 USDT


简化概念：

USDT 为一般等价物，相当于价格单位
BTC 为商品，但是价格是随着时间变动的
所以在相同的时间内 （假设交易在一瞬间完成）
若 xBTC = 1000USDT

那么 x BTC  和 a BTC 的价值差 其实就是   1000USDT *  0.1%

因此无论是买入还是卖出，手续费的计算公式均可以简化为：

手续费 = 买入/卖出总金额 * 手续费率

```
