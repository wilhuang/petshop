<!DOCTYPE html>
<html>

	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width,initial-scale=1,minimum-scale=1,maximum-scale=1,user-scalable=no" />
		<link rel="stylesheet" href="../static/css/amazeui.min.css" />
		<script type="text/javascript" src="../static/js/jquery-1.9.1.min.js"></script>
		<script type="text/javascript" src="../static/js/amazeui.min.js"></script>
		<link rel="stylesheet" href="../static/css/regedit.css" />
		<link rel="stylesheet" href="../static/css/sui.css" />
		<script type="text/javascript" src="../static/js/sui.js"></script>
		<title>宠物网</title>
		<style>
			/*.bottomLine {
				border-bottom: 1px solid #ccc;
				border-top: 1px solid #ccc;
				margin-bottom: 10px;
				/*padding: 10px;*/
				/*padding-top: 10px;
				padding-bottom: 10px;
				word-spacing: 0.05rem;
				text-align: center;
				margin-top: 82px;*/
				/*width:677px ;*/
				/*line-height: ;*/
			/*}*/
				.bottomLine{
				border-bottom: 1px solid #ccc;
				border-top: 1px solid #ccc;
				padding-top:8px;
				/*padding-bottom:10px;*/
				
			}
			.bot {
				margin-bottom: 50px;
				/*text-align: center;*/
				color: #7b6f5b;
			}
			
			.radio-pretty.checked > span:before {
				color: #f88600;
			}
			
			input {
				height: 62px;
			}
				.radio-pretty span:before {
		    margin-right: 2px;
		    vertical-align: -4px;
		    font-size: 20px;
		    color: #bdbdbd;
		    margin-left: -2px;
		}
		</style>
	</head>

	<body>
		<div class="am-g" style="margin-top: 25px;">
			<div class="logo"> <img src="../static/img/logo.png" /> </div>
		</div>
		<div class="am-g">
			<div  style="width: 999px;margin: 0 auto;font-size: 14px;"> <span style="float: right;color: #333333;">我已注册，马上<span style="color: #F88600;"><a href="/login">登录</a></span></span>
			</div>
		</div>
		<div class="am-g content" style="margin-top: 12px;">
			<div data-am-widget="tabs" class="am-tabs am-tabs-d2">
				<!-- 用户注册信息表 -->
				<div class="am-tabs-bd" style="margin-top: 59px;margin-left: -20px;">
					<form action="/register" method="post">
						<div class="am-g">
							<div class="am-u-sm-7 am-u-sm-offset-3 "> <span class="left4"></span> <input type="text" name="username" placeholder="请输入您的用户名" class="am-form-field" style="margin: auto;display: block;float: left;padding-left: 55px;width: 528px" /> </div>
							<div class="am-u-sm-1 am-u-end" style="padding: 10px;"> <span class="right" style="display: none;"></span> </div>
						</div>
						<div class="am-g" style="margin-top: 16px;">
							<div class="am-u-sm-7 am-u-sm-offset-3 "> <span class="left3"></span> <input type="password" name="password" placeholder="请输入您的密码" class="am-form-field" style="margin: auto;display: block;float: left;padding-left: 55px;width: 528px" /> </div>
							<div class="am-u-sm-1 am-u-end" style="padding: 10px;"> <span class="right" style="display: none;"></span> </div>
						</div>
						<div class="am-g" style="margin-top: 16px;">
							<div class="am-u-sm-7 am-u-sm-offset-3 "> <span class="left3"></span> <input type="password" placeholder="请再次输入密码" class="am-form-field" style="margin: auto;display: block;float: left;padding-left: 55px;width: 528px" /> </div>
							<div class="am-u-sm-1 am-u-end" style="padding: 10px;"> <span class="right" style="display: none;"></span> </div>
						</div>
						<div class="am-g" style="margin-top: 16px;">
							<div class="am-u-sm-7 am-u-sm-offset-3 "> <span class="left1"></span> <input type="email" name="email" placeholder="请输入您的邮箱" class="am-form-field" style="margin: auto;display: block;float: left;padding-left: 55px;width: 528px" /> </div>
							<div class="am-u-sm-1 am-u-end" style="padding: 10px;"> <span class="right" style="display: none;"></span> </div>
						</div>
						<div class="am-g" style="margin-top: 16px;">
							<div class="am-u-sm-4 am-u-sm-offset-3" style="padding-right: 0rem;"> <span class="left2"></span><input type="text" placeholder="请输入邮箱验证码" class="am-form-field" style="margin: auto;display: block;float: left;padding-left: 55px;width: 360px;" /> </div>
							<div class="am-u-sm-3 am-u-end" style="padding-left: 0rem;"> <button type="button" class="am-btn am-btn-default" style="width: 190px;    float: left;height: 62px;">获取邮箱验证码</button> </div>
						</div>
						<div class="am-g" style="margin-top: 16px;">
							<div class="am-u-sm-4 am-u-sm-offset-3" style="padding-right: 0rem;"> <span class="left2"></span> <input type="text" placeholder="请输入校验码" class="am-form-field" style="margin: auto;display: block;float: left;padding-left: 55px;width: 360px" /> </div>
							<div class="am-u-sm-3 am-u-end" style="padding-left: 0rem;"> <input type="image" style="width: 190px;float: left;height: 62px;" value="ee" src="../static/img/uwv6.png" /></div>
						</div>
						<div class="am-g" style="margin-top: 21px;">
								<div class="am-u-sm-7 am-u-sm-offset-3" style="padding-right: 0rem;"><input type="checkbox" checked="checked"><span style="font-size: 12px;color: #878787;">  同意<a>《宠物网会员章程》</a> <a>《支付宝协议》</a> 及 <a>《宠物网商业联盟在线协议》</a></span></div>
						</div>
						<div class="am-g" style="margin-top:70px;">
							<div class="am-u-sm-7 am-u-sm-offset-3 "> <button type="submit" class="am-btn am-btn-warning" style="width: 528px;padding: 20px;border: 1px rgba(187, 187, 187, 0.5) solid;"><font style="font-size: 20px;font-weight: bold;">同意协议并注册</font></button> </div>
						</div>
					</form> 
				</div>
				<div class="am-g" style="margin-top: 20px;margin-bottom: 135px;">
					<div class="am-u-sm-7 am-u-sm-offset-3" style="padding-right: 0rem;"> <a style="font-size: 14px;" href="#">您也可以使用手机号注册></a> </div>
				</div>
			</div>
		</div>
		  {{if .Finish}}{{ if .UserName }}<script>alert('用户{{.UserName}}注册成功');</script>{{ else }}<script>alert('注册失败请重试');</script>{{ end }}{{else}}{{end}}
			<div class="am-g " style="margin-top: 40px;text-align: center;font-size: 12px;">
			<div class="bottomLine" style="margin:0 auto;width:700px;height: 35px;">
			<span style="float: left;word-spacing:0.52rem;    text-align: center;    width: 100%;">关于我们 | 联系我们 | 商家入驻 | 友情链接 | 站点地图 | 宠物商城 | 销售联盟 | 商城社区 | 企业文化 | 帮助中心 
			</span>
			</div>
	    </div>	
			<div class="am-g">
			<div class="bot ">
					<div  style="margin:0 auto;width:700px;height: 30px;">
							<p style="font-size: 12px;float: left; line-height: 10px;">© &nbsp;2007-2018 宠物网 版权所有 ，并保留所有权利  <span  style="margin-left: 30px;">宠物 Tel ：4008125181 </span><span  style="margin-left: 20px;">E-mai：wilh@163.com</span>  </p>
					</div>
			</div>
		</div>
	</body>

</html>