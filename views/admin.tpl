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
    Your Information: 
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
    </table><br>
    All users' Information: 
    <table border="2px" cellspacing="0px" cellpadding="8px">
    <tr>
       <td>Id</td>
       <td>Name</td>
       <td>Password</td>
       <td>Role</td>
       <td>Email</td>
       <td colspan="2">Action</td>
    </tr>
     {{range $values:=.users}}
     <tr>
     <td>{{$values.Uid}}</td>
     <td>{{$values.UserName}}</td>
     <td>{{$values.Password}}</td>
     <td>{{$values.Role}}</td>
     <td>{{$values.Email}}</td>
     <td><button onclick="javascript:if(confirm('Are you sure you want to delete it?'))window.location.href='/admin/delete/{{$values.Uid}}'">Delete</button></td>
     <td><button onclick="window.location.href='/admin/update/{{$values.Uid}}'">Update</button></td>
     </tr>
     {{end}}
     <tr>
     <td colspan="2">{{.num}} records were found</td>
     <td colspan="3"><button onclick="window.location.href='/admin/update/0'">Create</button></td>
     </tr>
    </table>
</center>
</body>
</html>