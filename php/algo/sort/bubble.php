<?php

function bubble($arr) {
    $len = count($arr);
    if($len <= 0) {return false;}
    $falg = true;
    for($i = 0; $i<$len; $i++) {

        for($j = 0; $j<$len - $i -1; $j++) {
            if($arr[$j] > $arr[$j+1]) {
                $temp = $arr[$j+1];
                $arr[$j+1] = $arr[$j];
                $arr[$j] = $temp;
                $falg = false;
            }
        }
        if($falg) {
            break;
        }
    }

    return $arr;
}

$arr = [2,1,9,8,4,7];
$arr2 = [1,2,3,4,5,6];
$res = bubble($arr2);
print_r($res);