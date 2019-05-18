<html>
<head>
　　<script src="http://code.jquery.com/jquery-1.11.1.min.js" type="text/javascript"></script>
</head>
<body>
　　<input type="checkbox" id="price1" name="chk_list[]" value="1" />1
　　<input type="checkbox" id="price2" name="chk_list[]" value="2" />2
　　<input type="checkbox" id="price3" name="chk_list[]" value="3" />3
　　<input type="checkbox" id="price4" name="chk_list[]" value="4" />4
　　<input type="checkbox" id="price5" name="chk_all" id="chk_all" />全选/取消全选
<script type="text/javascript">
　　$("#chk_all").click(function(){
 　　// 使用prop则完美实现全选和反选
 　　$("input[name='chk_list[]']").prop("checked", $(this).prop("checked"));
　　});
var total=$('#price1').value
alert(total);
</script>
</body>
</html>