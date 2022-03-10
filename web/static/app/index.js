import{o as h,c as g,a as e,u as U,i as x,b as m,w as k,d as y,t as n,e as T,R as N,f as W,g as E,h as D,n as V,r as w,F as K,j as Q,k as ue,l as X,m as G,p as re,v as ce,q as ne,s as ae,x as de,y as ve,z as he,A as me}from"./vendor.js";const _e=function(){const l=document.createElement("link").relList;if(l&&l.supports&&l.supports("modulepreload"))return;for(const s of document.querySelectorAll('link[rel="modulepreload"]'))c(s);new MutationObserver(s=>{for(const i of s)if(i.type==="childList")for(const a of i.addedNodes)a.tagName==="LINK"&&a.rel==="modulepreload"&&c(a)}).observe(document,{childList:!0,subtree:!0});function u(s){const i={};return s.integrity&&(i.integrity=s.integrity),s.referrerpolicy&&(i.referrerPolicy=s.referrerpolicy),s.crossorigin==="use-credentials"?i.credentials="include":s.crossorigin==="anonymous"?i.credentials="omit":i.credentials="same-origin",i}function c(s){if(s.ep)return;s.ep=!0;const i=u(s);fetch(s.href,i)}};_e();var M=(o,l)=>{const u=o.__vccOpts||o;for(const[c,s]of l)u[c]=s;return u};const fe={},pe={width:"512",height:"512",version:"1.1",viewBox:"0 0 512 512",xmlns:"http://www.w3.org/2000/svg"},ge=e("path",{d:"m114.6 507.3c13.82-0.8729 90.06-3.083 106.2-3.079h13.49l2.639-5.293c5.191-10.41 10.35-29.57 13.13-48.74 1.811-12.5 1.488-16.68-1.826-23.56-2.528-5.249-11.04-13.06-20.54-18.86-28.38-17.32-51.21-34.77-76.05-58.12-6.673-6.272-9.377-8.209-16.08-11.53-12.37-6.124-20.36-12.88-29.19-24.7-1.45-1.941-1.795-2.038-4.695-1.334-1.722 0.4182-7.18 0.8166-12.13 0.8854-19.6 0.2722-38.18-7.715-52.49-22.56-15.97-16.57-23.43-35.49-23.41-59.37 0.01786-16.86 3.547-29.88 11.93-44 9.617-16.19 27.39-30.13 43.89-34.41 2.517-0.653 4.748-1.36 4.958-1.571 0.2103-0.2107-0.05169-3.078-0.5816-6.371-1.41-8.759-0.556-21.84 1.997-30.61 7.023-24.1 24.96-41.95 48.52-48.31 11.42-3.082 29.45-2.167 40.42 2.049 0.919 0.3533 1.637-0.7202 2.994-4.474 6.21-17.18 18.79-33.75 33.11-43.62 35.65-24.56 80.99-19.9 111.3 11.44l5.539 5.719 5.025-1.229c2.764-0.6762 9.075-1.428 14.02-1.672 16.17-0.7944 31.64 3.865 45.31 13.64l5.1 3.647 12.55 0.1171c9.984 0.09312 13.77 0.4491 18.49 1.739 29.1 7.945 50.37 32.58 56 64.86 1.354 7.764 1.096 21.6-0.5502 29.5l-0.916 4.394 5.367 7.086c20.1 26.54 25.52 61.05 14.68 93.52-7.862 23.55-24.99 43.53-46.38 54.11-10.16 5.023-23.1 8.587-31.28 8.613-3.79 0.0118-4.555 0.333-8.217 3.446-6.958 5.916-20.59 13.71-28.3 16.17-1.722 0.551-4.011 1.599-5.087 2.328-1.076 0.7297-8.119 4.802-15.65 9.05-18.57 10.47-37.53 22.69-48.72 31.39-13.27 10.32-16.83 22.63-14.72 50.94 2.648 35.64 8.059 59.41 16.33 71.75l2.89 4.313 29.4 0.4963c34.14 0.5764 92.53 2.291 93.01 2.731 0.1793 0.166-74.3 0.2503-165.5 0.1877-91.21-0.0631-161.4-0.392-156-0.732zm130.9-101.8c-0.1092-1.007-3.147-5.947-6.75-10.98-7.889-11.01-19.71-28.99-22.37-34.01-1.594-3.014-2.402-3.727-4.695-4.139-7.259-1.304-17.91-5.324-27.58-10.41-1.809-0.9509 5.669 20.15 9.032 25.49 4.859 7.714 27.69 24.94 45.08 34 5.177 2.7 7.575 2.715 7.285 0.0442zm6.168-19.46c0.8235-0.9976 1.074-4.638 1.074-15.62 0-7.875-0.1764-14.5-0.3923-14.71-0.2158-0.2163-2.307 0.1044-4.647 0.7128-2.34 0.6082-7 1.335-10.35 1.616-7.067 0.5907-6.896 0-2.984 10.28 2.73 7.17 6.328 12.81 10.4 16.3 3.513 3.013 5.293 3.38 6.908 1.424zm49.91-15.13c3.913-2 11.37-9.412 15.96-15.85 3.982-5.592 8.688-15 8.214-16.42-0.1628-0.49-3.275-1.75-6.916-2.799-3.641-1.049-9.568-3.501-13.17-5.449l-6.551-3.541-2.173 2.087c-2.148 2.062-2.179 2.234-2.645 14.37-0.2592 6.754-0.6282 15.17-0.8198 18.7-0.3026 5.572-0.1405 6.682 1.232 8.43 1.896 2.415 2.946 2.489 6.867 0.484zm-118.1-5.725c-2.19-4.729-6.256-17.66-6.256-19.89 0-0.8683-1.901-1.089-9.391-1.089-5.165 0-9.391 0.2195-9.391 0.488 0 1.561 25.33 25.39 26.99 25.39 0.1762 0-0.7011-2.206-1.949-4.901zm149.8-9.536c5.487-3.687 19.46-14.65 19.07-14.96-0.126-0.1032-4.197-0.3888-9.046-0.6338l-8.816-0.4452-3.967 5.891c-4.284 6.363-7.088 11.99-6.473 12.98 0.9066 1.47 4.434 0.3864 9.228-2.836z",fill:"#13ad73"},null,-1),$e=e("path",{d:"m253.9 167.7c-0.2579 0.0513-1.238 0.5454-4.691 0.5216-3.625-0.0248 4.941 9.227 6.934 34.4-8.031-0.4678-8.408-3.285-24.81-5.999-15.88-2.627-25.28 5.883-29.57 18.88-3.429 10.39-3.592 23.66-1.19 36.39 2.348 12.45 7.148 24.39 13.74 32.65 6.909 8.654 14.05 10.96 20.49 10.82 8.459-0.1842 15.7-4.599 19.58-4.407 6.726 0.3331 13.29 8.894 28.17 2.151 6.047-2.741 12.03-7.819 17.08-14.81 4.984-6.907 9.051-15.68 11.35-25.91 2.018-8.972 2.678-19.07 1.41-30-1.616-13.93-7.519-20.42-14.9-23.01-8.286-2.905-18.43-0.8896-26.46 1.032-4.71 1.127-8.69 2.223-11.14 2.273-0.155 3e-3 -0.2829-4.1e-4 -0.4335 2e-3 -0.01-0.0695-0.0197-0.1385-0.0323-0.2078-2.249-11.58-4.704-29.29-5.401-33.64-0.152-0.9486 0.0845-1.177-0.1162-1.137z",fill:"#f2ac05","fill-rule":"evenodd"},null,-1),we=[ge,$e];function be(o,l){return h(),g("svg",pe,we)}var ye=M(fe,[["render",be]]);const ke={},Ce={width:"24",height:"24",version:"1.1",viewBox:"0 0 24 24",xmlns:"http://www.w3.org/2000/svg"},Ie=e("path",{d:"m12 0c-6.624 0-12 5.376-12 12s5.376 12 12 12 12-5.376 12-12-5.376-12-12-12zm0 3.6c1.992 0 3.6 1.608 3.6 3.6s-1.608 3.6-3.6 3.6-3.6-1.608-3.6-3.6 1.608-3.6 3.6-3.6zm0 17.04c-3 0-5.652-1.536-7.2-3.864 0.036-2.388 4.8-3.696 7.2-3.696 2.388 0 7.164 1.308 7.2 3.696-1.548 2.328-4.2 3.864-7.2 3.864z"},null,-1),ze=[Ie];function Le(o,l){return h(),g("svg",Ce,ze)}var Fe=M(ke,[["render",Le]]);const Te={},Ee={width:"24",height:"24",version:"1.1",viewBox:"0 0 24 24",xmlns:"http://www.w3.org/2000/svg"},Me=e("path",{d:"m12 0-2.115 2.115 8.37 8.385h-18.255v3h18.255l-8.37 8.385 2.115 2.115 12-12z"},null,-1),Se=[Me];function xe(o,l){return h(),g("svg",Ee,Se)}var le=M(Te,[["render",xe]]);const De=e("div",{id:"app-name"},[e("span",null,"Hey"),e("span",null,"Apple")],-1),Re={id:"nav-main"},Be={id:"nav-user"},Ve={href:"https://docs.heyapple.org",target:"_blank"},Ne={setup(o){const{t:l}=U(),u=x("csrfToken");function c(i){i.preventDefault(),fetch("/auth/local",{method:"DELETE",headers:{"X-CSRF-Token":u}}).then(a=>{a.ok?window.location="/":window.dispatchEvent(new CustomEvent("error",{detail:{msg:l("signout.err"+a.status)}}))})}function s(i){i.stopPropagation(),document.querySelector("header nav").classList.toggle("open")}return document.addEventListener("click",function(){document.querySelector("header nav").classList.remove("open")}),(i,a)=>(h(),g("header",null,[m(ye,{id:"logo"}),De,e("nav",null,[e("button",{onClick:s},[m(le)]),e("ul",Re,[e("li",null,[m(T(N),{to:"/"},{default:k(()=>[y(n(i.$t("nav.food")),1)]),_:1})]),e("li",null,[m(T(N),{to:"/recipes"},{default:k(()=>[y(n(i.$t("nav.recipes")),1)]),_:1})]),e("li",null,[m(T(N),{to:"/diary"},{default:k(()=>[y(n(i.$t("nav.diary")),1)]),_:1})]),e("li",null,[m(T(N),{to:"/shopping"},{default:k(()=>[y(n(i.$t("nav.shopping")),1)]),_:1})])]),e("ul",Be,[e("li",null,[m(T(N),{to:"/profile"},{default:k(()=>[y(n(i.$t("nav.profile")),1)]),_:1})]),e("li",null,[m(T(N),{to:"/settings"},{default:k(()=>[y(n(i.$t("nav.settings")),1)]),_:1})]),e("li",null,[e("a",Ve,n(i.$t("nav.help")),1)]),e("li",null,[e("a",{href:"#",onClick:c},n(i.$t("nav.signout")),1)])])]),e("button",{onClick:s},[m(Fe)])]))}},Ae={},je={width:"24",height:"24",version:"1.1",viewBox:"0 0 24 24",xmlns:"http://www.w3.org/2000/svg"},Oe=e("path",{d:"m7.6364 17.318-5.7273-5.7273-1.9091 1.9091 7.6364 7.6364 16.364-16.364-1.9091-1.9091z"},null,-1),He=[Oe];function Xe(o,l){return h(),g("svg",je,He)}var Y=M(Ae,[["render",Xe]]);const qe={},Pe={width:"24",height:"24",version:"1.1",viewBox:"0 0 24 24",xmlns:"http://www.w3.org/2000/svg"},Ue=e("path",{d:"m12 1.6364-12 20.727h24zm0 4.353 8.2159 14.192h-16.43zm-1.0909 4.3743v5.4545h2.1818v-5.4545zm0 6.5455v2.1818h2.1818v-2.1818z"},null,-1),We=[Ue];function Ke(o,l){return h(),g("svg",Pe,We)}var Ge=M(qe,[["render",Ke]]);const Je={props:["msg"],emits:["timeout"],setup(o,{emit:l}){const u=o;return W(()=>{setTimeout(function(){l("timeout",u.msg.id)},u.msg.time)}),(c,s)=>(h(),g("div",{class:V(["message",[o.msg.type,o.msg.id]])},[o.msg.type=="message"?(h(),E(Y,{key:0})):D("",!0),o.msg.type!="message"?(h(),E(Ge,{key:1})):D("",!0),e("p",null,n(o.msg.msg),1)],2))}};const Qe={id:"messages"},Ye={setup(o){const l=w([]);let u=0;function c(i){l.value.push({id:u++,type:i.type,msg:i.detail.msg,time:i.detail.timeout})}function s(i){l.value=l.value.filter(a=>a.id!=i)}return W(()=>{window.addEventListener("message",c),window.addEventListener("warning",c),window.addEventListener("error",c)}),(i,a)=>(h(),g("div",Qe,[(h(!0),g(K,null,Q(l.value,d=>(h(),E(Je,{key:d.id,msg:d,onTimeout:s},null,8,["msg"]))),128))]))}};const Ze={setup(o){return W(()=>{document.querySelector("body > .spinner-container").remove()}),(l,u)=>(h(),g(K,null,[m(Ne),m(T(ue)),m(Ye)],64))}},et={},tt={version:"1.1",viewBox:"0 0 24 24",xmlns:"http://www.w3.org/2000/svg"},nt=e("path",{id:"path-top",d:"m0 3v2h24v-2z"},null,-1),at=e("path",{id:"path-mid",d:"m0 13h24v-2h-24z"},null,-1),lt=e("path",{id:"path-bottom",d:"m0 21h24v-2h-24z"},null,-1),ot=[nt,at,lt];function st(o,l){return h(),g("svg",tt,ot)}var it=M(et,[["render",st]]);const ut={},rt={width:"512",height:"512",version:"1.1",viewBox:"0 0 512 512",xmlns:"http://www.w3.org/2000/svg"},ct=e("path",{d:"m251.2 4.007c-17.24-0.2277-34.69 4.924-50.29 15.67-14.33 9.866-26.9 26.43-33.11 43.62-1.357 3.754-2.073 4.828-2.992 4.475-10.97-4.217-29-5.131-40.42-2.049-23.57 6.359-41.5 24.22-48.52 48.31-2.553 8.761-3.407 21.84-1.997 30.6 0.5299 3.293 0.7902 6.162 0.5799 6.372-0.2104 0.2108-2.439 0.9186-4.956 1.572-16.5 4.283-34.28 18.21-43.89 34.41-8.382 14.11-11.91 27.14-11.93 44-0.02401 23.88 7.433 42.79 23.41 59.36 14.31 14.85 32.89 22.84 52.49 22.56 4.95-0.0689 10.41-0.4681 12.13-0.8863 2.9-0.7045 3.245-0.6061 4.695 1.334 8.833 11.82 16.82 18.58 29.19 24.7 6.707 3.32 9.412 5.257 16.09 11.53 24.85 23.35 47.68 40.8 76.05 58.12 9.503 5.8 18.01 13.61 20.54 18.86 3.313 6.878 3.637 11.06 1.825 23.56-2.778 19.18-7.936 38.33-13.13 48.74l-2.639 5.295h-13.49c-16.17-3e-3 -92.41 2.205-106.2 3.077l0.01-3e-3c-5.381 0.3399 64.84 0.6684 156 0.7315 91.2 0.0625 165.7-0.0216 165.5-0.1878-0.4753-0.4399-58.86-2.155-93-2.731l-29.4-0.4942-2.89-4.313c-8.274-12.35-13.68-36.12-16.33-71.76-2.105-28.31 1.454-40.62 14.72-50.94 11.18-8.699 30.14-20.92 48.72-31.39 7.533-4.248 14.57-8.318 15.65-9.048 1.076-0.7296 3.366-1.779 5.087-2.33 7.71-2.467 21.34-10.26 28.3-16.17 3.662-3.113 4.427-3.435 8.218-3.446 8.183-0.0267 21.12-3.59 31.28-8.613 21.39-10.58 38.52-30.56 46.38-54.11 10.84-32.47 5.42-66.98-14.68-93.52l-5.364-7.084 0.916-4.395c1.646-7.896 1.904-21.74 0.5502-29.5-5.632-32.29-26.9-56.92-56-64.86-4.726-1.29-8.51-1.647-18.49-1.74l-12.55-0.1186-5.101-3.647c-13.67-9.776-29.15-14.43-45.31-13.64-4.95 0.2432-11.26 0.9944-14.02 1.671l-5.028 1.229-5.539-5.72c-17.07-17.63-38.89-26.81-61.05-27.1zm2.728 163.7c0.2008-0.0398-0.0367 0.1881 0.1154 1.137 0.6972 4.35 3.151 22.06 5.4 33.64 0.0118 0.0694 0.0228 0.138 0.0328 0.2076 0.1506-2e-3 0.2799 9e-5 0.4349-3e-3 2.448-0.0501 6.427-1.143 11.14-2.27 8.028-1.922 18.18-3.94 26.46-1.035 7.379 2.587 13.28 9.078 14.9 23.01 1.268 10.94 0.6106 21.03-1.407 30-2.301 10.23-6.37 19.01-11.35 25.91-5.045 6.991-11.03 12.07-17.07 14.81-14.87 6.743-21.44-1.818-28.16-2.152-3.88-0.1922-11.12 4.224-19.58 4.409-6.437 0.1402-13.58-2.169-20.49-10.82-6.596-8.262-11.39-20.2-13.74-32.65-2.402-12.73-2.24-26 1.189-36.39 4.29-13 13.69-21.51 29.58-18.88 16.41 2.714 16.78 5.532 24.81 6-1.993-25.18-10.56-34.43-6.936-34.4 3.453 0.0238 4.434-0.4726 4.692-0.5239zm45.16 159.1 6.55 3.542c3.603 1.947 9.529 4.401 13.17 5.45 3.641 1.049 6.753 2.307 6.916 2.797 0.4733 1.423-4.232 10.83-8.214 16.42-4.588 6.442-12.05 13.85-15.96 15.86-3.919 2.004-4.971 1.93-6.867-0.4843-1.372-1.748-1.532-2.86-1.229-8.432 0.1916-3.529 0.558-11.94 0.8171-18.7 0.4657-12.13 0.4983-12.3 2.646-14.37zm35.38 12.77 8.817 0.4448c4.849 0.245 8.919 0.5294 9.045 0.6326 0.3835 0.315-13.58 11.27-19.07 14.96-4.795 3.222-8.322 4.307-9.229 2.837-0.6151-0.9971 2.187-6.619 6.471-12.98zm-166.7 4.596c7.489 0 9.391 0.219 9.391 1.087 0 2.234 4.067 15.16 6.257 19.89 1.248 2.694 2.123 4.9 1.947 4.9-1.66 0-26.99-23.83-26.99-25.39 0-0.2685 4.226-0.4876 9.391-0.4876zm16.13 1.727c0.0388-0.0118 0.0851-6e-3 0.1417 0.0231 9.674 5.087 20.32 9.107 27.58 10.41 2.293 0.4121 3.101 1.124 4.695 4.138 2.657 5.026 14.48 23 22.37 34.02 3.603 5.03 6.639 9.968 6.748 10.98 0.2897 2.671-2.108 2.658-7.285-0.0429-17.39-9.068-40.22-26.29-45.07-34.01-3.258-5.172-10.38-25.14-9.173-25.52zm67.9 9.71c0.263-9e-3 0.4337 9e-3 0.4876 0.0625 0.2159 0.2154 0.3921 6.837 0.3921 14.71 0 10.98-0.2506 14.62-1.074 15.62-1.615 1.956-3.397 1.59-6.909-1.423-4.068-3.49-7.665-9.13-10.4-16.3-3.912-10.28-4.082-9.686 2.985-10.28 3.355-0.2804 8.013-1.006 10.35-1.615 1.755-0.4562 3.373-0.7502 4.162-0.7776z"},null,-1),dt=[ct];function vt(o,l){return h(),g("svg",rt,dt)}var ht=M(ut,[["render",vt]]);const mt={},_t={width:"24",height:"24",version:"1.1",viewBox:"0 0 24 24",xmlns:"http://www.w3.org/2000/svg"},ft=e("path",{d:"m12 6c1.65 0 3-1.35 3-3s-1.35-3-3-3-3 1.35-3 3 1.35 3 3 3zm0 3c-1.65 0-3 1.35-3 3s1.35 3 3 3 3-1.35 3-3-1.35-3-3-3zm0 9c-1.65 0-3 1.35-3 3s1.35 3 3 3 3-1.35 3-3-1.35-3-3-3z"},null,-1),pt=[ft];function gt(o,l){return h(),g("svg",_t,pt)}var $t=M(mt,[["render",gt]]);const wt={id:"filter"},bt=y(" This is the main search & filter area"),yt={id:"main"},kt={class:"controls"},Ct=e("span",{class:"spacer"},null,-1),It={class:"content"},zt=y("This is the main area"),Lt={id:"details"},Ft={class:"controls"},Tt=e("span",{class:"spacer"},null,-1),Et={class:"placeholder"},j={emits:["detailVisibility"],setup(o,{expose:l,emit:u}){const c=w(""),s=w("");function i(){c.value==""?(c.value="open-filter",s.value="",u("detailVisibility")):c.value=""}function a(){u("detailVisibility"),s.value==""?(s.value="open-details",c.value=""):s.value=""}function d(){u("detailVisibility"),s.value==""&&(s.value="open-details",c.value="")}return l({showDetails:d}),(v,p)=>(h(),g("main",{class:V([c.value,s.value])},[e("div",wt,[X(v.$slots,"filter",{},()=>[bt])]),e("div",yt,[e("div",kt,[e("button",{onClick:i,class:"open-filter icon"},[m(it)]),Ct,e("button",{onClick:a,class:"open-details icon"},[m($t)])]),e("div",It,[X(v.$slots,"main",{},()=>[zt])])]),e("div",Lt,[e("div",Ft,[X(v.$slots,"head-details"),Tt,e("button",{onClick:a,class:"open-details icon"},[m(le)])]),X(v.$slots,"details",{},()=>[e("div",Et,[m(ht),e("p",null,n(v.$t("details.noitem")),1)])])])],2))}},Mt=["placeholder"],St=y("Additional filters"),xt={props:["data","placeholder"],emits:["result"],setup(o,{emit:l}){const u=o;var c=void 0;function s(a){a.preventDefault(),clearTimeout(c),c=setTimeout(function(){i(a.target.closest("form"))},500)}function i(a){let d=new FormData(a),v=u.data.filter(p=>{for(let b of d.keys()){if(b=="name"){let f=d.get(b).toLowerCase();if(!p[b].toLowerCase().includes(f))return!1;continue}let[C,$]=d.getAll(b).map(f=>parseFloat(f));if(!isNaN(C)&&!isNaN($)&&(p[b]<C||$<p[b]))return!1}return!0});l("result",v)}return(a,d)=>(h(),g("form",null,[e("input",{type:"text",name:"name",placeholder:o.placeholder,onInput:s},null,40,Mt),X(a.$slots,"default",{confirm:s},()=>[St])]))}};const Dt={class:"slider"},Rt=["name","value"],Bt=y(" \xA0\u2013\xA0 "),Vt=["name","value"],Nt={class:"bar"},At={class:"slide"},jt=e("div",{class:"interact-area"},null,-1),Ot=[jt],Ht=e("div",{class:"interact-area"},null,-1),Xt=[Ht],S={props:["label","name","unit","min","max","frac"],emits:["input"],setup(o,{emit:l}){const u=o,c=w(parseFloat(u.min).toFixed(u.frac)),s=w(parseFloat(u.max).toFixed(u.frac)),i=w(0),a=w(100);function d($){$.target.blur();let f=parseFloat(u.min)||0,L=parseFloat($.target.value)||f;c.value=L,s.value=Math.max(c.value,s.value);let F=parseFloat(u.max)||0;i.value=(L-f)*100/(F-f),a.value=Math.max(i.value,a.value),l("input",$)}function v($){$.target.blur();let f=parseFloat(u.max)||0,L=parseFloat($.target.value)||f;s.value=L,c.value=Math.min(c.value,s.value);let F=parseFloat(u.min)||0;a.value=(L-F)*100/(f-F),i.value=Math.min(i.value,a.value),l("input",$)}function p($){let L=$.target.closest(".slide").getBoundingClientRect(),F=$.pageX!==void 0?$.pageX:$.changedTouches[0].pageX;F=Math.min(Math.max(F-L.left,0),L.width);let R=F*100/L.width,I=parseFloat(u.min)||0,r=parseFloat(u.max)||0,t=R/100*(r-I)+I;$.target.closest("button").classList.contains("min")?(c.value=t.toFixed(u.frac),s.value=Math.max(c.value,s.value),i.value=R,a.value=Math.max(i.value,a.value)):(s.value=t.toFixed(u.frac),c.value=Math.min(c.value,s.value),a.value=R,i.value=Math.min(i.value,a.value)),l("input",$)}function b($){let f=$.target.closest("button");f.addEventListener("mousemove",p),f.addEventListener("touchmove",p),f.addEventListener("mouseup",C),f.addEventListener("touchend",C),f.addEventListener("mouseleave",C),f.addEventListener("touchcancel",C),f.classList.add("active")}function C($){let f=$.target.closest("button");f.removeEventListener("mousemove",p),f.removeEventListener("touchmove",p),f.removeEventListener("mouseup",C),f.removeEventListener("touchend",C),f.removeEventListener("mouseleave",C),f.removeEventListener("touchcancel",C),f.classList.remove("active")}return($,f)=>(h(),g("div",Dt,[e("label",null,[e("span",null,n(o.label)+" ("+n($.$t("unit."+o.unit))+")",1),e("input",{type:"text",name:o.name,value:c.value,onChange:d},null,40,Rt),Bt,e("input",{type:"text",name:o.name,value:s.value,onChange:v},null,40,Vt)]),e("div",Nt,[e("div",At,[e("div",{class:"overlay",style:G({left:i.value+"%",right:100-a.value+"%"})},null,4),e("button",{type:"button",class:"handle min",style:G({left:i.value+"%"}),onMousedown:b,onTouchstart:b},Ot,36),e("button",{type:"button",class:"handle max",style:G({left:a.value+"%"}),onMousedown:b,onTouchstart:b},Xt,36)])])]))}};const qt={class:"clickable-input"},Pt=["placeholder"],Ut=["value"],oe={props:["label","placeholder"],emits:["confirm"],setup(o,{emit:l}){const u=w("");function c(s){s.preventDefault(),l("confirm",u.value),u.value=""}return(s,i)=>(h(),g("form",qt,[re(e("input",{type:"text","onUpdate:modelValue":i[0]||(i[0]=a=>u.value=a),placeholder:o.placeholder},null,8,Pt),[[ce,u.value]]),e("input",{type:"submit",onClick:c,value:o.label},null,8,Ut)]))}};const Wt={},Kt={class:"icon sort-arrow"};function Gt(o,l){return h(),g("span",Kt)}var H=M(Wt,[["render",Gt]]);const Jt=["onClick"],Qt={class:"name"},Yt={class:"num"},Zt={class:"unit"},en={class:"m num"},tn={class:"unit"},nn={class:"m num"},an={class:"unit"},ln={class:"m num"},on={class:"unit"},se={props:["items"],emits:"selected",setup(o,{emit:l}){const u=o,c=w("name"),s=w("asc"),i=new Intl.Collator(U().locale.value,{numeric:!0}),a=ne(()=>s.value=="asc"?[...u.items].sort((v,p)=>i.compare(v[c.value],p[c.value])):[...u.items].sort((v,p)=>-i.compare(v[c.value],p[c.value])));function d(v){let p=v.target.dataset.sort;c.value==p?s.value=s.value=="asc"?"desc":"asc":c.value=p}return(v,p)=>(h(),g("table",null,[e("thead",null,[e("tr",{class:V(s.value)},[e("th",{class:V(["name sort",{active:c.value=="name"}]),onClick:d,"data-sort":"name"},[y(n(v.$t("food.name"))+" ",1),m(H)],2),e("th",{class:V(["num sort",{active:c.value=="kcal"}]),onClick:d,"data-sort":"kcal"},[m(H),y(" "+n(v.$t("food.energy")),1)],2),e("th",{class:V(["m num sort",{active:c.value=="fat"}]),onClick:d,"data-sort":"fat"},[m(H),y(" "+n(v.$t("food.fat")),1)],2),e("th",{class:V(["m num sort",{active:c.value=="carb"}]),onClick:d,"data-sort":"carb"},[m(H),y(" "+n(v.$t("food.carbs2")),1)],2),e("th",{class:V(["m num sort",{active:c.value=="prot"}]),onClick:d,"data-sort":"prot"},[m(H),y(" "+n(v.$t("food.protein")),1)],2)],2)]),e("tbody",null,[(h(!0),g(K,null,Q(T(a),b=>(h(),g("tr",{key:b.id,onClick:C=>v.$emit("selected",b.id)},[e("td",Qt,n(b.name),1),e("td",Yt,[y(n(b.kcal)+" ",1),e("span",Zt,n(v.$t("unit.cal")),1)]),e("td",en,[y(n(b.fat)+" ",1),e("span",tn,n(v.$t("unit.g")),1)]),e("td",nn,[y(n(b.carb)+" ",1),e("span",an,n(v.$t("unit.g")),1)]),e("td",ln,[y(n(b.prot)+" ",1),e("span",on,n(v.$t("unit.g")),1)])],8,Jt))),128))])]))}},sn={},un={width:"24",height:"24",version:"1.1",viewBox:"0 0 24 24",xmlns:"http://www.w3.org/2000/svg"},rn=e("path",{d:"m-7.5e-8 19.001v4.9993h4.9993l14.745-14.745-4.9993-4.9993zm23.61-13.611c0.51993-0.51993 0.51993-1.3598 0-1.8797l-3.1196-3.1196c-0.51993-0.51993-1.3598-0.51993-1.8797 0l-2.4397 2.4397 4.9993 4.9993z"},null,-1),cn=[rn];function dn(o,l){return h(),g("svg",un,cn)}var ie=M(sn,[["render",dn]]);const vn={key:0,class:"new-item"},hn=e("section",{class:"subtitle"},"Some food category",-1),mn={class:"tags"},_n=e("span",{class:"tag"},"Tag 1",-1),fn=e("span",{class:"tag"},"Tag 2",-1),pn=e("span",{class:"tag"},"Tag 3",-1),gn=["disabled"],$n={class:"nutrient-block"},wn=["disabled"],bn=["value"],yn={class:"unit"},kn=["value"],Cn={class:"unit"},In=["value"],zn={class:"unit"},Ln=["value"],Fn={class:"unit"},Tn=["value"],En={class:"unit"},Mn=["disabled"],Sn=["value"],xn={class:"unit"},Dn=["value"],Rn={class:"unit"},Bn=["value"],Vn={class:"unit"},Nn=["value"],An={class:"unit"},jn=["value"],On={class:"unit"},Hn={setup(o){const{t:l}=U(),u=x("log"),c=x("csrfToken"),s=x("perms"),i=x("food"),a=w([]),d=w(null),v=w(!1),p=w(!1),b=w(null),C=w(null);function $(r){fetch("/api/v1/food",{method:"POST",headers:{"X-CSRF-Token":c}}).then(t=>{if(!t.ok)throw l("createfood.err"+t.status);return t.json()}).then(t=>{t.name=r,i.value.push(t),a.value.push(t),u.msg(l("createfood.ok")),F(t.id)}).catch(t=>u.err(t))}function f(){p.value=!0;let r=d.value.id;fetch("/api/v1/food/"+r,{method:"PUT",headers:{"Content-Type":"application/x-www-form-urlencoded","X-CSRF-Token":c},body:new URLSearchParams(new FormData(C.value))}).then(t=>{if(!t.ok)throw l("savefood.err"+t.status);return v.value=!1,fetch("/api/v1/food/"+r)}).then(t=>t.json()).then(t=>{t.name=l(t.id.toString()),i.value=i.value.map(_=>t.id==_.id?t:_),a.value=a.value.map(_=>t.id==_.id?t:_),d.value=d.value.id==t.id?t:d.value,u.msg(l("savefood.ok"))}).catch(t=>u.err(t)).finally(()=>{setTimeout(function(){p.value=!1},500)})}function L(r){a.value=r,d.value&&a.value.filter(t=>t.id==d.value.id).length==0&&(d.value=null)}function F(r){d.value=a.value.filter(t=>t.id==r)[0],b.value.showDetails()}function R(){v.value?f():v.value=!0}function I(r){r.target.blur(),isNaN(parseFloat(r.target.value))&&(r.target.value=d.value[r.target.name])}return(r,t)=>(h(),E(j,{ref_key:"main",ref:b,onDetailVisibility:t[0]||(t[0]=_=>v.value=!1)},ae({filter:k(()=>[T(s).canCreateFood?(h(),g("section",vn,[e("h2",null,n(r.$t("aria.headnew")),1),m(oe,{label:r.$t("btn.new"),placeholder:r.$t("food.hintnew"),onConfirm:$},null,8,["label","placeholder"])])):D("",!0),e("section",null,[e("h2",null,n(r.$t("aria.headsearch")),1),m(xt,{data:T(i),placeholder:r.$t("food.hintsearch"),onResult:L},{default:k(_=>[e("fieldset",null,[e("legend",null,n(r.$t("aria.headmacro1")),1),m(S,{label:r.$t("food.energy"),onInput:_.confirm,name:"kcal",unit:"cal",min:"0",max:"900",frac:"0"},null,8,["label","onInput"]),m(S,{label:r.$t("food.fat"),onInput:_.confirm,name:"fat",unit:"g",min:"0",max:"100",frac:"0"},null,8,["label","onInput"]),m(S,{label:r.$t("food.carbs"),onInput:_.confirm,name:"carb",unit:"g",min:"0",max:"100",frac:"0"},null,8,["label","onInput"]),m(S,{label:r.$t("food.protein"),onInput:_.confirm,name:"prot",unit:"g",min:"0",max:"89",frac:"0"},null,8,["label","onInput"]),m(S,{label:r.$t("food.fiber"),onInput:_.confirm,name:"fib",unit:"g",min:"0",max:"71",frac:"0"},null,8,["label","onInput"])]),e("fieldset",null,[e("legend",null,n(r.$t("aria.headmacro2")),1),m(S,{label:r.$t("food.fatsat"),onInput:_.confirm,name:"fatsat",unit:"g",min:"0",max:"83",frac:"0"},null,8,["label","onInput"]),m(S,{label:r.$t("food.fato3"),onInput:_.confirm,name:"fato3",unit:"g",min:"0",max:"54",frac:"0"},null,8,["label","onInput"]),m(S,{label:r.$t("food.fato6"),onInput:_.confirm,name:"fato6",unit:"g",min:"0",max:"70",frac:"0"},null,8,["label","onInput"]),m(S,{label:r.$t("food.sugar"),onInput:_.confirm,name:"sug",unit:"g",min:"0",max:"100",frac:"0"},null,8,["label","onInput"]),m(S,{label:r.$t("food.salt"),onInput:_.confirm,name:"salt",unit:"g",min:"0",max:"100",frac:"0"},null,8,["label","onInput"])])]),_:1},8,["data","placeholder"])])]),main:k(()=>[m(se,{items:a.value,onSelected:F},null,8,["items"])]),_:2},[d.value?{name:"head-details",fn:k(()=>[e("h2",null,n(d.value.name),1)])}:void 0,d.value?{name:"details",fn:k(()=>[hn,e("section",mn,[_n,fn,pn,T(s).canCreateFood||T(s).canEditFood?(h(),g("button",{key:0,class:"icon async",disabled:p.value,onClick:R},[v.value?D("",!0):(h(),E(ie,{key:0})),v.value?(h(),E(Y,{key:1})):D("",!0)],8,gn)):D("",!0)]),e("section",null,[e("h2",null,n(r.$t("aria.headnutrients")),1),e("form",{ref_key:"form",ref:C},[e("div",$n,[e("fieldset",{disabled:!v.value,class:"col50"},[e("div",null,[e("label",null,n(r.$t("food.energy")),1),e("input",{type:"text",value:d.value.kcal,name:"kcal",onChange:I},null,40,bn),e("span",yn,n(r.$t("unit.cal")),1)]),e("div",null,[e("label",null,n(r.$t("food.fat")),1),e("input",{type:"text",value:d.value.fat,name:"fat",onChange:I},null,40,kn),e("span",Cn,n(r.$t("unit.g")),1)]),e("div",null,[e("label",null,n(r.$t("food.carbs2")),1),e("input",{type:"text",value:d.value.carb,name:"carb",onChange:I},null,40,In),e("span",zn,n(r.$t("unit.g")),1)]),e("div",null,[e("label",null,n(r.$t("food.protein")),1),e("input",{type:"text",value:d.value.prot,name:"prot",onChange:I},null,40,Ln),e("span",Fn,n(r.$t("unit.g")),1)]),e("div",null,[e("label",null,n(r.$t("food.fiber")),1),e("input",{type:"text",value:d.value.fib,name:"fib",onChange:I},null,40,Tn),e("span",En,n(r.$t("unit.g")),1)])],8,wn),e("fieldset",{disabled:!v.value,class:"col50"},[e("div",null,[e("label",null,n(r.$t("food.fatsat")),1),e("input",{type:"text",value:d.value.fatsat,name:"fatsat",onChange:I},null,40,Sn),e("span",xn,n(r.$t("unit.g")),1)]),e("div",null,[e("label",null,n(r.$t("food.fato3")),1),e("input",{type:"text",value:d.value.fato3,name:"fato3",onChange:I},null,40,Dn),e("span",Rn,n(r.$t("unit.g")),1)]),e("div",null,[e("label",null,n(r.$t("food.fato6")),1),e("input",{type:"text",value:d.value.fato6,name:"fato6",onChange:I},null,40,Bn),e("span",Vn,n(r.$t("unit.g")),1)]),e("div",null,[e("label",null,n(r.$t("food.sugar")),1),e("input",{type:"text",value:d.value.sug,name:"sug",onChange:I},null,40,Nn),e("span",An,n(r.$t("unit.g")),1)]),e("div",null,[e("label",null,n(r.$t("food.salt")),1),e("input",{type:"text",value:d.value.salt,name:"salt",onChange:I},null,40,jn),e("span",On,n(r.$t("unit.g")),1)])],8,Mn)])],512)])])}:void 0]),1536))}};const Xn={class:"ingredients"},qn={disabled:""},Pn=["value"],Un={class:"unit"},Wn=["value"],Kn={props:["items"],setup(o){const l=o,u=x("food"),c=w("name"),s=new Intl.Collator(U().locale.value,{numeric:!0}),i=ne(()=>{let a=l.items.map(v=>v.id);return u.value.filter(v=>a.includes(v.id)).sort((v,p)=>s.compare(v[c.value],p[c.value]))});return(a,d)=>(h(),g("form",Xn,[e("fieldset",qn,[(h(!0),g(K,null,Q(T(i),v=>(h(),g("div",{key:v.id},[e("label",null,n(v.name),1),e("input",{type:"text",name:"amount",value:v.amount},null,8,Pn),e("span",Un,n(a.$t("unit.g")),1),e("input",{type:"hidden",name:"id",value:v.id},null,8,Wn)]))),128))])]))}},Gn={},Jn={width:"24",height:"24",version:"1.1",viewBox:"0 0 24 24",xmlns:"http://www.w3.org/2000/svg"},Qn=e("path",{d:"m0 18.316h2.5263v0.63158h-1.2632v1.2632h1.2632v0.63158h-2.5263v1.2632h3.7895v-5.0526h-3.7895zm1.2632-11.368h1.2632v-5.0526h-2.5263v1.2632h1.2632zm-1.2632 3.7895h2.2737l-2.2737 2.6526v1.1368h3.7895v-1.2632h-2.2737l2.2737-2.6526v-1.1368h-3.7895zm6.3158-7.5789v2.5263h17.684v-2.5263zm0 17.684h17.684v-2.5263h-17.684zm0-7.5789h17.684v-2.5263h-17.684z"},null,-1),Yn=[Qn];function Zn(o,l){return h(),g("svg",Jn,Yn)}var ea=M(Gn,[["render",Zn]]);const ta={class:"new-item"},na=["disabled"],aa=["value"],la=["innerHTML"],oa={class:"tags"},sa=e("span",{class:"tag"},"Tag 1",-1),ia=e("span",{class:"tag"},"Tag 2",-1),ua=e("span",{class:"tag"},"Tag 3",-1),ra=["disabled"],ca=y(" Add to diary here "),da={class:"nutrient-block"},va={class:"col50"},ha={class:"unit"},ma={class:"unit"},_a={class:"col50"},fa={class:"unit"},pa={class:"unit"},ga={class:"prep"},$a=["disabled"],wa=["value"],ba=["value"],ya={class:"unit"},ka=["disabled"],Ca=["value"],Ia={class:"unit"},za=["value"],La={class:"unit"},Fa=["value"],Ta={class:"unit"},Ea={class:"placeholder"},Ma={setup(o){const{t:l}=U(),u=x("log"),c=x("csrfToken");x("perms");const s=x("recipes"),i=w([]),a=w(null),d=w(!1),v=w(!1),p=w("&nbsp;"),b=w(null),C=w(null);function $(t){fetch("/api/v1/recipe",{method:"POST",headers:{"Content-Type":"application/x-www-form-urlencoded","X-CSRF-Token":c},body:new URLSearchParams({name:t})}).then(_=>{if(!_.ok)throw l("createrec.err"+_.status);return _.json()}).then(_=>{_.isowner=!0,s.value.push(_),i.value.push(_),u.msg(l("createrec.ok")),L(_.id)}).catch(_=>u.err(_))}function f(){v.value=!0;let t=a.value.id,_=a.value.owner,ee=a.value.isowner;fetch("/api/v1/recipe/"+t,{method:"PUT",headers:{"Content-Type":"application/x-www-form-urlencoded","X-CSRF-Token":c},body:new URLSearchParams(new FormData(C.value))}).then(z=>{if(!z.ok)throw l("saverec.err"+z.status);return d.value=!1,fetch("/api/v1/recipe/"+t)}).then(z=>z.json()).then(z=>{z.owner=_,z.isowner=ee,s.value=s.value.map(O=>z.id==O.id?z:O),i.value=i.value.map(O=>z.id==O.id?z:O),a.value=a.value.id==z.id?z:a.value,u.msg(l("saverec.ok"))}).catch(z=>u.err(z)).finally(()=>{setTimeout(function(){v.value=!1},500)})}function L(t){a.value=i.value.filter(_=>_.id==t)[0],b.value.showDetails(),"isowner"in a.value?R():F()}function F(){fetch(`/api/v1/recipe/${a.value.id}/owner`).then(t=>{if(!t.ok)throw t;return t.json()}).then(t=>{a.value.isowner=t.isowner,a.value.owner=t.owner,R()}).catch(()=>{u.err(l("recowner.err")),p.value="&nbsp;"})}function R(){a.value.isowner?p.value=l("recipe.isowner"):a.value.owner?p.value=l("recipe.owner",{name:a.value.owner}):p.value=l("recipe.ispublic")}function I(){d.value?f():d.value=!0}function r(t){t.target.blur(),isNaN(parseFloat(t.target.value))&&(t.target.value=a.value[t.target.name])}return W(()=>i.value=s.value),(t,_)=>(h(),E(j,{ref_key:"main",ref:b,onDetailVisibility:_[0]||(_[0]=ee=>d.value=!1)},ae({filter:k(()=>[e("section",ta,[e("h2",null,n(t.$t("aria.headnew")),1),m(oe,{label:t.$t("btn.new"),placeholder:t.$t("recipe.hintnew"),onConfirm:$},null,8,["label","placeholder"])]),e("section",null,[e("h2",null,n(t.$t("aria.headsearch")),1)])]),main:k(()=>[m(se,{items:i.value,onSelected:L},null,8,["items"])]),_:2},[a.value?{name:"head-details",fn:k(()=>[e("form",{ref_key:"form",ref:C,autocomplete:"off",id:"form-recipe"},[e("fieldset",{disabled:!d.value},[e("input",{type:"text",name:"name",value:a.value.name},null,8,aa)],8,na)],512)])}:void 0,a.value?{name:"details",fn:k(()=>[e("section",{class:"subtitle",innerHTML:p.value},null,8,la),e("section",oa,[sa,ia,ua,a.value.isowner?(h(),g("button",{key:0,class:"icon async",disabled:v.value,onClick:I},[d.value?D("",!0):(h(),E(ie,{key:0})),d.value?(h(),E(Y,{key:1})):D("",!0)],8,ra)):D("",!0)]),e("section",null,[e("h2",null,n(t.$t("aria.headtrack")),1),ca]),e("section",null,[e("h2",null,n(t.$t("aria.headingred")),1),m(Kn,{items:a.value.items},null,8,["items"])]),e("section",null,[e("h2",null,n(t.$t("aria.headnutrients")),1),e("div",da,[e("div",va,[e("div",null,[e("label",null,n(t.$t("food.energy")),1),e("span",null,n(a.value.kcal),1),e("span",ha,n(t.$t("unit.cal")),1)]),e("div",null,[e("label",null,n(t.$t("food.fat")),1),e("span",null,n(a.value.fat),1),e("span",ma,n(t.$t("unit.g")),1)])]),e("div",_a,[e("div",null,[e("label",null,n(t.$t("food.carbs2")),1),e("span",null,n(a.value.carb),1),e("span",fa,n(t.$t("unit.g")),1)]),e("div",null,[e("label",null,n(t.$t("food.protein")),1),e("span",null,n(a.value.prot),1),e("span",pa,n(t.$t("unit.g")),1)])])])]),e("section",ga,[e("h2",null,n(t.$t("aria.headprep")),1),e("div",null,[e("fieldset",{disabled:!d.value,class:"col50"},[e("div",null,[e("label",null,n(t.$t("recipe.size")),1),e("input",{type:"text",name:"size",form:"form-recipe",value:a.value.size,onChange:r},null,40,wa)]),e("div",null,[e("label",null,n(t.$t("recipe.time")),1),e("input",{type:"text",disabled:"",value:a.value.preptime+a.value.cooktime+a.value.misctime},null,8,ba),e("span",ya,n(t.$t("unit.min")),1)])],8,$a),e("fieldset",{disabled:!d.value,class:"col50"},[e("div",null,[e("label",null,n(t.$t("recipe.preptime")),1),e("input",{type:"text",name:"preptime",form:"form-recipe",value:a.value.preptime,onChange:r},null,40,Ca),e("span",Ia,n(t.$t("unit.min")),1)]),e("div",null,[e("label",null,n(t.$t("recipe.cooktime")),1),e("input",{type:"text",name:"cooktime",form:"form-recipe",value:a.value.cooktime,onChange:r},null,40,za),e("span",La,n(t.$t("unit.min")),1)]),e("div",null,[e("label",null,n(t.$t("recipe.misctime")),1),e("input",{type:"text",name:"misctime",form:"form-recipe",value:a.value.misctime,onChange:r},null,40,Fa),e("span",Ta,n(t.$t("unit.min")),1)])],8,ka)]),e("div",Ea,[m(ea),e("p",null,n(t.$t("todo.instructions")),1)])])])}:void 0]),1536))}},Sa=y(" Diary "),xa={setup(o){return(l,u)=>(h(),E(j,null,{main:k(()=>[Sa]),_:1}))}},Da=y(" Shopping Lists "),Ra={setup(o){return(l,u)=>(h(),E(j,null,{main:k(()=>[Da]),_:1}))}},Ba=y(" Profile "),Va={setup(o){return(l,u)=>(h(),E(j,null,{main:k(()=>[Ba]),_:1}))}},Na=y(" Settings "),Aa={setup(o){return(l,u)=>(h(),E(j,null,{main:k(()=>[Na]),_:1}))}},ja=de({history:ve(),routes:[{path:"/",name:"food",component:Hn},{path:"/recipes",name:"recipes",component:Ma},{path:"/diary",name:"diary",component:xa},{path:"/shopping",name:"shopping",component:Ra},{path:"/profile",name:"profile",component:Va},{path:"/settings",name:"settings",component:Aa}]});let J=document.documentElement.lang||navigator.language;!J&&navigator.languages!=null&&(J=navigator.languages[0]);const te=document.querySelector("meta[name='_csrf']"),Oa=te?te.content:"",Ha=function(){const o=document.documentElement.dataset.perm||1,l=65536,u=131072;function c(s){return(o&s)==s}return{canCreateFood:c(l),canEditFood:c(u)}}(),Xa=function(){function o(l){return typeof l=="string"?l:"message"in l?l.message:A.global.t("err.err")}return{msg:function(l){let u={msg:o(l),timeout:3e3};window.dispatchEvent(new CustomEvent("message",{detail:u}))},warn:function(l){let u={msg:o(l),timeout:4e3};window.dispatchEvent(new CustomEvent("warning",{detail:u}))},err:function(l){let u={msg:o(l),timeout:5e3};window.dispatchEvent(new CustomEvent("error",{detail:u}))}}}(),B=he(Ze);let A,q,P;function qa(o){A=me({locale:J.split("-")[0],fallbackLocale:"en",messages:o}),q&&P&&Z()}function Pa(o){q=o,A&&P&&Z()}function Ua(o){P=o,console.log(P),A&&q&&Z()}function Z(){q.forEach(o=>{o.name=A.global.t(o.id.toString())}),B.provide("csrfToken",Oa),B.provide("perms",Ha),B.provide("food",w(q)),B.provide("recipes",w(P)),B.provide("log",Xa),B.use(ja),B.use(A),B.mount("#app")}fetch("/app/l10n.json").then(o=>o.json()).then(qa);fetch("api/v1/foods").then(o=>o.json()).then(Pa);fetch("api/v1/recipes").then(o=>o.json()).then(Ua);
