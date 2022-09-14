
/**
 * 可组合的散列值
 * Q: 权限分成 页面级别权限和组件级别权限
 * Q: 如何实现组件级别权限？
 *      - 表达权限
 *      - 判定权限
 * 
 * 其他参考:
 * 
 * React的fiber中的fiberflex有关于对可组合的散列值
 */

// const CREATE = 1

// const DELETE = 2

// const UPDATE = 3

// const DETAIL = 4

const CREATE = 0b0001

const DELETE = 0b0010

const UPDATE = 0b0100

const DETAIL = 0b1000

// 通过或运算来确定拥有的权限值
const result = CREATE | DELETE
console.log(result.toString(2))
// 通过且运算来判断当前组件是否拥有权限
console.log( (result & DELETE) === DELETE )