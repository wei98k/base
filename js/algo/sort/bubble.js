/**
 * 冒泡排序 按从小到大排序
 * 
 * 1. 两层循环
 * 2. 相邻两个值比较，如果左边大于右边，就先把左边的值赋予临时变量。
 * 3. 把右边的值，赋予左边的key位置
 * 4. 临时变量赋值给右边的key位置 
 * 5. 循环直到所有左边的值小于右边
 */

function bubble(arr) {
    if(arr.length <= 0) { return}
    let flag = true
    for(let i=0; i<arr.length; i++) {
        
        for(let j=0; j<arr.length - i -1; j++) {
            if(arr[j] > arr[j+1]) {
                let tmpVar = arr[j+1]
                arr[j+1] = arr[j]
                arr[j] = tmpVar
                flag = false
            }
        }
        if(flag) {
            break
        }
    }
    return arr
}

let arr1 = [1,3,4,9,2,3,5] 
let res = bubble(arr1)
console.log("bubble", res)