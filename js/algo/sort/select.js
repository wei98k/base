
let arr = [3,2,9,2,4,8,7,5,6]
console.log("排序前: ", arr)
let res = selectSort(arr)
console.log("选择排序后: ", res)
let res2 = bubble(arr)
console.log("冒泡排序后: ", res2)
function selectSort(arr) {
    // 必须有数组 并且是有元素的
    for(let i=0; i<arr.length; i++) {
        let minKey = i
        for(let j = i+1; j<arr.length; j++) {
            // console.log(j)
            if(arr[j] < arr[minKey]) {
                minKey = j
            }
        }
        let temp = arr[i]
        arr[i] = arr[minKey]
        arr[minKey] = temp
    }
    return arr
}

//冒泡排序思路:
/**
 * 1. 循环比较相邻两个元素 元素小的在左边，大的在右边
 */

function bubble(arr) {
    if(arr.length <= 0) {
        return false
    }
    for(let i = 0; i<arr.length; i++) {
        for(let j = 0; j<arr.length; j++) {
            if(arr[j] > arr[j+1]) {
                let temp = arr[j+1]
                arr[j+1] = arr[j]
                arr[j] = temp
            }
        }
    }
    return arr
}
// 插入排序思路
function insertSort(arr) {
    
}