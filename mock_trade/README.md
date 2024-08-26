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
action 包含交易需要 的各种信息，包括手续费率，余额等。

action.SetNewPositionList
// SetNewPositionList 仓位相关的信息，需要判断参数，需要根据币种状态进行过滤，需要计算单项手续费。
//

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
