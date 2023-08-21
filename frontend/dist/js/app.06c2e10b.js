(function(){"use strict";var e={5064:function(e,t,o){var a=o(9242),n=o(3396);function r(e,t,o,a,r,s){const i=(0,n.up)("NavBar"),l=(0,n.up)("router-view");return(0,n.wg)(),(0,n.iD)(n.HY,null,[(0,n.Wm)(i),(0,n.Wm)(l)],64)}const s={class:"navbar navbar-expand-lg fixed-top bg-body-tertiary bg-primary","data-bs-theme":"dark"},i={class:"container-fluid"},l=(0,n._)("button",{class:"navbar-toggler",type:"button","data-bs-toggle":"collapse","data-bs-target":"#navbarNavDropdown","aria-controls":"navbarNavDropdown","aria-expanded":"false","aria-label":"Toggle navigation"},[(0,n._)("span",{class:"navbar-toggler-icon"})],-1),c={class:"collapse navbar-collapse",id:"navbarNavDropdown"},u={class:"navbar-nav"},d={class:"nav-item"},p={class:"nav-item"},m=(0,n._)("li",{class:"nav-item dropdown"},[(0,n._)("a",{class:"nav-link dropdown-toggle",href:"#",role:"button","data-bs-toggle":"dropdown","aria-expanded":"false"}," 其他 "),(0,n._)("ul",{class:"dropdown-menu"},[(0,n._)("li",null,[(0,n._)("a",{class:"dropdown-item",href:"https://cv.pjmcode.top/"},[(0,n.Uk)("开发者信息 "),(0,n._)("svg",{preserveAspectRatio:"xMidYMid meet",viewBox:"0 0 24 24",width:"1.2em",height:"1.2em",class:"link-icon"},[(0,n._)("path",{fill:"currentColor",d:"M10 6v2H5v11h11v-5h2v6a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V7a1 1 0 0 1 1-1h6zm11-3v8h-2V6.413l-7.793 7.794l-1.414-1.414L17.585 5H13V3h8z"})])])]),(0,n._)("li",null,[(0,n._)("a",{class:"dropdown-item",href:"https://github.com/pjimming/baize"},[(0,n.Uk)("开源代码 "),(0,n._)("svg",{preserveAspectRatio:"xMidYMid meet",viewBox:"0 0 24 24",width:"1.2em",height:"1.2em",class:"link-icon"},[(0,n._)("path",{fill:"currentColor",d:"M10 6v2H5v11h11v-5h2v6a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V7a1 1 0 0 1 1-1h6zm11-3v8h-2V6.413l-7.793 7.794l-1.414-1.414L17.585 5H13V3h8z"})])])]),(0,n._)("li",null,[(0,n._)("a",{class:"dropdown-item",href:"#"},"赞助")])])],-1);function f(e,t,o,a,r,f){const g=(0,n.up)("router-link");return(0,n.wg)(),(0,n.iD)("nav",s,[(0,n._)("div",i,[(0,n.Wm)(g,{class:"navbar-brand",to:{name:"home"}},{default:(0,n.w5)((()=>[(0,n.Uk)("白泽")])),_:1}),l,(0,n._)("div",c,[(0,n._)("ul",u,[(0,n._)("li",d,[(0,n.Wm)(g,{class:"nav-link","aria-current":"page",to:{name:"home"}},{default:(0,n.w5)((()=>[(0,n.Uk)("首页")])),_:1})]),(0,n._)("li",p,[(0,n.Wm)(g,{class:"nav-link",to:{name:"about"}},{default:(0,n.w5)((()=>[(0,n.Uk)("关于白泽")])),_:1})]),m])])])])}var g={name:"NavBar"},h=o(89);const v=(0,h.Z)(g,[["render",f]]);var b=v,w={name:"app",components:{NavBar:b}};const k=(0,h.Z)(w,[["render",r]]);var _=k,j=o(2483);const P=(0,n._)("h3",null,"🛠golang项目包依赖可视化工具😼",-1),I=(0,n._)("hr",null,null,-1);function y(e,t,o,a,r,s){const i=(0,n.up)("InputVue"),l=(0,n.up)("ProjectInfo"),c=(0,n.up)("ListInfo"),u=(0,n.up)("BackToTop"),d=(0,n.up)("ContentBase");return(0,n.wg)(),(0,n.j4)(d,null,{default:(0,n.w5)((()=>[P,I,(0,n.Wm)(i),(0,n.Wm)(l),(0,n.Wm)(c),(0,n.Wm)(u)])),_:1})}const C={class:"home"},z={class:"container"},L={class:"card"},W={class:"card-body"};function V(e,t,o,a,r,s){return(0,n.wg)(),(0,n.iD)("div",C,[(0,n._)("div",z,[(0,n._)("div",L,[(0,n._)("div",W,[(0,n.WI)(e.$slots,"default",{},void 0,!0)])])])])}var N={name:"ContentBase"};const x=(0,h.Z)(N,[["render",V],["__scopeId","data-v-0c4ac83c"]]);var M=x;const O={class:"input-group input-group-lg"},B=(0,n._)("span",{class:"input-group-text",id:"inputGroup-sizing-default"},"请输入项目的绝对路径",-1);function F(e,t,o,r,s,i){return(0,n.wg)(),(0,n.iD)("div",O,[B,(0,n.wy)((0,n._)("input",{type:"text",class:"form-control","onUpdate:modelValue":t[0]||(t[0]=e=>s.inputValue=e)},null,512),[[a.nr,s.inputValue]]),(0,n._)("button",{class:"btn btn-primary",onClick:t[1]||(t[1]=(...e)=>i.getModulePath&&i.getModulePath(...e))},"Submit")])}var D=o(65);const T={modulePath:"",dir:"",projectInfo:{modulePath:"",moduleName:"",moduleVersion:"",otherPkgList:[],otherPkgCount:0,projectPkgList:[],projectPkgCount:0,goFileList:[],goFileCount:0}};var Z=(0,D.MT)({state:{...T},getters:{},mutations:{resetState(e){Object.assign(e,T)},setModuleInfo(e,t){e.projectInfo.modulePath=t.modulePath,e.projectInfo.moduleName=t.moduleName,e.projectInfo.moduleVersion="go "+t.moduleVersion},setDir(e,t){e.dir=t},setPackages(e,t){e.projectInfo.otherPkgList=t.otherPkgList,e.projectInfo.otherPkgCount=t.otherPkgCount,e.projectInfo.projectPkgList=t.projectPkgList,e.projectInfo.projectPkgCount=t.projectPkgCount},setGoFiles(e,t){e.projectInfo.goFileList=t.goFileList,e.projectInfo.goFileCount=t.goFileCount},setModuleName(e,t){e.projectInfo.moduleName=t.moduleName},setProjectInfo(e,t){e.projectInfo.modulePath=e.modulePath,e.projectInfo.moduleName=t.moduleName,e.projectInfo.packageCount=t.packageCount,e.projectInfo.packageList=t.packageList}},actions:{},modules:{}}),U=o(4161),$=o(7178);const H="http://localhost:8888/baize/v1",S=async e=>{try{const t=await U.Z.get(`${H}/local/module`,{params:e});if(200===t.status)return Z.commit("setModuleInfo",t.data.result),await R(e),await A(e),t.data.modulePath;throw new Error(`Request fail with status: ${t}`)}catch(t){"ERR_NETWORK"===t.code?(0,$.z8)({showClose:!0,message:"网络错误",type:"error"}):(0,$.z8)({showClose:!0,message:`${t.response.data.desc}`,type:"error"})}},R=async e=>{try{const t=await U.Z.get(`${H}/local/packages`,{params:e});if(200===t.status)return Z.commit("setPackages",t.data.result),t.data;throw new Error(`Request fail with status: ${t.status}`)}catch(t){(0,$.z8)({showClose:!0,message:`${t.response.data.desc}`,type:"error"})}},A=async e=>{try{const t=await U.Z.get(`${H}/local/file`,{params:e});if(200===t.status)return Z.commit("setGoFiles",t.data.result),t.data;throw new Error(`Request fail with status: ${t.status}`)}catch(t){(0,$.z8)({showClose:!0,message:`${t.response.data.desc}`,type:"error"})}};var G={name:"InputVue",data(){return{inputValue:"/Users/panjiangming/Project/baize/backend"}},methods:{async getModulePath(){const e={dir:this.inputValue};Z.commit("setDir",this.inputValue);try{S(e)}catch(t){console.error(t.message)}}}};const q=(0,h.Z)(G,[["render",F]]);var E=q,Y=o(7139);const K=e=>((0,n.dD)("data-v-48ab32f3"),e=e(),(0,n.Cn)(),e),J={class:"container"},Q=K((()=>(0,n._)("h5",null,"项目信息 / Project Info",-1)));function X(e,t,o,a,r,s){const i=(0,n.up)("el-descriptions-item"),l=(0,n.up)("el-descriptions");return(0,n.wg)(),(0,n.iD)("div",J,[Q,(0,n.Wm)(l,{direction:"vertical",column:3,size:r.size,border:""},{default:(0,n.w5)((()=>[(0,n.Wm)(i,{label:"模块名称 / Module Name"},{default:(0,n.w5)((()=>[(0,n.Uk)((0,Y.zw)(s.projectInfo.moduleName),1)])),_:1}),(0,n.Wm)(i,{label:"模块版本 / Module Version"},{default:(0,n.w5)((()=>[(0,n.Uk)((0,Y.zw)(s.projectInfo.moduleVersion),1)])),_:1}),(0,n.Wm)(i,{label:"Go文件数量 / Go File Count"},{default:(0,n.w5)((()=>[(0,n.Uk)((0,Y.zw)(s.projectInfo.goFileCount),1)])),_:1}),(0,n.Wm)(i,{label:"项目包数量 / Project Packages Count"},{default:(0,n.w5)((()=>[(0,n.Uk)((0,Y.zw)(s.projectInfo.projectPkgCount),1)])),_:1}),(0,n.Wm)(i,{label:"第三方包数量 / Other Packages Count"},{default:(0,n.w5)((()=>[(0,n.Uk)((0,Y.zw)(s.projectInfo.otherPkgCount),1)])),_:1})])),_:1},8,["size"])])}var ee={name:"ProjectInfo",data(){return{size:"large"}},computed:{projectInfo(){return this.$store.state.projectInfo}}};const te=(0,h.Z)(ee,[["render",X],["__scopeId","data-v-48ab32f3"]]);var oe=te;const ae=e=>((0,n.dD)("data-v-677490c2"),e=e(),(0,n.Cn)(),e),ne={class:"container"},re=ae((()=>(0,n._)("h5",null,"Go 文件列表 / Go File List",-1))),se={class:"container"},ie=ae((()=>(0,n._)("h5",null,"包列表 / Package List",-1))),le=ae((()=>(0,n._)("h6",null,"工程包列表 / Project Package List",-1))),ce=ae((()=>(0,n._)("h6",null,"外部包列表 / Other Package List",-1)));function ue(e,t,o,a,r,s){const i=(0,n.up)("el-table-column"),l=(0,n.up)("el-table");return(0,n.wg)(),(0,n.iD)(n.HY,null,[(0,n._)("div",ne,[re,(0,n.Wm)(l,{data:s.pkgList.goFileList,stripe:"",border:"","max-height":"280",style:{width:"100%"}},{default:(0,n.w5)((()=>[(0,n.Wm)(i,{prop:"name",label:"文件名 / File Name",style:{width:"50%"}}),(0,n.Wm)(i,{prop:"size",label:"文件大小 / File Size"})])),_:1},8,["data"])]),(0,n._)("div",se,[ie,le,(0,n.Wm)(l,{data:s.formattedProjectList,stripe:"",border:"","max-height":"280",style:{width:"100%","margin-bottom":"10px"}},{default:(0,n.w5)((()=>[(0,n.Wm)(i,{prop:"packageName",label:"包名 / Package Name"})])),_:1},8,["data"]),ce,(0,n.Wm)(l,{data:s.pkgList.otherPkgList,stripe:"",border:"","max-height":"280",style:{width:"100%"}},{default:(0,n.w5)((()=>[(0,n.Wm)(i,{prop:"name",label:"包名 / Package Name"}),(0,n.Wm)(i,{prop:"version",label:"包版本 / Package Version"})])),_:1},8,["data"])])],64)}var de={name:"ListInfo",data(){return{size:"large"}},computed:{pkgList(){return this.$store.state.projectInfo},formattedProjectList(){return this.pkgList.projectPkgList.map((e=>({packageName:e})))}}};const pe=(0,h.Z)(de,[["render",ue],["__scopeId","data-v-677490c2"]]);var me=pe;function fe(e,t,o,a,r,s){const i=(0,n.up)("el-backtop");return(0,n.wg)(),(0,n.j4)(i,{right:100,bottom:100})}var ge={name:"BackToTop"};const he=(0,h.Z)(ge,[["render",fe]]);var ve=he,be={name:"HomeView",components:{ContentBase:M,InputVue:E,ProjectInfo:oe,ListInfo:me,BackToTop:ve}};const we=(0,h.Z)(be,[["render",y]]);var ke=we;function _e(e,t,o,a,r,s){const i=(0,n.up)("AboutBaize"),l=(0,n.up)("ContentBase");return(0,n.wg)(),(0,n.j4)(l,null,{default:(0,n.w5)((()=>[(0,n.Wm)(i)])),_:1})}const je={class:"text-center"},Pe=(0,n._)("blockquote",{class:"blockquote h1"},[(0,n._)("p",null,"东望山有兽，名曰白泽，能言语。王者有德，明照幽远则至。")],-1),Ie=(0,n._)("figcaption",{class:"blockquote-footer"},[(0,n._)("cite",{title:"Source Title"},"《山海经》")],-1),ye=[Pe,Ie];function Ce(e,t,o,a,r,s){return(0,n.wg)(),(0,n.iD)("figure",je,ye)}var ze={name:"AboutBaize"};const Le=(0,h.Z)(ze,[["render",Ce]]);var We=Le,Ve={name:"AboutView",components:{ContentBase:M,AboutBaize:We}};const Ne=(0,h.Z)(Ve,[["render",_e]]);var xe=Ne;const Me=[{path:"/",name:"home",component:ke},{path:"/about",name:"about",component:xe}],Oe=(0,j.p7)({history:(0,j.PO)(),routes:Me,linkActiveClass:"active"});var Be=Oe,Fe=(o(5654),o(8458)),De=(o(4415),o(2748));const Te=(0,a.ri)(_);for(let Ze in De)Te.component(Ze,De[Ze]);Te.use(Z).use(Be).use(Fe.Z).mount("#app")}},t={};function o(a){var n=t[a];if(void 0!==n)return n.exports;var r=t[a]={exports:{}};return e[a].call(r.exports,r,r.exports,o),r.exports}o.m=e,function(){var e=[];o.O=function(t,a,n,r){if(!a){var s=1/0;for(u=0;u<e.length;u++){a=e[u][0],n=e[u][1],r=e[u][2];for(var i=!0,l=0;l<a.length;l++)(!1&r||s>=r)&&Object.keys(o.O).every((function(e){return o.O[e](a[l])}))?a.splice(l--,1):(i=!1,r<s&&(s=r));if(i){e.splice(u--,1);var c=n();void 0!==c&&(t=c)}}return t}r=r||0;for(var u=e.length;u>0&&e[u-1][2]>r;u--)e[u]=e[u-1];e[u]=[a,n,r]}}(),function(){o.n=function(e){var t=e&&e.__esModule?function(){return e["default"]}:function(){return e};return o.d(t,{a:t}),t}}(),function(){o.d=function(e,t){for(var a in t)o.o(t,a)&&!o.o(e,a)&&Object.defineProperty(e,a,{enumerable:!0,get:t[a]})}}(),function(){o.g=function(){if("object"===typeof globalThis)return globalThis;try{return this||new Function("return this")()}catch(e){if("object"===typeof window)return window}}()}(),function(){o.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)}}(),function(){o.r=function(e){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})}}(),function(){var e={143:0};o.O.j=function(t){return 0===e[t]};var t=function(t,a){var n,r,s=a[0],i=a[1],l=a[2],c=0;if(s.some((function(t){return 0!==e[t]}))){for(n in i)o.o(i,n)&&(o.m[n]=i[n]);if(l)var u=l(o)}for(t&&t(a);c<s.length;c++)r=s[c],o.o(e,r)&&e[r]&&e[r][0](),e[r]=0;return o.O(u)},a=self["webpackChunkbaize"]=self["webpackChunkbaize"]||[];a.forEach(t.bind(null,0)),a.push=t.bind(null,a.push.bind(a))}();var a=o.O(void 0,[998],(function(){return o(5064)}));a=o.O(a)})();
//# sourceMappingURL=app.06c2e10b.js.map