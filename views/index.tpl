<!doctype html>
<html class="no-js" lang="zxx">

<head>
    <meta charset="utf-8">
    <title>宠物网</title>
    <meta http-equiv="x-ua-compatible" content="ie=edge">
    <meta name="description" content="">
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <!-- Favicon -->
    <link rel="shortcut icon" type="image/x-icon" href="../static/img/favicon.png">

    <!-- All css files are included here -->
    <!-- Bootstrap fremwork main css -->
    <link rel="stylesheet" href="../static/css/bootstrap.min.css">
    <!-- This core.css file contents all plugings css file. -->
    <link rel="stylesheet" href="../static/css/core.css">
    <!-- Theme shortcodes/elements style -->
    <link rel="stylesheet" href="../static/css/shortcode/shortcodes.css">
    <!-- Theme main style -->
    <link rel="stylesheet" href="../static/css/style.css">
    <!-- Responsive css -->
    <link rel="stylesheet" href="../static/css/responsive.css">
    <!-- User style -->
    <link rel="stylesheet" href="../static/css/custom.css">

    <!-- Modernizr JS -->
    <script src="../static/js/vendor/modernizr-2.8.3.min.js"></script>
</head>
<body>
    <!-- top start -->
    <div class="header-top gray-bg">
        <div class="container">
            <div class="row">
                <div class="col-sm-5 hidden-xs">
                    <div class="header-top-left">
                        <ul class="header-top-style text-capitalize mr-25">
                            <li><a href="#"><span class="mr-10">{{ if .UserName }}您好，{{.UserName}}{{ else }}<a href="/login">请登录</a>{{ end }}</span><i class="fa fa-angle-down"></i></a>
                                <ul class="ul-style my-account box-shadow white-bg">     
                                    <li><a href="/user/order">我的订单</a></li>
                                    <li><a href="/user">我的资料</a></li>
                                    <li><a href="/logout">登出</a></li>
                                </ul>
                            </li>
                        </ul>
                        <ul class="header-top-style text-capitalize mr-25">
                            <li><a href="/register"><span class="mr-10">免费注册</span></a>
                            </li>
                        </ul>
                        <ul class="header-top-style text-capitalize mr-25">
                            <li><a href="#"><span class="mr-10">手机逛petshop</span></a>
                            </li>
                        </ul>
                        <ul class="header-top-style pl-10">
                            <li>
                                <span class="mr-10"><img alt="" src="../static/img/header/1-min.jpg"></span>
                                <a href="#"><span class="mr-10">中文</span><i class="fa fa-angle-down"></i></a>
                                <ul class="ul-style language box-shadow white-bg">
                                    <li><a href="#"><img alt="" src="../static/img/header/1-min.jpg"><span>Chinese</span></a></li>
                                    <li><a href="#"><img alt="" src="../static/img/header/2-min.jpg"><span>English</span></a></li>
                                </ul>
                            </li>
                        </ul>
                    </div> 
                </div>
                <div class="header-top-right">
                    <ul class="header-top-style text-capitalize mr-25">
                        <li><a href="/"><span class="mr-10">首页</span></a>
                        </li>
                    </ul>
                    <ul class="header-top-style text-capitalize mr-25">
                        <li><a href="/cart"><span class="mr-10">我的购物车</span></a>
                        </li>
                    </ul>
                </div>
            </div>
        </div>
    </div>

    <center style="margin: 30px 300px;">
        <form class="subscribe-box" action="#">
            <div class="subscribe-box">
                <input type="txt" required="" placeholder="请输入关键词...">
                <button type="submit" id="search-letter-btn" style="color:#EEE;position: absolute;right: 0;top: -2px;">搜索</button>
            </div>  
        </form> 
    </center>
    <!-- top end -->
    <!-- body start -->
    <div class="new-arrival-area"> 
        <div class="container">
        
            <div class="row rp-style">
                <div class="col-md-12">
                    <div class="section-title text-center mb-35">
                        <h2 class="text-uppercase"><strong>超高人气宠物出没</strong></h2>
                        <p class="text-defualt">快拿出你的精灵球</p>
                        <img src="../static/img/section-border.png">
                    </div>
                </div>
            </div>

            <div class="row rp-style-2">
                <div class="featured-carousel indicator-style">
                        {{range $Num,$values:=.ProductList1}}
                    {{if eq $Num 0 2 4 6 8 10 12 14 16}}<div class="product-container cp-style-2">{{end}}

                        <div class="product-inner">
                            <a href="/product/{{$values.Pid}}">
                                <div class="product-img b-img">
                                    <img src="{{$values.Pimg1}}">
                                </div>
                            </a>
                            <ul class="quick-veiw text-center">
                                <li><a href="/cart/add/{{$values.Pid}}"><i class="fa fa-shopping-cart"></i></a></li>
                                <li><a href="/product/{{$values.Pid}}"><i class="fa fa-eye"></i></a></li>
                                <li><a href="/"><i class="fa fa-refresh"></i></a></li>
                                <li><a href="#"><i class="fa fa-heart"></i></a></li>
                            </ul>
                            <div class="product-text">
                                <ul class="pull-left list-inline ratings">
                                    <li><i class="rated fa fa-star"></i></li>
                                    <li><i class="rated fa fa-star"></i></li>
                                    <li><i class="rated fa fa-star"></i></li>
                                    <li><i class="rated fa fa-star"></i></li>
                                    <li><i class="rated fa fa-star-half-o"></i></li>
                                </ul>
                                <ul class="pricing list-inline pull-right">
                                    <li class="text-right c-price">￥{{$values.Oprice}}</li>
                                    <li class="text-right p-price">￥{{$values.Pprice}}</li>
                                </ul>
                                <div class="clear"></div>
                                <h6 class="product-name">
                                    <a href="#" title="Eletria ostma">{{$values.Pname}}</a>
                                </h6>
                            </div>
                        </div>
                    {{if eq $Num 1 3 5 7 9 11 13 15 17}}</div>{{end}}

                    {{end}}
                </div>
            </div>

        </div>
    </div> 
    <!-- body1 end -->
    <!--  -->
    <div class="best-saller-iteam-area pb-90">
        <div class="container">
            <div class="row">
                <div class="col-md-6 cp-style-2">

                    <div class="section-title text-center mb-35">
                        <h2 class="text-uppercase"><strong>BEST SELLER</strong></h2>
                        <p class="text-defualt">TOP COLLECTION</p>
                        <img src="../static/img/section-border.png" alt="">
                    </div>

                    <div class="best-seller-carousel indicator-style-two">
                        {{range $Num,$values:=.ProductList2}}
                        {{if eq $Num 0 2 4 6 8 10 12 14 16}}<div class="product-container">{{end}}

                                        <div class="best-product-inner mb-35 row rp-style-3 clearfix">
                                            <div class="col-xs-5 cp-style-3">
                                                <div class="best-product-img b-img">
                                                    <a href="/product/{{$values.Pid}}">
                                                        <img src="{{$values.Pimg1}}" alt="">
                                                    </a>
                                                </div>
                                            </div>
                                            <div class="col-xs-7 cp-style-3">
                                                <div class="best-product-text">
                                                    <h6 class="product-name m-0 p-0">
                                                        <a href="#">{{$values.Pname}}</a>
                                                    </h6>
                                                    <ul class="list-inline ratings">
                                                        <li><i class="rated fa fa-star"></i></li>
                                                        <li><i class="rated fa fa-star"></i></li>
                                                        <li><i class="rated fa fa-star"></i></li>
                                                        <li><i class="rated fa fa-star"></i></li>
                                                        <li><i class="rated fa fa-star"></i></li>
                                                    </ul>
                                                    <ul class="pricing">
                                                        <li class="c-price">￥{{$values.Oprice}}</li>
                                                    </ul>
                                                </div>
                                            </div>
                                        </div>

                        {{if eq $Num 1 3 5 7 9 11 13 15 17}}</div>{{end}}
                        {{end}}
                    </div>
                </div>
                <div class="col-md-6 cp-style-2">

                    <div class="section-title text-center mb-35">
                        <h2 class="text-uppercase"><strong>TRENDY ITEAMS</strong></h2>
                        <p class="text-defualt">NEW TREND</p>
                        <img src="../static/img/section-border.png" alt="">
                    </div>
                                
                    <div class="best-seller-carousel indicator-style">
                            {{range $Num,$values:=.ProductList3}}
                            {{if eq $Num 0 2 4 6 8 10 12 14 16}}<div class="product-container">{{end}}
    
                                            <div class="best-product-inner mb-35 row rp-style-3 clearfix">
                                                <div class="col-xs-5 cp-style-3">
                                                    <div class="best-product-img b-img">
                                                        <a href="/product/{{$values.Pid}}">
                                                            <img src="{{$values.Pimg1}}" alt="">
                                                        </a>
                                                    </div>
                                                </div>
                                                <div class="col-xs-7 cp-style-3">
                                                    <div class="best-product-text">
                                                        <h6 class="product-name m-0 p-0">
                                                            <a href="#">{{$values.Pname}}</a>
                                                        </h6>
                                                        <ul class="list-inline ratings">
                                                            <li><i class="rated fa fa-star"></i></li>
                                                            <li><i class="rated fa fa-star"></i></li>
                                                            <li><i class="rated fa fa-star"></i></li>
                                                            <li><i class="rated fa fa-star"></i></li>
                                                            <li><i class="rated fa fa-star"></i></li>
                                                        </ul>
                                                        <ul class="pricing">
                                                            <li class="c-price">￥{{$values.Oprice}}</li>
                                                        </ul>
                                                    </div>
                                                </div>
                                            </div>
    
                            {{if eq $Num 1 3 5 7 9 11 13 15 17}}</div>{{end}}
                            {{end}}
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <!-- End Of Best Seller Iteams Area -->
            
    <!-- jquery latest version -->
    <script src="../static/js/vendor/jquery-1.12.4.min.js"></script>
    <!-- Bootstrap framework js -->
    <script src="../static/js/bootstrap.min.js"></script>
    <!-- All js plugins included in this file. -->
    <script src="../static/js/plugins.js"></script>
    <!-- Main js file that contents all jQuery plugins activation. -->
    <script src="../static/js/main.js"></script>

    <div style="margin-top: 81px;text-align: center;font-size: 12px;">
        <div style="margin:0 auto;width:700px;height: 35px;">
        <span style="float: left;word-spacing:0.52rem;    text-align: center;    width: 100%;">关于我们 | 联系我们 | 商家入驻 | 友情链接 | 站点地图 | 宠物商城 | 销售联盟 | 商城社区 | 企业文化 | 帮助中心 </span>
        </div>
    </div>	
    <div>
            <div>
                <div  style="margin:0 auto;width:700px;height: 30px;">
                    <p style="color:#999;font-size: 12px;float: left; line-height: 10px;">© &nbsp;2007-2018 宠物网 版权所有 ，并保留所有权利  <span  style="margin-left: 30px;">宠物 Tel ：4008125181 </span><span  style="margin-left: 20px;">E-mai：wilh@163.com</span>  </p>
                </div>
            </div>
    </div>
</body>
</html>