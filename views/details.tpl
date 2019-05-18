
<!DOCTYPE html>
<html>
<head>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
        <title>商品详情</title>
        <link rel="stylesheet" type="text/css" href="../static/css/normalize.css"/>
        <script type="text/javascript" src="../static/js/jquery-1.9.1.min.js"></script>
        <script src="../static/js/common.js" type="text/javascript" charset="utf-8"></script>
        <script type="text/javascript">
          $(document).ready(function(){
        	  var showproduct = {
        		  "boxid":"showbox",
        		  "sumid":"showsum",
        		  "boxw":400,//宽度,该版本中请把宽高填写成一样
        		  "boxh":400,//高度,该版本中请把宽高填写成一样
        		  "sumw":60,//列表每个宽度,该版本中请把宽高填写成一样
        		  "sumh":60,//列表每个高度,该版本中请把宽高填写成一样
        		  "sumi":7,//列表间隔
        		  "sums":5,//列表显示个数
        		  "sumsel":"sel",
        		  "sumborder":1,//列表边框，没有边框填写0，边框在css中修改
        		  "lastid":"showlast",
        		  "nextid":"shownext"
        		  };//参数定义	  
        	 $.ljsGlasses.pcGlasses(showproduct);//方法调用，务必在加载完后执行
          });
        </script>
