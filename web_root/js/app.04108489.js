(function(t){function a(a){for(var n,r,i=a[0],c=a[1],l=a[2],f=0,p=[];f<i.length;f++)r=i[f],s[r]&&p.push(s[r][0]),s[r]=0;for(n in c)Object.prototype.hasOwnProperty.call(c,n)&&(t[n]=c[n]);u&&u(a);while(p.length)p.shift()();return o.push.apply(o,l||[]),e()}function e(){for(var t,a=0;a<o.length;a++){for(var e=o[a],n=!0,i=1;i<e.length;i++){var c=e[i];0!==s[c]&&(n=!1)}n&&(o.splice(a--,1),t=r(r.s=e[0]))}return t}var n={},s={app:0},o=[];function r(a){if(n[a])return n[a].exports;var e=n[a]={i:a,l:!1,exports:{}};return t[a].call(e.exports,e,e.exports,r),e.l=!0,e.exports}r.m=t,r.c=n,r.d=function(t,a,e){r.o(t,a)||Object.defineProperty(t,a,{enumerable:!0,get:e})},r.r=function(t){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},r.t=function(t,a){if(1&a&&(t=r(t)),8&a)return t;if(4&a&&"object"===typeof t&&t&&t.__esModule)return t;var e=Object.create(null);if(r.r(e),Object.defineProperty(e,"default",{enumerable:!0,value:t}),2&a&&"string"!=typeof t)for(var n in t)r.d(e,n,function(a){return t[a]}.bind(null,n));return e},r.n=function(t){var a=t&&t.__esModule?function(){return t["default"]}:function(){return t};return r.d(a,"a",a),a},r.o=function(t,a){return Object.prototype.hasOwnProperty.call(t,a)},r.p="/";var i=window["webpackJsonp"]=window["webpackJsonp"]||[],c=i.push.bind(i);i.push=a,i=i.slice();for(var l=0;l<i.length;l++)a(i[l]);var u=c;o.push([0,"chunk-vendors"]),e()})({0:function(t,a,e){t.exports=e("56d7")},"00d4":function(t,a,e){},"034f":function(t,a,e){"use strict";var n=e("64a9"),s=e.n(n);s.a},"0803":function(t,a,e){t.exports=e.p+"img/avatar.5032714e.jpg"},"3f1b":function(t,a,e){"use strict";var n=e("ee86"),s=e.n(n);s.a},"4f9f":function(t,a,e){},"52f7":function(t,a,e){"use strict";var n=e("4f9f"),s=e.n(n);s.a},"56d7":function(t,a,e){"use strict";e.r(a);e("cadf"),e("551c"),e("097d");var n=e("2b0e"),s=function(){var t=this,a=t.$createElement,e=t._self._c||a;return e("div",{attrs:{id:"app"}},[e("div",{staticClass:"header"},[e("NavBar")],1),e("div",{staticClass:"main container"},[e("router-view")],1)])},o=[],r=function(){var t=this,a=t.$createElement,e=t._self._c||a;return e("div",{staticClass:"navbar-container"},[e("b-navbar",{staticClass:"navbar-custom",attrs:{toggleable:"lg",type:"light",variant:"light"}},[e("b-navbar-brand",{attrs:{to:"/"}},[t._v("SEUNGJUN OH")]),e("b-navbar-toggle",{attrs:{target:"nav_collapse"}}),e("b-collapse",{attrs:{"is-nav":"",id:"nav_collapse"}},[e("b-navbar-nav",{staticClass:"ml-auto"},[e("b-nav-item",{attrs:{to:"/blog"}},[t._v("Blog")]),e("b-nav-item",{attrs:{to:"/about"}},[t._v("About")]),e("b-nav-form",{staticClass:"navbar-search"},[e("b-form-input",{staticClass:"mr-sm-2",attrs:{size:"sm",type:"text",placeholder:"Not supported",disabled:""}}),e("b-button",{staticClass:"my-2 my-sm-0",attrs:{size:"sm",type:"submit",disabled:""}},[t._v("Search")])],1),e("br")],1)],1)],1),t._m(0)],1)},i=[function(){var t=this,a=t.$createElement,n=t._self._c||a;return n("div",{staticClass:"avatar-container"},[n("div",{staticClass:"avatar-img-border"},[n("a",{attrs:{title:"Seungjun Oh",to:"/"}},[n("img",{staticClass:"avatar-img",attrs:{src:e("0803"),alt:"Seungjun Oh"}})])])])}],c={name:"NavBar",props:{}},l=c,u=(e("876c"),e("2877")),f=Object(u["a"])(l,r,i,!1,null,"2da0350c",null);f.options.__file="NavBar.vue";var p=f.exports,v={name:"app",components:{NavBar:p}},d=v,m=(e("034f"),Object(u["a"])(d,s,o,!1,null,null,null));m.options.__file="App.vue";var h=m.exports,b=e("8c4f"),_=function(){var t=this,a=t.$createElement,e=t._self._c||a;return e("div",{staticClass:"about"},[e("img"),e("i",{staticClass:"fab fa-linkedin"}),t._m(0),e("div",{staticClass:"profile"},[e("div",{staticClass:"record"},[e("span",{staticClass:"icon"},[e("font-awesome-icon",{staticClass:"fa-fw",attrs:{icon:"laptop"}})],1),e("span",{staticClass:"text"},[t._v("Software Engineer")])]),e("div",{staticClass:"record"},[e("span",{staticClass:"icon"},[e("font-awesome-icon",{staticClass:"fa-fw",attrs:{icon:"user-graduate"}})],1),e("span",{staticClass:"text"},[t._v("Hanyang Univ.")])]),e("div",{staticClass:"record"},[e("span",{staticClass:"icon"},[e("font-awesome-icon",{staticClass:"fa-fw",attrs:{icon:"home"}})],1),e("span",{staticClass:"text"},[t._v("Seoul, South Korea")])]),e("div",{staticClass:"footer"},[e("span",{staticClass:"icon"},[e("a",{attrs:{title:"resume",target:"_blank"}},[e("font-awesome-layers",{staticClass:"fa-3x"},[e("a",{attrs:{href:"https://www.linkedin.com/in/seungjun-oh",title:"linkedin",target:"_blank"}},[e("font-awesome-icon",{attrs:{icon:"circle"}}),e("font-awesome-icon",{staticClass:"fa-inverse",attrs:{icon:["fab","linkedin"],transform:"shrink-7"}}),e("i",{staticClass:"fab fa-linkedin"})],1)])],1)]),e("span",{staticClass:"icon"},[e("a",{attrs:{href:"https://github.com/sjnov11",title:"github",target:"_blank"}},[e("font-awesome-layers",{staticClass:"fa-3x"},[e("font-awesome-icon",{attrs:{icon:"circle"}}),e("font-awesome-icon",{staticClass:"fa-inverse",attrs:{icon:["fab","github"],transform:"shrink-7"}})],1)],1)]),e("span",{staticClass:"icon"},[e("a",{attrs:{href:"mailto:sjnov11@gmail.com",title:"mail"}},[e("font-awesome-layers",{staticClass:"fa-3x"},[e("font-awesome-icon",{attrs:{icon:"circle"}}),e("font-awesome-icon",{staticClass:"fa-inverse",attrs:{icon:"envelope",transform:"shrink-7"}})],1)],1)])])])])},g=[function(){var t=this,a=t.$createElement,e=t._self._c||a;return e("div",{staticClass:"title-container"},[e("h1",[t._v("Seung-jun Oh")]),e("hr",{staticClass:"hr-small"})])}],C={name:"About"},w=C,y=(e("742a"),Object(u["a"])(w,_,g,!1,null,"75f6e208",null));y.options.__file="About.vue";var x=y.exports,j=function(){var t=this,a=t.$createElement,e=t._self._c||a;return e("div",[this.$route.params.post?t._e():e("div",{staticClass:"post-list"},t._l(t.postMetaData.index,function(a){return e("ul",{key:a.title},[e("router-link",{staticClass:"post-title header",attrs:{to:a.uri}},[t._v(t._s(a.title))]),e("p",[t._v(t._s(a.date))]),e("div",{staticClass:"post-tags"},t._l(a.tags,function(a){return e("router-link",{key:a,staticClass:"post-tag",attrs:{to:"/"}},[t._v(t._s(a))])}))],1)})),e("router-view",{attrs:{"post-meta":t.postMetaData}})],1)},O=[],k={name:"Blog",created:function(){this.getPostMetaData()},data:function(){return{postMetaData:{}}},methods:{getPostMetaData:function(){var t=this,a="/blog/list.json";fetch(a).then(function(t){return t.json()}).then(function(a){return t.postMetaData=a})}}},P=k,S=(e("ea38"),e("3f1b"),Object(u["a"])(P,j,O,!1,null,"f36a0fee",null));S.options.__file="Blog.vue";var M=S.exports,$=function(){var t=this,a=t.$createElement,e=t._self._c||a;return e("div",[e("div",[t._v(t._s(t.post-t.meta))]),e("div",{staticClass:"post",domProps:{innerHTML:t._s(t.post)}}),t.loading?e("div",{staticClass:"loading"},[t._v("\n    Loading...\n  ")]):t._e(),t.error?e("div",{staticClass:"error"},[t._v("\n    There is some problem. "+t._s(t.error)+"\n  ")]):t._e()])},E=[],T=function(){var t=this,a=t.$createElement,e=t._self._c||a;return e("div",{staticClass:"title-container"},[e("h1",[t._v(t._s(t.title))]),e("hr",{staticClass:"hr-small"})])},B=[],N={name:"PageTitle",data:function(){return{title:"SeungJun Oh"}}},D=N,H=(e("b684"),Object(u["a"])(D,T,B,!1,null,"ef893aea",null));H.options.__file="Title.vue";var J=H.exports,A={name:"Post",props:["post-meta"],components:J,data:function(){return{loading:!1,post:"",error:null}},created:function(){this.fetchPost()},watch:{$route:function(){this.fetchPost()}},methods:{fetchPost:function(){var t=this;this.error=this.post=null,this.loading=!0;var a="/blog/"+this.$route.params.post+".html";fetch(a).then(function(t){return t.text()}).then(function(a){t.post=a,t.loading=null,t.$nextTick().then(function(){window.MathJax.Hub.Queue(["Typeset",window.MathJax.Hub])})}).catch(function(a){t.error=a})}}},U=A,z=(e("52f7"),Object(u["a"])(U,$,E,!1,null,"278d0fa0",null));z.options.__file="Post.vue";var L=z.exports;n["a"].use(b["a"]);var G=new b["a"]({routes:[{path:"/",name:"Home",component:x},{path:"/about",name:"About",component:x},{path:"/blog",name:"Blog",component:M,children:[{path:":post",component:L}]}]}),K=e("9f7b"),Q=(e("f9e3"),e("2dd8"),e("ecee")),q=e("c074"),F=e("f2d1"),I=e("ad3d");n["a"].use(K["a"]),n["a"].config.productionTip=!1,Q["c"].add(q["b"],q["f"],q["a"],q["g"],q["e"],q["d"],q["c"]),Q["c"].add(F["a"],F["c"],F["b"]),n["a"].component("font-awesome-icon",I["a"]),n["a"].component("font-awesome-layers",I["b"]),n["a"].config.productionTip=!1,new n["a"]({router:G,render:function(t){return t(h)}}).$mount("#app")},"64a9":function(t,a,e){},"742a":function(t,a,e){"use strict";var n=e("00d4"),s=e.n(n);s.a},"876c":function(t,a,e){"use strict";var n=e("b6d8"),s=e.n(n);s.a},b3d4:function(t,a,e){},b684:function(t,a,e){"use strict";var n=e("ddce"),s=e.n(n);s.a},b6d8:function(t,a,e){},ddce:function(t,a,e){},ea38:function(t,a,e){"use strict";var n=e("b3d4"),s=e.n(n);s.a},ee86:function(t,a,e){}});
//# sourceMappingURL=app.04108489.js.map