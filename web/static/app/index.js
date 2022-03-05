import{o as p,c as $,a as e,u as O,i as R,b as m,w as g,d as _,t as d,e as b,R as F,f as W,g as G,F as P,r as y,h as J,j as D,n as L,k as N,l as Q,v as Y,m as Z,p as ee,q as z,s as te,x as ne,y as ae,z as oe,A as se,B as le}from"./vendor.js";const ie=function(){const i=document.createElement("link").relList;if(i&&i.supports&&i.supports("modulepreload"))return;for(const t of document.querySelectorAll('link[rel="modulepreload"]'))n(t);new MutationObserver(t=>{for(const s of t)if(s.type==="childList")for(const l of s.addedNodes)l.tagName==="LINK"&&l.rel==="modulepreload"&&n(l)}).observe(document,{childList:!0,subtree:!0});function a(t){const s={};return t.integrity&&(s.integrity=t.integrity),t.referrerpolicy&&(s.referrerPolicy=t.referrerpolicy),t.crossorigin==="use-credentials"?s.credentials="include":t.crossorigin==="anonymous"?s.credentials="omit":s.credentials="same-origin",s}function n(t){if(t.ep)return;t.ep=!0;const s=a(t);fetch(t.href,s)}};ie();var E=(o,i)=>{const a=o.__vccOpts||o;for(const[n,t]of i)a[n]=t;return a};const re={},ce={width:"512",height:"512",version:"1.1",viewBox:"0 0 512 512",xmlns:"http://www.w3.org/2000/svg"},ue=e("path",{d:"m114.6 507.3c13.82-0.8729 90.06-3.083 106.2-3.079h13.49l2.639-5.293c5.191-10.41 10.35-29.57 13.13-48.74 1.811-12.5 1.488-16.68-1.826-23.56-2.528-5.249-11.04-13.06-20.54-18.86-28.38-17.32-51.21-34.77-76.05-58.12-6.673-6.272-9.377-8.209-16.08-11.53-12.37-6.124-20.36-12.88-29.19-24.7-1.45-1.941-1.795-2.038-4.695-1.334-1.722 0.4182-7.18 0.8166-12.13 0.8854-19.6 0.2722-38.18-7.715-52.49-22.56-15.97-16.57-23.43-35.49-23.41-59.37 0.01786-16.86 3.547-29.88 11.93-44 9.617-16.19 27.39-30.13 43.89-34.41 2.517-0.653 4.748-1.36 4.958-1.571 0.2103-0.2107-0.05169-3.078-0.5816-6.371-1.41-8.759-0.556-21.84 1.997-30.61 7.023-24.1 24.96-41.95 48.52-48.31 11.42-3.082 29.45-2.167 40.42 2.049 0.919 0.3533 1.637-0.7202 2.994-4.474 6.21-17.18 18.79-33.75 33.11-43.62 35.65-24.56 80.99-19.9 111.3 11.44l5.539 5.719 5.025-1.229c2.764-0.6762 9.075-1.428 14.02-1.672 16.17-0.7944 31.64 3.865 45.31 13.64l5.1 3.647 12.55 0.1171c9.984 0.09312 13.77 0.4491 18.49 1.739 29.1 7.945 50.37 32.58 56 64.86 1.354 7.764 1.096 21.6-0.5502 29.5l-0.916 4.394 5.367 7.086c20.1 26.54 25.52 61.05 14.68 93.52-7.862 23.55-24.99 43.53-46.38 54.11-10.16 5.023-23.1 8.587-31.28 8.613-3.79 0.0118-4.555 0.333-8.217 3.446-6.958 5.916-20.59 13.71-28.3 16.17-1.722 0.551-4.011 1.599-5.087 2.328-1.076 0.7297-8.119 4.802-15.65 9.05-18.57 10.47-37.53 22.69-48.72 31.39-13.27 10.32-16.83 22.63-14.72 50.94 2.648 35.64 8.059 59.41 16.33 71.75l2.89 4.313 29.4 0.4963c34.14 0.5764 92.53 2.291 93.01 2.731 0.1793 0.166-74.3 0.2503-165.5 0.1877-91.21-0.0631-161.4-0.392-156-0.732zm130.9-101.8c-0.1092-1.007-3.147-5.947-6.75-10.98-7.889-11.01-19.71-28.99-22.37-34.01-1.594-3.014-2.402-3.727-4.695-4.139-7.259-1.304-17.91-5.324-27.58-10.41-1.809-0.9509 5.669 20.15 9.032 25.49 4.859 7.714 27.69 24.94 45.08 34 5.177 2.7 7.575 2.715 7.285 0.0442zm6.168-19.46c0.8235-0.9976 1.074-4.638 1.074-15.62 0-7.875-0.1764-14.5-0.3923-14.71-0.2158-0.2163-2.307 0.1044-4.647 0.7128-2.34 0.6082-7 1.335-10.35 1.616-7.067 0.5907-6.896 0-2.984 10.28 2.73 7.17 6.328 12.81 10.4 16.3 3.513 3.013 5.293 3.38 6.908 1.424zm49.91-15.13c3.913-2 11.37-9.412 15.96-15.85 3.982-5.592 8.688-15 8.214-16.42-0.1628-0.49-3.275-1.75-6.916-2.799-3.641-1.049-9.568-3.501-13.17-5.449l-6.551-3.541-2.173 2.087c-2.148 2.062-2.179 2.234-2.645 14.37-0.2592 6.754-0.6282 15.17-0.8198 18.7-0.3026 5.572-0.1405 6.682 1.232 8.43 1.896 2.415 2.946 2.489 6.867 0.484zm-118.1-5.725c-2.19-4.729-6.256-17.66-6.256-19.89 0-0.8683-1.901-1.089-9.391-1.089-5.165 0-9.391 0.2195-9.391 0.488 0 1.561 25.33 25.39 26.99 25.39 0.1762 0-0.7011-2.206-1.949-4.901zm149.8-9.536c5.487-3.687 19.46-14.65 19.07-14.96-0.126-0.1032-4.197-0.3888-9.046-0.6338l-8.816-0.4452-3.967 5.891c-4.284 6.363-7.088 11.99-6.473 12.98 0.9066 1.47 4.434 0.3864 9.228-2.836z",fill:"#13ad73"},null,-1),de=e("path",{d:"m253.9 167.7c-0.2579 0.0513-1.238 0.5454-4.691 0.5216-3.625-0.0248 4.941 9.227 6.934 34.4-8.031-0.4678-8.408-3.285-24.81-5.999-15.88-2.627-25.28 5.883-29.57 18.88-3.429 10.39-3.592 23.66-1.19 36.39 2.348 12.45 7.148 24.39 13.74 32.65 6.909 8.654 14.05 10.96 20.49 10.82 8.459-0.1842 15.7-4.599 19.58-4.407 6.726 0.3331 13.29 8.894 28.17 2.151 6.047-2.741 12.03-7.819 17.08-14.81 4.984-6.907 9.051-15.68 11.35-25.91 2.018-8.972 2.678-19.07 1.41-30-1.616-13.93-7.519-20.42-14.9-23.01-8.286-2.905-18.43-0.8896-26.46 1.032-4.71 1.127-8.69 2.223-11.14 2.273-0.155 3e-3 -0.2829-4.1e-4 -0.4335 2e-3 -0.01-0.0695-0.0197-0.1385-0.0323-0.2078-2.249-11.58-4.704-29.29-5.401-33.64-0.152-0.9486 0.0845-1.177-0.1162-1.137z",fill:"#f2ac05","fill-rule":"evenodd"},null,-1),me=[ue,de];function _e(o,i){return p(),$("svg",ce,me)}var he=E(re,[["render",_e]]);const pe={},fe={width:"24",height:"24",version:"1.1",viewBox:"0 0 24 24",xmlns:"http://www.w3.org/2000/svg"},ve=e("path",{d:"m12 0c-6.624 0-12 5.376-12 12s5.376 12 12 12 12-5.376 12-12-5.376-12-12-12zm0 3.6c1.992 0 3.6 1.608 3.6 3.6s-1.608 3.6-3.6 3.6-3.6-1.608-3.6-3.6 1.608-3.6 3.6-3.6zm0 17.04c-3 0-5.652-1.536-7.2-3.864 0.036-2.388 4.8-3.696 7.2-3.696 2.388 0 7.164 1.308 7.2 3.696-1.548 2.328-4.2 3.864-7.2 3.864z"},null,-1),ge=[ve];function $e(o,i){return p(),$("svg",fe,ge)}var we=E(pe,[["render",$e]]);const ye={},be={width:"24",height:"24",version:"1.1",viewBox:"0 0 24 24",xmlns:"http://www.w3.org/2000/svg"},ke=e("path",{d:"m12 0-2.115 2.115 8.37 8.385h-18.255v3h18.255l-8.37 8.385 2.115 2.115 12-12z"},null,-1),xe=[ke];function Le(o,i){return p(),$("svg",be,xe)}var Fe=E(ye,[["render",Le]]);const Ce=e("div",{id:"app-name"},[e("span",null,"Hey"),e("span",null,"Apple")],-1),ze={id:"nav-main"},Ee={id:"nav-user"},Me={href:"https://docs.heyapple.org",target:"_blank"},Ie={setup(o){const{t:i}=O(),a=R("csrfToken");function n(s){s.preventDefault(),fetch("/auth/local",{method:"DELETE",headers:{"X-CSRF-Token":a}}).then(l=>{l.ok?window.location="/":window.dispatchEvent(new CustomEvent("error",{detail:{msg:i("signout.err"+l.status)}}))})}function t(s){s.stopPropagation(),document.querySelector("header nav").classList.toggle("open")}return document.addEventListener("click",function(){document.querySelector("header nav").classList.remove("open")}),(s,l)=>(p(),$("header",null,[m(he,{id:"logo"}),Ce,e("nav",null,[e("button",{onClick:t},[m(Fe)]),e("ul",ze,[e("li",null,[m(b(F),{to:"/"},{default:g(()=>[_(d(s.$t("nav.food")),1)]),_:1})]),e("li",null,[m(b(F),{to:"/recipes"},{default:g(()=>[_(d(s.$t("nav.recipes")),1)]),_:1})]),e("li",null,[m(b(F),{to:"/diary"},{default:g(()=>[_(d(s.$t("nav.diary")),1)]),_:1})]),e("li",null,[m(b(F),{to:"/shopping"},{default:g(()=>[_(d(s.$t("nav.shopping")),1)]),_:1})])]),e("ul",Ee,[e("li",null,[m(b(F),{to:"/profile"},{default:g(()=>[_(d(s.$t("nav.profile")),1)]),_:1})]),e("li",null,[m(b(F),{to:"/settings"},{default:g(()=>[_(d(s.$t("nav.settings")),1)]),_:1})]),e("li",null,[e("a",Me,d(s.$t("nav.help")),1)]),e("li",null,[e("a",{href:"#",onClick:n},d(s.$t("nav.signout")),1)])])]),e("button",{onClick:t},[m(we)])]))}};const Se={setup(o){return W(()=>{document.querySelector("body > .spinner-container").remove()}),(i,a)=>(p(),$(P,null,[m(Ie),m(b(G))],64))}},De={},Te={version:"1.1",viewBox:"0 0 24 24",xmlns:"http://www.w3.org/2000/svg"},Be=e("path",{id:"path-top",d:"m0 3v2h24v-2z"},null,-1),Ae=e("path",{id:"path-mid",d:"m0 13h24v-2h-24z"},null,-1),Ne=e("path",{id:"path-bottom",d:"m0 21h24v-2h-24z"},null,-1),Re=[Be,Ae,Ne];function Ve(o,i){return p(),$("svg",Te,Re)}var qe=E(De,[["render",Ve]]);const He={},je={width:"512",height:"512",version:"1.1",viewBox:"0 0 512 512",xmlns:"http://www.w3.org/2000/svg"},Oe=e("path",{d:"m251.2 4.007c-17.24-0.2277-34.69 4.924-50.29 15.67-14.33 9.866-26.9 26.43-33.11 43.62-1.357 3.754-2.073 4.828-2.992 4.475-10.97-4.217-29-5.131-40.42-2.049-23.57 6.359-41.5 24.22-48.52 48.31-2.553 8.761-3.407 21.84-1.997 30.6 0.5299 3.293 0.7902 6.162 0.5799 6.372-0.2104 0.2108-2.439 0.9186-4.956 1.572-16.5 4.283-34.28 18.21-43.89 34.41-8.382 14.11-11.91 27.14-11.93 44-0.02401 23.88 7.433 42.79 23.41 59.36 14.31 14.85 32.89 22.84 52.49 22.56 4.95-0.0689 10.41-0.4681 12.13-0.8863 2.9-0.7045 3.245-0.6061 4.695 1.334 8.833 11.82 16.82 18.58 29.19 24.7 6.707 3.32 9.412 5.257 16.09 11.53 24.85 23.35 47.68 40.8 76.05 58.12 9.503 5.8 18.01 13.61 20.54 18.86 3.313 6.878 3.637 11.06 1.825 23.56-2.778 19.18-7.936 38.33-13.13 48.74l-2.639 5.295h-13.49c-16.17-3e-3 -92.41 2.205-106.2 3.077l0.01-3e-3c-5.381 0.3399 64.84 0.6684 156 0.7315 91.2 0.0625 165.7-0.0216 165.5-0.1878-0.4753-0.4399-58.86-2.155-93-2.731l-29.4-0.4942-2.89-4.313c-8.274-12.35-13.68-36.12-16.33-71.76-2.105-28.31 1.454-40.62 14.72-50.94 11.18-8.699 30.14-20.92 48.72-31.39 7.533-4.248 14.57-8.318 15.65-9.048 1.076-0.7296 3.366-1.779 5.087-2.33 7.71-2.467 21.34-10.26 28.3-16.17 3.662-3.113 4.427-3.435 8.218-3.446 8.183-0.0267 21.12-3.59 31.28-8.613 21.39-10.58 38.52-30.56 46.38-54.11 10.84-32.47 5.42-66.98-14.68-93.52l-5.364-7.084 0.916-4.395c1.646-7.896 1.904-21.74 0.5502-29.5-5.632-32.29-26.9-56.92-56-64.86-4.726-1.29-8.51-1.647-18.49-1.74l-12.55-0.1186-5.101-3.647c-13.67-9.776-29.15-14.43-45.31-13.64-4.95 0.2432-11.26 0.9944-14.02 1.671l-5.028 1.229-5.539-5.72c-17.07-17.63-38.89-26.81-61.05-27.1zm2.728 163.7c0.2008-0.0398-0.0367 0.1881 0.1154 1.137 0.6972 4.35 3.151 22.06 5.4 33.64 0.0118 0.0694 0.0228 0.138 0.0328 0.2076 0.1506-2e-3 0.2799 9e-5 0.4349-3e-3 2.448-0.0501 6.427-1.143 11.14-2.27 8.028-1.922 18.18-3.94 26.46-1.035 7.379 2.587 13.28 9.078 14.9 23.01 1.268 10.94 0.6106 21.03-1.407 30-2.301 10.23-6.37 19.01-11.35 25.91-5.045 6.991-11.03 12.07-17.07 14.81-14.87 6.743-21.44-1.818-28.16-2.152-3.88-0.1922-11.12 4.224-19.58 4.409-6.437 0.1402-13.58-2.169-20.49-10.82-6.596-8.262-11.39-20.2-13.74-32.65-2.402-12.73-2.24-26 1.189-36.39 4.29-13 13.69-21.51 29.58-18.88 16.41 2.714 16.78 5.532 24.81 6-1.993-25.18-10.56-34.43-6.936-34.4 3.453 0.0238 4.434-0.4726 4.692-0.5239zm45.16 159.1 6.55 3.542c3.603 1.947 9.529 4.401 13.17 5.45 3.641 1.049 6.753 2.307 6.916 2.797 0.4733 1.423-4.232 10.83-8.214 16.42-4.588 6.442-12.05 13.85-15.96 15.86-3.919 2.004-4.971 1.93-6.867-0.4843-1.372-1.748-1.532-2.86-1.229-8.432 0.1916-3.529 0.558-11.94 0.8171-18.7 0.4657-12.13 0.4983-12.3 2.646-14.37zm35.38 12.77 8.817 0.4448c4.849 0.245 8.919 0.5294 9.045 0.6326 0.3835 0.315-13.58 11.27-19.07 14.96-4.795 3.222-8.322 4.307-9.229 2.837-0.6151-0.9971 2.187-6.619 6.471-12.98zm-166.7 4.596c7.489 0 9.391 0.219 9.391 1.087 0 2.234 4.067 15.16 6.257 19.89 1.248 2.694 2.123 4.9 1.947 4.9-1.66 0-26.99-23.83-26.99-25.39 0-0.2685 4.226-0.4876 9.391-0.4876zm16.13 1.727c0.0388-0.0118 0.0851-6e-3 0.1417 0.0231 9.674 5.087 20.32 9.107 27.58 10.41 2.293 0.4121 3.101 1.124 4.695 4.138 2.657 5.026 14.48 23 22.37 34.02 3.603 5.03 6.639 9.968 6.748 10.98 0.2897 2.671-2.108 2.658-7.285-0.0429-17.39-9.068-40.22-26.29-45.07-34.01-3.258-5.172-10.38-25.14-9.173-25.52zm67.9 9.71c0.263-9e-3 0.4337 9e-3 0.4876 0.0625 0.2159 0.2154 0.3921 6.837 0.3921 14.71 0 10.98-0.2506 14.62-1.074 15.62-1.615 1.956-3.397 1.59-6.909-1.423-4.068-3.49-7.665-9.13-10.4-16.3-3.912-10.28-4.082-9.686 2.985-10.28 3.355-0.2804 8.013-1.006 10.35-1.615 1.755-0.4562 3.373-0.7502 4.162-0.7776z"},null,-1),Pe=[Oe];function Xe(o,i){return p(),$("svg",je,Pe)}var Ke=E(He,[["render",Xe]]);const Ue={id:"filter"},We=_(" This is the main search & filter area"),Ge={id:"main"},Je={class:"controls"},Qe={class:"content"},Ye=_("This is the main area"),Ze={id:"details"},e1={class:"placeholder"},M={props:["current"],setup(o){const i=o,a=y(""),n=y("");J(()=>i.current,l=>{l&&s()});function t(){a.value==""?(a.value="open-filter",n.value=""):a.value=""}function s(){n.value==""&&(n.value="open-details",a.value="")}return(l,c)=>(p(),$("main",{class:L([a.value,n.value])},[e("div",Ue,[D(l.$slots,"filter",{},()=>[We])]),e("div",Ge,[e("div",Je,[e("button",{onClick:t,class:"open-close"},[m(qe)])]),e("div",Qe,[D(l.$slots,"main",{},()=>[Ye])])]),e("div",Ze,[D(l.$slots,"details",{},()=>[e("div",e1,[m(Ke),e("p",null,d(l.$t("details.noitem")),1)])])])],2))}},t1=["placeholder"],n1=_("Additional filters"),a1={props:["data","placeholder"],emits:["result"],setup(o,{emit:i}){const a=o;var n=void 0;function t(l){l.preventDefault(),clearTimeout(n),n=setTimeout(function(){s(l.target.closest("form"))},500)}function s(l){let c=new FormData(l),r=a.data.filter(f=>{for(let v of c.keys()){if(v=="name"){let u=c.get(v).toLowerCase();if(!f[v].toLowerCase().includes(u))return!1;continue}let[w,h]=c.getAll(v).map(u=>parseFloat(u));if(!isNaN(w)&&!isNaN(h)&&(f[v]<w||h<f[v]))return!1}return!0});i("result",r)}return(l,c)=>(p(),$("form",null,[e("input",{type:"text",name:"name",placeholder:o.placeholder,onInput:t},null,40,t1),D(l.$slots,"default",{confirm:t},()=>[n1])]))}};const o1={class:"slider"},s1=["name","value"],l1=_(" \xA0\u2013\xA0 "),i1=["name","value"],r1={class:"bar"},c1={class:"slide"},u1=e("div",{class:"interact-area"},null,-1),d1=[u1],m1=e("div",{class:"interact-area"},null,-1),_1=[m1],S={props:["label","name","unit","min","max","frac"],emits:["input"],setup(o,{emit:i}){const a=o,n=y(parseFloat(a.min).toFixed(a.frac)),t=y(parseFloat(a.max).toFixed(a.frac)),s=y(0),l=y(100);function c(h){h.target.blur();let u=parseFloat(a.min)||0,k=parseFloat(h.target.value)||u;n.value=k,t.value=Math.max(n.value,t.value);let x=parseFloat(a.max)||0;s.value=(k-u)*100/(x-u),l.value=Math.max(s.value,l.value),i("input",h)}function r(h){h.target.blur();let u=parseFloat(a.max)||0,k=parseFloat(h.target.value)||u;t.value=k,n.value=Math.min(n.value,t.value);let x=parseFloat(a.min)||0;l.value=(k-x)*100/(u-x),s.value=Math.min(s.value,l.value),i("input",h)}function f(h){let k=h.target.closest(".slide").getBoundingClientRect(),x=h.pageX!==void 0?h.pageX:h.changedTouches[0].pageX;x=Math.min(Math.max(x-k.left,0),k.width);let A=x*100/k.width,q=parseFloat(a.min)||0,U=parseFloat(a.max)||0,H=A/100*(U-q)+q;h.target.closest("button").classList.contains("min")?(n.value=H.toFixed(a.frac),t.value=Math.max(n.value,t.value),s.value=A,l.value=Math.max(s.value,l.value)):(t.value=H.toFixed(a.frac),n.value=Math.min(n.value,t.value),l.value=A,s.value=Math.min(s.value,l.value)),i("input",h)}function v(h){let u=h.target.closest("button");u.addEventListener("mousemove",f),u.addEventListener("touchmove",f),u.addEventListener("mouseup",w),u.addEventListener("touchend",w),u.addEventListener("mouseleave",w),u.addEventListener("touchcancel",w),u.classList.add("active")}function w(h){let u=h.target.closest("button");u.removeEventListener("mousemove",f),u.removeEventListener("touchmove",f),u.removeEventListener("mouseup",w),u.removeEventListener("touchend",w),u.removeEventListener("mouseleave",w),u.removeEventListener("touchcancel",w),u.classList.remove("active")}return(h,u)=>(p(),$("div",o1,[e("label",null,[e("span",null,d(o.label)+" ("+d(h.$t("unit."+o.unit))+")",1),e("input",{type:"text",name:o.name,value:n.value,onChange:c},null,40,s1),l1,e("input",{type:"text",name:o.name,value:t.value,onChange:r},null,40,i1)]),e("div",r1,[e("div",c1,[e("div",{class:"overlay",style:N({left:s.value+"%",right:100-l.value+"%"})},null,4),e("button",{type:"button",class:"handle min",style:N({left:s.value+"%"}),onMousedown:v,onTouchstart:v},d1,36),e("button",{type:"button",class:"handle max",style:N({left:l.value+"%"}),onMousedown:v,onTouchstart:v},_1,36)])])]))}};const h1={class:"clickable-input"},p1=["placeholder"],f1=["value"],X={props:["label","placeholder"],emits:["confirm"],setup(o,{emit:i}){const a=y("");function n(t){t.preventDefault(),i("confirm",a.value),a.value=""}return(t,s)=>(p(),$("form",h1,[Q(e("input",{type:"text","onUpdate:modelValue":s[0]||(s[0]=l=>a.value=l),placeholder:o.placeholder},null,8,p1),[[Y,a.value]]),e("input",{type:"submit",onClick:n,value:o.label},null,8,f1)]))}};const v1={},g1={class:"icon sort-arrow"};function $1(o,i){return p(),$("span",g1)}var I=E(v1,[["render",$1]]);const w1=["onClick"],y1={class:"name"},b1={class:"num"},k1={class:"unit"},x1={class:"m num"},L1={class:"unit"},F1={class:"m num"},C1={class:"unit"},z1={class:"m num"},E1={class:"unit"},M1={props:["items"],emits:"selected",setup(o,{emit:i}){const a=o,n=y("name"),t=y("asc"),s=Z(()=>t.value=="asc"?[...a.items].sort((c,r)=>c[n.value]<r[n.value]?-1:c[n.value]>r[n.value]?1:0):[...a.items].sort((c,r)=>c[n.value]>r[n.value]?-1:c[n.value]<r[n.value]?1:0));function l(c){let r=c.target.dataset.sort;n.value==r?t.value=t.value=="asc"?"desc":"asc":n.value=r}return(c,r)=>(p(),$("table",null,[e("thead",null,[e("tr",{class:L(t.value)},[e("th",{class:L(["name sort",{active:n.value=="name"}]),onClick:l,"data-sort":"name"},[_(d(c.$t("food.name"))+" ",1),m(I)],2),e("th",{class:L(["num sort",{active:n.value=="kcal"}]),onClick:l,"data-sort":"kcal"},[m(I),_(" "+d(c.$t("food.energy")),1)],2),e("th",{class:L(["m num sort",{active:n.value=="fat"}]),onClick:l,"data-sort":"fat"},[m(I),_(" "+d(c.$t("food.fat")),1)],2),e("th",{class:L(["m num sort",{active:n.value=="carb"}]),onClick:l,"data-sort":"carb"},[m(I),_(" "+d(c.$t("food.carbs2")),1)],2),e("th",{class:L(["m num sort",{active:n.value=="prot"}]),onClick:l,"data-sort":"prot"},[m(I),_(" "+d(c.$t("food.protein")),1)],2)],2)]),e("tbody",null,[(p(!0),$(P,null,ee(b(s),f=>(p(),$("tr",{key:f.id,onClick:v=>c.$emit("selected",f.id)},[e("td",y1,d(f.name),1),e("td",b1,[_(d(f.kcal)+" ",1),e("span",k1,d(c.$t("unit.cal")),1)]),e("td",x1,[_(d(f.fat)+" ",1),e("span",L1,d(c.$t("unit.g")),1)]),e("td",F1,[_(d(f.carb)+" ",1),e("span",C1,d(c.$t("unit.g")),1)]),e("td",z1,[_(d(f.prot)+" ",1),e("span",E1,d(c.$t("unit.g")),1)])],8,w1))),128))])]))}},I1={key:0},S1={setup(o){O();const i=R("perms"),a=R("food"),n=y([]),t=y(null);function s(r){console.log(r)}function l(r){n.value=r,t.value&&n.value.filter(f=>f.id==t.value.id).length==0&&(t.value=null)}function c(r){t.value=n.value.filter(f=>f.id==r)[0]}return(r,f)=>(p(),z(M,{current:t.value},te({filter:g(()=>[b(i).canCreateFood?(p(),$("section",I1,[e("h2",null,d(r.$t("aria.headnew")),1),m(X,{label:r.$t("btn.new"),placeholder:r.$t("food.hintnew"),onConfirm:s},null,8,["label","placeholder"])])):ne("",!0),e("section",null,[e("h2",null,d(r.$t("aria.headsearch")),1),m(a1,{data:b(a),placeholder:r.$t("food.hintsearch"),onResult:l},{default:g(v=>[m(S,{label:r.$t("food.energy"),onInput:v.confirm,name:"kcal",unit:"cal",min:"0",max:"900",frac:"0"},null,8,["label","onInput"]),m(S,{label:r.$t("food.fat"),onInput:v.confirm,name:"fat",unit:"g",min:"0",max:"100",frac:"0"},null,8,["label","onInput"]),m(S,{label:r.$t("food.carbs"),onInput:v.confirm,name:"carb",unit:"g",min:"0",max:"100",frac:"0"},null,8,["label","onInput"]),m(S,{label:r.$t("food.protein"),onInput:v.confirm,name:"prot",unit:"g",min:"0",max:"100",frac:"0"},null,8,["label","onInput"])]),_:1},8,["data","placeholder"])])]),main:g(()=>[m(M1,{items:n.value,onSelected:c},null,8,["items"])]),_:2},[t.value?{name:"details",fn:g(()=>[e("section",null,d(t.value.name),1)])}:void 0]),1032,["current"]))}},D1=_(" Recipes "),T1={setup(o){function i(a){console.log(a)}return(a,n)=>(p(),z(M,null,{filter:g(()=>[e("section",null,[m(X,{label:a.$t("btn.new"),placeholder:a.$t("recipe.hintnew"),onConfirm:i},null,8,["label","placeholder"])])]),main:g(()=>[D1]),_:1}))}},B1=_(" Diary "),A1={setup(o){return(i,a)=>(p(),z(M,null,{main:g(()=>[B1]),_:1}))}},N1=_(" Shopping Lists "),R1={setup(o){return(i,a)=>(p(),z(M,null,{main:g(()=>[N1]),_:1}))}},V1=_(" Profile "),q1={setup(o){return(i,a)=>(p(),z(M,null,{main:g(()=>[V1]),_:1}))}},H1=_(" Settings "),j1={setup(o){return(i,a)=>(p(),z(M,null,{main:g(()=>[H1]),_:1}))}},O1=ae({history:oe(),routes:[{path:"/",name:"food",component:S1},{path:"/recipes",name:"recipes",component:T1},{path:"/diary",name:"diary",component:A1},{path:"/shopping",name:"shopping",component:R1},{path:"/profile",name:"profile",component:q1},{path:"/settings",name:"settings",component:j1}]});let V=document.documentElement.lang||navigator.language;!V&&navigator.languages!=null&&(V=navigator.languages[0]);const j=document.querySelector("meta[name='_csrf']"),P1=j?j.content:"",X1=function(){const o=document.documentElement.dataset.perm||1,i=65536,a=131072;function n(t){return(o&t)==t}return{canCreateFood:n(i),canEditFood:n(a)}}(),C=se(Se);let T,B;function K1(o){T=le({locale:V.split("-")[0],fallbackLocale:"en",messages:o}),B&&K()}function U1(o){B=o,T&&K()}function K(){B.forEach(o=>{o.name=T.global.t(o.id.toString())}),C.provide("csrfToken",P1),C.provide("perms",X1),C.provide("food",B),C.use(O1),C.use(T),C.mount("#app")}fetch("/app/l10n.json").then(o=>o.json()).then(K1);fetch("api/v1/foods").then(o=>o.json()).then(U1);
