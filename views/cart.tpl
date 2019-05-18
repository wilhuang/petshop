<!DOCTYPE html>
<html>
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <title>我的购物车</title>
    <link href="../static/css/cart.css" rel="stylesheet" type="text/css" />
    <script type="text/javascript" src="../static/js/jquery-1.11.1.min.js"></script>


</head>

<body>

<div class="gwc" style=" margin:auto;">
	<table cellpadding="0" cellspacing="0" class="gwc_tb1">
		<tr>
			<td class="tb1_td1"><input name="checkAll" id="checkAll" type="checkbox"  class="allselect"/></td>
			<td class="tb1_td1">全选</td>
			<td class="tb1_td3">商品</td>
			<td class="tb1_td4">商品信息</td>
			<td class="tb1_td5">数量</td>
			<td class="tb1_td6">单价</td>
			<td class="tb1_td7">操作</td>
		</tr>
	</table>
    {{range $Num,$values:=.CartList}}
	<table cellpadding="0" cellspacing="0" class="gwc_tb2">
		<tr>
			<td class="tb2_td1"><input type="checkbox" value="{{$Num}}" name="list[]"/></td>
			<td class="tb2_td2"><a href="/product/{{$values.Pid}}"><img src="{{$values.Pimg1}}"/></a></td>
			<td class="tb2_td3"><a href="/product/{{$values.Pid}}">{{$values.Pname}}</a></td>
			<td class="tb1_td4">{{$values.Pnum}}件</td>
			<td class="tb1_td5">
				<a href="/cart/del/{{$values.Pid}}"><button style="width:20px; height:18px;border:1px solid #ccc;"> - </button></a>
				<input name="" type="text" value="{{$values.Pnum}}" style="width:30px; text-align:center; border:1px solid #ccc;" />
				<a href="/cart/add/{{$values.Pid}}"><button style="width:20px; height:18px;border:1px solid #ccc;"> + </button></a>
			</td>
			<td class="tb1_td6"><label id="total1" class="tot" style="color:#ff5500;font-size:14px; font-weight:bold;"><span>￥</span>{{$values.Oprice}}</label></td>
			<td class="tb1_td7"><a href="/cart/delp/{{$values.Pid}}">删除</a></td>
		</tr>
	</table>
{{end}}
<a href="/"><br>点击这里继续购物...</a>
	<table cellpadding="0" cellspacing="0" class="gwc_tb3">
		<tr>
			<td class="tb3_td2">已选商品 <label style="color:#ff5500;font-size:14px; font-weight:bold;">{{.AllNum}}</label> 件</td>
			<td class="tb3_td3">合计(不含运费):<span>￥</span><span style="color:#ff5500;"><label id="zong1" style="color:#ff5500;font-size:14px; font-weight:bold;">{{.AllPrice}}</label></span></td>
			<td class="tb3_td4"><a href="#"  class="jz2" >结算</a></td>
		</tr>
	</table>
</div>
<script type="text/javascript">
　　$("#checkAll").click(function(){
 　　// 使用prop则完美实现全选和反选
 　　$("input[name='list[]']").prop("checked", $(this).prop("checked"));
　　});
</script>
</body>
</html>
