<?php
$para1 = $_REQUEST["input1"];
$para2 = $_REQUEST["input2"];
$prog = '.\\prog\\app.exe';
exec("$prog $para1 $para2", $outarr ,$rv);
foreach($outarr as $v)
{
	echo $v."<br />";
}
echo "return: ".$rv;
//var_dump($outarr);
?>
