
### 冒泡算法
```
const bubbleSort = (arr) => {
    if (arr.length <= 1) return
    for (let i = 0; i < arr.length; i++) {
        let hasChange = false

        for (let j = 0; j < arr.length - i - 1; j++) {

            if (arr[j] > arr[j + 1]) {

                const temp = arr[j]

                arr[j] = arr[j + 1]

                arr[j + 1] = temp

                hasChange = true
            }
        }
        // 如果false 说明所有元素已经到位
        if (!hasChange) break
    }
    console.log(arr)
}
```
### 插入排序

```
const insertionSort = (arr) => {
    if (arr.length <= 1) return

    for (let i = 1; i < arr.length; i++) {
        const temp = arr[i]

        let j = i - 1

        // 若arr[i]前有大于arr[i]的值的化，向后移位，腾出空间，直到一个<=arr[i]的值
        for (j; j >= 0; j--) {

            if (arr[j] > temp) {

                arr[j + 1] = arr[j]

            } else {
                break
            }
        }
        arr[j + 1] = temp
    }
    console.log(arr)
}
```
### 选择排序

```
const selectionSort = (arr) => {

    if (arr.length <= 1) return

    // 需要注意这里的边界, 因为需要在内层进行 i+1后的循环，所以外层需要 数组长度-1

    for (let i = 0; i < arr.length - 1; i++) {

        let minIndex = i

        for (let j = i + 1; j < arr.length; j++) {

            if (arr[j] < arr[minIndex]) {

                minIndex = j // 找到整个数组的最小值
            }
        }

        const temp = arr[i]

        arr[i] = arr[minIndex]

        arr[minIndex] = temp
    }
    
    console.log(arr)
}
```
### 快速排序

```
const swap = (arr, i, j) => {
    const temp = arr[i]
    arr[i] = arr[j]
    arr[j] = temp
}

// 获取 pivot 交换完后的index
const partition = (arr, pivot, left, right) => {
    const pivotVal = arr[pivot]
    let startIndex = left
    for (let i = left; i < right; i++) {
        if (arr[i] < pivotVal) {
            swap(arr, i, startIndex)
            startIndex++
        }
    }
    swap(arr, startIndex, pivot)
    return startIndex
}

const quickSort = (arr, left, right) => {
    if (left < right) {
        let pivot = right
        let partitionIndex = partition(arr, pivot, left, right)
        quickSort(arr, left, partitionIndex - 1 < left ? left : partitionIndex - 1)
        quickSort(arr, partitionIndex + 1 > right ? right : partitionIndex + 1, right)
    }

}


const testArr = []
let i = 0
while (i < 10) {
    testArr.push(Math.floor(Math.random() * 1000))
    i++
}
console.log('unsort', testArr)
quickSort(testArr, 0, testArr.length - 1);
console.log('sort', testArr)
```
### 归并排序

```
const mergeArr = (left, right) => {
    let temp = []
    let leftIndex = 0
    let rightIndex = 0
    // 判断2个数组中元素大小，依次插入数组
    while (left.length > leftIndex && right.length > rightIndex) {
        if (left[leftIndex] <= right[rightIndex]) {
            temp.push(left[leftIndex])
            leftIndex++
        } else {
            temp.push(right[rightIndex])
            rightIndex++
        }
    }
    // 合并 多余数组
    return temp.concat(left.slice(leftIndex)).concat(right.slice(rightIndex))
}

const mergeSort = (arr) => {
    // 当任意数组分解到只有一个时返回。
    if (arr.length <= 1) return arr
    const middle = Math.floor(arr.length / 2) // 找到中间值
    const left = arr.slice(0, middle) // 分割数组
    const right = arr.slice(middle)
    // 递归 分解 合并
    return mergeArr(mergeSort(left), mergeSort(right))
}

const testArr = []
let i = 0
while (i < 100) {
    testArr.push(Math.floor(Math.random() * 1000))
    i++
}

const res = mergeSort(testArr)
console.log(res)
```
### 第k大的数

```
function kthNum(arr, k) {
  const len = arr.length;
  if (k > len) {
    return -1;
  }
  let p = partition(arr, 0, len - 1);
  while (p + 1 !== k) {
    if (p + 1 > k) {
      p = partition(arr, 0, p - 1);
    } else {
      p = partition(arr, p + 1, len - 1);
    }
  }
  return arr[p];
}

function partition(arr, start, end) {
  let i = start;
  let pivot = arr[end];
  for (let j = start; j < end; j++) {
    if (arr[j] < pivot) {
      swap(arr, i, j);
      i += 1;
    }
  }
  swap(arr, i, end);
  return i;
}

function swap(arr, i, j) {
  if (i === j) return;
  let tmp = arr[i];
  arr[i] = arr[j];
  arr[j] = tmp;
}
```
### 每个桶用插入排序算法

```
function bucketSort(array, bucketSize = 5) {
    if (array.length < 2) {
        return array
    }
    const buckets = createBuckets(array, bucketSize)
    return sortBuckets(buckets)
}

function createBuckets(array, bucketSize) {
    let minValue = array[0]
    let maxValue = array[0]
    // 遍历数组，找到数组最小值与数组最大值
    for (let i = 1; i < array.length; i++) {
        if (array[i] < minValue) {
            minValue = array[i]
        } else if (array[i] > maxValue) {
            maxValue = array[i]
        }
    }
    // 根据最小值、最大值、桶的大小，计算得到桶的个数
    const bucketCount = Math.floor((maxValue - minValue) / bucketSize) + 1
    // 建立一个二维数组，将桶放入buckets中
    const buckets = []
    for (let i = 0; i < bucketCount; i++) {
        buckets[i] = []
    }
    // 计算每一个值应该放在哪一个桶中
    for (let i = 0; i < array.length; i++) {
        const bucketIndex = Math.floor((array[i] - minValue) / bucketSize)
        buckets[bucketIndex].push(array[i])
    }
    return buckets
}

function sortBuckets(buckets) {
    const sortedArray = []
    for (let i = 0; i < buckets.length; i++) {
        if (buckets[i] != null) {
            insertionSort(buckets[i])
            sortedArray.push(...buckets[i])
        }
    }
    return sortedArray
}

// 插入排序
function insertionSort(array) {
    const { length } = array
    if (length <= 1) return

    for (let i = 1; i < length; i++) {
        let value = array[i]
        let j = i - 1

        while (j >= 0) {
            if (array[j] > value) {
                array[j + 1] = array[j] // 移动
                j--
            } else {
                break
            }
        }
        array[j + 1] = value // 插入数据
    }
}
```
### countingSort

```
const countingSort = array => {
    if (array.length <= 1) return

    const max = findMaxValue(array)
    const counts = new Array(max + 1)

    // 计算每个元素的个数，放入到counts桶中
    // counts下标是元素，值是元素个数
    array.forEach(element => {
        if (!counts[element]) {
            counts[element] = 0
        }
        counts[element]++
    })

    // counts下标是元素，值是元素个数
    // 例如： array: [6, 4, 3, 1], counts: [empty, 1, empty, 1, 1, empty, 1]
    // i是元素, count是元素个数
    let sortedIndex = 0
    counts.forEach((count, i) => {
        while (count > 0) {
            array[sortedIndex] = i
            sortedIndex++
            count--
        }
    })
    // return array
}

function findMaxValue(array) {
    let max = array[0]
    for (let i = 1; i < array.length; i++) {
        if (array[i] > max) {
            max = array[i]
        }
    }
    return max
}
```