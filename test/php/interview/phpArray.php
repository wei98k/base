<?php

// 代码测试环境 env/php/5.6  1天5个函数 
// docker run -idt -v $PWD:/app --name php-array mybasephp:5.6-fpm

// array_change_key_case 将数组中的所有键名修改为全大写或小写, 默认将key转为消息, CASE_LOWER(小写) CASE_UPPER(大写)

/*

$a1 = array('a' => 'this is a', 'b' => 'this is b');

$a2 = array_change_key_case($a1, CASE_UPPER);

var_dump($a2);

*/

// array_chunk 将一个数组分割成多个 

/*

$a1 = array('a', 'b', 'c', 'd');

$b2 = array_chunk($a1, 3, true);

var_dump($b2);

*/

// array_column 返回数组中指定的一列 

/*

$records = array(
    array(
        'id' => 2135,
        'first_name' => 'John',
        'last_name' => 'Doe',
    ),
    array(
        'id' => 3245,
        'first_name' => 'Sally',
        'last_name' => 'Smith',
    ),
    array(
        'id' => 5342,
        'first_name' => 'Sally',
        'last_name' => 'Jones',
    ),
    array(
        'id' => 5623,
        'first_name' => 'Peter',
        'last_name' => 'Doe',
    )
);

$id = array_column($records, 'last_name', 'first_name');

var_dump($id);

*/

// array_combine 接收两个数组 数组1值为键值  数组2值为值

/* 

$a = array('green', 'red', 'yellow');
$b = array('avocado', 'apple', 'banana');
$c = array_combine($a, $b);
var_dump($c);

*/

// array_count_values 统计数组中所值的值

/*

$array = array(1, "hello", 1, "world", "hello");

$a = array_count_values($array);

var_dump($a);

*/

// ?? array_diff_assoc 带索引检查计算数组的差值##对比数组返回不相同的键元素

/*

$array1 = array("a" => "green", "b" => "brown", "c" => "blue", "red");
$array2 = array("a" => "green", "yellow", "red");

$result = array_diff_assoc($array1, $array2);
print_r($result);

*/

// ?? array_diff_key 使用键名比较计算差值
/*
$array1 = array('blue'  => 1, 'red'  => 2, 'green'  => 3, 'purple' => 4);
$array2 = array('green' => 5, 'blue' => 6, 'yellow' => 7, 'cyan'   => 8);

var_dump(array_diff_key($array1, $array2));
*/

// ?? array_diff_uassoc 用用户提供的回调函数做索引检查来计算数组的差值 
/*
function key_compare_func($a, $b)
{
    if ($a === $b) {
        return 0;
    }
    return ($a > $b)? 1:-1;
}

$array1 = array("a" => "green", "b" => "brown", "c" => "blue", "red");
$array2 = array("a" => "green", "yellow", "red");
$result = array_diff_uassoc($array1, $array2, "key_compare_func");
print_r($result);
*/

// ?? array_diff_ukey 用回调函数对键名比较计算数组的差集 
/*
function key_compare_func($key1, $key2)
{
	echo $key1 .'#' . $key2 . '#';
    if ($key1 == $key2) {
    	echo "0\n";
        return 0;
    }else if ($key1 > $key2){
    	echo "1\n";
        return 1;
    }else {
    	echo "-1\n";
        return -1;
    }

}

$array1 = array('blue'  => 1, 'red'  => 2, 'green'  => 3, 'purple' => 4);
$array2 = array('green' => 5, 'blue' => 6, 'yellow' => 7, 'cyan'   => 8);

var_dump(array_diff_ukey($array1, $array2, 'key_compare_func'));

$a = 'A';
$b = 'b';
// 字符是怎么比较大小的?  http://t.cn/RDcEqXc http://suo.im/4S2rlF
if($a == $b) {
	echo "0\n";
} else if ($a > $b) {
	echo "1\n";
} else {
	echo "-1\n";
}
*/

// $a = "aabbzz";
// $a++;
// echo $a."\n";

// if ('1e3' == '1000') echo 'LOL'."\n";



# 2018年08月08日09:37:00 start 

// array_diff 计算数组差集##比较元素中的值和键无关s
/*
$array1 = array("a" => "green", "red", "blue", "red", "a");
$array2 = array("b" => "green", "yellow", "red");
$array3 = array("b" => "green", "yellow", "red", "blue");
$result = array_diff($array1, $array2, $array3 );

print_r($result);
*/

// array_fill_keys 使用指定的键和值填充数组
/*
$keys = array('foo', 5, 10, 'bar');
$b = 'banana';
$c = array('a', 'b');

$a = array_fill_keys($keys, $c);
print_r($a);
*/

