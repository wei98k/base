
// 找出一个基准值
// 缩小问题规模，知道符合基准值条件

function quickSort(arr, left, right) {
    var len = arr.length,
        partitionIndex,
        left = typeof letf != 'number' ? 0 : left,
        right = typeof right != 'number' ? len - 1 : right

    if (left < right) {
        let pivot = left,
            index = pivot + 1;
        for(let i = index; i<= right; i++) {
            if(arr[i] < arr[pivot]) {
                let temp = arr[i]
                arr[i] = arr[index];
                arr[index] = temp;
                index++;
            }
        }
        let temp2 = arr[pivot]
        arr[pivot] = arr[index-1]
        arr[index-1] = temp2
        partitionIndex = index - 1

        quickSort(arr, left, partitionIndex-1)
        console.log(partitionIndex+1, right)
        quickSort(arr, partitionIndex+1, right)
    }

    return arr;
}

// function quickSort(arr, left, right) {
//     var len = arr.length,
//         partitionIndex,
//         left = typeof left != 'number' ? 0 : left,
//         right = typeof right != 'number' ? len - 1 : right;

//     if (left < right) {
//         partitionIndex = partition(arr, left, right);
//         quickSort(arr, left, partitionIndex-1);
//         // console.log(partitionIndex+1, right)
//         quickSort(arr, partitionIndex+1, right);
//     }
//     return arr;
// }

// function partition(arr, left ,right) {     // 分区操作
//     var pivot = left,                      // 设定基准值（pivot）
//         index = pivot + 1;
//     for (var i = index; i <= right; i++) {
//         if (arr[i] < arr[pivot]) {
//             swap(arr, i, index);
//             index++;
//         }        
//     }
//     swap(arr, pivot, index - 1);
//     return index-1;
// }

// function swap(arr, i, j) {
//     var temp = arr[i];
//     arr[i] = arr[j];
//     arr[j] = temp;
// }




const testArr = [9,4,5,8,7]
// let i = 0
// while (i < 10) {
//     testArr.push(Math.floor(Math.random() * 1000))
//     i++
// }
console.log('unsort', testArr)
quickSort(testArr, 0, testArr.length - 1);
console.log('sort', testArr)