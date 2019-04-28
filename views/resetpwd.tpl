<!DOCTYPE html>
<html lang="en">
<head>
  <title>LAM</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>

<body>
<center>
    <h1>You can reset your password</h1><br>
    <h4>User's Name:{{.UserName}}</h4>
    <h4>Login Time:{{.Time}}</h4>
    <form action="/resetpwd" method="post">
            Old Password:<input type="password" name="OldPassword"><br>
            New Password:<input type="password" name="NewPassword"><br>
            Enter Again:<input type="password" name="SurePassword"><br>
            <button type="submit">submit</button>
    </form>
</center>
</body>
</html>