// array_fill 用给定的值填充数组
/*
$a1 = 'banana';
$a2 = array('a' => 'banana');
$b1 = 'pear';

$a = array_fill(5, 6, $a2);
$b = array_fill(-2, 4, $b1);
print_r($a);
print_r($b);
*/


// array_filter 用回调函数过滤数组中的单元 什么情况可以用到?  在接口中过滤null值 
/*
function odd($var)
{
    // returns whether the input integer is odd
    // echo $var."\n";
    return($var & 1);
}

function even($var)
{
    // returns whether the input integer is even
    return(!($var & 1));
}

$array1 = array("a"=>1, "b"=>2, "c"=>3, "d"=>4, "e"=>5);
$array2 = array(6, 7, 8, 9, 10, 11, 12);

// echo "Odd :\n";
// print_r(array_filter($array1, "odd"));
// echo "Even:\n";
// print_r(array_filter($array2, "even"));
// $var = 3;
// var_dump($var & 1);
// & 按位与 引用  
// 过滤 null '' 0 false 值
$entry = array(
             0 => 'foo',
             1 => false,
             2 => -1,
             3 => null,
             4 => '',
             5 => 0
          );
function myfilter ($var)
{
    if($var === false || $var === null) {
        return false;
    } 
    return true;

}

print_r(array_filter($entry, "myfilter"));
*/

// array_flip 交换数组中的键和值 什么情况可以用到? 
/*
$input = array("oranges", "apples", "pears", null);
$flipped = array_flip($input);

print_r($flipped);
*/

# 2018年08月08日10:06:40 end

# 2018年08月09日09:29:14 start 

// array_intersect_assoc 带索引检查计算数组的差集##对比数组中有键和值都一样的元素
/*
$array1 = array("a" => "green", "b" => "brown", "c" => "blue", "red");
$array2 = array("a" => "green", "b" => "yellow", "blue", "red");
$result_array = array_intersect_assoc($array1, $array2);
print_r($result_array);
*/

// array_intersect_key 用键名比较计算数组的交集##返回键值一样的元素, 是array1中的. 
/*
$array1 = array('blue'  => 1, 'red'  => 2, 'green'  => 3, 'purple' => 4);
$array2 = array('green' => 5, 'blue' => 6, 'yellow' => 7, 'cyan'   => 8);

var_dump(array_intersect_key($array1, $array2));
*/

// array_intersect_uassoc 带索引计算数组的交集, 用回调函数比较索引##回调函数返回0为交集
// 使用自定义函数接收的键值比较. 系统函数 strcasecmp 会比较键和值, 键不区分大小写, 值会区分大小.  该数组包含了所有在 array1 中也同时出现在所有其它参数数组中的值。注意和 array_intersect() 不同的是(array_intersect_uassoc)键名也用于比较
/*
$array1 = array("a" => "green", "b" => "brown", "c" => "blue", "red");
$array2 = array("A" => "green", "c" => "brown", "yellow", "red");

print_r(array_intersect_uassoc($array1, $array2, "strcasecmp")); // "a" => "green" 为什么不输出?


function myfunction($a,$b)
{
    echo $a .'#'. $b . "\n";
    if ($a===$b)
    {
        return 0;
    }
    return ($a>$b)?1:-1;
}

$a1=array("a"=>"red","b"=>"green","c"=>"blue");
$a2=array("a"=>"red","b"=>"green","e"=>"blue");

// $result=array_intersect_uassoc($array1,$array2,"myfunction"); // "a" => "green" 和 "a" => "GREEN" 为什么没有比较? 这个不是应该输出是 "a" => "green" ? 
print_r($result);

// $a = 'a';
// $b = 'a';
// $c = strcasecmp($a, $b);
// var_dump($c);
*/

// array_intersect_ukey 用回调函数比较建名来计算数组的交集##回调函数返回0时认为元素交集
/*
function key_compare_func($key1, $key2)
{
    echo $key1 .'#'. $key2 . "\n";
    if ($key1 == $key2)
        return 0;
    else if ($key1 > $key2)
        return 1;
    else
        return -1;
}

$array1 = array('blue'  => 1, 'red'  => 2, 'green'  => 3, 'purple' => 4);
$array2 = array('green' => 5, 'blue' => 6, 'yellow' => 7, 'cyan'   => 8);

var_dump(array_intersect_ukey($array1, $array2, 'key_compare_func'));
*/

