<!DOCTYPE html>
<html>
<head>
  <title>宠物网</title>
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
    <h1 class="logo">我的资料</h1><br>
    <h4>昵称:{{.UserName}}</h4>
    <h4>登录时间:{{.Time}}</h4>
    <table border="2px" cellspacing="0px" cellpadding="8px">
    <tr>
       <td>用户ID</td>
       <td>角色</td>
       <td>昵称</td>
       <td>密码</td>
       <td>邮箱</td>
    </tr>
    <tr>
       <td>{{.Uid}}</td>
       <td>{{if .Role}}普通会员{{else}}管理员{{end}}</td>
       <form action="/user/updateinfo" method="post">
       <td><input type="text" name="username" value="{{.UserName}}"/></td>
       <td><input type="password" name="password" placeholder="必填"/></td>
       <td><input type="text" name="email" value="{{.Email}}"/></td> 
    </tr>
    <tr>
       <td colspan="5"><button type="submit">确认修改并提交个人资料</button></td>
    </tr>
    </form>
    </table>
</center>
</body>
</html>