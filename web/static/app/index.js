import{o as d,c as p,a as e,u as H,i as B,b as r,w as M,d as z,t as l,e as n,R as X,f as J,g as D,h as V,n as N,r as y,F as Q,j as ne,k as ve,l as K,m as oe,p as ee,q as he,v as fe,s as ie,x as ce,y as pe,z as ge,A as we,B as $e}from"./vendor.js";const be=function(){const t=document.createElement("link").relList;if(t&&t.supports&&t.supports("modulepreload"))return;for(const c of document.querySelectorAll('link[rel="modulepreload"]'))_(c);new MutationObserver(c=>{for(const o of c)if(o.type==="childList")for(const a of o.addedNodes)a.tagName==="LINK"&&a.rel==="modulepreload"&&_(a)}).observe(document,{childList:!0,subtree:!0});function i(c){const o={};return c.integrity&&(o.integrity=c.integrity),c.referrerpolicy&&(o.referrerPolicy=c.referrerpolicy),c.crossorigin==="use-credentials"?o.credentials="include":c.crossorigin==="anonymous"?o.credentials="omit":o.credentials="same-origin",o}function _(c){if(c.ep)return;c.ep=!0;const o=i(c);fetch(c.href,o)}};be();var R=(s,t)=>{const i=s.__vccOpts||s;for(const[_,c]of t)i[_]=c;return i};const ye={},xe={width:"512",height:"512",version:"1.1",viewBox:"0 0 512 512",xmlns:"http://www.w3.org/2000/svg"},ke=e("path",{d:"m114.6 507.3c13.82-0.8729 90.06-3.083 106.2-3.079h13.49l2.639-5.293c5.191-10.41 10.35-29.57 13.13-48.74 1.811-12.5 1.488-16.68-1.826-23.56-2.528-5.249-11.04-13.06-20.54-18.86-28.38-17.32-51.21-34.77-76.05-58.12-6.673-6.272-9.377-8.209-16.08-11.53-12.37-6.124-20.36-12.88-29.19-24.7-1.45-1.941-1.795-2.038-4.695-1.334-1.722 0.4182-7.18 0.8166-12.13 0.8854-19.6 0.2722-38.18-7.715-52.49-22.56-15.97-16.57-23.43-35.49-23.41-59.37 0.01786-16.86 3.547-29.88 11.93-44 9.617-16.19 27.39-30.13 43.89-34.41 2.517-0.653 4.748-1.36 4.958-1.571 0.2103-0.2107-0.05169-3.078-0.5816-6.371-1.41-8.759-0.556-21.84 1.997-30.61 7.023-24.1 24.96-41.95 48.52-48.31 11.42-3.082 29.45-2.167 40.42 2.049 0.919 0.3533 1.637-0.7202 2.994-4.474 6.21-17.18 18.79-33.75 33.11-43.62 35.65-24.56 80.99-19.9 111.3 11.44l5.539 5.719 5.025-1.229c2.764-0.6762 9.075-1.428 14.02-1.672 16.17-0.7944 31.64 3.865 45.31 13.64l5.1 3.647 12.55 0.1171c9.984 0.09312 13.77 0.4491 18.49 1.739 29.1 7.945 50.37 32.58 56 64.86 1.354 7.764 1.096 21.6-0.5502 29.5l-0.916 4.394 5.367 7.086c20.1 26.54 25.52 61.05 14.68 93.52-7.862 23.55-24.99 43.53-46.38 54.11-10.16 5.023-23.1 8.587-31.28 8.613-3.79 0.0118-4.555 0.333-8.217 3.446-6.958 5.916-20.59 13.71-28.3 16.17-1.722 0.551-4.011 1.599-5.087 2.328-1.076 0.7297-8.119 4.802-15.65 9.05-18.57 10.47-37.53 22.69-48.72 31.39-13.27 10.32-16.83 22.63-14.72 50.94 2.648 35.64 8.059 59.41 16.33 71.75l2.89 4.313 29.4 0.4963c34.14 0.5764 92.53 2.291 93.01 2.731 0.1793 0.166-74.3 0.2503-165.5 0.1877-91.21-0.0631-161.4-0.392-156-0.732zm130.9-101.8c-0.1092-1.007-3.147-5.947-6.75-10.98-7.889-11.01-19.71-28.99-22.37-34.01-1.594-3.014-2.402-3.727-4.695-4.139-7.259-1.304-17.91-5.324-27.58-10.41-1.809-0.9509 5.669 20.15 9.032 25.49 4.859 7.714 27.69 24.94 45.08 34 5.177 2.7 7.575 2.715 7.285 0.0442zm6.168-19.46c0.8235-0.9976 1.074-4.638 1.074-15.62 0-7.875-0.1764-14.5-0.3923-14.71-0.2158-0.2163-2.307 0.1044-4.647 0.7128-2.34 0.6082-7 1.335-10.35 1.616-7.067 0.5907-6.896 0-2.984 10.28 2.73 7.17 6.328 12.81 10.4 16.3 3.513 3.013 5.293 3.38 6.908 1.424zm49.91-15.13c3.913-2 11.37-9.412 15.96-15.85 3.982-5.592 8.688-15 8.214-16.42-0.1628-0.49-3.275-1.75-6.916-2.799-3.641-1.049-9.568-3.501-13.17-5.449l-6.551-3.541-2.173 2.087c-2.148 2.062-2.179 2.234-2.645 14.37-0.2592 6.754-0.6282 15.17-0.8198 18.7-0.3026 5.572-0.1405 6.682 1.232 8.43 1.896 2.415 2.946 2.489 6.867 0.484zm-118.1-5.725c-2.19-4.729-6.256-17.66-6.256-19.89 0-0.8683-1.901-1.089-9.391-1.089-5.165 0-9.391 0.2195-9.391 0.488 0 1.561 25.33 25.39 26.99 25.39 0.1762 0-0.7011-2.206-1.949-4.901zm149.8-9.536c5.487-3.687 19.46-14.65 19.07-14.96-0.126-0.1032-4.197-0.3888-9.046-0.6338l-8.816-0.4452-3.967 5.891c-4.284 6.363-7.088 11.99-6.473 12.98 0.9066 1.47 4.434 0.3864 9.228-2.836z",fill:"#13ad73"},null,-1),Ce=e("path",{d:"m253.9 167.7c-0.2579 0.0513-1.238 0.5454-4.691 0.5216-3.625-0.0248 4.941 9.227 6.934 34.4-8.031-0.4678-8.408-3.285-24.81-5.999-15.88-2.627-25.28 5.883-29.57 18.88-3.429 10.39-3.592 23.66-1.19 36.39 2.348 12.45 7.148 24.39 13.74 32.65 6.909 8.654 14.05 10.96 20.49 10.82 8.459-0.1842 15.7-4.599 19.58-4.407 6.726 0.3331 13.29 8.894 28.17 2.151 6.047-2.741 12.03-7.819 17.08-14.81 4.984-6.907 9.051-15.68 11.35-25.91 2.018-8.972 2.678-19.07 1.41-30-1.616-13.93-7.519-20.42-14.9-23.01-8.286-2.905-18.43-0.8896-26.46 1.032-4.71 1.127-8.69 2.223-11.14 2.273-0.155 3e-3 -0.2829-4.1e-4 -0.4335 2e-3 -0.01-0.0695-0.0197-0.1385-0.0323-0.2078-2.249-11.58-4.704-29.29-5.401-33.64-0.152-0.9486 0.0845-1.177-0.1162-1.137z",fill:"#f2ac05","fill-rule":"evenodd"},null,-1),ze=[ke,Ce];function Ie(s,t){return d(),p("svg",xe,ze)}var Fe=R(ye,[["render",Ie]]);const Le={},Me={width:"24",height:"24",version:"1.1",viewBox:"0 0 24 24",xmlns:"http://www.w3.org/2000/svg"},Te=e("path",{d:"m12 0c-6.624 0-12 5.376-12 12s5.376 12 12 12 12-5.376 12-12-5.376-12-12-12zm0 3.6c1.992 0 3.6 1.608 3.6 3.6s-1.608 3.6-3.6 3.6-3.6-1.608-3.6-3.6 1.608-3.6 3.6-3.6zm0 17.04c-3 0-5.652-1.536-7.2-3.864 0.036-2.388 4.8-3.696 7.2-3.696 2.388 0 7.164 1.308 7.2 3.696-1.548 2.328-4.2 3.864-7.2 3.864z"},null,-1),Ee=[Te];function Se(s,t){return d(),p("svg",Me,Ee)}var De=R(Le,[["render",Se]]);const Re={},Be={width:"24",height:"24",version:"1.1",viewBox:"0 0 24 24",xmlns:"http://www.w3.org/2000/svg"},Ve=e("path",{d:"m12 0-2.115 2.115 8.37 8.385h-18.255v3h18.255l-8.37 8.385 2.115 2.115 12-12z"},null,-1),Ae=[Ve];function Ne(s,t){return d(),p("svg",Be,Ae)}var ue=R(Re,[["render",Ne]]);const je=e("div",{id:"app-name"},[e("span",null,"Hey"),e("span",null,"Apple")],-1),Oe={id:"nav-main"},He={id:"nav-user"},Xe={href:"https://docs.heyapple.org",target:"_blank"},qe={setup(s){const{t}=H(),i=B("csrfToken");function _(o){o.preventDefault(),fetch("/auth/local",{method:"DELETE",headers:{"X-CSRF-Token":i}}).then(a=>{a.ok?window.location="/":window.dispatchEvent(new CustomEvent("error",{detail:{msg:t("signout.err"+a.status)}}))})}function c(o){o.stopPropagation(),document.querySelector("header nav").classList.toggle("open")}return document.addEventListener("click",function(){document.querySelector("header nav").classList.remove("open")}),(o,a)=>(d(),p("header",null,[r(Fe,{id:"logo"}),je,e("nav",null,[e("button",{onClick:c},[r(ue)]),e("ul",Oe,[e("li",null,[r(n(X),{to:"/"},{default:M(()=>[z(l(n(t)("nav.food")),1)]),_:1})]),e("li",null,[r(n(X),{to:"/recipes"},{default:M(()=>[z(l(n(t)("nav.recipes")),1)]),_:1})]),e("li",null,[r(n(X),{to:"/diary"},{default:M(()=>[z(l(n(t)("nav.diary")),1)]),_:1})]),e("li",null,[r(n(X),{to:"/shopping"},{default:M(()=>[z(l(n(t)("nav.shopping")),1)]),_:1})])]),e("ul",He,[e("li",null,[r(n(X),{to:"/profile"},{default:M(()=>[z(l(n(t)("nav.profile")),1)]),_:1})]),e("li",null,[r(n(X),{to:"/settings"},{default:M(()=>[z(l(n(t)("nav.settings")),1)]),_:1})]),e("li",null,[e("a",Xe,l(n(t)("nav.help")),1)]),e("li",null,[e("a",{href:"#",onClick:_},l(n(t)("nav.signout")),1)])])]),e("button",{onClick:c},[r(De)])]))}},Ue={},Pe={width:"24",height:"24",version:"1.1",viewBox:"0 0 24 24",xmlns:"http://www.w3.org/2000/svg"},We=e("path",{d:"m7.6364 17.318-5.7273-5.7273-1.9091 1.9091 7.6364 7.6364 16.364-16.364-1.9091-1.9091z"},null,-1),Ke=[We];function Ge(s,t){return d(),p("svg",Pe,Ke)}var ae=R(Ue,[["render",Ge]]);const Je={},Qe={width:"24",height:"24",version:"1.1",viewBox:"0 0 24 24",xmlns:"http://www.w3.org/2000/svg"},Ye=e("path",{d:"m12 1.6364-12 20.727h24zm0 4.353 8.2159 14.192h-16.43zm-1.0909 4.3743v5.4545h2.1818v-5.4545zm0 6.5455v2.1818h2.1818v-2.1818z"},null,-1),Ze=[Ye];function et(s,t){return d(),p("svg",Qe,Ze)}var tt=R(Je,[["render",et]]);const nt={props:["msg"],emits:["timeout"],setup(s,{emit:t}){const i=s;return J(()=>{setTimeout(function(){t("timeout",i.msg.id)},i.msg.time)}),(_,c)=>(d(),p("div",{class:N(["message",[s.msg.type,s.msg.id]])},[s.msg.type=="message"?(d(),D(ae,{key:0})):V("",!0),s.msg.type!="message"?(d(),D(tt,{key:1})):V("",!0),e("p",null,l(s.msg.msg),1)],2))}};const at={id:"messages"},lt={setup(s){const t=y([]);let i=0;function _(o){t.value.push({id:i++,type:o.type,msg:o.detail.msg,time:o.detail.timeout})}function c(o){t.value=t.value.filter(a=>a.id!=o)}return J(()=>{window.addEventListener("message",_),window.addEventListener("warning",_),window.addEventListener("error",_)}),(o,a)=>(d(),p("div",at,[(d(!0),p(Q,null,ne(t.value,u=>(d(),D(nt,{key:u.id,msg:u,onTimeout:c},null,8,["msg"]))),128))]))}};const ot={setup(s){return J(()=>{document.querySelector("body > .spinner-container").remove()}),(t,i)=>(d(),p(Q,null,[r(qe),r(n(ve)),r(lt)],64))}},st={},it={version:"1.1",viewBox:"0 0 24 24",xmlns:"http://www.w3.org/2000/svg"},ct=e("path",{id:"path-top",d:"m0 3v2h24v-2z"},null,-1),ut=e("path",{id:"path-mid",d:"m0 13h24v-2h-24z"},null,-1),rt=e("path",{id:"path-bottom",d:"m0 21h24v-2h-24z"},null,-1),dt=[ct,ut,rt];function _t(s,t){return d(),p("svg",it,dt)}var mt=R(st,[["render",_t]]);const vt={},ht={width:"512",height:"512",version:"1.1",viewBox:"0 0 512 512",xmlns:"http://www.w3.org/2000/svg"},ft=e("path",{d:"m251.2 4.007c-17.24-0.2277-34.69 4.924-50.29 15.67-14.33 9.866-26.9 26.43-33.11 43.62-1.357 3.754-2.073 4.828-2.992 4.475-10.97-4.217-29-5.131-40.42-2.049-23.57 6.359-41.5 24.22-48.52 48.31-2.553 8.761-3.407 21.84-1.997 30.6 0.5299 3.293 0.7902 6.162 0.5799 6.372-0.2104 0.2108-2.439 0.9186-4.956 1.572-16.5 4.283-34.28 18.21-43.89 34.41-8.382 14.11-11.91 27.14-11.93 44-0.02401 23.88 7.433 42.79 23.41 59.36 14.31 14.85 32.89 22.84 52.49 22.56 4.95-0.0689 10.41-0.4681 12.13-0.8863 2.9-0.7045 3.245-0.6061 4.695 1.334 8.833 11.82 16.82 18.58 29.19 24.7 6.707 3.32 9.412 5.257 16.09 11.53 24.85 23.35 47.68 40.8 76.05 58.12 9.503 5.8 18.01 13.61 20.54 18.86 3.313 6.878 3.637 11.06 1.825 23.56-2.778 19.18-7.936 38.33-13.13 48.74l-2.639 5.295h-13.49c-16.17-3e-3 -92.41 2.205-106.2 3.077l0.01-3e-3c-5.381 0.3399 64.84 0.6684 156 0.7315 91.2 0.0625 165.7-0.0216 165.5-0.1878-0.4753-0.4399-58.86-2.155-93-2.731l-29.4-0.4942-2.89-4.313c-8.274-12.35-13.68-36.12-16.33-71.76-2.105-28.31 1.454-40.62 14.72-50.94 11.18-8.699 30.14-20.92 48.72-31.39 7.533-4.248 14.57-8.318 15.65-9.048 1.076-0.7296 3.366-1.779 5.087-2.33 7.71-2.467 21.34-10.26 28.3-16.17 3.662-3.113 4.427-3.435 8.218-3.446 8.183-0.0267 21.12-3.59 31.28-8.613 21.39-10.58 38.52-30.56 46.38-54.11 10.84-32.47 5.42-66.98-14.68-93.52l-5.364-7.084 0.916-4.395c1.646-7.896 1.904-21.74 0.5502-29.5-5.632-32.29-26.9-56.92-56-64.86-4.726-1.29-8.51-1.647-18.49-1.74l-12.55-0.1186-5.101-3.647c-13.67-9.776-29.15-14.43-45.31-13.64-4.95 0.2432-11.26 0.9944-14.02 1.671l-5.028 1.229-5.539-5.72c-17.07-17.63-38.89-26.81-61.05-27.1zm2.728 163.7c0.2008-0.0398-0.0367 0.1881 0.1154 1.137 0.6972 4.35 3.151 22.06 5.4 33.64 0.0118 0.0694 0.0228 0.138 0.0328 0.2076 0.1506-2e-3 0.2799 9e-5 0.4349-3e-3 2.448-0.0501 6.427-1.143 11.14-2.27 8.028-1.922 18.18-3.94 26.46-1.035 7.379 2.587 13.28 9.078 14.9 23.01 1.268 10.94 0.6106 21.03-1.407 30-2.301 10.23-6.37 19.01-11.35 25.91-5.045 6.991-11.03 12.07-17.07 14.81-14.87 6.743-21.44-1.818-28.16-2.152-3.88-0.1922-11.12 4.224-19.58 4.409-6.437 0.1402-13.58-2.169-20.49-10.82-6.596-8.262-11.39-20.2-13.74-32.65-2.402-12.73-2.24-26 1.189-36.39 4.29-13 13.69-21.51 29.58-18.88 16.41 2.714 16.78 5.532 24.81 6-1.993-25.18-10.56-34.43-6.936-34.4 3.453 0.0238 4.434-0.4726 4.692-0.5239zm45.16 159.1 6.55 3.542c3.603 1.947 9.529 4.401 13.17 5.45 3.641 1.049 6.753 2.307 6.916 2.797 0.4733 1.423-4.232 10.83-8.214 16.42-4.588 6.442-12.05 13.85-15.96 15.86-3.919 2.004-4.971 1.93-6.867-0.4843-1.372-1.748-1.532-2.86-1.229-8.432 0.1916-3.529 0.558-11.94 0.8171-18.7 0.4657-12.13 0.4983-12.3 2.646-14.37zm35.38 12.77 8.817 0.4448c4.849 0.245 8.919 0.5294 9.045 0.6326 0.3835 0.315-13.58 11.27-19.07 14.96-4.795 3.222-8.322 4.307-9.229 2.837-0.6151-0.9971 2.187-6.619 6.471-12.98zm-166.7 4.596c7.489 0 9.391 0.219 9.391 1.087 0 2.234 4.067 15.16 6.257 19.89 1.248 2.694 2.123 4.9 1.947 4.9-1.66 0-26.99-23.83-26.99-25.39 0-0.2685 4.226-0.4876 9.391-0.4876zm16.13 1.727c0.0388-0.0118 0.0851-6e-3 0.1417 0.0231 9.674 5.087 20.32 9.107 27.58 10.41 2.293 0.4121 3.101 1.124 4.695 4.138 2.657 5.026 14.48 23 22.37 34.02 3.603 5.03 6.639 9.968 6.748 10.98 0.2897 2.671-2.108 2.658-7.285-0.0429-17.39-9.068-40.22-26.29-45.07-34.01-3.258-5.172-10.38-25.14-9.173-25.52zm67.9 9.71c0.263-9e-3 0.4337 9e-3 0.4876 0.0625 0.2159 0.2154 0.3921 6.837 0.3921 14.71 0 10.98-0.2506 14.62-1.074 15.62-1.615 1.956-3.397 1.59-6.909-1.423-4.068-3.49-7.665-9.13-10.4-16.3-3.912-10.28-4.082-9.686 2.985-10.28 3.355-0.2804 8.013-1.006 10.35-1.615 1.755-0.4562 3.373-0.7502 4.162-0.7776z"},null,-1),pt=[ft];function gt(s,t){return d(),p("svg",ht,pt)}var wt=R(vt,[["render",gt]]);const $t={},bt={width:"24",height:"24",version:"1.1",viewBox:"0 0 24 24",xmlns:"http://www.w3.org/2000/svg"},yt=e("path",{d:"m12 6c1.65 0 3-1.35 3-3s-1.35-3-3-3-3 1.35-3 3 1.35 3 3 3zm0 3c-1.65 0-3 1.35-3 3s1.35 3 3 3 3-1.35 3-3-1.35-3-3-3zm0 9c-1.65 0-3 1.35-3 3s1.35 3 3 3 3-1.35 3-3-1.35-3-3-3z"},null,-1),xt=[yt];function kt(s,t){return d(),p("svg",bt,xt)}var Ct=R($t,[["render",kt]]);const zt={id:"filter"},It=z(" This is the main search & filter area"),Ft={id:"main"},Lt={class:"controls"},Mt=e("span",{class:"spacer"},null,-1),Tt={class:"content"},Et=z("This is the main area"),St={id:"details"},Dt={class:"controls"},Rt=e("span",{class:"spacer"},null,-1),Bt={class:"placeholder"},U={emits:["detailVisibility"],setup(s,{expose:t,emit:i}){const{t:_}=H(),c=y(""),o=y("");function a(){c.value==""?(c.value="open-filter",o.value="",i("detailVisibility")):c.value=""}function u(){i("detailVisibility"),o.value==""?(o.value="open-details",c.value=""):o.value=""}function g(){i("detailVisibility"),o.value==""&&(o.value="open-details",c.value="")}return t({showDetails:g}),(v,T)=>(d(),p("main",{class:N([c.value,o.value])},[e("div",zt,[K(v.$slots,"filter",{},()=>[It])]),e("div",Ft,[e("div",Lt,[e("button",{onClick:a,class:"open-filter icon"},[r(mt)]),Mt,e("button",{onClick:u,class:"open-details icon"},[r(Ct)])]),e("div",Tt,[K(v.$slots,"main",{},()=>[Et])])]),e("div",St,[e("div",Dt,[K(v.$slots,"head-details"),Rt,e("button",{onClick:u,class:"open-details icon"},[r(ue)])]),K(v.$slots,"details",{},()=>[e("div",Bt,[r(wt),e("p",null,l(n(_)("details.noitem")),1)])])])],2))}},Vt=["placeholder"],At=z("Additional filters"),re={props:["data","placeholder"],emits:["result"],setup(s,{emit:t}){const i=s;var _=void 0;function c(a){a.preventDefault(),clearTimeout(_),_=setTimeout(function(){o(a.target.closest("form"))},500)}function o(a){let u=new FormData(a),g=i.data.filter(v=>{let T=v.size||1;for(let x of u.keys()){if(x=="name"){let L=u.get(x).toLowerCase();if(!v[x].toLowerCase().includes(L))return!1;continue}let[k,w]=u.getAll(x).map(L=>parseFloat(L));if(!isNaN(k)&&!isNaN(w)&&(v[x]/T<k||w<v[x]/T))return!1}return!0});t("result",g)}return(a,u)=>(d(),p("form",null,[e("input",{type:"text",name:"name",placeholder:s.placeholder,onInput:c},null,40,Vt),K(a.$slots,"default",{confirm:c},()=>[At])]))}};const Nt={class:"slider"},jt=["name","value"],Ot=z(" \xA0\u2013\xA0 "),Ht=["name","value"],Xt={class:"bar"},qt={class:"slide"},Ut=e("div",{class:"interact-area"},null,-1),Pt=[Ut],Wt=e("div",{class:"interact-area"},null,-1),Kt=[Wt],S={props:["label","name","unit","min","max","frac"],emits:["input"],setup(s,{emit:t}){const i=s,{t:_}=H(),c=y(parseFloat(i.min).toFixed(i.frac)),o=y(parseFloat(i.max).toFixed(i.frac)),a=y(0),u=y(100);oe(()=>i.min,(h,f)=>g(h)),oe(()=>i.max,(h,f)=>v(h));function g(h){let f=parseFloat(i.min)||0,I=parseFloat(h)||f;c.value=I,o.value=Math.max(c.value,o.value);let C=parseFloat(i.max)||0;a.value=(I-f)*100/(C-f),u.value=Math.max(a.value,u.value)}function v(h){let f=parseFloat(i.max)||0,I=parseFloat(h)||f;o.value=I,c.value=Math.min(c.value,o.value);let C=parseFloat(i.min)||0;u.value=(I-C)*100/(f-C),a.value=Math.min(a.value,u.value)}function T(h){h.target.blur(),g(h.target.value),t("input",h)}function x(h){h.target.blur(),v(h.target.value),t("input",h)}function k(h){let I=h.target.closest(".slide").getBoundingClientRect(),C=h.pageX!==void 0?h.pageX:h.changedTouches[0].pageX;C=Math.min(Math.max(C-I.left,0),I.width);let m=C*100/I.width,F=parseFloat(i.min)||0,Z=parseFloat(i.max)||0,j=m/100*(Z-F)+F;h.target.closest("button").classList.contains("min")?(c.value=j.toFixed(i.frac),o.value=Math.max(c.value,o.value),a.value=m,u.value=Math.max(a.value,u.value)):(o.value=j.toFixed(i.frac),c.value=Math.min(c.value,o.value),u.value=m,a.value=Math.min(a.value,u.value)),t("input",h)}function w(h){let f=h.target.closest("button");f.addEventListener("mousemove",k),f.addEventListener("touchmove",k),f.addEventListener("mouseup",L),f.addEventListener("touchend",L),f.addEventListener("mouseleave",L),f.addEventListener("touchcancel",L),f.classList.add("active")}function L(h){let f=h.target.closest("button");f.removeEventListener("mousemove",k),f.removeEventListener("touchmove",k),f.removeEventListener("mouseup",L),f.removeEventListener("touchend",L),f.removeEventListener("mouseleave",L),f.removeEventListener("touchcancel",L),f.classList.remove("active")}return(h,f)=>(d(),p("div",Nt,[e("label",null,[e("span",null,l(s.label)+" ("+l(n(_)("unit."+s.unit))+")",1),e("input",{type:"text",name:s.name,value:c.value,onChange:T},null,40,jt),Ot,e("input",{type:"text",name:s.name,value:o.value,onChange:x},null,40,Ht)]),e("div",Xt,[e("div",qt,[e("div",{class:"overlay",style:ee({left:a.value+"%",right:100-u.value+"%"})},null,4),e("button",{type:"button",class:"handle min",style:ee({left:a.value+"%"}),onMousedown:w,onTouchstart:w},Pt,36),e("button",{type:"button",class:"handle max",style:ee({left:u.value+"%"}),onMousedown:w,onTouchstart:w},Kt,36)])])]))}};const Gt={class:"clickable-input"},Jt=["placeholder"],Qt=["value"],de={props:["label","placeholder"],emits:["confirm"],setup(s,{emit:t}){const i=y("");function _(c){c.preventDefault(),t("confirm",i.value),i.value=""}return(c,o)=>(d(),p("form",Gt,[he(e("input",{type:"text","onUpdate:modelValue":o[0]||(o[0]=a=>i.value=a),placeholder:s.placeholder},null,8,Jt),[[fe,i.value]]),e("input",{type:"submit",onClick:_,value:s.label},null,8,Qt)]))}};const Yt={},Zt={class:"icon sort-arrow"};function en(s,t){return d(),p("span",Zt)}var W=R(Yt,[["render",en]]);const tn=["onClick"],nn={class:"name"},an={class:"num"},ln={class:"unit"},on={class:"m num"},sn={class:"unit"},cn={class:"m num"},un={class:"unit"},rn={class:"m num"},dn={class:"unit"},_e={props:["items"],emits:"selected",setup(s,{emit:t}){const i=s,{t:_,locale:c}=H(),o=y("name"),a=y("asc"),u=new Intl.Collator(c.value,{numeric:!0}),g=ie(()=>a.value=="asc"?[...i.items].sort((x,k)=>u.compare(x[o.value],k[o.value])):[...i.items].sort((x,k)=>-u.compare(x[o.value],k[o.value])));function v(x,k,w=1){return k?+parseFloat(x/k).toFixed(w):x}function T(x){let k=x.target.dataset.sort;o.value==k?a.value=a.value=="asc"?"desc":"asc":o.value=k}return(x,k)=>(d(),p("table",null,[e("thead",null,[e("tr",{class:N(a.value)},[e("th",{class:N(["name sort",{active:o.value=="name"}]),onClick:T,"data-sort":"name"},[z(l(n(_)("food.name"))+" ",1),r(W)],2),e("th",{class:N(["num sort",{active:o.value=="kcal"}]),onClick:T,"data-sort":"kcal"},[r(W),z(" "+l(n(_)("food.energy")),1)],2),e("th",{class:N(["m num sort",{active:o.value=="fat"}]),onClick:T,"data-sort":"fat"},[r(W),z(" "+l(n(_)("food.fat")),1)],2),e("th",{class:N(["m num sort",{active:o.value=="carb"}]),onClick:T,"data-sort":"carb"},[r(W),z(" "+l(n(_)("food.carbs2")),1)],2),e("th",{class:N(["m num sort",{active:o.value=="prot"}]),onClick:T,"data-sort":"prot"},[r(W),z(" "+l(n(_)("food.protein")),1)],2)],2)]),e("tbody",null,[(d(!0),p(Q,null,ne(n(g),w=>(d(),p("tr",{key:w.id,onClick:L=>x.$emit("selected",w.id)},[e("td",nn,l(w.name),1),e("td",an,[z(l(v(w.kcal,w.size))+" ",1),e("span",ln,l(n(_)("unit.cal")),1)]),e("td",on,[z(l(v(w.fat,w.size))+" ",1),e("span",sn,l(n(_)("unit.g")),1)]),e("td",cn,[z(l(v(w.carb,w.size))+" ",1),e("span",un,l(n(_)("unit.g")),1)]),e("td",rn,[z(l(v(w.prot,w.size))+" ",1),e("span",dn,l(n(_)("unit.g")),1)])],8,tn))),128))])]))}},_n={},mn={width:"24",height:"24",version:"1.1",viewBox:"0 0 24 24",xmlns:"http://www.w3.org/2000/svg"},vn=e("path",{d:"m-7.5e-8 19.001v4.9993h4.9993l14.745-14.745-4.9993-4.9993zm23.61-13.611c0.51993-0.51993 0.51993-1.3598 0-1.8797l-3.1196-3.1196c-0.51993-0.51993-1.3598-0.51993-1.8797 0l-2.4397 2.4397 4.9993 4.9993z"},null,-1),hn=[vn];function fn(s,t){return d(),p("svg",mn,hn)}var me=R(_n,[["render",fn]]);const pn={key:0,class:"new-item"},gn=e("section",{class:"subtitle"},"Some food category",-1),wn={class:"tags"},$n=e("span",{class:"tag"},"Tag 1",-1),bn=e("span",{class:"tag"},"Tag 2",-1),yn=e("span",{class:"tag"},"Tag 3",-1),xn=["disabled"],kn={class:"nutrient-block"},Cn=["disabled"],zn=["value"],In={class:"unit"},Fn=["value"],Ln={class:"unit"},Mn=["value"],Tn={class:"unit"},En=["value"],Sn={class:"unit"},Dn=["value"],Rn={class:"unit"},Bn=["disabled"],Vn=["value"],An={class:"unit"},Nn=["value"],jn={class:"unit"},On=["value"],Hn={class:"unit"},Xn=["value"],qn={class:"unit"},Un=["value"],Pn={class:"unit"},Wn={setup(s){const{t}=H(),i=B("log"),_=B("csrfToken"),c=B("perms"),o=B("food"),a=y([]),u=y(null),g=y(!1),v=y(!1),T=y(null),x=y(null);function k(C){fetch("/api/v1/food",{method:"POST",headers:{"X-CSRF-Token":_}}).then(m=>{if(!m.ok)throw t("createfood.err"+m.status);return m.json()}).then(m=>{m.name=C,o.value.push(m),a.value.push(m),i.msg(t("createfood.ok")),h(m.id)}).catch(m=>i.err(m))}function w(){v.value=!0;let C=u.value.id;fetch("/api/v1/food/"+C,{method:"PUT",headers:{"Content-Type":"application/x-www-form-urlencoded","X-CSRF-Token":_},body:new URLSearchParams(new FormData(x.value))}).then(m=>{if(!m.ok)throw t("savefood.err"+m.status);return g.value=!1,fetch("/api/v1/food/"+C)}).then(m=>m.json()).then(m=>{m.name=t(m.id.toString()),o.value=o.value.map(F=>m.id==F.id?m:F),a.value=a.value.map(F=>m.id==F.id?m:F),u.value=u.value.id==m.id?m:u.value,i.msg(t("savefood.ok"))}).catch(m=>i.err(m)).finally(()=>{setTimeout(function(){v.value=!1},500)})}function L(C){a.value=C,u.value&&a.value.filter(m=>m.id==u.value.id).length==0&&(u.value=null)}function h(C){u.value=a.value.filter(m=>m.id==C)[0],T.value.showDetails()}function f(){g.value?w():g.value=!0}function I(C){C.target.blur(),isNaN(parseFloat(C.target.value))&&(C.target.value=u.value[C.target.name])}return(C,m)=>(d(),D(U,{ref_key:"main",ref:T,onDetailVisibility:m[0]||(m[0]=F=>g.value=!1)},ce({filter:M(()=>[n(c).canCreateFood?(d(),p("section",pn,[e("h2",null,l(n(t)("aria.headnew")),1),r(de,{label:n(t)("btn.new"),placeholder:n(t)("food.hintnew"),onConfirm:k},null,8,["label","placeholder"])])):V("",!0),e("section",null,[e("h2",null,l(n(t)("aria.headsearch")),1),r(re,{data:n(o),placeholder:n(t)("food.hintsearch"),onResult:L},{default:M(F=>[e("fieldset",null,[e("legend",null,l(n(t)("aria.headmacro1")),1),r(S,{label:n(t)("food.energy"),onInput:F.confirm,name:"kcal",unit:"cal",min:"0",max:"900",frac:"0"},null,8,["label","onInput"]),r(S,{label:n(t)("food.fat"),onInput:F.confirm,name:"fat",unit:"g",min:"0",max:"100",frac:"0"},null,8,["label","onInput"]),r(S,{label:n(t)("food.carbs"),onInput:F.confirm,name:"carb",unit:"g",min:"0",max:"100",frac:"0"},null,8,["label","onInput"]),r(S,{label:n(t)("food.protein"),onInput:F.confirm,name:"prot",unit:"g",min:"0",max:"89",frac:"0"},null,8,["label","onInput"]),r(S,{label:n(t)("food.fiber"),onInput:F.confirm,name:"fib",unit:"g",min:"0",max:"71",frac:"0"},null,8,["label","onInput"])]),e("fieldset",null,[e("legend",null,l(n(t)("aria.headmacro2")),1),r(S,{label:n(t)("food.fatsat"),onInput:F.confirm,name:"fatsat",unit:"g",min:"0",max:"83",frac:"0"},null,8,["label","onInput"]),r(S,{label:n(t)("food.fato3"),onInput:F.confirm,name:"fato3",unit:"g",min:"0",max:"54",frac:"0"},null,8,["label","onInput"]),r(S,{label:n(t)("food.fato6"),onInput:F.confirm,name:"fato6",unit:"g",min:"0",max:"70",frac:"0"},null,8,["label","onInput"]),r(S,{label:n(t)("food.sugar"),onInput:F.confirm,name:"sug",unit:"g",min:"0",max:"100",frac:"0"},null,8,["label","onInput"]),r(S,{label:n(t)("food.salt"),onInput:F.confirm,name:"salt",unit:"g",min:"0",max:"100",frac:"0"},null,8,["label","onInput"])])]),_:1},8,["data","placeholder"])])]),main:M(()=>[r(_e,{items:a.value,onSelected:h},null,8,["items"])]),_:2},[u.value?{name:"head-details",fn:M(()=>[e("h2",null,l(u.value.name),1)])}:void 0,u.value?{name:"details",fn:M(()=>[gn,e("section",wn,[$n,bn,yn,n(c).canCreateFood||n(c).canEditFood?(d(),p("button",{key:0,class:"icon async",disabled:v.value,onClick:f},[g.value?V("",!0):(d(),D(me,{key:0})),g.value?(d(),D(ae,{key:1})):V("",!0)],8,xn)):V("",!0)]),e("section",null,[e("h2",null,l(n(t)("aria.headnutrients")),1),e("form",{ref_key:"form",ref:x},[e("div",kn,[e("fieldset",{disabled:!g.value,class:"col50"},[e("div",null,[e("label",null,l(n(t)("food.energy")),1),e("input",{type:"text",value:u.value.kcal,name:"kcal",onChange:I},null,40,zn),e("span",In,l(n(t)("unit.cal")),1)]),e("div",null,[e("label",null,l(n(t)("food.fat")),1),e("input",{type:"text",value:u.value.fat,name:"fat",onChange:I},null,40,Fn),e("span",Ln,l(n(t)("unit.g")),1)]),e("div",null,[e("label",null,l(n(t)("food.carbs2")),1),e("input",{type:"text",value:u.value.carb,name:"carb",onChange:I},null,40,Mn),e("span",Tn,l(n(t)("unit.g")),1)]),e("div",null,[e("label",null,l(n(t)("food.protein")),1),e("input",{type:"text",value:u.value.prot,name:"prot",onChange:I},null,40,En),e("span",Sn,l(n(t)("unit.g")),1)]),e("div",null,[e("label",null,l(n(t)("food.fiber")),1),e("input",{type:"text",value:u.value.fib,name:"fib",onChange:I},null,40,Dn),e("span",Rn,l(n(t)("unit.g")),1)])],8,Cn),e("fieldset",{disabled:!g.value,class:"col50"},[e("div",null,[e("label",null,l(n(t)("food.fatsat")),1),e("input",{type:"text",value:u.value.fatsat,name:"fatsat",onChange:I},null,40,Vn),e("span",An,l(n(t)("unit.g")),1)]),e("div",null,[e("label",null,l(n(t)("food.fato3")),1),e("input",{type:"text",value:u.value.fato3,name:"fato3",onChange:I},null,40,Nn),e("span",jn,l(n(t)("unit.g")),1)]),e("div",null,[e("label",null,l(n(t)("food.fato6")),1),e("input",{type:"text",value:u.value.fato6,name:"fato6",onChange:I},null,40,On),e("span",Hn,l(n(t)("unit.g")),1)]),e("div",null,[e("label",null,l(n(t)("food.sugar")),1),e("input",{type:"text",value:u.value.sug,name:"sug",onChange:I},null,40,Xn),e("span",qn,l(n(t)("unit.g")),1)]),e("div",null,[e("label",null,l(n(t)("food.salt")),1),e("input",{type:"text",value:u.value.salt,name:"salt",onChange:I},null,40,Un),e("span",Pn,l(n(t)("unit.g")),1)])],8,Bn)])],512)])])}:void 0]),1536))}};const Kn={class:"ingredients"},Gn={disabled:""},Jn=["value"],Qn={class:"unit"},Yn=["value"],Zn={props:["items"],setup(s){const t=s,{t:i,locale:_}=H();B("food");const c=y("name"),o=new Intl.Collator(_.value,{numeric:!0}),a=ie(()=>t.items.map(g=>({id:g.id,amount:g.amount,name:i(g.id.toString())})).sort((g,v)=>o.compare(g[c.value],v[c.value])));return(u,g)=>(d(),p("form",Kn,[e("fieldset",Gn,[(d(!0),p(Q,null,ne(n(a),v=>(d(),p("div",{key:v.id},[e("label",null,l(v.name),1),e("input",{type:"text",name:"amount",value:v.amount},null,8,Jn),e("span",Qn,l(n(i)("unit.g")),1),e("input",{type:"hidden",name:"id",value:v.id},null,8,Yn)]))),128))])]))}},ea={},ta={width:"24",height:"24",version:"1.1",viewBox:"0 0 24 24",xmlns:"http://www.w3.org/2000/svg"},na=e("path",{d:"m0 18.316h2.5263v0.63158h-1.2632v1.2632h1.2632v0.63158h-2.5263v1.2632h3.7895v-5.0526h-3.7895zm1.2632-11.368h1.2632v-5.0526h-2.5263v1.2632h1.2632zm-1.2632 3.7895h2.2737l-2.2737 2.6526v1.1368h3.7895v-1.2632h-2.2737l2.2737-2.6526v-1.1368h-3.7895zm6.3158-7.5789v2.5263h17.684v-2.5263zm0 17.684h17.684v-2.5263h-17.684zm0-7.5789h17.684v-2.5263h-17.684z"},null,-1),aa=[na];function la(s,t){return d(),p("svg",ta,aa)}var oa=R(ea,[["render",la]]);const sa={class:"new-item"},ia=["disabled"],ca=["value"],ua=["innerHTML"],ra={class:"tags"},da=e("span",{class:"tag"},"Tag 1",-1),_a=e("span",{class:"tag"},"Tag 2",-1),ma=e("span",{class:"tag"},"Tag 3",-1),va=["disabled"],ha=z(" Add to diary here "),fa={class:"nutrient-block"},pa={class:"col50"},ga={class:"unit"},wa={class:"unit"},$a={class:"col50"},ba={class:"unit"},ya={class:"unit"},xa={class:"prep"},ka=["disabled"],Ca={class:"prep-size"},za=["value"],Ia=["disabled"],Fa=["value"],La={class:"unit"},Ma=["value"],Ta={class:"unit"},Ea=["value"],Sa={class:"unit"},Da=["value"],Ra={class:"unit"},Ba={class:"placeholder"},Va={setup(s){const{t}=H(),i=B("log"),_=B("csrfToken");B("perms");const c=B("recipes"),o=y([]),a=y(null),u=y(!1),g=y(!1),v=y("&nbsp;"),T=y(null),x=y(null);function k($,b=2){return+parseFloat($/(a.value.size||1)).toFixed(b)}function w($){return Math.min.apply(Math,c.value.map(b=>Math.floor(b[$]/b.size)))}function L($){return Math.max.apply(Math,c.value.map(b=>Math.ceil(b[$]/b.size)))}function h($){fetch("/api/v1/recipe",{method:"POST",headers:{"Content-Type":"application/x-www-form-urlencoded","X-CSRF-Token":_},body:new URLSearchParams({name:$})}).then(b=>{if(!b.ok)throw t("createrec.err"+b.status);return b.json()}).then(b=>{b.isowner=!0,c.value.push(b),o.value.push(b),i.msg(t("createrec.ok")),I(b.id)}).catch(b=>i.err(b))}function f(){g.value=!0;let $=a.value.id,b=a.value.owner,O=a.value.isowner;fetch("/api/v1/recipe/"+$,{method:"PUT",headers:{"Content-Type":"application/x-www-form-urlencoded","X-CSRF-Token":_},body:new URLSearchParams(new FormData(x.value))}).then(E=>{if(!E.ok)throw t("saverec.err"+E.status);return u.value=!1,fetch("/api/v1/recipe/"+$)}).then(E=>E.json()).then(E=>{E.owner=b,E.isowner=O,c.value=c.value.map(P=>E.id==P.id?E:P),o.value=o.value.map(P=>E.id==P.id?E:P),a.value=a.value.id==E.id?E:a.value,i.msg(t("saverec.ok"))}).catch(E=>i.err(E)).finally(()=>{setTimeout(function(){g.value=!1},500)})}function I($){a.value=o.value.filter(b=>b.id==$)[0],T.value.showDetails(),"isowner"in a.value?m():C()}function C(){fetch(`/api/v1/recipe/${a.value.id}/owner`).then($=>{if(!$.ok)throw $;return $.json()}).then($=>{a.value.isowner=$.isowner,a.value.owner=$.owner,m()}).catch(()=>{i.err(t("recowner.err")),v.value="&nbsp;"})}function m(){a.value.isowner?v.value=t("recipe.isowner"):a.value.owner?v.value=t("recipe.owner",{name:a.value.owner}):v.value=t("recipe.ispublic")}function F($){o.value=$,a.value&&o.value.filter(b=>b.id==a.value.id).length==0&&(a.value=null)}function Z(){u.value?f():u.value=!0}function j($){$.target.blur(),isNaN(parseFloat($.target.value))&&($.target.value=a.value[$.target.name])}return J(()=>o.value=[...c.value]),($,b)=>(d(),D(U,{ref_key:"main",ref:T,onDetailVisibility:b[0]||(b[0]=O=>u.value=!1)},ce({filter:M(()=>[e("section",sa,[e("h2",null,l(n(t)("aria.headnewrec")),1),r(de,{label:n(t)("btn.new"),placeholder:n(t)("recipe.hintnew"),onConfirm:h},null,8,["label","placeholder"])]),e("section",null,[e("h2",null,l(n(t)("aria.headsearch")),1),r(re,{data:n(c),placeholder:n(t)("recipe.hintsearch"),onResult:F},{default:M(O=>[e("fieldset",null,[r(S,{label:n(t)("food.energy"),min:w("kcal"),max:L("kcal"),onInput:O.confirm,name:"kcal",unit:"cal",frac:"0"},null,8,["label","min","max","onInput"]),r(S,{label:n(t)("food.fat"),min:w("fat"),max:L("fat"),onInput:O.confirm,name:"fat",unit:"g",frac:"0"},null,8,["label","min","max","onInput"]),r(S,{label:n(t)("food.carbs"),min:w("carb"),max:L("carb"),onInput:O.confirm,name:"carb",unit:"g",frac:"0"},null,8,["label","min","max","onInput"]),r(S,{label:n(t)("food.protein"),min:w("prot"),max:L("prot"),onInput:O.confirm,name:"prot",unit:"g",frac:"0"},null,8,["label","min","max","onInput"])])]),_:1},8,["data","placeholder"])])]),main:M(()=>[r(_e,{items:o.value,onSelected:I},null,8,["items"])]),_:2},[a.value?{name:"head-details",fn:M(()=>[e("form",{ref_key:"form",ref:x,autocomplete:"off",id:"form-recipe"},[e("fieldset",{disabled:!u.value},[e("input",{type:"text",name:"name",value:a.value.name},null,8,ca)],8,ia)],512)])}:void 0,a.value?{name:"details",fn:M(()=>[e("section",{class:"subtitle",innerHTML:v.value},null,8,ua),e("section",ra,[da,_a,ma,a.value.isowner?(d(),p("button",{key:0,class:"icon async",disabled:g.value,onClick:Z},[u.value?V("",!0):(d(),D(me,{key:0})),u.value?(d(),D(ae,{key:1})):V("",!0)],8,va)):V("",!0)]),e("section",null,[e("h2",null,l(n(t)("aria.headtrack")),1),ha]),e("section",null,[e("h2",null,l(n(t)("aria.headingred")),1),r(Zn,{items:a.value.items},null,8,["items"])]),e("section",null,[e("h2",null,l(n(t)("aria.headnutrients")),1),e("div",fa,[e("div",pa,[e("div",null,[e("label",null,l(n(t)("food.energy")),1),e("span",null,l(k(a.value.kcal,1)),1),e("span",ga,l(n(t)("unit.cal")),1)]),e("div",null,[e("label",null,l(n(t)("food.fat")),1),e("span",null,l(k(a.value.fat,1)),1),e("span",wa,l(n(t)("unit.g")),1)])]),e("div",$a,[e("div",null,[e("label",null,l(n(t)("food.carbs2")),1),e("span",null,l(k(a.value.carb,1)),1),e("span",ba,l(n(t)("unit.g")),1)]),e("div",null,[e("label",null,l(n(t)("food.protein")),1),e("span",null,l(k(a.value.prot,1)),1),e("span",ya,l(n(t)("unit.g")),1)])])])]),e("section",xa,[e("h2",null,l(n(t)("aria.headprep")),1),e("div",null,[e("fieldset",{disabled:!u.value,class:"col50"},[e("div",Ca,[e("label",null,l(n(t)("recipe.size",2)),1),e("input",{type:"text",name:"size",form:"form-recipe",value:a.value.size,onChange:j},null,40,za),e("label",null,l(n(t)("recipe.size",a.value.size)),1)])],8,ka),e("fieldset",{disabled:!u.value,class:"col50"},[e("div",null,[e("label",null,l(n(t)("recipe.time")),1),e("input",{type:"text",disabled:"",value:a.value.preptime+a.value.cooktime+a.value.misctime},null,8,Fa),e("span",La,l(n(t)("unit.min")),1)]),e("div",null,[e("label",null,l(n(t)("recipe.preptime")),1),e("input",{type:"text",name:"preptime",form:"form-recipe",value:a.value.preptime,onChange:j},null,40,Ma),e("span",Ta,l(n(t)("unit.min")),1)]),e("div",null,[e("label",null,l(n(t)("recipe.cooktime")),1),e("input",{type:"text",name:"cooktime",form:"form-recipe",value:a.value.cooktime,onChange:j},null,40,Ea),e("span",Sa,l(n(t)("unit.min")),1)]),e("div",null,[e("label",null,l(n(t)("recipe.misctime")),1),e("input",{type:"text",name:"misctime",form:"form-recipe",value:a.value.misctime,onChange:j},null,40,Da),e("span",Ra,l(n(t)("unit.min")),1)])],8,Ia)]),e("div",Ba,[r(oa),e("p",null,l(n(t)("todo.instructions")),1)])])])}:void 0]),1536))}},Aa=z(" Diary "),Na={setup(s){return(t,i)=>(d(),D(U,null,{main:M(()=>[Aa]),_:1}))}},ja=z(" Shopping Lists "),Oa={setup(s){return(t,i)=>(d(),D(U,null,{main:M(()=>[ja]),_:1}))}},Ha=z(" Profile "),Xa={setup(s){return(t,i)=>(d(),D(U,null,{main:M(()=>[Ha]),_:1}))}},qa=z(" Settings "),Ua={setup(s){return(t,i)=>(d(),D(U,null,{main:M(()=>[qa]),_:1}))}},Pa=pe({history:ge(),routes:[{path:"/",name:"food",component:Wn},{path:"/recipes",name:"recipes",component:Va},{path:"/diary",name:"diary",component:Na},{path:"/shopping",name:"shopping",component:Oa},{path:"/profile",name:"profile",component:Xa},{path:"/settings",name:"settings",component:Ua}]});let te=document.documentElement.lang||navigator.language;!te&&navigator.languages!=null&&(te=navigator.languages[0]);const se=document.querySelector("meta[name='_csrf']"),Wa=se?se.content:"",Ka=function(){const s=document.documentElement.dataset.perm||1,t=65536,i=131072;function _(c){return(s&c)==c}return{canCreateFood:_(t),canEditFood:_(i)}}(),Ga=function(){function s(t){return typeof t=="string"?t:"message"in t?t.message:q.global.t("err.err")}return{msg:function(t){let i={msg:s(t),timeout:3e3};window.dispatchEvent(new CustomEvent("message",{detail:i}))},warn:function(t){let i={msg:s(t),timeout:4e3};window.dispatchEvent(new CustomEvent("warning",{detail:i}))},err:function(t){let i={msg:s(t),timeout:5e3};window.dispatchEvent(new CustomEvent("error",{detail:i}))}}}(),A=we(ot);let q,G,Y;function Ja(s){q=$e({legacy:!1,locale:te.split("-")[0],fallbackLocale:"en",messages:s}),G&&Y&&le()}function Qa(s){G=s,q&&Y&&le()}function Ya(s){Y=s,q&&G&&le()}function le(){G.forEach(s=>{s.name=q.global.t(s.id.toString())}),A.provide("csrfToken",Wa),A.provide("perms",Ka),A.provide("food",y(G)),A.provide("recipes",y(Y)),A.provide("log",Ga),A.use(Pa),A.use(q),A.mount("#app")}fetch("/app/l10n.json").then(s=>s.json()).then(Ja);fetch("api/v1/foods").then(s=>s.json()).then(Qa);fetch("api/v1/recipes").then(s=>s.json()).then(Ya);