</head>
<body>
    <div class="showall">
	    <!--left -->
	    <div class="showbot">
            <div id="showbox">
                <img src="{{.Pimg1}}" width="400" height="400" />
                <img src="{{.Pimg2}}" width="400" height="400" />
                <img src="{{.Pimg3}}" width="400" height="400" />
                <img src="{{.Pimg4}}" width="400" height="400" />
                <img src="{{.Pimg5}}" width="400" height="400" />
                <img src="{{.Pimg6}}" width="400" height="400" />
			</div>
			<!--展示图片盒子-->
            <div id="showsum">
            <!--展示图片里边-->
            </div>
            <p class="showpage">
                <a href="javascript:void(0);" id="showlast"> < </a>
                <a href="javascript:void(0);" id="shownext"> > </a>
            </p>
        </div>
        <!--conet -->
        <div class="tb-property">
        	<div class="tr-nobdr">
        		<h3>{{.Pname}}</h3>
        	</div>
        	<div class="txt">
        		<span class="nowprice">￥<a>{{.Oprice}}</a></span>
        		<div class="cumulative">
        			<span class="number ty1">累计售出<br /><em id="add_sell_num_363">370</em></span>
        			<span class="number tyu">累计评价<br /><em id="add_sell_num_363">25</em></span>
        		</div>
        	</div>
        	<div class="txt-h">
        		<span class="tex-o">分类</span>
        		<ul class="glist" id="glist" data-monitor="goodsdetails_fenlei_click">
        			<li><a title="幼体" href="">幼体</a></li>
        			<li><a title="成熟体" href="">成熟体</a></li>
        		</ul>
        	</div>
        	<script>
        	$(document).ready(function(){
          	    var t = $("#text_box");
        	    $('#min').attr('disabled',true);
            	$("#add").click(function(){    
               	    t.val(parseInt(t.val())+1)
               	    if (parseInt(t.val())!=1){
                   	    $('#min').attr('disabled',false);
                	}
      
                }) 
                $("#min").click(function(){
                    t.val(parseInt(t.val())-1);
                    if (parseInt(t.val())==1){
                        $('#min').attr('disabled',true);
                    }
      
                })  
            });
        	</script>  
        	<div class="gcIpt">
        		<span class="guT">数量</span>
        		<input id="min" name="" type="button" value="-"/>  
        		<input id="text_box" name="" type="text" value="1"style="width:30px; text-align: center; color: #0F0F0F;"/>  
        		<input id="add" name="" type="button" value="+"/>
        		<span class="Hgt">库存（{{.Onhand}}）</span>
        	</div>
        	<div class="nobdr-btns">
        		<a href="/cart/add/{{.Pid}}"><button class="addcart hu"><img src="../static/img/details/shop.png" width="25" height="25"/>加入购物车</button></a>
        		<button class="addcart yh"><img src="../static/img/details/ht.png" width="25" height="25"/>立即购买</button>
        	</div>
        	<div class="guarantee">
        		<span>邮费：包邮&nbsp;&nbsp;支持货到付款 <a href=""><img src="../static/img/details/me.png"/></a></span>
        	</div>
        </div>
        <!--right -->
        <div class="extInfo">
			<a href=""><img src="../static/img/details/shop.png"/></a>
            <div class="brand-logo"><img src="../static/img/logo.png"/></div>
            <div class="seller-pop-box">
            	<span class="tr">商家名称：宠物在线商城</span>
            	<span class="tr">商家等级：初级店铺</span>
            	<span class="tr hoh">
            		<a title="nanchang" href="">所在地区：江西南昌</a>
            	</span>
            </div>
            <div class="seller-phone">
            	<span class="pop im">
            		<a href="" data-name="联系卖家"><img src="../static/img/details/phon.png"/>联系卖家</a>
            	</span>
            	<span class="pop in">
            		<a href="" data-name="咨询卖家"><img src="../static/img/details/qq.png"/>咨询客服</a>
            	</span>
            	<span class="pop in">
            		<a href="" data-name="进店逛逛"><img src="../static/img/details/shop-line.png"/>进店逛逛</a>
            	</span>
            	<span class="pop in">
            		<a href="" data-name="关注店铺"><img src="../static/img/details/staar.png"/>关注店铺</a>
            	</span>
            </div>
            <div class="suport-icons">
            	<h4>扫一扫<br/>手机下单更优惠</h4>
            	<img src="../static/img/erweimas.png"/>
            </div>
        </div>
    </div>
          
    <!-- 商品介紹 -->                
    <div class="gdetail">
        <!-- left -->
        <div class="aside">
            <h3>看了还看<span><img src="../static/img/details/fod.png"width="22" height="22"/>换一批</span></h3>
            <dl class="ac-mod-list">
                <dt><a href="/product/008"><img src="{{.Pimga}}"/></a></dt>
            	<dd>{{.Pnamea}}<span>￥{{.Opricea}}</span></dd>
            </dl>
            <dl class="ac-mod-list">
            	<dt><a href="/product/018"><img src="{{.Pimgb}}"/></a></dt>
            	<dd>{{.Pnameb}}<span>￥{{.Opriceb}}</span></dd>
            </dl>
        </div>
        <!-- right -->
        <script>
            var detail = document.querySelector('.detail');
            var origOffsetY = detail.offsetTop;
            function onScroll(e) {
                window.scrollY >= origOffsetY ? detail.classList.add('sticky') :
                detail.classList.remove('sticky');
            }
            document.addEventListener('scroll', onScroll); 
		</script>
		
        <div class="detail">
            <div class="active_tab" id="outer">
				<ul class="act_title_left" id="tab">
					<li class="act_active"><a href="#">规格参数</a></li>
					<li><a href="#">商品介绍</a></li>
					<li><a href="#">商品评价</a></li>
					<li><a href="#">售后保障</a></li>
				</ul>
				<ul class="act_title_right">
					<li class="mui-ac">
						<a href="/cart"><img src="../static/img/details/shop1.png"/>我的购物车</a>
					</li>
				</ul>
			    <div class="clear"></div>
		    </div>
		    <div id="content" class="active_list"> 
		    	<!--0-->
		    	<div id="ui-a" class="ui-a">
		    	    <ul style="display:block;">
		    	        <li>商品名称：活体宠物猫狗 宠物猫狗活体宠物狗宠物猫幼体 宠物AA级</li>
		    	        <li>商品编号：ECS001983</li>
		    			<li>上架时间：2019-05-10</li>
		    			<li>产地:英国</li>
		    	        <li>库存： {{.Onhand}}</li>
		    	        <li><img src="{{.Pimg4}}"/></li>
		    	    </ul>
		    	</div>
		    	<!--商品规格-->
		    	<div id="bit" class="bit">
		    	    <ul style="display:none;">
		    	        <span><img src="{{.Pimg3}}"/></span>
		    	    </ul>
		    	</div>   
		    	<!--商品评价-->      
		    	<div id="ui-c" class="ui-c">
		    	    <ul style="display:none;">
					<span><img src="{{.Pimg2}}"/></span>
		    	    </ul>
		    	</div>
		    	<!--售后保障-->
		    	<div id="uic" class="uic">
		    	    <ul style="display:none;">
		    	        <p>商品名称：活体宠物猫狗 宠物猫狗活体宠物狗宠物猫幼体 宠物AA级</p>
		    	        <p>商品编号：ECS001983</p>
		    	        <p>产地:英国</p>
		    	        <p>上架时间：2019-05-10</p>
		    	        <span><img src="{{.Pimg1}}"/></span>
		    	    </ul>
		    	</div>
		    </div>
		    <script>
		    	$(function(){
		    		window.onload = function()
		    		{
		    			var $li = $('#tab li');
		    			var $ul = $('#content ul');
		    						
		    			$li.mouseover(function(){
		    				var $this = $(this);
		    				var $t = $this.index();
		    				$li.removeClass();
		    				$this.addClass('act_active');
		    				$ul.css('display','none');
		    				$ul.eq($t).css('display','block');
		    			})
		    		}
		    	});  
		    </script>
	    </div>
	</div>

	<!-- last -->
	<div style="margin-top: 600px;text-align: center;font-size: 12px;">
        <div style="margin:0 auto;width:700px;height: 35px;">
        <span style="float: left;word-spacing:0.52rem;text-align: center;width: 100%;">关于我们 | 联系我们 | 商家入驻 | 友情链接 | 站点地图 | 宠物商城 | 销售联盟 | 商城社区 | 企业文化 | 帮助 </span>
        </div>
    </div>
    <div>
        <div>
            <div style="margin:0 auto;width:700px;height: 30px;">
                <p style="color:#999;font-size: 12px;float: left; line-height: 10px;">© &nbsp;2007-2018 宠物网 版权所有 ，并保留所有权利  <span  style="margin-left: 30px;">宠物 Tel ：4008125181 </span><span  style="margin-left: 20px;">E-mai：wilh@163.com</span></p>
            </div>
        </div>
	</div>
	<!-- last end -->
</body>
</html>
