package gate

// 【已弃用】

// // Gate 是控制访问安全的门对象
// type Gate interface {
// 	Handler() gin.HandlerFunc
// }

// // Builder 用于创建 Gate 对象
// type Builder interface {

// 	// 设置目标 Handler
// 	ForHandler(fn gin.HandlerFunc) Builder

// 	// 复位 Builder
// 	Reset() Builder

// 	// 创建 Gate 对象
// 	Gate() Gate

// 	// 创建 Gate 对象, 并返回代理 Handler
// 	Handler() gin.HandlerFunc

// 	// 设置目标 Handler，创建 Gate 对象, 并返回代理 Handler
// 	H(fn gin.HandlerFunc) gin.HandlerFunc
// }

// // // NewBuilder 新建一个 gate.Builder
// // func NewBuilder() Builder {
// // 	return nil
// // }

// ////////////////////////////////////////////////////////////////////////////////

// type myGate struct {
// 	target gin.HandlerFunc
// }

// func (inst *myGate) _Impl() Gate {
// 	return inst
// }

// func (inst *myGate) Handler() gin.HandlerFunc {
// 	return inst.handle
// }

// func (inst *myGate) handle(c *gin.Context) {
// 	// todo...

// 	t := inst.target
// 	t(c)
// }

// ////////////////////////////////////////////////////////////////////////////////

// type myGateBuilder struct {
// 	target gin.HandlerFunc
// }

// func (inst *myGateBuilder) _Impl() Builder {
// 	return inst
// }

// func (inst *myGateBuilder) Reset() Builder {
// 	inst.target = nil
// 	return inst
// }

// func (inst *myGateBuilder) Gate() Gate {

// 	// todo ...
// 	g := &myGate{}
// 	inst.Reset()
// 	return g
// }

// func (inst *myGateBuilder) Handler() gin.HandlerFunc {
// 	return inst.Gate().Handler()
// }

// func (inst *myGateBuilder) H(target gin.HandlerFunc) gin.HandlerFunc {
// 	inst.setTarget(target)
// 	g := inst.Gate()
// 	return g.Handler()
// }

// func (inst *myGateBuilder) setTarget(target gin.HandlerFunc) Builder {
// 	inst.target = target
// 	return inst
// }

// func (inst *myGateBuilder) ForHandler(target gin.HandlerFunc) Builder {
// 	inst.Reset()
// 	return inst.setTarget(target)
// }

// ////////////////////////////////////////////////////////////////////////////////
