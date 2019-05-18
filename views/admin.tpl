<!DOCTYPE html>
<html>
<head>
  <title>LAM</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>
<style>
input{
   background:none;
   outline:none;
   border:0px;
 }
</style>
<body>
{{if .Finish}}{{ if .OK }}<script>alert('用户资料修改成功');</script>{{ else }}<script>alert('用户资料修改失败请重试');</script>{{ end }}{{else}}{{end}}

<center>
   <h1 class="logo">欢迎来到宠物商城后台</h1><br>
   <h2 class="logo">--我的资料--</h2><br>
   <h4>昵称:{{.UserName}}</h4>
   <h4>登录时间:{{.Time}}</h4>
    <table border="2px" cellspacing="0px" cellpadding="8px">
    <tr>
       <td>用户ID</td>
       <td>角色</td>
       <td>昵称</td>
       <td>密码</td>
       <td>邮箱</td>
       <td>操作</td>
    </tr>
    <tr>
       <td>{{.Uid}}</td>
       <td>{{if .Role}}管理员{{else}}普通会员{{end}}</td>
       <form action="/user/updateinfo" method="post">
       <td><input type="text" name="username" value="{{.UserName}}"/></td>
       <td><input type="password" name="password" placeholder="必填"/></td>
       <td><input type="text" name="email" value="{{.Email}}"/></td> 
       <td><button type="submit">确认修改并提交</button></td>
    </tr>
    </form>
    </table>

    <h2 class="logo">--所有用户信息--</h2><br>
    <table border="2px" cellspacing="0px" cellpadding="8px">
     <tr>
       <td>用户ID</td>
       <td>昵称</td>
       <td>邮箱</td>
       <td colspan="2">操作</td>
    </tr>
     {{range $values:=.UserList}}
    <tr>
       <td><input type="text" name="uid" onfocus=this.blur() value="{{$values.Uid}}"/></td>
       <form action="/admin/updateuser/{{$values.Uid}}" method="post">
       <td><input type="text" name="username" value="{{$values.UserName}}"/></td>
       <td><input type="text" name="email" value="{{$values.Email}}"/></td> 
       <td><button type="submit">确认修改并提交</button></td>
       </form>
       <td><button onclick="window.location.href='/admin/deluser/{{$values.Uid}}'">删除该账号</button></td>
    </tr>
     {{end}}
     <tr>
     <td colspan="2">{{.usernum}}条记录被找到</td>
     <td colspan="2"><button onclick="window.location.href='/register'">创建新用户</button></td>
     </tr>
    </table>

        <h2 class="logo">--所有产品信息--</h2><br>
    <table border="2px" cellspacing="0px" cellpadding="8px">
     <tr>
      <td>产品ID</td>
      <td>产品名</td>
      <td>当前售价</td>
      <td>库存</td>
      <td>图片一</td>
      <td>图片二</td>
      <td>图片三</td>
      <td>图片四</td>
      <td>图片五</td>
      <td>图片六</td>
      <td colspan="2">操作</td>
    </tr>
     {{range $values:=.ProductList}}
    <tr>
      <td><input type="text" name="pid" onfocus=this.blur() value="{{$values.Pid}}"/></td>
      <form action="/user/updateinfo/{{$values.Pid}}" method="post">
      <td><input type="text" name="pname" value="{{$values.Pname}}"/></td>
      <td><input type="text" name="oprice" value="{{$values.Oprice}}"/></td>
      <td><input type="text" name="onhand" value="{{$values.Onhand}}"/></td>
      
      <td><a href="/product/{{$values.Pid}}"><img src="{{$values.Pimg1}}" width="65" height="65"/></a></td>
      <td><a href="/product/{{$values.Pid}}"><img src="{{$values.Pimg2}}" width="65" height="65"/></a></td>
      <td><a href="/product/{{$values.Pid}}"><img src="{{$values.Pimg3}}" width="65" height="65"/></a></td>
      <td><a href="/product/{{$values.Pid}}"><img src="{{$values.Pimg4}}" width="65" height="65"/></a></td>
      <td><a href="/product/{{$values.Pid}}"><img src="{{$values.Pimg5}}" width="65" height="65"/></a></td>
      <td><a href="/product/{{$values.Pid}}"><img src="{{$values.Pimg6}}" width="65" height="65"/></a></td>
      <td><button type="submit">确认修改并提交</button></td>
      </form>
      <td><button onclick="window.location.href='/admin/delproduct/{{$values.Pid}}'">删除该商品</button></td>
    </tr>
     {{end}}
     <tr>
     <td colspan="4">{{.productnum}}条记录被找到</td>
     <td colspan="7"><button onclick="window.location.href='/admin'">添加新产品</button></td>
     </tr>
    </table>
</center>
</body>
</html>