// array_interscet 计算数组交集, 键名保留不变##以array1为主值其他数组出现array1的值则为交集元素
/*
$array1 = array("a" => "green", "red", "blue", 'ddd' => 1);
$array2 = array("b" => "green", "yellow", "red", 1);
$result = array_intersect($array1, $array2);
print_r($result);
*/

// 2018年08月09日10:41:09 end

// 2018年08月11日16:34:27 start

// array_key_exists 指定的数组是否有指定的键名或索引##只能是一维数组
// 多维数组的key如何检测?
/*
$search_array = array('firsta' => null, 'second' => 4);
// if (array_key_exists('first', $search_array)) {
//     echo "The 'first' element is in the array\n";
// }

$a = checkKey('first', $search_array);
var_dump($a);

function checkKey($key, $arr) {
    return (isset($arr [$key]) || array_key_exists($key, $arr));
    // 速度比直接使用array_key_exists要快 http://php.net/manual/zh/function.array-key-exists.php#107786
}
*/

// array_key_first
// array_key_last php7新加函数


// array_keys 返回数组中部分或全部的键名
/*
// $array = array(0 => 100, "color" => "red");
// print_r(array_keys($array));

// $array = array("blue", "red", "green", "blue", "blue");
// print_r(array_keys($array, "blue")); // 利用这个函数检测数组中的值是否存在


$array = array("color" => array( 'a' => "blue", "red", "green"),
               "size"  => array("small", "medium", "large"));
print_r(array_keys($array));
*/

// array_map 为数组的每个元素应用回调函数 TODO::有哪些好玩的用法
/*
function cube($n)
{
    return($n * $n * $n);
}

$a = array(1, 2, 3, 4, 5);
$b = array_map("cube", $a);
print_r($b);
*/

// array_merge_reursive 递归合并一个或多个数组#相同的键会覆盖但值不会
/*
$ar1 = array("color" => array("favorite" => "red"), 5);
$ar2 = array(10, "color" => array("favorite" => "green", "blue"));
$result = array_merge_recursive($ar1, $ar2);
print_r($result);
*/

// array_merge 合并一个或多个数组
/*
$array1 = array("color" => "red", 2, 4, array('d' => 111));
$array2 = array("a", "b", "color" => "green", "shape" => "trapezoid", 4);
$result = array_merge($array1, $array2);
print_r($result);
*/

//date:2018年08月11日17:34:27 end


//date:2018年08月12日13:55:59 start
/*
// array_multisort 对多个数组或多维数组进行排序

// $ar1 = array(10, 100, 100, 0);
// $ar2 = array(5, 3, 2, 4);
// array_multisort($ar1, $ar2); //第二个数组排序位置对应第一组数组排序后的位置

// var_dump($ar1);
// var_dump($ar2);


// $ar = array(
//        array("10", 11, 100, 100, "a"),
//        array(   1,  2, "2",   3,   1)
//       );
// array_multisort($ar[0], SORT_ASC, SORT_STRING,
//                 $ar[1], SORT_NUMERIC, SORT_DESC);
// var_dump($ar);

$data[] = array('volume' => 67, 'edition' => 2);
$data[] = array('volume' => 86, 'edition' => 1);
$data[] = array('volume' => 85, 'edition' => 6);
$data[] = array('volume' => 98, 'edition' => 2);
$data[] = array('volume' => 86, 'edition' => 6);
$data[] = array('volume' => 67, 'edition' => 7);

// 取得列的列表
foreach ($data as $key => $row) {
    $volume[$key]  = $row['volume'];
    $edition[$key] = $row['edition'];
}

// 将数据根据 volume 降序排列，根据 edition 升序排列
// 把 $data 作为最后一个参数，以通用键排序
array_multisort($volume, SORT_DESC, $edition, SORT_ASC, $data);

var_dump($data);
*/

// array_pad 以指定长度将一个值填充进数组##从1开始
/*
$input = array(12, 10, 9);

$result = array_pad($input, 5, 0);
var_dump($result);
// result is array(12, 10, 9, 0, 0)

$result = array_pad($input, -7, -1);
// result is array(-1, -1, -1, -1, 12, 10, 9)

$result = array_pad($input, 2, "noop");
*/

// array_pop 弹出数组最后一个单元, 弹出并返回 array 数组的最后一个单元，并将数组 array 的长度减一
/*
$stack = array("orange", "banana", "apple", "raspberry");
$fruit = array_pop($stack);
print_r($stack);
*/
// array_product 计算数组中所有值的积, 空数组现在会产生 1，而之前此函数处理空数组会产生 0。
/*
$a = array(2, 4, 6, 8);
echo "product(a) = " . array_product($a) . "\n";
echo "product(array()) = " . array_product(array()) . "\n";
*/

