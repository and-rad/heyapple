import{o as f,c as $,a as e,u as P,i as j,b as d,w as b,d as v,t as s,e as F,R as T,f as Q,g as Y,F as K,r as y,h as A,n as S,j as O,k as Z,v as ee,l as te,m as ne,p as M,q as ae,s as R,x as oe,y as le,z as se,A as ie}from"./vendor.js";const ue=function(){const c=document.createElement("link").relList;if(c&&c.supports&&c.supports("modulepreload"))return;for(const n of document.querySelectorAll('link[rel="modulepreload"]'))a(n);new MutationObserver(n=>{for(const l of n)if(l.type==="childList")for(const t of l.addedNodes)t.tagName==="LINK"&&t.rel==="modulepreload"&&a(t)}).observe(document,{childList:!0,subtree:!0});function i(n){const l={};return n.integrity&&(l.integrity=n.integrity),n.referrerpolicy&&(l.referrerPolicy=n.referrerpolicy),n.crossorigin==="use-credentials"?l.credentials="include":n.crossorigin==="anonymous"?l.credentials="omit":l.credentials="same-origin",l}function a(n){if(n.ep)return;n.ep=!0;const l=i(n);fetch(n.href,l)}};ue();var E=(u,c)=>{const i=u.__vccOpts||u;for(const[a,n]of c)i[a]=n;return i};const ce={},re={width:"512",height:"512",version:"1.1",viewBox:"0 0 512 512",xmlns:"http://www.w3.org/2000/svg"},de=e("path",{d:"m114.6 507.3c13.82-0.8729 90.06-3.083 106.2-3.079h13.49l2.639-5.293c5.191-10.41 10.35-29.57 13.13-48.74 1.811-12.5 1.488-16.68-1.826-23.56-2.528-5.249-11.04-13.06-20.54-18.86-28.38-17.32-51.21-34.77-76.05-58.12-6.673-6.272-9.377-8.209-16.08-11.53-12.37-6.124-20.36-12.88-29.19-24.7-1.45-1.941-1.795-2.038-4.695-1.334-1.722 0.4182-7.18 0.8166-12.13 0.8854-19.6 0.2722-38.18-7.715-52.49-22.56-15.97-16.57-23.43-35.49-23.41-59.37 0.01786-16.86 3.547-29.88 11.93-44 9.617-16.19 27.39-30.13 43.89-34.41 2.517-0.653 4.748-1.36 4.958-1.571 0.2103-0.2107-0.05169-3.078-0.5816-6.371-1.41-8.759-0.556-21.84 1.997-30.61 7.023-24.1 24.96-41.95 48.52-48.31 11.42-3.082 29.45-2.167 40.42 2.049 0.919 0.3533 1.637-0.7202 2.994-4.474 6.21-17.18 18.79-33.75 33.11-43.62 35.65-24.56 80.99-19.9 111.3 11.44l5.539 5.719 5.025-1.229c2.764-0.6762 9.075-1.428 14.02-1.672 16.17-0.7944 31.64 3.865 45.31 13.64l5.1 3.647 12.55 0.1171c9.984 0.09312 13.77 0.4491 18.49 1.739 29.1 7.945 50.37 32.58 56 64.86 1.354 7.764 1.096 21.6-0.5502 29.5l-0.916 4.394 5.367 7.086c20.1 26.54 25.52 61.05 14.68 93.52-7.862 23.55-24.99 43.53-46.38 54.11-10.16 5.023-23.1 8.587-31.28 8.613-3.79 0.0118-4.555 0.333-8.217 3.446-6.958 5.916-20.59 13.71-28.3 16.17-1.722 0.551-4.011 1.599-5.087 2.328-1.076 0.7297-8.119 4.802-15.65 9.05-18.57 10.47-37.53 22.69-48.72 31.39-13.27 10.32-16.83 22.63-14.72 50.94 2.648 35.64 8.059 59.41 16.33 71.75l2.89 4.313 29.4 0.4963c34.14 0.5764 92.53 2.291 93.01 2.731 0.1793 0.166-74.3 0.2503-165.5 0.1877-91.21-0.0631-161.4-0.392-156-0.732zm130.9-101.8c-0.1092-1.007-3.147-5.947-6.75-10.98-7.889-11.01-19.71-28.99-22.37-34.01-1.594-3.014-2.402-3.727-4.695-4.139-7.259-1.304-17.91-5.324-27.58-10.41-1.809-0.9509 5.669 20.15 9.032 25.49 4.859 7.714 27.69 24.94 45.08 34 5.177 2.7 7.575 2.715 7.285 0.0442zm6.168-19.46c0.8235-0.9976 1.074-4.638 1.074-15.62 0-7.875-0.1764-14.5-0.3923-14.71-0.2158-0.2163-2.307 0.1044-4.647 0.7128-2.34 0.6082-7 1.335-10.35 1.616-7.067 0.5907-6.896 0-2.984 10.28 2.73 7.17 6.328 12.81 10.4 16.3 3.513 3.013 5.293 3.38 6.908 1.424zm49.91-15.13c3.913-2 11.37-9.412 15.96-15.85 3.982-5.592 8.688-15 8.214-16.42-0.1628-0.49-3.275-1.75-6.916-2.799-3.641-1.049-9.568-3.501-13.17-5.449l-6.551-3.541-2.173 2.087c-2.148 2.062-2.179 2.234-2.645 14.37-0.2592 6.754-0.6282 15.17-0.8198 18.7-0.3026 5.572-0.1405 6.682 1.232 8.43 1.896 2.415 2.946 2.489 6.867 0.484zm-118.1-5.725c-2.19-4.729-6.256-17.66-6.256-19.89 0-0.8683-1.901-1.089-9.391-1.089-5.165 0-9.391 0.2195-9.391 0.488 0 1.561 25.33 25.39 26.99 25.39 0.1762 0-0.7011-2.206-1.949-4.901zm149.8-9.536c5.487-3.687 19.46-14.65 19.07-14.96-0.126-0.1032-4.197-0.3888-9.046-0.6338l-8.816-0.4452-3.967 5.891c-4.284 6.363-7.088 11.99-6.473 12.98 0.9066 1.47 4.434 0.3864 9.228-2.836z",fill:"#13ad73"},null,-1),_e=e("path",{d:"m253.9 167.7c-0.2579 0.0513-1.238 0.5454-4.691 0.5216-3.625-0.0248 4.941 9.227 6.934 34.4-8.031-0.4678-8.408-3.285-24.81-5.999-15.88-2.627-25.28 5.883-29.57 18.88-3.429 10.39-3.592 23.66-1.19 36.39 2.348 12.45 7.148 24.39 13.74 32.65 6.909 8.654 14.05 10.96 20.49 10.82 8.459-0.1842 15.7-4.599 19.58-4.407 6.726 0.3331 13.29 8.894 28.17 2.151 6.047-2.741 12.03-7.819 17.08-14.81 4.984-6.907 9.051-15.68 11.35-25.91 2.018-8.972 2.678-19.07 1.41-30-1.616-13.93-7.519-20.42-14.9-23.01-8.286-2.905-18.43-0.8896-26.46 1.032-4.71 1.127-8.69 2.223-11.14 2.273-0.155 3e-3 -0.2829-4.1e-4 -0.4335 2e-3 -0.01-0.0695-0.0197-0.1385-0.0323-0.2078-2.249-11.58-4.704-29.29-5.401-33.64-0.152-0.9486 0.0845-1.177-0.1162-1.137z",fill:"#f2ac05","fill-rule":"evenodd"},null,-1),he=[de,_e];function fe(u,c){return f(),$("svg",re,he)}var me=E(ce,[["render",fe]]);const ve={},pe={width:"24",height:"24",version:"1.1",viewBox:"0 0 24 24",xmlns:"http://www.w3.org/2000/svg"},ge=e("path",{d:"m12 0c-6.624 0-12 5.376-12 12s5.376 12 12 12 12-5.376 12-12-5.376-12-12-12zm0 3.6c1.992 0 3.6 1.608 3.6 3.6s-1.608 3.6-3.6 3.6-3.6-1.608-3.6-3.6 1.608-3.6 3.6-3.6zm0 17.04c-3 0-5.652-1.536-7.2-3.864 0.036-2.388 4.8-3.696 7.2-3.696 2.388 0 7.164 1.308 7.2 3.696-1.548 2.328-4.2 3.864-7.2 3.864z"},null,-1),$e=[ge];function be(u,c){return f(),$("svg",pe,$e)}var ye=E(ve,[["render",be]]);const we={},ke={width:"24",height:"24",version:"1.1",viewBox:"0 0 24 24",xmlns:"http://www.w3.org/2000/svg"},Ce=e("path",{d:"m12 0-2.115 2.115 8.37 8.385h-18.255v3h18.255l-8.37 8.385 2.115 2.115 12-12z"},null,-1),Fe=[Ce];function Ie(u,c){return f(),$("svg",ke,Fe)}var W=E(we,[["render",Ie]]);const xe=e("div",{id:"app-name"},[e("span",null,"Hey"),e("span",null,"Apple")],-1),Le={id:"nav-main"},ze={id:"nav-user"},Ee={href:"https://docs.heyapple.org",target:"_blank"},Me={setup(u){const{t:c}=P(),i=j("csrfToken");function a(l){l.preventDefault(),fetch("/auth/local",{method:"DELETE",headers:{"X-CSRF-Token":i}}).then(t=>{t.ok?window.location="/":window.dispatchEvent(new CustomEvent("error",{detail:{msg:c("signout.err"+t.status)}}))})}function n(l){l.stopPropagation(),document.querySelector("header nav").classList.toggle("open")}return document.addEventListener("click",function(){document.querySelector("header nav").classList.remove("open")}),(l,t)=>(f(),$("header",null,[d(me,{id:"logo"}),xe,e("nav",null,[e("button",{onClick:n},[d(W)]),e("ul",Le,[e("li",null,[d(F(T),{to:"/"},{default:b(()=>[v(s(l.$t("nav.food")),1)]),_:1})]),e("li",null,[d(F(T),{to:"/recipes"},{default:b(()=>[v(s(l.$t("nav.recipes")),1)]),_:1})]),e("li",null,[d(F(T),{to:"/diary"},{default:b(()=>[v(s(l.$t("nav.diary")),1)]),_:1})]),e("li",null,[d(F(T),{to:"/shopping"},{default:b(()=>[v(s(l.$t("nav.shopping")),1)]),_:1})])]),e("ul",ze,[e("li",null,[d(F(T),{to:"/profile"},{default:b(()=>[v(s(l.$t("nav.profile")),1)]),_:1})]),e("li",null,[d(F(T),{to:"/settings"},{default:b(()=>[v(s(l.$t("nav.settings")),1)]),_:1})]),e("li",null,[e("a",Ee,s(l.$t("nav.help")),1)]),e("li",null,[e("a",{href:"#",onClick:a},s(l.$t("nav.signout")),1)])])]),e("button",{onClick:n},[d(ye)])]))}};const Se={setup(u){return Q(()=>{document.querySelector("body > .spinner-container").remove()}),(c,i)=>(f(),$(K,null,[d(Me),d(F(Y))],64))}},Te={},De={version:"1.1",viewBox:"0 0 24 24",xmlns:"http://www.w3.org/2000/svg"},Be=e("path",{id:"path-top",d:"m0 3v2h24v-2z"},null,-1),Ne=e("path",{id:"path-mid",d:"m0 13h24v-2h-24z"},null,-1),Ve=e("path",{id:"path-bottom",d:"m0 21h24v-2h-24z"},null,-1),Ae=[Be,Ne,Ve];function Re(u,c){return f(),$("svg",De,Ae)}var je=E(Te,[["render",Re]]);const qe={},He={width:"512",height:"512",version:"1.1",viewBox:"0 0 512 512",xmlns:"http://www.w3.org/2000/svg"},Oe=e("path",{d:"m251.2 4.007c-17.24-0.2277-34.69 4.924-50.29 15.67-14.33 9.866-26.9 26.43-33.11 43.62-1.357 3.754-2.073 4.828-2.992 4.475-10.97-4.217-29-5.131-40.42-2.049-23.57 6.359-41.5 24.22-48.52 48.31-2.553 8.761-3.407 21.84-1.997 30.6 0.5299 3.293 0.7902 6.162 0.5799 6.372-0.2104 0.2108-2.439 0.9186-4.956 1.572-16.5 4.283-34.28 18.21-43.89 34.41-8.382 14.11-11.91 27.14-11.93 44-0.02401 23.88 7.433 42.79 23.41 59.36 14.31 14.85 32.89 22.84 52.49 22.56 4.95-0.0689 10.41-0.4681 12.13-0.8863 2.9-0.7045 3.245-0.6061 4.695 1.334 8.833 11.82 16.82 18.58 29.19 24.7 6.707 3.32 9.412 5.257 16.09 11.53 24.85 23.35 47.68 40.8 76.05 58.12 9.503 5.8 18.01 13.61 20.54 18.86 3.313 6.878 3.637 11.06 1.825 23.56-2.778 19.18-7.936 38.33-13.13 48.74l-2.639 5.295h-13.49c-16.17-3e-3 -92.41 2.205-106.2 3.077l0.01-3e-3c-5.381 0.3399 64.84 0.6684 156 0.7315 91.2 0.0625 165.7-0.0216 165.5-0.1878-0.4753-0.4399-58.86-2.155-93-2.731l-29.4-0.4942-2.89-4.313c-8.274-12.35-13.68-36.12-16.33-71.76-2.105-28.31 1.454-40.62 14.72-50.94 11.18-8.699 30.14-20.92 48.72-31.39 7.533-4.248 14.57-8.318 15.65-9.048 1.076-0.7296 3.366-1.779 5.087-2.33 7.71-2.467 21.34-10.26 28.3-16.17 3.662-3.113 4.427-3.435 8.218-3.446 8.183-0.0267 21.12-3.59 31.28-8.613 21.39-10.58 38.52-30.56 46.38-54.11 10.84-32.47 5.42-66.98-14.68-93.52l-5.364-7.084 0.916-4.395c1.646-7.896 1.904-21.74 0.5502-29.5-5.632-32.29-26.9-56.92-56-64.86-4.726-1.29-8.51-1.647-18.49-1.74l-12.55-0.1186-5.101-3.647c-13.67-9.776-29.15-14.43-45.31-13.64-4.95 0.2432-11.26 0.9944-14.02 1.671l-5.028 1.229-5.539-5.72c-17.07-17.63-38.89-26.81-61.05-27.1zm2.728 163.7c0.2008-0.0398-0.0367 0.1881 0.1154 1.137 0.6972 4.35 3.151 22.06 5.4 33.64 0.0118 0.0694 0.0228 0.138 0.0328 0.2076 0.1506-2e-3 0.2799 9e-5 0.4349-3e-3 2.448-0.0501 6.427-1.143 11.14-2.27 8.028-1.922 18.18-3.94 26.46-1.035 7.379 2.587 13.28 9.078 14.9 23.01 1.268 10.94 0.6106 21.03-1.407 30-2.301 10.23-6.37 19.01-11.35 25.91-5.045 6.991-11.03 12.07-17.07 14.81-14.87 6.743-21.44-1.818-28.16-2.152-3.88-0.1922-11.12 4.224-19.58 4.409-6.437 0.1402-13.58-2.169-20.49-10.82-6.596-8.262-11.39-20.2-13.74-32.65-2.402-12.73-2.24-26 1.189-36.39 4.29-13 13.69-21.51 29.58-18.88 16.41 2.714 16.78 5.532 24.81 6-1.993-25.18-10.56-34.43-6.936-34.4 3.453 0.0238 4.434-0.4726 4.692-0.5239zm45.16 159.1 6.55 3.542c3.603 1.947 9.529 4.401 13.17 5.45 3.641 1.049 6.753 2.307 6.916 2.797 0.4733 1.423-4.232 10.83-8.214 16.42-4.588 6.442-12.05 13.85-15.96 15.86-3.919 2.004-4.971 1.93-6.867-0.4843-1.372-1.748-1.532-2.86-1.229-8.432 0.1916-3.529 0.558-11.94 0.8171-18.7 0.4657-12.13 0.4983-12.3 2.646-14.37zm35.38 12.77 8.817 0.4448c4.849 0.245 8.919 0.5294 9.045 0.6326 0.3835 0.315-13.58 11.27-19.07 14.96-4.795 3.222-8.322 4.307-9.229 2.837-0.6151-0.9971 2.187-6.619 6.471-12.98zm-166.7 4.596c7.489 0 9.391 0.219 9.391 1.087 0 2.234 4.067 15.16 6.257 19.89 1.248 2.694 2.123 4.9 1.947 4.9-1.66 0-26.99-23.83-26.99-25.39 0-0.2685 4.226-0.4876 9.391-0.4876zm16.13 1.727c0.0388-0.0118 0.0851-6e-3 0.1417 0.0231 9.674 5.087 20.32 9.107 27.58 10.41 2.293 0.4121 3.101 1.124 4.695 4.138 2.657 5.026 14.48 23 22.37 34.02 3.603 5.03 6.639 9.968 6.748 10.98 0.2897 2.671-2.108 2.658-7.285-0.0429-17.39-9.068-40.22-26.29-45.07-34.01-3.258-5.172-10.38-25.14-9.173-25.52zm67.9 9.71c0.263-9e-3 0.4337 9e-3 0.4876 0.0625 0.2159 0.2154 0.3921 6.837 0.3921 14.71 0 10.98-0.2506 14.62-1.074 15.62-1.615 1.956-3.397 1.59-6.909-1.423-4.068-3.49-7.665-9.13-10.4-16.3-3.912-10.28-4.082-9.686 2.985-10.28 3.355-0.2804 8.013-1.006 10.35-1.615 1.755-0.4562 3.373-0.7502 4.162-0.7776z"},null,-1),Xe=[Oe];function Ue(u,c){return f(),$("svg",He,Xe)}var Pe=E(qe,[["render",Ue]]);const Ke={},We={width:"24",height:"24",version:"1.1",viewBox:"0 0 24 24",xmlns:"http://www.w3.org/2000/svg"},Ge=e("path",{d:"m12 6c1.65 0 3-1.35 3-3s-1.35-3-3-3-3 1.35-3 3 1.35 3 3 3zm0 3c-1.65 0-3 1.35-3 3s1.35 3 3 3 3-1.35 3-3-1.35-3-3-3zm0 9c-1.65 0-3 1.35-3 3s1.35 3 3 3 3-1.35 3-3-1.35-3-3-3z"},null,-1),Je=[Ge];function Qe(u,c){return f(),$("svg",We,Je)}var Ye=E(Ke,[["render",Qe]]);const Ze={id:"filter"},et=v(" This is the main search & filter area"),tt={id:"main"},nt={class:"controls"},at=e("span",{class:"spacer"},null,-1),ot={class:"content"},lt=v("This is the main area"),st={id:"details"},it={class:"controls"},ut=e("span",{class:"spacer"},null,-1),ct={class:"placeholder"},B={emits:["detailVisibility"],setup(u,{expose:c,emit:i}){const a=y(""),n=y("");function l(){a.value==""?(a.value="open-filter",n.value="",i("detailVisibility")):a.value=""}function t(){i("detailVisibility"),n.value==""?(n.value="open-details",a.value=""):n.value=""}function r(){i("detailVisibility"),n.value==""&&(n.value="open-details",a.value="")}return c({showDetails:r}),(p,g)=>(f(),$("main",{class:S([a.value,n.value])},[e("div",Ze,[A(p.$slots,"filter",{},()=>[et])]),e("div",tt,[e("div",nt,[e("button",{onClick:l,class:"open-filter icon"},[d(je)]),at,e("button",{onClick:t,class:"open-details icon"},[d(Ye)])]),e("div",ot,[A(p.$slots,"main",{},()=>[lt])])]),e("div",st,[e("div",it,[A(p.$slots,"head-details"),ut,e("button",{onClick:t,class:"open-details icon"},[d(W)])]),A(p.$slots,"details",{},()=>[e("div",ct,[d(Pe),e("p",null,s(p.$t("details.noitem")),1)])])])],2))}},rt=["placeholder"],dt=v("Additional filters"),_t={props:["data","placeholder"],emits:["result"],setup(u,{emit:c}){const i=u;var a=void 0;function n(t){t.preventDefault(),clearTimeout(a),a=setTimeout(function(){l(t.target.closest("form"))},500)}function l(t){let r=new FormData(t),p=i.data.filter(g=>{for(let w of r.keys()){if(w=="name"){let h=r.get(w).toLowerCase();if(!g[w].toLowerCase().includes(h))return!1;continue}let[C,m]=r.getAll(w).map(h=>parseFloat(h));if(!isNaN(C)&&!isNaN(m)&&(g[w]<C||m<g[w]))return!1}return!0});c("result",p)}return(t,r)=>(f(),$("form",null,[e("input",{type:"text",name:"name",placeholder:u.placeholder,onInput:n},null,40,rt),A(t.$slots,"default",{confirm:n},()=>[dt])]))}};const ht={class:"slider"},ft=["name","value"],mt=v(" \xA0\u2013\xA0 "),vt=["name","value"],pt={class:"bar"},gt={class:"slide"},$t=e("div",{class:"interact-area"},null,-1),bt=[$t],yt=e("div",{class:"interact-area"},null,-1),wt=[yt],z={props:["label","name","unit","min","max","frac"],emits:["input"],setup(u,{emit:c}){const i=u,a=y(parseFloat(i.min).toFixed(i.frac)),n=y(parseFloat(i.max).toFixed(i.frac)),l=y(0),t=y(100);function r(m){m.target.blur();let h=parseFloat(i.min)||0,I=parseFloat(m.target.value)||h;a.value=I,n.value=Math.max(a.value,n.value);let x=parseFloat(i.max)||0;l.value=(I-h)*100/(x-h),t.value=Math.max(l.value,t.value),c("input",m)}function p(m){m.target.blur();let h=parseFloat(i.max)||0,I=parseFloat(m.target.value)||h;n.value=I,a.value=Math.min(a.value,n.value);let x=parseFloat(i.min)||0;t.value=(I-x)*100/(h-x),l.value=Math.min(l.value,t.value),c("input",m)}function g(m){let I=m.target.closest(".slide").getBoundingClientRect(),x=m.pageX!==void 0?m.pageX:m.changedTouches[0].pageX;x=Math.min(Math.max(x-I.left,0),I.width);let k=x*100/I.width,o=parseFloat(i.min)||0,L=parseFloat(i.max)||0,_=k/100*(L-o)+o;m.target.closest("button").classList.contains("min")?(a.value=_.toFixed(i.frac),n.value=Math.max(a.value,n.value),l.value=k,t.value=Math.max(l.value,t.value)):(n.value=_.toFixed(i.frac),a.value=Math.min(a.value,n.value),t.value=k,l.value=Math.min(l.value,t.value)),c("input",m)}function w(m){let h=m.target.closest("button");h.addEventListener("mousemove",g),h.addEventListener("touchmove",g),h.addEventListener("mouseup",C),h.addEventListener("touchend",C),h.addEventListener("mouseleave",C),h.addEventListener("touchcancel",C),h.classList.add("active")}function C(m){let h=m.target.closest("button");h.removeEventListener("mousemove",g),h.removeEventListener("touchmove",g),h.removeEventListener("mouseup",C),h.removeEventListener("touchend",C),h.removeEventListener("mouseleave",C),h.removeEventListener("touchcancel",C),h.classList.remove("active")}return(m,h)=>(f(),$("div",ht,[e("label",null,[e("span",null,s(u.label)+" ("+s(m.$t("unit."+u.unit))+")",1),e("input",{type:"text",name:u.name,value:a.value,onChange:r},null,40,ft),mt,e("input",{type:"text",name:u.name,value:n.value,onChange:p},null,40,vt)]),e("div",pt,[e("div",gt,[e("div",{class:"overlay",style:O({left:l.value+"%",right:100-t.value+"%"})},null,4),e("button",{type:"button",class:"handle min",style:O({left:l.value+"%"}),onMousedown:w,onTouchstart:w},bt,36),e("button",{type:"button",class:"handle max",style:O({left:t.value+"%"}),onMousedown:w,onTouchstart:w},wt,36)])])]))}};const kt={class:"clickable-input"},Ct=["placeholder"],Ft=["value"],G={props:["label","placeholder"],emits:["confirm"],setup(u,{emit:c}){const i=y("");function a(n){n.preventDefault(),c("confirm",i.value),i.value=""}return(n,l)=>(f(),$("form",kt,[Z(e("input",{type:"text","onUpdate:modelValue":l[0]||(l[0]=t=>i.value=t),placeholder:u.placeholder},null,8,Ct),[[ee,i.value]]),e("input",{type:"submit",onClick:a,value:u.label},null,8,Ft)]))}};const It={},xt={class:"icon sort-arrow"};function Lt(u,c){return f(),$("span",xt)}var V=E(It,[["render",Lt]]);const zt=["onClick"],Et={class:"name"},Mt={class:"num"},St={class:"unit"},Tt={class:"m num"},Dt={class:"unit"},Bt={class:"m num"},Nt={class:"unit"},Vt={class:"m num"},At={class:"unit"},Rt={props:["items"],emits:"selected",setup(u,{emit:c}){const i=u,a=y("name"),n=y("asc"),l=te(()=>n.value=="asc"?[...i.items].sort((r,p)=>r[a.value]<p[a.value]?-1:r[a.value]>p[a.value]?1:0):[...i.items].sort((r,p)=>r[a.value]>p[a.value]?-1:r[a.value]<p[a.value]?1:0));function t(r){let p=r.target.dataset.sort;a.value==p?n.value=n.value=="asc"?"desc":"asc":a.value=p}return(r,p)=>(f(),$("table",null,[e("thead",null,[e("tr",{class:S(n.value)},[e("th",{class:S(["name sort",{active:a.value=="name"}]),onClick:t,"data-sort":"name"},[v(s(r.$t("food.name"))+" ",1),d(V)],2),e("th",{class:S(["num sort",{active:a.value=="kcal"}]),onClick:t,"data-sort":"kcal"},[d(V),v(" "+s(r.$t("food.energy")),1)],2),e("th",{class:S(["m num sort",{active:a.value=="fat"}]),onClick:t,"data-sort":"fat"},[d(V),v(" "+s(r.$t("food.fat")),1)],2),e("th",{class:S(["m num sort",{active:a.value=="carb"}]),onClick:t,"data-sort":"carb"},[d(V),v(" "+s(r.$t("food.carbs2")),1)],2),e("th",{class:S(["m num sort",{active:a.value=="prot"}]),onClick:t,"data-sort":"prot"},[d(V),v(" "+s(r.$t("food.protein")),1)],2)],2)]),e("tbody",null,[(f(!0),$(K,null,ne(F(l),g=>(f(),$("tr",{key:g.id,onClick:w=>r.$emit("selected",g.id)},[e("td",Et,s(g.name),1),e("td",Mt,[v(s(g.kcal)+" ",1),e("span",St,s(r.$t("unit.cal")),1)]),e("td",Tt,[v(s(g.fat)+" ",1),e("span",Dt,s(r.$t("unit.g")),1)]),e("td",Bt,[v(s(g.carb)+" ",1),e("span",Nt,s(r.$t("unit.g")),1)]),e("td",Vt,[v(s(g.prot)+" ",1),e("span",At,s(r.$t("unit.g")),1)])],8,zt))),128))])]))}},jt={},qt={width:"24",height:"24",version:"1.1",viewBox:"0 0 24 24",xmlns:"http://www.w3.org/2000/svg"},Ht=e("path",{d:"m-7.5e-8 19.001v4.9993h4.9993l14.745-14.745-4.9993-4.9993zm23.61-13.611c0.51993-0.51993 0.51993-1.3598 0-1.8797l-3.1196-3.1196c-0.51993-0.51993-1.3598-0.51993-1.8797 0l-2.4397 2.4397 4.9993 4.9993z"},null,-1),Ot=[Ht];function Xt(u,c){return f(),$("svg",qt,Ot)}var Ut=E(jt,[["render",Xt]]);const Pt={},Kt={width:"24",height:"24",version:"1.1",viewBox:"0 0 24 24",xmlns:"http://www.w3.org/2000/svg"},Wt=e("path",{d:"m7.6364 17.318-5.7273-5.7273-1.9091 1.9091 7.6364 7.6364 16.364-16.364-1.9091-1.9091z"},null,-1),Gt=[Wt];function Jt(u,c){return f(),$("svg",Kt,Gt)}var Qt=E(Pt,[["render",Jt]]);const Yt={key:0,class:"new-item"},Zt=e("legend",null,"Primary Macronutrients",-1),en=e("legend",null,"Secondary Macronutrients",-1),tn=e("section",{class:"subtitle"},"Some food category",-1),nn={class:"tags"},an=e("span",{class:"tag"},"Tag 1",-1),on=e("span",{class:"tag"},"Tag 2",-1),ln=e("span",{class:"tag"},"Tag 3",-1),sn=["disabled"],un={class:"nutrients"},cn=e("h2",null,"Nutrients",-1),rn=["disabled"],dn=["value"],_n={class:"unit"},hn=["value"],fn={class:"unit"},mn=["value"],vn={class:"unit"},pn=["value"],gn={class:"unit"},$n=["value"],bn={class:"unit"},yn=["disabled"],wn=["value"],kn={class:"unit"},Cn=["value"],Fn={class:"unit"},In=["value"],xn={class:"unit"},Ln=["value"],zn={class:"unit"},En=["value"],Mn={class:"unit"},Sn={setup(u){const{t:c}=P(),i=j("csrfToken"),a=j("perms"),n=j("food"),l=y([]),t=y(null),r=y(!1),p=y(!1),g=y(null),w=y(null);function C(o){console.log(o)}function m(){p.value=!0;let o=t.value.id,L=new FormData(w.value);fetch("/api/v1/food/"+o,{method:"PUT",headers:{"Content-Type":"application/x-www-form-urlencoded","X-CSRF-Token":i},body:new URLSearchParams(L)}).then(_=>{if(!_.ok)throw c("savefood.err"+_.status);return r.value=!1,fetch("/api/v1/food/"+o)}).then(_=>_.json()).then(_=>{_.name=c(_.id.toString()),n.value=n.value.map(N=>_.id==N.id?_:N),l.value=l.value.map(N=>_.id==N.id?_:N),t.value=t.value.id==_.id?_:t.value,console.log(c("savefood.msg204"))}).catch(_=>{typeof _=="string"?console.log(_):"message"in _?console.log(_.message):console.log(c("err.err"))}).finally(()=>{setTimeout(function(){p.value=!1},500)})}function h(o){l.value=o,t.value&&l.value.filter(L=>L.id==t.value.id).length==0&&(t.value=null)}function I(o){t.value=l.value.filter(L=>L.id==o)[0],g.value.showDetails()}function x(){r.value?m():r.value=!0}function k(o){o.target.blur(),isNaN(parseFloat(o.target.value))&&(o.target.value=t.value[o.target.name])}return(o,L)=>(f(),M(B,{ref_key:"main",ref:g,onDetailVisibility:L[0]||(L[0]=_=>r.value=!1)},ae({filter:b(()=>[F(a).canCreateFood?(f(),$("section",Yt,[e("h2",null,s(o.$t("aria.headnew")),1),d(G,{label:o.$t("btn.new"),placeholder:o.$t("food.hintnew"),onConfirm:C},null,8,["label","placeholder"])])):R("",!0),e("section",null,[e("h2",null,s(o.$t("aria.headsearch")),1),d(_t,{data:F(n),placeholder:o.$t("food.hintsearch"),onResult:h},{default:b(_=>[e("fieldset",null,[Zt,d(z,{label:o.$t("food.energy"),onInput:_.confirm,name:"kcal",unit:"cal",min:"0",max:"900",frac:"0"},null,8,["label","onInput"]),d(z,{label:o.$t("food.fat"),onInput:_.confirm,name:"fat",unit:"g",min:"0",max:"100",frac:"0"},null,8,["label","onInput"]),d(z,{label:o.$t("food.carbs"),onInput:_.confirm,name:"carb",unit:"g",min:"0",max:"100",frac:"0"},null,8,["label","onInput"]),d(z,{label:o.$t("food.protein"),onInput:_.confirm,name:"prot",unit:"g",min:"0",max:"89",frac:"0"},null,8,["label","onInput"]),d(z,{label:o.$t("food.fiber"),onInput:_.confirm,name:"fib",unit:"g",min:"0",max:"71",frac:"0"},null,8,["label","onInput"])]),e("fieldset",null,[en,d(z,{label:o.$t("food.fatsat"),onInput:_.confirm,name:"fatsat",unit:"g",min:"0",max:"83",frac:"0"},null,8,["label","onInput"]),d(z,{label:o.$t("food.fato3"),onInput:_.confirm,name:"fato3",unit:"g",min:"0",max:"54",frac:"0"},null,8,["label","onInput"]),d(z,{label:o.$t("food.fato6"),onInput:_.confirm,name:"fato6",unit:"g",min:"0",max:"70",frac:"0"},null,8,["label","onInput"]),d(z,{label:o.$t("food.sugar"),onInput:_.confirm,name:"sug",unit:"g",min:"0",max:"100",frac:"0"},null,8,["label","onInput"]),d(z,{label:o.$t("food.salt"),onInput:_.confirm,name:"salt",unit:"g",min:"0",max:"100",frac:"0"},null,8,["label","onInput"])])]),_:1},8,["data","placeholder"])])]),main:b(()=>[d(Rt,{items:l.value,onSelected:I},null,8,["items"])]),_:2},[t.value?{name:"head-details",fn:b(()=>[e("h2",null,s(t.value.name),1)])}:void 0,t.value?{name:"details",fn:b(()=>[tn,e("section",nn,[an,on,ln,F(a).canCreateFood||F(a).canEditFood?(f(),$("button",{key:0,class:"icon async",disabled:p.value,onClick:x},[r.value?R("",!0):(f(),M(Ut,{key:0})),r.value?(f(),M(Qt,{key:1})):R("",!0)],8,sn)):R("",!0)]),e("section",un,[cn,e("form",{ref_key:"form",ref:w},[e("div",null,[e("fieldset",{disabled:!r.value},[e("div",null,[e("label",null,s(o.$t("food.energy")),1),e("input",{type:"text",value:t.value.kcal,name:"kcal",onChange:k},null,40,dn),e("span",_n,s(o.$t("unit.cal")),1)]),e("div",null,[e("label",null,s(o.$t("food.fat")),1),e("input",{type:"text",value:t.value.fat,name:"fat",onChange:k},null,40,hn),e("span",fn,s(o.$t("unit.g")),1)]),e("div",null,[e("label",null,s(o.$t("food.carbs2")),1),e("input",{type:"text",value:t.value.carb,name:"carb",onChange:k},null,40,mn),e("span",vn,s(o.$t("unit.g")),1)]),e("div",null,[e("label",null,s(o.$t("food.protein")),1),e("input",{type:"text",value:t.value.prot,name:"prot",onChange:k},null,40,pn),e("span",gn,s(o.$t("unit.g")),1)]),e("div",null,[e("label",null,s(o.$t("food.fiber")),1),e("input",{type:"text",value:t.value.fib,name:"fib",onChange:k},null,40,$n),e("span",bn,s(o.$t("unit.g")),1)])],8,rn),e("fieldset",{disabled:!r.value},[e("div",null,[e("label",null,s(o.$t("food.fatsat")),1),e("input",{type:"text",value:t.value.fatsat,name:"fatsat",onChange:k},null,40,wn),e("span",kn,s(o.$t("unit.g")),1)]),e("div",null,[e("label",null,s(o.$t("food.fato3")),1),e("input",{type:"text",value:t.value.fato3,name:"fato3",onChange:k},null,40,Cn),e("span",Fn,s(o.$t("unit.g")),1)]),e("div",null,[e("label",null,s(o.$t("food.fato6")),1),e("input",{type:"text",value:t.value.fato6,name:"fato6",onChange:k},null,40,In),e("span",xn,s(o.$t("unit.g")),1)]),e("div",null,[e("label",null,s(o.$t("food.sugar")),1),e("input",{type:"text",value:t.value.sug,name:"sug",onChange:k},null,40,Ln),e("span",zn,s(o.$t("unit.g")),1)]),e("div",null,[e("label",null,s(o.$t("food.salt")),1),e("input",{type:"text",value:t.value.salt,name:"salt",onChange:k},null,40,En),e("span",Mn,s(o.$t("unit.g")),1)])],8,yn)])],512)])])}:void 0]),1536))}},Tn={class:"new-item"},Dn=v(" Recipes "),Bn={setup(u){function c(i){console.log(i)}return(i,a)=>(f(),M(B,null,{filter:b(()=>[e("section",Tn,[e("h2",null,s(i.$t("aria.headnew")),1),d(G,{label:i.$t("btn.new"),placeholder:i.$t("recipe.hintnew"),onConfirm:c},null,8,["label","placeholder"])]),e("section",null,[e("h2",null,s(i.$t("aria.headsearch")),1)])]),main:b(()=>[Dn]),_:1}))}},Nn=v(" Diary "),Vn={setup(u){return(c,i)=>(f(),M(B,null,{main:b(()=>[Nn]),_:1}))}},An=v(" Shopping Lists "),Rn={setup(u){return(c,i)=>(f(),M(B,null,{main:b(()=>[An]),_:1}))}},jn=v(" Profile "),qn={setup(u){return(c,i)=>(f(),M(B,null,{main:b(()=>[jn]),_:1}))}},Hn=v(" Settings "),On={setup(u){return(c,i)=>(f(),M(B,null,{main:b(()=>[Hn]),_:1}))}},Xn=oe({history:le(),routes:[{path:"/",name:"food",component:Sn},{path:"/recipes",name:"recipes",component:Bn},{path:"/diary",name:"diary",component:Vn},{path:"/shopping",name:"shopping",component:Rn},{path:"/profile",name:"profile",component:qn},{path:"/settings",name:"settings",component:On}]});let X=document.documentElement.lang||navigator.language;!X&&navigator.languages!=null&&(X=navigator.languages[0]);const U=document.querySelector("meta[name='_csrf']"),Un=U?U.content:"",Pn=function(){const u=document.documentElement.dataset.perm||1,c=65536,i=131072;function a(n){return(u&n)==n}return{canCreateFood:a(c),canEditFood:a(i)}}(),D=se(Se);let q,H;function Kn(u){q=ie({locale:X.split("-")[0],fallbackLocale:"en",messages:u}),H&&J()}function Wn(u){H=u,q&&J()}function J(){H.forEach(u=>{u.name=q.global.t(u.id.toString())}),D.provide("csrfToken",Un),D.provide("perms",Pn),D.provide("food",y(H)),D.use(Xn),D.use(q),D.mount("#app")}fetch("/app/l10n.json").then(u=>u.json()).then(Kn);fetch("api/v1/foods").then(u=>u.json()).then(Wn);
