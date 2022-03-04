import{o as _,c as p,a as o,u as H,i as V,b as u,w as m,d,t as h,e as $,R as y,f as P,F as j,r as b,g as M,n as q,h as O,v as X,j as T,k as w,l as U,m as K,p as W,q as G,s as J}from"./vendor.js";const Q=function(){const a=document.createElement("link").relList;if(a&&a.supports&&a.supports("modulepreload"))return;for(const e of document.querySelectorAll('link[rel="modulepreload"]'))s(e);new MutationObserver(e=>{for(const n of e)if(n.type==="childList")for(const i of n.addedNodes)i.tagName==="LINK"&&i.rel==="modulepreload"&&s(i)}).observe(document,{childList:!0,subtree:!0});function t(e){const n={};return e.integrity&&(n.integrity=e.integrity),e.referrerpolicy&&(n.referrerPolicy=e.referrerpolicy),e.crossorigin==="use-credentials"?n.credentials="include":e.crossorigin==="anonymous"?n.credentials="omit":n.credentials="same-origin",n}function s(e){if(e.ep)return;e.ep=!0;const n=t(e);fetch(e.href,n)}};Q();var C=(l,a)=>{const t=l.__vccOpts||l;for(const[s,e]of a)t[s]=e;return t};const Y={},Z={width:"512",height:"512",version:"1.1",viewBox:"0 0 512 512",xmlns:"http://www.w3.org/2000/svg"},ee=o("path",{d:"m114.6 507.3c13.82-0.8729 90.06-3.083 106.2-3.079h13.49l2.639-5.293c5.191-10.41 10.35-29.57 13.13-48.74 1.811-12.5 1.488-16.68-1.826-23.56-2.528-5.249-11.04-13.06-20.54-18.86-28.38-17.32-51.21-34.77-76.05-58.12-6.673-6.272-9.377-8.209-16.08-11.53-12.37-6.124-20.36-12.88-29.19-24.7-1.45-1.941-1.795-2.038-4.695-1.334-1.722 0.4182-7.18 0.8166-12.13 0.8854-19.6 0.2722-38.18-7.715-52.49-22.56-15.97-16.57-23.43-35.49-23.41-59.37 0.01786-16.86 3.547-29.88 11.93-44 9.617-16.19 27.39-30.13 43.89-34.41 2.517-0.653 4.748-1.36 4.958-1.571 0.2103-0.2107-0.05169-3.078-0.5816-6.371-1.41-8.759-0.556-21.84 1.997-30.61 7.023-24.1 24.96-41.95 48.52-48.31 11.42-3.082 29.45-2.167 40.42 2.049 0.919 0.3533 1.637-0.7202 2.994-4.474 6.21-17.18 18.79-33.75 33.11-43.62 35.65-24.56 80.99-19.9 111.3 11.44l5.539 5.719 5.025-1.229c2.764-0.6762 9.075-1.428 14.02-1.672 16.17-0.7944 31.64 3.865 45.31 13.64l5.1 3.647 12.55 0.1171c9.984 0.09312 13.77 0.4491 18.49 1.739 29.1 7.945 50.37 32.58 56 64.86 1.354 7.764 1.096 21.6-0.5502 29.5l-0.916 4.394 5.367 7.086c20.1 26.54 25.52 61.05 14.68 93.52-7.862 23.55-24.99 43.53-46.38 54.11-10.16 5.023-23.1 8.587-31.28 8.613-3.79 0.0118-4.555 0.333-8.217 3.446-6.958 5.916-20.59 13.71-28.3 16.17-1.722 0.551-4.011 1.599-5.087 2.328-1.076 0.7297-8.119 4.802-15.65 9.05-18.57 10.47-37.53 22.69-48.72 31.39-13.27 10.32-16.83 22.63-14.72 50.94 2.648 35.64 8.059 59.41 16.33 71.75l2.89 4.313 29.4 0.4963c34.14 0.5764 92.53 2.291 93.01 2.731 0.1793 0.166-74.3 0.2503-165.5 0.1877-91.21-0.0631-161.4-0.392-156-0.732zm130.9-101.8c-0.1092-1.007-3.147-5.947-6.75-10.98-7.889-11.01-19.71-28.99-22.37-34.01-1.594-3.014-2.402-3.727-4.695-4.139-7.259-1.304-17.91-5.324-27.58-10.41-1.809-0.9509 5.669 20.15 9.032 25.49 4.859 7.714 27.69 24.94 45.08 34 5.177 2.7 7.575 2.715 7.285 0.0442zm6.168-19.46c0.8235-0.9976 1.074-4.638 1.074-15.62 0-7.875-0.1764-14.5-0.3923-14.71-0.2158-0.2163-2.307 0.1044-4.647 0.7128-2.34 0.6082-7 1.335-10.35 1.616-7.067 0.5907-6.896 0-2.984 10.28 2.73 7.17 6.328 12.81 10.4 16.3 3.513 3.013 5.293 3.38 6.908 1.424zm49.91-15.13c3.913-2 11.37-9.412 15.96-15.85 3.982-5.592 8.688-15 8.214-16.42-0.1628-0.49-3.275-1.75-6.916-2.799-3.641-1.049-9.568-3.501-13.17-5.449l-6.551-3.541-2.173 2.087c-2.148 2.062-2.179 2.234-2.645 14.37-0.2592 6.754-0.6282 15.17-0.8198 18.7-0.3026 5.572-0.1405 6.682 1.232 8.43 1.896 2.415 2.946 2.489 6.867 0.484zm-118.1-5.725c-2.19-4.729-6.256-17.66-6.256-19.89 0-0.8683-1.901-1.089-9.391-1.089-5.165 0-9.391 0.2195-9.391 0.488 0 1.561 25.33 25.39 26.99 25.39 0.1762 0-0.7011-2.206-1.949-4.901zm149.8-9.536c5.487-3.687 19.46-14.65 19.07-14.96-0.126-0.1032-4.197-0.3888-9.046-0.6338l-8.816-0.4452-3.967 5.891c-4.284 6.363-7.088 11.99-6.473 12.98 0.9066 1.47 4.434 0.3864 9.228-2.836z",fill:"#13ad73"},null,-1),te=o("path",{d:"m253.9 167.7c-0.2579 0.0513-1.238 0.5454-4.691 0.5216-3.625-0.0248 4.941 9.227 6.934 34.4-8.031-0.4678-8.408-3.285-24.81-5.999-15.88-2.627-25.28 5.883-29.57 18.88-3.429 10.39-3.592 23.66-1.19 36.39 2.348 12.45 7.148 24.39 13.74 32.65 6.909 8.654 14.05 10.96 20.49 10.82 8.459-0.1842 15.7-4.599 19.58-4.407 6.726 0.3331 13.29 8.894 28.17 2.151 6.047-2.741 12.03-7.819 17.08-14.81 4.984-6.907 9.051-15.68 11.35-25.91 2.018-8.972 2.678-19.07 1.41-30-1.616-13.93-7.519-20.42-14.9-23.01-8.286-2.905-18.43-0.8896-26.46 1.032-4.71 1.127-8.69 2.223-11.14 2.273-0.155 3e-3 -0.2829-4.1e-4 -0.4335 2e-3 -0.01-0.0695-0.0197-0.1385-0.0323-0.2078-2.249-11.58-4.704-29.29-5.401-33.64-0.152-0.9486 0.0845-1.177-0.1162-1.137z",fill:"#f2ac05","fill-rule":"evenodd"},null,-1),ne=[ee,te];function oe(l,a){return _(),p("svg",Z,ne)}var ae=C(Y,[["render",oe]]);const le={},se={width:"24",height:"24",version:"1.1",viewBox:"0 0 24 24",xmlns:"http://www.w3.org/2000/svg"},ie=o("path",{d:"m12 0c-6.624 0-12 5.376-12 12s5.376 12 12 12 12-5.376 12-12-5.376-12-12-12zm0 3.6c1.992 0 3.6 1.608 3.6 3.6s-1.608 3.6-3.6 3.6-3.6-1.608-3.6-3.6 1.608-3.6 3.6-3.6zm0 17.04c-3 0-5.652-1.536-7.2-3.864 0.036-2.388 4.8-3.696 7.2-3.696 2.388 0 7.164 1.308 7.2 3.696-1.548 2.328-4.2 3.864-7.2 3.864z"},null,-1),re=[ie];function ce(l,a){return _(),p("svg",se,re)}var ue=C(le,[["render",ce]]);const de={},_e={width:"24",height:"24",version:"1.1",viewBox:"0 0 24 24",xmlns:"http://www.w3.org/2000/svg"},me=o("path",{d:"m12 0-2.115 2.115 8.37 8.385h-18.255v3h18.255l-8.37 8.385 2.115 2.115 12-12z"},null,-1),he=[me];function pe(l,a){return _(),p("svg",_e,he)}var ve=C(de,[["render",pe]]);const fe=o("div",{id:"app-name"},[o("span",null,"Hey"),o("span",null,"Apple")],-1),ge={id:"nav-main"},$e={id:"nav-user"},be={href:"https://docs.heyapple.org",target:"_blank"},ye={setup(l){const{t:a}=H(),t=V("csrfToken");function s(n){n.preventDefault(),fetch("/auth/local",{method:"DELETE",headers:{"X-CSRF-Token":t}}).then(i=>{i.ok?window.location="/":window.dispatchEvent(new CustomEvent("error",{detail:{msg:a("signout.err"+i.status)}}))})}function e(n){n.stopPropagation(),document.querySelector("header nav").classList.toggle("open")}return document.addEventListener("click",function(){document.querySelector("header nav").classList.remove("open")}),(n,i)=>(_(),p("header",null,[u(ae,{id:"logo"}),fe,o("nav",null,[o("button",{onClick:e},[u(ve)]),o("ul",ge,[o("li",null,[u($(y),{to:"/"},{default:m(()=>[d(h(n.$t("nav.food")),1)]),_:1})]),o("li",null,[u($(y),{to:"/recipes"},{default:m(()=>[d(h(n.$t("nav.recipes")),1)]),_:1})]),o("li",null,[u($(y),{to:"/diary"},{default:m(()=>[d(h(n.$t("nav.diary")),1)]),_:1})]),o("li",null,[u($(y),{to:"/shopping"},{default:m(()=>[d(h(n.$t("nav.shopping")),1)]),_:1})])]),o("ul",$e,[o("li",null,[u($(y),{to:"/profile"},{default:m(()=>[d(h(n.$t("nav.profile")),1)]),_:1})]),o("li",null,[u($(y),{to:"/settings"},{default:m(()=>[d(h(n.$t("nav.settings")),1)]),_:1})]),o("li",null,[o("a",be,h(n.$t("nav.help")),1)]),o("li",null,[o("a",{href:"#",onClick:s},h(n.$t("nav.signout")),1)])])]),o("button",{onClick:e},[u(ue)])]))}};const we={setup(l){return(a,t)=>(_(),p(j,null,[u(ye),u($(P))],64))}},xe={},Le={version:"1.1",viewBox:"0 0 24 24",xmlns:"http://www.w3.org/2000/svg"},Fe=o("path",{id:"path-top",d:"m0 3v2h24v-2z"},null,-1),ke=o("path",{id:"path-mid",d:"m0 13h24v-2h-24z"},null,-1),Ee=o("path",{id:"path-bottom",d:"m0 21h24v-2h-24z"},null,-1),Me=[Fe,ke,Ee];function Ce(l,a){return _(),p("svg",Le,Me)}var ze=C(xe,[["render",Ce]]);const Ie={id:"filter"},Te=d(" This is the main search & filter area"),Se={id:"main"},Re={class:"controls"},Be=d("This is the main area"),De={id:"details"},Ve=d("This is the details area"),x={setup(l){const a=b(""),t=b("");function s(){a.value==""?(a.value="open-filter",t.value=""):a.value=""}return(e,n)=>(_(),p("main",{class:q([a.value,t.value])},[o("div",Ie,[M(e.$slots,"filter",{},()=>[Te])]),o("div",Se,[o("div",Re,[o("button",{onClick:s,class:"open-close"},[u(ze)])]),M(e.$slots,"main",{},()=>[Be])]),o("div",De,[M(e.$slots,"details",{},()=>[Ve])])],2))}};const Ne={class:"clickable-input"},Ae=["placeholder"],He=["value"],N={props:["label","placeholder"],emits:["confirm"],setup(l,{emit:a}){const t=b("");function s(e){e.preventDefault(),a("confirm",t.value),t.value=""}return(e,n)=>(_(),p("form",Ne,[O(o("input",{type:"text","onUpdate:modelValue":n[0]||(n[0]=i=>t.value=i),placeholder:l.placeholder},null,8,Ae),[[X,t.value]]),o("input",{type:"submit",onClick:s,value:l.label},null,8,He)]))}},Pe=["placeholder"],je=d("Additional filters"),qe={props:["placeholder"],emits:["result"],setup(l,{emit:a}){var t=void 0;function s(n){n.preventDefault(),clearTimeout(t),t=setTimeout(function(){e(n.target.closest("form"))},500)}function e(n){let i=new URLSearchParams(new FormData(n));fetch(n.action+"?"+i).then(L=>{L.json().then(z=>a("result",L.status,z))})}return(n,i)=>(_(),p("form",null,[o("input",{type:"text",name:"name",placeholder:l.placeholder,onInput:s},null,40,Pe),M(n.$slots,"default",{confirm:s},()=>[je])]))}};const Oe={class:"slider"},Xe=["name","value"],Ue=d(" \xA0\u2013\xA0 "),Ke=["name","value"],We={class:"bar"},Ge={class:"slide"},Je=o("div",{class:"interact-area"},null,-1),Qe=[Je],Ye=o("div",{class:"interact-area"},null,-1),Ze=[Ye],E={props:["label","name","unit","min","max","frac"],emits:["input"],setup(l,{emit:a}){const t=l,s=b(parseFloat(t.min).toFixed(t.frac)),e=b(parseFloat(t.max).toFixed(t.frac)),n=b(0),i=b(100);function L(c){c.target.blur();let r=parseFloat(t.min)||0,f=parseFloat(c.target.value)||r;s.value=f,e.value=Math.max(s.value,e.value);let v=parseFloat(t.max)||0;n.value=(f-r)*100/(v-r),i.value=Math.max(n.value,i.value),a("input",c)}function z(c){c.target.blur();let r=parseFloat(t.max)||0,f=parseFloat(c.target.value)||r;e.value=f,s.value=Math.min(s.value,e.value);let v=parseFloat(t.min)||0;i.value=(f-v)*100/(r-v),n.value=Math.min(n.value,i.value),a("input",c)}function F(c){let f=c.target.closest(".slide").getBoundingClientRect(),v=c.pageX!==void 0?c.pageX:c.changedTouches[0].pageX;console.log(v),v=Math.min(Math.max(v-f.left,0),f.width),console.log(v);let I=v*100/f.width,R=parseFloat(t.min)||0,A=parseFloat(t.max)||0,B=I/100*(A-R)+R;c.target.closest("button").classList.contains("min")?(s.value=B.toFixed(t.frac),e.value=Math.max(s.value,e.value),n.value=I,i.value=Math.max(n.value,i.value)):(e.value=B.toFixed(t.frac),s.value=Math.min(s.value,e.value),i.value=I,n.value=Math.min(n.value,i.value)),a("input",c)}function k(c){let r=c.target.closest("button");r.addEventListener("mousemove",F),r.addEventListener("touchmove",F),r.addEventListener("mouseup",g),r.addEventListener("touchend",g),r.addEventListener("mouseleave",g),r.addEventListener("touchcancel",g),r.classList.add("active")}function g(c){let r=c.target.closest("button");r.removeEventListener("mousemove",F),r.removeEventListener("touchmove",F),r.removeEventListener("mouseup",g),r.removeEventListener("touchend",g),r.removeEventListener("mouseleave",g),r.removeEventListener("touchcancel",g),r.classList.remove("active")}return(c,r)=>(_(),p("div",Oe,[o("label",null,[o("span",null,h(l.label)+" ("+h(c.$t("unit."+l.unit))+")",1),o("input",{type:"text",name:l.name,value:s.value,onChange:L},null,40,Xe),Ue,o("input",{type:"text",name:l.name,value:e.value,onChange:z},null,40,Ke)]),o("div",We,[o("div",Ge,[o("div",{class:"overlay",style:T({left:n.value+"%",right:100-i.value+"%"})},null,4),o("button",{type:"button",class:"handle min",style:T({left:n.value+"%"}),onMousedown:k,onTouchstart:k},Qe,36),o("button",{type:"button",class:"handle max",style:T({left:i.value+"%"}),onMousedown:k,onTouchstart:k},Ze,36)])])]))}},et={key:0},tt=d(" Food "),nt={setup(l){const a=V("perms");function t(e){console.log(e)}function s(e,n){console.log(e,n)}return(e,n)=>(_(),w(x,null,{filter:m(()=>[$(a).canCreateFood?(_(),p("section",et,[o("h2",null,h(e.$t("aria.headnew")),1),u(N,{label:e.$t("btn.new"),placeholder:e.$t("food.hintnew"),onConfirm:t},null,8,["label","placeholder"])])):U("",!0),o("section",null,[o("h2",null,h(e.$t("aria.headsearch")),1),u(qe,{action:"/api/v1/foods",placeholder:e.$t("food.hintsearch"),onResult:s},{default:m(i=>[u(E,{label:e.$t("food.energy"),onInput:i.confirm,name:"kcal",unit:"cal",min:"0",max:"900",frac:"0"},null,8,["label","onInput"]),u(E,{label:e.$t("food.fat"),onInput:i.confirm,name:"fat",unit:"g",min:"0",max:"100",frac:"0"},null,8,["label","onInput"]),u(E,{label:e.$t("food.carbs"),onInput:i.confirm,name:"carb",unit:"g",min:"0",max:"100",frac:"0"},null,8,["label","onInput"]),u(E,{label:e.$t("food.protein"),onInput:i.confirm,name:"prot",unit:"g",min:"0",max:"100",frac:"0"},null,8,["label","onInput"])]),_:1},8,["placeholder"])])]),main:m(()=>[tt]),_:1}))}},ot=d(" Recipes "),at={setup(l){function a(t){console.log(t)}return(t,s)=>(_(),w(x,null,{filter:m(()=>[o("section",null,[u(N,{label:t.$t("btn.new"),placeholder:t.$t("recipe.hintnew"),onConfirm:a},null,8,["label","placeholder"])])]),main:m(()=>[ot]),_:1}))}},lt=d(" Diary "),st={setup(l){return(a,t)=>(_(),w(x,null,{main:m(()=>[lt]),_:1}))}},it=d(" Shopping Lists "),rt={setup(l){return(a,t)=>(_(),w(x,null,{main:m(()=>[it]),_:1}))}},ct=d(" Profile "),ut={setup(l){return(a,t)=>(_(),w(x,null,{main:m(()=>[ct]),_:1}))}},dt=d(" Settings "),_t={setup(l){return(a,t)=>(_(),w(x,null,{main:m(()=>[dt]),_:1}))}},mt=K({history:W(),routes:[{path:"/",name:"food",component:nt},{path:"/recipes",name:"recipes",component:at},{path:"/diary",name:"diary",component:st},{path:"/shopping",name:"shopping",component:rt},{path:"/profile",name:"profile",component:ut},{path:"/settings",name:"settings",component:_t}]});let S=document.documentElement.lang||navigator.language;!S&&navigator.languages!=null&&(S=navigator.languages[0]);const D=document.querySelector("meta[name='_csrf']"),ht=D?D.content:"",pt=function(){const l=document.documentElement.dataset.perm||1,a=65536,t=131072;function s(e){return(l&e)==e}return{canCreateFood:s(a),canEditFood:s(t)}}();fetch("/app/l10n.json").then(l=>l.json()).then(l=>{const a=G({locale:S.split("-")[0],fallbackLocale:"en",messages:l}),t=J(we);t.provide("csrfToken",ht),t.provide("perms",pt),t.use(mt),t.use(a),t.mount("#app")});
