<!DOCTYPE html>
<html>
<head>
  <title>LAM</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>

<body>
<center>
    <h1 class="logo">Welcome to LAM</h1><br>
    <h4>User's Name:{{.UserName}}</h4>
    <h4>Login Time:{{.Time}}</h4>
    <table border="2px" cellspacing="0px" cellpadding="8px">
    <tr>
       <td>Id</td>
       <td>Name</td>
       <td>Password</td>
       <td>Role</td>
       <td>Email</td>
    </tr>
    <tr>
       <td>{{.Uid}}</td>
       <td>{{.UserName}}</td>
       <td>{{.Password}}</td>
       <td>{{.Role}}</td>
       <td>{{.Email}}</td>
    </tr>
    <tr>
       <td colspan="2"><button onclick="javascript:window.location.href='/resetemail'">Reset Email</button></td>
       <td colspan="3"><button onclick="javascript:window.location.href='/resetpwd'">Reset Password</button></td>
    </tr>
    </table>
</center>
</body>
</html>