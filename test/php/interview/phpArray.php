<?php

// 代码测试环境 env/php/5.6  1天5个函数 
// docker run -idt -v $PWD:/app mybasephp:5.6-fpm

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









