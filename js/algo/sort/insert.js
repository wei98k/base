
let arr = [3,2,9,2,4,8,7,5,6]

let res = inserSort(arr)
//let res = insertionSort(arr)
console.log(res)

// function inserSort(arr) {
//     // 循环全部需要对比的元素
//     // N+1比较前面X的全部元素
//     // N+1位置需要比N位置的元素大就在X+1位置插入
//     // 循环上面操作找到结束
//     // Tip: 0等于temp  循环1开始
//     for(let i = 1; i<arr.length; i++) {
//         let j = i - 1;
//         let temp = arr[i]
//         for(j;j >=0; j--) {
//             // 比较之前已经比较过的元素
//             // 依次比较，比元素小，之前的元素就后移位置
//             if(temp < arr[j]) {
//                 arr[j+1] = arr[j]
//             } else {
//                 break;
//             }
//         }
//         // 当之前的元素没有比当前元素小的了，就在这个位置插入当前的元素
//         arr[j+1] = temp
//     }
//     return arr
// }

function inserSort(arr) {
    for(let i = 1; i < arr.length; i++) {
        let temp = arr[i];
        let j = i - 1
        while(j >= 0 && temp < arr[j]) {
            arr[j+1] = arr[j]
            --j
        }
        // console.log(temp)
        arr[j+1] = temp
    }
    return arr
}