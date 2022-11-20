<?php

$arr = array();

$arr1 = $arr;

$arr[0] = 1;
$arr[1] = 2;
$arr[2] = 9;


var_dump($arr, $arr1);