// array_push 将一个或多个单元压入数组
/*
$stack = array("orange", "banana");
array_push($stack, "apple", "raspberry");
print_r($stack);
*/

//date:2018年08月12日14:39:08 end


//date:2018年08月13日14:51:32 start

// array_rand() 从数组中随机取出一个或多个单元# 键值不返回
/*
$input = array("Neo", 'a' => "Morpheus", "Trinity", "Cypher", "Tank");
$rand_keys = array_rand($input, 3);
echo $input[$rand_keys[0]] . "\n";
echo $input[$rand_keys[1]] . "\n";
*/

//TODO::array_reduce() 用回调函数迭代地将数组简化为单一的值#应该还有其他玩法
/*
function sum($carry, $item)
{
    $carry += $item;
    return $carry;
}

function product($carry, $item)
{
    $carry *= $item;
    return $carry;
}

$a = array(1, 2, 3, 4, 5);
$x = array();

var_dump(array_reduce($a, "sum")); // int(15)
var_dump(array_reduce($a, "product", 10)); // int(1200), because: 10*1*2*3*4*5
var_dump(array_reduce($x, "sum", "No data to reduce")); // string(17) "No data to reduce"
*/

//TODO::array_replace_recursive() 使用传递的数组递归替换第一个数组的元素
/*
$base = array('citrus' => array( "orange") , 'berries' => array("blackberry", "raspberry"), );
$replacements = array('citrus' => array('pineapple'), 'berries' => array('blueberry'));

$basket = array_replace_recursive($base, $replacements);
print_r($basket);

$basket = array_replace($base, $replacements);
print_r($basket);
*/

// array_replace() 使用传递的数组替换第一个数组的元素
/*
$base = array("orange", "banana", "apple", "raspberry");
$replacements = array(0 => "pineapple", 4 => "cherry");
$replacements2 = array(0 => "grape");

$basket = array_replace($base, $replacements, $replacements2);
print_r($basket);
*/

// array_reverse() 返回单元顺序的数组
/*
$input  = array("php", 4.0, array("green", "red"));
$reversed = array_reverse($input);
$preserved = array_reverse($input, true);

print_r($input);
print_r($reversed);
print_r($preserved);
*/

//date:2018年08月13日15:18:47 end 


//date:2018年08月14日11:19:17 start 

// array_search 在数组中搜索给定的值, 如果成功返回对应值的键
/*
$array = array(0 => 'blue', 1 => 'red', 2 => 'green', 3 => 'red');

$key = array_search('blue', $array); // $key = 2;
var_dump($key);
$key = array_search('red', $array);   // $key = 1;

var_dump($key);
*/

// array_shift 将数组开头的单元移出数组#key会重新排序
/*
$stack = array(1 => "orange", 2 => "banana", "apple", "raspberry");
$fruit = array_shift($stack);
print_r($stack);
*/

//TODO:array_slice 从数组中取一段
/*
$input = array("a", "b", "c", "d", "e");

$output = array_slice($input, 2);      // returns "c", "d", and "e"
$output = array_slice($input, -2, 1);  // returns "d"
$output = array_slice($input, 0, 3);   // returns "a", "b", and "c"

// note the differences in the array keys
print_r(array_slice($input, 2, -1));
print_r(array_slice($input, 2, -1, true));
*/

//TODO:array_splice 去掉数组中的某一部分并用其他值代替#手册中的key从1开始, 实际是从0开始
/*
$input = array("red", "green", "blue", "yellow");
$a = array_splice($input, 2);
var_dump($a);
// $input is now array("red", "green")

$input = array("red", "green", "blue", "yellow");
$b = array_splice($input, 1, -1);
var_dump($b);
// $input is now array("red", "yellow")

$input = array("red", "green", "blue", "yellow");
array_splice($input, 1, count($input), "orange");
// $input is now array("red", "orange")

$input = array("red", "green", "blue", "yellow");
array_splice($input, -1, 1, array("black", "maroon"));
// $input is now array("red", "green",
//          "blue", "black", "maroon")

$input = array("red", "green", "blue", "yellow");
array_splice($input, 3, 0, "purple");
// $input is now array("red", "green",
//          "blue", "purple", "yellow");
*/

// array_sum 对数组中所有值求和
/*
$a = array(2, 4, 6, 8);
echo "sum(a) = " . array_sum($a) . "\n";

$b = array("a" => 1.2, "b" => 2.3, "c" => 3.4);
echo "sum(b) = " . array_sum($b) . "\n";
*/

//date:2018年08月14日11:37:50 end
