(this.webpackJsonplordkhanh=this.webpackJsonplordkhanh||[]).push([[0],{246:function(e,t){},317:function(e){e.exports=JSON.parse('{"ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2":"Cosmos Hub","ibc/1480B8FD20AD5FCAE81EA87584D269547DD4D436843C1D20F15E00EB64743EF4":"Akash","ibc/9712DBB13B9631EDFA9BF61B55F1B2D290B2ADB67E3A4EB3A875F3B6081B3B84":"Sentinel","ibc/7C4D60AA95E5A7558B0A364860979CA34B7FF8AAF255B87AF9E879374470CEC0":"IRISnet","ibc/E6931F78057F7CC5DA0FD6CEF82FF39373A6E0452BF1FD76910B93292CF356C1":"Crypto.org","ibc/A0CC0CF735BFB30E730C70019D4218A1244FF383503FF7579C9201AB93CA9293":"Persistence","ibc/1DCC8A6CB5689018431323953344A9F6CC4D0BFB261E88C9F7777372C10CD076":"Regen Network","ibc/52B1AA623B34EB78FD767CEA69E8D7FA6C9CFE1FBF49C5406268FD325E2CC2AC":"Starname","ibc/1DC495FCEFDA068A3820F903EDBD78B942FBD204D7E93D3BA2B432E9669D1A59":"e-Money-NGM","ibc/5973C068568365FFF40DEDCF1A1CB7582B6116B731CD31A12231AE25E20B871F":"e-Money-EEUR","ibc/46B44899322F3CD854D2D46DEEF881958467CDD4B3B10086DA49296BBED94BED":"Juno","ibc/9989AD6CCA39D1131523DB0617B50F6442081162294B4795E26746292467B525":"LikeCoin","ibc/F3FF7A84A73B62921538642F9797C423D2B4C4ACB3C7FCFFCE7F12AA69909C4B":"IXO","ibc/BE1BB42D4BE3C30D50B68D7C41DB4DFCE9678E8EF8C539F6E6A9345048894FCC":"terrausd","ibc/0EF15DF2F02480ADE0BB6E85D9EBB5DAEA2836D3860E9F97F9AADE4F57A31AA0":"terra-luna","ibc/D805F1DA50D31B96E4282C1D4181EDDFB1A44A598BFF5666F4B43E4B8BEA95A5":"BitCana","ibc/B547DC9B897E7C3AA5B824696110B8E3D2C31E3ED3F02FF363DCBAD82457E07E":"Ki","ibc/0954E1C28EB7AF5B72D24F3BC2B47BBB2FDF91BDDFD57B74B99E133AED40972A":"Secret Network","ibc/3BCCC93AD5DF58D11A6F8A05FA8BC801CBA0BA61A981F57E91B8B598BF8061CB":"MediBloc"}')},337:function(e,t,n){},338:function(e,t,n){},349:function(e,t){},351:function(e,t){},361:function(e,t){},363:function(e,t){},413:function(e,t){},414:function(e,t){},419:function(e,t){},421:function(e,t){},428:function(e,t){},447:function(e,t){},604:function(e,t,n){},614:function(e,t,n){"use strict";n.r(t);var r=n(0),o=n.n(r),c=n(33),a=n.n(c),i=(n(337),n(17)),s=n(14),d=n.n(s),l=n(39),u=n(30),b=(n(338),n(622)),f=n(624),p=n(162),j=function(){var e=Object(l.a)(d.a.mark((function e(){var t,n,r,o=arguments;return d.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:if(t=o.length>0&&void 0!==o[0]?o[0]:"osmosis-1",window.getOfflineSigner&&window.keplr){e.next=5;break}alert("Keplr Wallet not detected, please install extension"),e.next=12;break;case 5:return e.next=7,window.keplr.enable(t);case 7:return n=window.keplr.getOfflineSigner(t),e.next=10,n.getAccounts();case 10:return r=e.sent,e.abrupt("return",{accounts:r,offlineSigner:n});case 12:case"end":return e.stop()}}),e)})));return function(){return e.apply(this,arguments)}}(),m=function(e,t){return new p.SigningCosmosClient("https://lcd-osmosis.keplr.app",e[0].address,t)},x=function(){var e=Object(l.a)(d.a.mark((function e(t){var n,r,o,c=arguments;return d.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return c.length>1&&void 0!==c[1]?c[1]:1e6,n=c.length>2&&void 0!==c[2]?c[2]:"osmo1cptdzpwjc5zh6nm00dvetlg24rv9j3tjh7wnnz",r="Deposit",e.next=5,t.sendTokens(n,Object(p.coins)(1e6,"uosmo"),r);case 5:return o=e.sent,console.log(o),e.abrupt("return",o);case 8:case"end":return e.stop()}}),e)})));return function(t){return e.apply(this,arguments)}}(),h=n(6),g=b.a.Title,O=b.a.Paragraph,C={div:{display:"flex",flexDirection:"column",justifyContent:"center space-between",alignContent:"center",backgroundColor:"#ffc369",borderStyle:"solid",borderWidth:"20px",borderColor:"#ffb459",height:"35em",width:"30%",borderRadius:"10px",marginLeft:"50em",boxShadow:"0 4px 8px 0 rgba(0, 0, 0, 0.5), 0 6px 20px 0 rgba(0, 0, 0, 0.5)"},buttonDiv:{alignSelf:"stretch",marginTop:"5em"},button:{borderWidth:"0px",borderRadius:"10px",size:"10em",backgroundColor:"#ff9e61",color:"#ffffff",boxShadow:"0 4px 8px 0 rgba(0, 0, 0, 0.5), 0 6px 20px 0 rgba(0, 0, 0, 0.5)",width:"80%",height:"10%",padding:"4em",paddingTop:"2em"},p:{marginLeft:"1em"},content:{justifyContent:"center",alignItems:"center",marginBotom:"10px",fontSize:"30px"},addrDiv:{marginTop:"10em",display:"flex",flexDirection:"column",justifyContent:"left",alignContent:"left",marginBottom:"10em"},addrContent:{backgroundColor:"#ffffff",alignContent:"left",margin:"5em",marginTop:"1px",alignItems:"left",borderRadius:"10px",overflowWrap:"break-word",padding:"1em"}},v=function(){var e=Object(r.useState)("hello"),t=Object(u.a)(e,2),n=t[0],o=(t[1],function(){var e=Object(l.a)(d.a.mark((function e(){var t,n,r,o;return d.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,j();case 2:t=e.sent,n=t.accounts,r=t.offlineSigner,null!=(o=m(n,r))&&x(o).then((function(e){})).catch((function(e){}));case 7:case"end":return e.stop()}}),e)})));return function(){return e.apply(this,arguments)}}());return Object(h.jsxs)("div",{claasName:"container-fluid",style:C.div,children:[Object(h.jsx)("div",{style:C.buttonDiv,children:Object(h.jsx)("button",{onClick:Object(l.a)(d.a.mark((function e(){return d.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,o();case 2:return e.abrupt("return",e.sent);case 3:case"end":return e.stop()}}),e)}))),size:"large",style:C.button,onMouseEnter:function(e){e.target.style.transform="scale(1.01)"},onMouseLeave:function(e){e.target.style.transform="scale(1)"},children:Object(h.jsxs)("div",{style:C.content,children:[Object(h.jsx)(f.a,{}),Object(h.jsx)("span",{style:C.p,children:"Register"})]})})}),""!==n&&Object(h.jsxs)("div",{style:C.addrDiv,children:[Object(h.jsx)(g,{level:3,children:"Generated wallet address"}),Object(h.jsx)("div",{style:C.addrContent,children:Object(h.jsx)(O,{copyable:{text:n},children:n.length>100?"".concat(n.substring(0,100),"... "):"".concat(n," ")})})]}),Object(h.jsx)("div",{})]})},B=n(142),y=n(98),D=n.n(y),F=n(77),A=n(317),w=function(e){var t={denom:"",name:"",amount:"",logo:""};if("ibc/"===e.denom.substring(0,4)){var n=A["".concat(e.denom)];if(F["".concat(n)]){t.denom=F["".concat(n)].denom;var r=F["".concat(n)].chain_name;t.name=n+" - "+r.toUpperCase(),t.amount=(parseInt(e.amount)/1e6).toString(),t.logo=F["".concat(n)].logo}}else{t.denom=e.denom;var o="uosmo"===e.denom?"osmo":"ion";t.name=o.toUpperCase(),t.amount=(parseInt(e.amount)/1e6).toString(),t.logo="uosmo"===e.denom?"https://dl.airtable.com/.attachments/4ef30ec4008bc86cc3c0f74a6bb84050/0eeb4d64/aQ5W3zaT_400x400.jpg":"https://app.osmosis.zone/public/assets/tokens/ion.png"}return t},E=function(){var e=Object(l.a)(d.a.mark((function e(){var t,n,r,o,c,a=arguments;return d.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return t=a.length>0&&void 0!==a[0]?a[0]:"https://lcd-osmosis.keplr.app",n=a.length>1?a[1]:void 0,r="".concat(t,"/bank/balances/").concat(n),o=[],e.next=6,D.a.get(r);case 6:return(c=e.sent).data&&c.data.result&&c.data.result.map((function(e){var t=w(e);o.push(t)})),e.abrupt("return",o);case 9:case"end":return e.stop()}}),e)})));return function(){return e.apply(this,arguments)}}(),S=function(){var e=Object(l.a)(d.a.mark((function e(t){var n,r;return d.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,E(void 0,t);case 2:if(0!==(n=e.sent).length){e.next=5;break}return e.abrupt("return",0);case 5:return r=n.filter((function(e){return"uosmo"===e.denom})),e.abrupt("return",r[0].amount);case 7:case"end":return e.stop()}}),e)})));return function(t){return e.apply(this,arguments)}}(),k=n(326),T=n(621),z=n(619),R=(n(604),b.a.Title),I={divTitle:{display:"flex",justifyContent:"left",alignContent:"left",textAlign:"left"},div:{marginBottom:"8rem",borderRadius:"10px"}},W=function(e){var t=e.address,n=Object(r.useState)([]),o=Object(u.a)(n,2),c=o[0],a=o[1],i=function(){k.b.success("Fetching success",1)},s=function(){k.b.error("Fetching failed",1)};Object(r.useEffect)((function(){console.log(t),""!==t?E(void 0,t).then((function(e){e.length>0&&(a(Object(B.a)(e)),i()),k.b.warning("No Assets Yet",1)})).catch((function(){s()})):E(void 0,"osmo1cy2fkq04yh7zm6v52dm525pvx0fph7ed75lnz7").then((function(e){e.length>0&&(a(Object(B.a)(e)),i())})).catch((function(){s()}))}),[]);var d=[{dataIndex:"logo",key:"logo",fixed:"left",render:function(e){T.a}},{title:"Asset/Chain",dataIndex:"name",key:"name",fixed:"left"},{title:"Amount",dataIndex:"amount",key:"amount",fixed:"left"}];return Object(h.jsxs)("div",{style:I.div,children:[Object(h.jsx)("div",{style:I.divTitle,children:Object(h.jsx)(R,{children:"Osmosis Assets"})}),Object(h.jsx)("hr",{}),Object(h.jsx)(z.a,{dataSource:c,columns:d,style:{marginTop:"3rem",borderRadius:"5px"}})]})},N=(b.a.Title,b.a.Paragraph,n(625)),M=n(75),L=function(e){var t=e.current,n=e.wrapSetter,o=Object(r.useState)("pending"),c=Object(u.a)(o,2),a=c[0],i=c[1];Object(r.useEffect)((function(){1===t?(i("running"),s()):0===t&&i("pending")}),[t]);var s=function(){setTimeout((function(){n(2)}),3e3)};return Object(h.jsx)("div",{style:{height:"10rem",width:"20%",borderRadius:"10px",borderWidth:"30px",backgroundColor:"running"===a?"#8abf80":"#e6e6e6",paddingTop:"4.2rem",boxShadow:"0 4px 8px 0 rgba(0, 0, 0, 0.5), 0 6px 20px 0 rgba(0, 0, 0, 0.5)"},children:"Pulling Reward"})},_=n(325),P=function(e){var t=e.current,n=e.wrapSetter,o=Object(r.useState)("pending"),c=Object(u.a)(o,2),a=c[0],i=c[1];Object(r.useEffect)((function(){2===t?(i("running"),s()):0===t&&i("pending")}),[t]);var s=function(){setTimeout((function(){var e,t;e="Pool found",t="Pick Pool #1",_.a.open({message:e,description:t}),n(3)}),2e3)};return Object(h.jsx)("div",{style:{height:"10rem",width:"20%",borderRadius:"10px",borderWidth:"30px",backgroundColor:"running"===a?"#8abf80":"#e6e6e6",paddingTop:"4.2rem",boxShadow:"0 4px 8px 0 rgba(0, 0, 0, 0.5), 0 6px 20px 0 rgba(0, 0, 0, 0.5)"},children:"Identify Best Returns"})},H=function(e){var t=e.current,n=e.wrapSetter,o=Object(r.useState)("pending"),c=Object(u.a)(o,2),a=c[0],i=c[1];Object(r.useEffect)((function(){3===t?(i("running"),s()):0===t&&i("pending")}),[t]);var s=function(){setTimeout((function(){var e,t;e="Auto Invest",t="Success",_.a.open({message:e,description:t}),n(4)}),3e3)};return Object(h.jsx)("div",{style:{height:"10rem",width:"20%",borderRadius:"10px",borderWidth:"30px",backgroundColor:"running"===a?"#8abf80":"#e6e6e6",paddingTop:"4.2rem",boxShadow:"0 4px 8px 0 rgba(0, 0, 0, 0.5), 0 6px 20px 0 rgba(0, 0, 0, 0.5)"},children:"Auto Invest"})},J=function(e){var t=e.current,n=e.wrapSetter,o=Object(r.useState)("pending"),c=Object(u.a)(o,2),a=c[0],i=c[1];Object(r.useEffect)((function(){4===t?(i("running"),s()):0===t&&i("pending")}),[t]);var s=function(){setTimeout((function(){n(0)}),3e3)};return Object(h.jsx)("div",{style:{height:"10rem",width:"20%",borderRadius:"10px",borderWidth:"30px",backgroundColor:"running"===a?"#8abf80":"#e6e6e6",paddingTop:"4.2rem",boxShadow:"0 4px 8px 0 rgba(0, 0, 0, 0.5), 0 6px 20px 0 rgba(0, 0, 0, 0.5)"},children:"New Balances"})},U=b.a.Title,q={flexDiv:{display:"flex",flexDirection:"row",justifyContent:"space-between"},divTitle:{display:"flex",justifyContent:"left",alignContent:"left",textAlign:"left"},div:{display:"flex",flexDirection:"column",justifyContent:"space-between",alignContent:"center"},button:{borderWidth:"0px",borderRadius:"10px",backgroundColor:"#8abf80",color:"#ffffff",boxShadow:"0 4px 8px 0 rgba(0, 0, 0, 0.5), 0 6px 20px 0 rgba(0, 0, 0, 0.5)",width:"30%",height:"5rem",padding:"2rem",paddingTop:"1rem",marginTop:"10rem"},buttonText:{fontSize:"30px",display:"flex",flexDirection:"row",justifyContent:"center"}},G=function(){var e=Object(r.useState)(0),t=Object(u.a)(e,2),n=t[0],o=t[1],c=Object(r.useCallback)((function(e){o(e)}),[o]);return Object(h.jsxs)("div",{style:q.div,children:[Object(h.jsx)("div",{style:q.divTitle,children:Object(h.jsx)(U,{children:"BeanStalk"})}),Object(h.jsx)("hr",{}),Object(h.jsxs)("div",{style:Object(i.a)(Object(i.a)({},q.flexDiv),{},{marginTop:"2rem"}),children:[Object(h.jsx)(L,{current:n,wrapSetter:c}),Object(h.jsx)(N.a,{style:{fontSize:"10rem",color:"#7a7a7a"}}),Object(h.jsx)(P,{current:n,wrapSetter:c}),Object(h.jsx)(N.a,{style:{fontSize:"10rem",color:"#7a7a7a"}}),Object(h.jsx)(H,{current:n,wrapSetter:c}),Object(h.jsx)(N.a,{style:{fontSize:"10rem",color:"#7a7a7a"}}),Object(h.jsx)(J,{current:n,wrapSetter:c})]}),Object(h.jsx)("div",{children:Object(h.jsx)("button",{style:q.button,onClick:function(){o(1)},onMouseEnter:function(e){e.target.style.transform="scale(1.01)"},onMouseLeave:function(e){e.target.style.transform="scale(1)"},children:Object(h.jsxs)("div",{style:q.buttonText,children:[Object(h.jsx)("p",{children:"Run Manually"})," ",Object(h.jsx)(M.a,{style:{marginTop:"0.6rem"}})]})})})]})},K=n(626),Q={borderWidth:"0px",borderRadius:"5px",size:"10em",backgroundColor:"#9c8bad",color:"#ffffff",width:"80%",height:"3rem"},X=function(e){var t=e.collapsed,n=e.wrapSetter;return Object(h.jsxs)("div",{style:{marginTop:"34rem",marginBottom:"0.3rem"},children:[Object(h.jsx)("hr",{}),Object(h.jsx)("button",{style:Object(i.a)(Object(i.a)({},Q),{},{fontSize:t?"10px":"15px"}),onClick:function(){n(!0)},children:t?Object(h.jsx)(K.a,{style:{fontSize:"1.5rem"}}):Object(h.jsxs)("div",{children:["Deposit ",Object(h.jsx)(K.a,{style:{fontSize:"1.5rem"}})]})})]})},V=n(623),Y=n(627),Z={transfer:{display:"flex",flexDirection:"row",justifyContent:"space-between",marginBottom:"2rem"},transferInfo:{padding:"10px",border:"2px solid #c4c4c4",borderRadius:"10px",width:"12rem"},container:{display:"flex",flexDirection:"column",justifyContent:"space-between"},form:{display:"flex",flexDirection:"column",justifyContent:"center"},button:{display:"flex",flexDirection:"column",justifyContent:"center",marginTop:"2rem",marginBottom:"2rem"}},$=function(e){var t=e.account,n=e.wrapSetter,r=e.cookieMonster,o=function(){var e=Object(l.a)(d.a.mark((function e(){var t,o,c;return d.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,j();case 2:t=e.sent,o=t.accounts,c=t.offlineSigner,null!=m(o,c)&&x(r).then((function(e){console.log(e),n(!1),k.b.success("Deposit success",1)})).catch((function(){k.b.error("Deposit failed",1),n(!1)}));case 7:case"end":return e.stop()}}),e)})));return function(){return e.apply(this,arguments)}}();return Object(h.jsxs)("div",{children:[Object(h.jsxs)("div",{style:Z.transfer,children:[Object(h.jsxs)("div",{style:Z.transferInfo,children:[Object(h.jsx)("p",{children:"From"}),Object(h.jsx)("p",{children:t.address})]}),Object(h.jsx)(Y.a,{style:{fontSize:"2rem",marginTop:"15px"}}),Object(h.jsxs)("div",{style:Z.transferInfo,children:[Object(h.jsx)("p",{children:"To"}),Object(h.jsx)("p",{children:r})]})]}),Object(h.jsxs)("div",{style:Z.form,children:[Object(h.jsx)("div",{style:{marginBottom:"1rem"},children:"Amount To Deposit"}),Object(h.jsx)(V.a,{style:{width:"100%",height:"60px",borderRadius:"10px",border:"2px solid #c4c4c4",fontSize:"2rem"},min:0,size:"large"})]}),Object(h.jsx)("div",{style:Z.button,children:Object(h.jsx)("button",{onClick:o,style:{borderRadius:"10px",height:"4rem",fontSize:"1.5rem",backgroundColor:"#9b8da6",color:"#ffffff"},children:"Deposit"})})]})},ee=(n(610),n(618)),te=n(100),ne=n(628),re=n(629),oe=n(620),ce=n(184),ae=n(26),ie=n.p+"static/media/logo.d4e302ed.png",se="http://192.168.1.29:8000",de=function(){var e=Object(l.a)(d.a.mark((function e(t){var n;return d.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,D.a.post("".concat(se,"/check-account"),{address:t});case 2:return n=e.sent,e.abrupt("return",n);case 4:case"end":return e.stop()}}),e)})));return function(t){return e.apply(this,arguments)}}(),le=(ee.a.Header,ee.a.Content),ue=(ee.a.Footer,ee.a.Sider),be={borderWidth:"0px",borderRadius:"5px",size:"10em",backgroundColor:"#8abf80",color:"#ffffff",width:"80%",height:"3rem"};var fe=function(){var e=Object(r.useState)({address:"",amount:""}),t=Object(u.a)(e,2),n=t[0],o=t[1],c=Object(r.useState)(""),a=Object(u.a)(c,2),s=a[0],b=a[1],p=Object(r.useState)(!1),m=Object(u.a)(p,2),x=m[0],g=m[1],O=Object(r.useState)(!1),C=Object(u.a)(O,2),B=C[0],y=C[1],D=Object(r.useState)(!1),F=Object(u.a)(D,2),A=F[0],w=F[1],E=Object(r.useCallback)((function(e){w(e)}),[w]),z=function(){var e=Object(l.a)(d.a.mark((function e(){var t,n,r;return d.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,j();case 2:return t=e.sent,n=t.accounts,t.offlineSigner,e.next=7,S(n[0].address);case 7:r=e.sent,o({address:n[0].address,amount:(parseInt(r)/1e6).toString()}),de("osmo1cy2fkq04yh7zm6v52dm525pvx0fph7ed75lnz7").then((function(e){""!==e.data.Address&&(k.b.success("Connect",1),b(e.data.Address))})).catch((function(){k.b.error("Connect failed",1)}));case 10:case"end":return e.stop()}}),e)})));return function(){return e.apply(this,arguments)}}();return Object(h.jsxs)("div",{className:"App container-fluid",children:[Object(h.jsx)(ee.a,{style:{minHeight:"100vh",marginLeft:"-12px",marginRight:"-12px"},children:Object(h.jsxs)(ce.a,{children:[Object(h.jsxs)(ue,{theme:"light",collapsible:!0,collapsed:x,onCollapse:function(){g(!x),setTimeout((function(){y(!B)}),100)},width:256,style:{backgroundColor:"#ffffff"},children:[Object(h.jsx)("div",{className:"logo",style:{marginRight:"0.1rem",marginTop:"1rem",marginBottom:"1rem"},children:Object(h.jsx)(T.a,{width:B?50:100,src:ie,preview:!1})}),Object(h.jsx)("hr",{}),Object(h.jsxs)(te.a,{theme:"light",style:{backgroundColor:"#ffffff"},mode:"inline",children:[Object(h.jsxs)(te.a.Item,{icon:Object(h.jsx)(ne.a,{style:{marginLeft:x?"-0.3rem":"1.5rem",fontSize:"1rem"}}),style:{margin:0,marginTop:"10px",fontSize:"1.3rem",color:"#2b2b2b",fontWeight:300},className:"modified-item",children:["Home",Object(h.jsx)(ce.b,{to:"/"})]},"home"),Object(h.jsxs)(te.a.Item,{icon:Object(h.jsx)(f.a,{style:{marginLeft:x?"-0.3rem":"1.5rem",fontSize:"1rem"}}),style:{margin:0,marginTop:"10px",fontSize:"1.3rem",color:"#2b2b2b",fontWeight:300},className:"modified-item",children:["Asset",Object(h.jsx)(ce.b,{to:"/asset"})]},"asset")]}),""===s?Object(h.jsxs)("div",{style:{marginTop:"34rem",marginBottom:"0.3rem"},children:[Object(h.jsx)("hr",{}),Object(h.jsx)("button",{style:Object(i.a)(Object(i.a)({},be),{},{fontSize:x?"10px":"15px"}),onClick:Object(l.a)(d.a.mark((function e(){return d.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,z();case 2:case"end":return e.stop()}}),e)}))),children:x?Object(h.jsx)(re.a,{style:{fontSize:"1.5rem"}}):"Connect To BeanStalk"})]}):Object(h.jsx)(X,{collapsed:x,wrapSetter:E})]}),Object(h.jsx)(ee.a,{className:"site-layout",style:{backgroundColor:"#c5e6be"},children:Object(h.jsx)(le,{style:{margin:"2rem"},children:Object(h.jsx)("div",{className:"site-layout-background",style:{padding:24,paddingTop:"2rem",paddingBottom:"17rem",minHeight:360,marginTop:"10px"},children:Object(h.jsxs)(ae.c,{children:[Object(h.jsx)(ae.a,{exact:!0,path:"/",element:Object(h.jsx)(G,{})}),Object(h.jsx)(ae.a,{exact:!0,path:"/asset",element:Object(h.jsx)(W,{address:s})}),Object(h.jsx)(ae.a,{exact:!0,path:"/account",element:Object(h.jsx)(v,{})})]})})})})]})}),Object(h.jsx)(h.Fragment,{children:Object(h.jsxs)(oe.a,{show:A,onHide:function(){w(!1)},children:[Object(h.jsx)(oe.a.Header,{closeButton:!0,children:Object(h.jsx)(oe.a.Title,{children:"Deposit"})}),Object(h.jsx)(oe.a.Body,{children:Object(h.jsx)($,{cookieMonster:s,account:n,wrapSetter:E})})]})})]})},pe=function(e){e&&e instanceof Function&&n.e(3).then(n.bind(null,630)).then((function(t){var n=t.getCLS,r=t.getFID,o=t.getFCP,c=t.getLCP,a=t.getTTFB;n(e),r(e),o(e),c(e),a(e)}))};n(613);a.a.render(Object(h.jsx)(o.a.StrictMode,{children:Object(h.jsx)(fe,{})}),document.getElementById("root")),pe()},77:function(e){e.exports=JSON.parse('{"Osmosis":{"chain_id":"osmosis-1","coingecko":"osmosis","apiUrl":"https://lcd-osmosis.keplr.app","logo":"https://dl.airtable.com/.attachments/4ef30ec4008bc86cc3c0f74a6bb84050/0eeb4d64/aQ5W3zaT_400x400.jpg","denom":"uosmo","chain_name":"osmosis"},"Cosmos Hub":{"chain_id":"cosmoshub-4","coingecko":"cosmos","apiUrl":"https://cosmos.api.ping.pub","logo":"https://dl.airtable.com/.attachments/e54f814bba8c0f9af8a3056020210de0/2d1155fb/cosmos-hub.svg","denom":"uatom","chain_name":"atom"},"Juno":{"chain_id":"juno-1","chain_name":"juno","coingecko":"juno-network","apiUrl":"https://juno.api.ping.pub","logo":"https://dl.airtable.com/.attachments/0f66137c6fb2868000d5a1e214c9ae3d/75a9c5bc/S3c2V3Xd_400x400.jpg","denom":"ujuno"}}')}},[[614,1,2]]]);
//# sourceMappingURL=main.0f734cf3.chunk.js.map