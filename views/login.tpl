<!DOCTYPE html>
<html>

	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width,initial-scale=1,minimum-scale=1,maximum-scale=1,user-scalable=no" />
		<link rel="stylesheet" href="../static/css/amazeui.min.css" />
		<script type="text/javascript" src="../static/js/jquery-1.9.1.min.js"></script>
		<script type="text/javascript" src="../static/js/amazeui.min.js"></script>
		<link rel="stylesheet" href="../static/css/login.css" />
		<link rel="stylesheet" href="../static/css/mui.min.css" />
		<script type="text/javascript" src="../static/js/mui.min.js" ></script>
		<script type="text/javascript" src="../static/js/login.js" ></script>
		<title></title>
		<style>
			.bottomLine{
				border-bottom: 1px solid #ccc;
				border-top: 1px solid #ccc;
				padding-top:8px;
				/*padding-bottom:10px;*/
				
			}
			.bot{
				/*margin-top: 5px;*/
				margin-bottom: 49px;
				/*text-align: center;*/
				color: #7b6f5b;
			}
		</style>
	</head>
     
	<body>
		<div class="am-g"  style="margin-top:5px;">
			<div style="width: 1200px;margin: 0 auto;">
			<div class="logo" >
				<img src="../static/img/logo.png" />
			</div>
			</div>
		</div>
		<div class="am-g content"  style="margin-top: 10px;">
			<div style="width:1200px;margin:0 auto;position: relative;z-index: 999;">
				<div class="loginDiv " style="display: block;"  id="login" >
					<div class="am-u-sm-12" style="padding-right: 0rem">
						<img src="../static/img/erweima.png" style="float: right;height: 58px;" id="loginimg"/ >						
					</div>
					<div class="am-u-sm-12" style="padding-left: 30px;padding-right: 30px;">
					<h4 style="font-weight: normal;">欢迎登录</h4>
				 	<span class="left1"></span>	<input type="text" class="am-form-field"  name="username" placeholder="邮箱/手机号/用户名" style="padding-left: 50px;margin-top: 48px;font-size: 12px;"/>
				    </div>	
					<form action="/" method="post">
				        <div class="am-u-sm-12" style="margin-top: 21px;padding-left: 30px;padding-right: 30px;">
				            <span class="left2"></span>	  <input type="password" class="am-form-field"  name="password" placeholder="请输入密码"style="padding-left: 50px;font-size: 12px;"/>
				        </div>	
				        <div class="am-u-sm-12" style="margin-top: 12px;padding-left: 30px;padding-right: 30px;font-size: 12px;color: #757575;">
				            <input type="checkbox" style="margin-top:1px;vertical-align:middle;"/><span style="margin-bottom:-1px;margin-left: 10px;font-size: 12px;vertical-align: middle;">自动登录</span>
				        </div>	
				        <div class="am-u-sm-12" style="margin-top: 45px;padding-left: 30px;padding-right: 30px;">
				            <button type="submit" class="am-btn am-btn-warning" style="width: 100%;border-radius: 5px;font-size: 16px;" >登录</button>	
						</div>
			        </form>
				<div class="am-u-sm-12" style="padding-left: 30px;padding-right: 30px;">
					<p style="width: 100%;margin-top: 15px;font-size: 12px;color: #333333;"><a href="/building">忘记登录密码</a><span style="float: right;"><a href="/regist">免费注册</a></span></p>
				</div>
				
				</div>
				<!--扫码页面-->
				<div class="loginDiv" style="display: none;" id="erm"  >
						<div class="am-u-sm-12" style="padding-right: 0rem">
						    <img src="../static/img/pc_2.png" style="float: right;" id="ermimg" />						
						</div>
					<div class="am-u-sm-12" style="padding-left: 30px;padding-right: 30px;margin-top: 30px;text-align: center;">
					<h4 >手机扫码，安全登录</h4>
				</div>	
				 <div class="am-u-sm-12" style="padding-left: 30px;padding-right: 30px;margin-top: 50px;text-align: center;">
				 <img src="../static/img/aerweima.png" />
				 </div>	
				 
				 <div class="am-u-sm-12" style="margin-top: 50px;padding-left: 30px;padding-right: 30px;text-align: center;">
				  <p style="font-size: 14px;">请使用宠物网客户端扫描二维码登录</p>
				    <p style="text-align: center;"><span style="background-image: url(../static/img/Download-512.png);float: left;width: 20px;height: 20px;margin-top:5px;background-repeat: no-repeat;margin-left: 100px;"></span>
				<a style="float: left;font-size: 12px;" href="/building" >下载宠物客户端</a></p> 
				 </div>	
				<div class="am-u-sm-12" style="margin-top: 10px;padding-left: 30px;padding-right: 30px;margin-bottom: 100px;">
				</div>
				
				</div>
			</div>
		</div>
		<div class="am-g " style="margin-top: 81px;text-align: center;font-size: 12px;">
			<div class="bottomLine" style="margin:0 auto;width:700px;height: 35px;">
			<span style="float: left;word-spacing:0.52rem;    text-align: center;    width: 100%;">关于我们 | 联系我们 | 商家入驻 | 友情链接 | 站点地图 | 宠物商城 | 销售联盟 | 商城社区 | 企业文化 | 帮助中心 </span>
			</div>
	    </div>	
	    <div class="am-g" >
				<div class="bot ">
					<div  style="margin:0 auto;width:700px;height: 30px;">
				        <p style="font-size: 12px;float: left; line-height: 10px;">© &nbsp;2007-2018 宠物网 版权所有 ，并保留所有权利  <span  style="margin-left: 30px;">宠物 Tel ：4008125181 </span><span  style="margin-left: 20px;">E-mai：wilh@163.com</span>  </p>
					</div>
					</div>
		</div>
		
	</body>

</html>