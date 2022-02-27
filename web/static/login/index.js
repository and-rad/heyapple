import{o as g,c as U,u as v,R as D,a as h,F as k,b as e,d as z,r as p,e as B,t as i,f as m,n as G,g as R,w,v as y,h as f,i as x,j as b,k as M,l as O,m as P,p as T,q as V,s as C,x as N}from"./vendor.js";const F=function(){const s=document.createElement("link").relList;if(s&&s.supports&&s.supports("modulepreload"))return;for(const t of document.querySelectorAll('link[rel="modulepreload"]'))r(t);new MutationObserver(t=>{for(const a of t)if(a.type==="childList")for(const l of a.addedNodes)l.tagName==="LINK"&&l.rel==="modulepreload"&&r(l)}).observe(document,{childList:!0,subtree:!0});function u(t){const a={};return t.integrity&&(a.integrity=t.integrity),t.referrerpolicy&&(a.referrerPolicy=t.referrerpolicy),t.crossorigin==="use-credentials"?a.credentials="include":t.crossorigin==="anonymous"?a.credentials="omit":a.credentials="same-origin",a}function r(t){if(t.ep)return;t.ep=!0;const a=u(t);fetch(t.href,a)}};F();const E={setup(n){return(s,u)=>(g(),U(v(D)))}};var L=(n,s)=>{const u=n.__vccOpts||n;for(const[r,t]of s)u[r]=t;return u};const I={},j=e("svg",{version:"1.1",viewBox:"0 0 512 1024",xmlns:"http://www.w3.org/2000/svg","xmlns:xlink":"http://www.w3.org/1999/xlink"},[e("defs",{id:"defs2"},[e("filter",{id:"filter1988",x:"-.15817",y:"-.14127",width:"1.3163",height:"1.2825","color-interpolation-filters":"sRGB"},[e("feGaussianBlur",{id:"feGaussianBlur1990",stdDeviation:"7.5344244"})]),e("filter",{id:"filter3416",x:"-.044885",y:"-.043258",width:"1.0898",height:"1.0865","color-interpolation-filters":"sRGB"},[e("feGaussianBlur",{id:"feGaussianBlur3418",stdDeviation:"9.0841236"})]),e("linearGradient",{id:"linearGradient14145",x1:"-512",x2:"260.58",y1:"-256",y2:"-256",gradientTransform:"translate(512)",gradientUnits:"userSpaceOnUse"},[e("stop",{id:"stop14139","stop-color":"#cad0db",offset:"0"}),e("stop",{id:"stop14141","stop-color":"#eeeef1",offset:"1"})])]),e("g",{id:"day"},[e("path",{transform:"rotate(90)",d:"m0-512h1024v512h-1024z",fill:"url(#linearGradient14145)"}),e("path",{transform:"rotate(90)",d:"m768-512h256v512h-256z",fill:"#dadfcb"}),e("path",{transform:"matrix(.59152 0 0 .59277 104.57 630.63)",d:"m114.29 507.25c13.845-0.87292 90.246-3.0826 106.45-3.0786h13.517l2.6446-5.2934c5.202-10.413 10.372-29.566 13.156-48.742 1.8152-12.503 1.4908-16.681-1.8294-23.56-2.5338-5.249-11.059-13.062-20.582-18.862-28.438-17.323-51.315-34.767-76.215-58.12-6.6875-6.2717-9.3966-8.2094-16.118-11.53-12.397-6.1236-20.402-12.884-29.253-24.702-1.4532-1.9405-1.7986-2.0384-4.7052-1.334-1.7252 0.41819-7.1951 0.8166-12.155 0.8854-19.638 0.27216-38.261-7.7154-52.601-22.562-16.008-16.574-23.48-35.485-23.456-59.366 0.0176-16.864 3.5545-29.885 11.954-43.999 9.6372-16.195 27.448-30.128 43.987-34.411 2.5219-0.65305 4.7576-1.3598 4.9684-1.5706 0.2108-0.21067-0.05166-3.0776-0.58275-6.3707-1.4128-8.7588-0.55717-21.845 2.0015-30.605 7.0381-24.096 25.011-41.954 48.627-48.314 11.443-3.0816 29.517-2.1672 40.505 2.0494 0.92093 0.3533 1.6409-0.72022 3.0004-4.4745 6.2228-17.184 18.826-33.752 33.182-43.618 35.73-24.556 81.166-19.899 111.58 11.438l5.5509 5.7188 5.0358-1.2294c2.7696-0.67611 9.0941-1.4285 14.054-1.6716 16.2-0.79443 31.709 3.8647 45.409 13.641l5.1106 3.6468 12.578 0.11718c10.005 0.0932 13.796 0.44906 18.532 1.7393 29.165 7.9446 50.477 32.576 56.12 64.865 1.3569 7.7636 1.0982 21.604-0.55135 29.5l-0.91799 4.3935 5.3784 7.086c20.145 26.541 25.574 61.055 14.712 93.524-7.8782 23.548-25.045 43.533-46.478 54.11-10.18 5.0232-23.149 8.5867-31.349 8.6134-3.7983 0.0126-4.5646 0.33301-8.2342 3.4464-6.9726 5.9156-20.63 13.705-28.356 16.172-1.7253 0.551-4.0192 1.5988-5.0973 2.3285-1.0783 0.72968-8.1361 4.8024-15.685 9.0501-18.613 10.475-37.611 22.69-48.819 31.389-13.296 10.319-16.861 22.632-14.752 50.944 2.6541 35.638 8.0757 59.405 16.367 71.755l2.8962 4.3131 29.46 0.49632c34.213 0.57645 92.726 2.2911 93.202 2.7311 0.17965 0.16581-74.457 0.25023-165.85 0.18774-91.402-0.063-161.77-0.39199-156.38-0.73194zm131.15-101.76c-0.10962-1.0073-3.1533-5.9466-6.7638-10.976-7.906-11.014-19.751-28.989-22.413-34.015-1.5972-3.0144-2.4072-3.727-4.7054-4.1391-7.2745-1.3045-17.945-5.3239-27.639-10.411-1.8124-0.95092 5.6807 20.154 9.0507 25.492 4.8697 7.714 27.748 24.937 45.171 34.005 5.1883 2.7004 7.5909 2.7147 7.3006 0.0441zm6.1807-19.462c0.8253-0.99768 1.0762-4.638 1.0762-15.619 0-7.8748-0.17691-14.495-0.39312-14.711-0.21622-0.21634-2.3117 0.10458-4.6568 0.71278-2.345 0.6082-7.0144 1.3354-10.377 1.6157-7.0821 0.59068-6.9109 0-2.9902 10.277 2.7358 7.1695 6.3413 12.81 10.418 16.301 3.5201 3.0132 5.3042 3.3801 6.9224 1.4242zm50.012-15.128c3.9214-2.0005 11.399-9.4118 15.997-15.855 3.9906-5.5923 8.7061-15 8.2318-16.423-0.1633-0.49002-3.2822-1.7496-6.9306-2.799-3.6486-1.0492-9.5879-3.5012-13.199-5.4486l-6.565-3.5409-2.178 2.0868c-2.1521 2.0617-2.1838 2.2341-2.6505 14.367-0.25968 6.7542-0.62949 15.168-0.82152 18.697-0.30328 5.5719-0.14075 6.6822 1.2342 8.4302 1.8996 2.4149 2.9518 2.489 6.8818 0.48396zm-118.34-5.7248c-2.1947-4.7287-6.2691-17.655-6.2691-19.889 0-0.86826-1.9055-1.0886-9.4104-1.0886-5.1757 0-9.4104 0.21949-9.4104 0.488 0 1.5611 25.38 25.39 27.043 25.39 0.17665 0-0.70258-2.2056-1.9536-4.9013zm150.12-9.5361c5.4984-3.687 19.499-14.645 19.114-14.96-0.12651-0.10333-4.2056-0.38884-9.0648-0.63378l-8.8346-0.44516-3.9754 5.8913c-4.2936 6.3628-7.1031 11.986-6.4868 12.983 0.90859 1.4702 4.443 0.38643 9.2479-2.8356z",filter:"url(#filter3416)",opacity:".27123","stroke-width":"1.6888"}),e("path",{d:"m172.17 931.32c8.1895-0.51743 53.382-1.8273 62.965-1.8249h7.9957l1.5643-3.1378c3.0771-6.1722 6.1351-17.526 7.7818-28.893 1.0737-7.4114 0.88184-9.888-1.0821-13.965-1.4988-3.1115-6.5415-7.743-12.175-11.181-16.822-10.268-30.354-20.609-45.082-34.452-3.9558-3.7177-5.5582-4.8663-9.534-6.8346-7.333-3.6299-12.068-7.6369-17.304-14.643-0.85959-1.1503-1.0639-1.2083-2.7832-0.79073-1.0205 0.24792-4.256 0.48407-7.19 0.52485-11.616 0.16135-22.632-4.5734-31.115-13.374-9.4692-9.8246-13.889-21.035-13.875-35.19 0.0106-9.9964 2.1025-17.715 7.0709-26.081 5.7006-9.5998 16.236-17.859 26.019-20.398 1.4917-0.38711 2.8142-0.80605 2.9389-0.931 0.12468-0.12487-0.0306-1.8243-0.34473-3.7764-0.83571-5.1919-0.32958-12.949 1.1839-18.142 4.1632-14.284 14.794-24.869 28.764-28.639 6.769-1.8267 17.46-1.2846 23.96 1.2148 0.54474 0.20943 0.97062-0.42691 1.7748-2.6523 3.6809-10.186 11.136-20.007 19.628-25.855 21.135-14.556 48.011-11.796 66.003 6.7801l3.2835 3.3899 2.9788-0.72873c1.6383-0.40081 5.3793-0.84675 8.3132-0.9909 9.5825-0.47092 18.757 2.2908 26.86 8.0858l3.023 2.1617 7.4399 0.0694c5.9182 0.0552 8.1604 0.26619 10.962 1.031 17.252 4.7093 29.858 19.31 33.196 38.45 0.80264 4.602 0.64963 12.806-0.32612 17.487l-0.543 2.6043 3.1814 4.2004c11.916 15.732 15.128 36.191 8.7023 55.438-4.6601 13.959-14.815 25.805-27.492 32.074-6.0214 2.9776-13.693 5.0899-18.543 5.1058-2.2467 7e-3 -2.7 0.1974-4.8707 2.0429-4.1244 3.5066-12.203 8.1239-16.773 9.5863-1.0206 0.32661-2.3774 0.94773-3.0152 1.3802-0.63782 0.43253-4.8126 2.8467-9.2776 5.3646-11.01 6.209-22.248 13.45-28.877 18.606-7.8645 6.1168-9.9738 13.416-8.7261 30.198 1.5699 21.125 4.7769 35.214 9.6816 42.534l1.7132 2.5567 17.426 0.2942c20.237 0.3417 54.849 1.3581 55.131 1.6189 0.10628 0.0984-44.043 0.14835-98.105 0.11127-54.066-0.0374-95.691-0.23237-92.501-0.43387zm77.58-60.321c-0.0647-0.59708-1.8652-3.525-4.0009-6.5064-4.6765-6.5285-11.683-17.184-13.258-20.163-0.94477-1.7869-1.4239-2.2092-2.7833-2.4536-4.303-0.77326-10.615-3.1558-16.349-6.1712-1.072-0.56368 3.3603 11.946 5.3537 15.111 2.8805 4.5726 16.413 14.782 26.719 20.157 3.069 1.6008 4.4901 1.6092 4.3184 0.0262zm3.656-11.536c0.48817-0.59138 0.63656-2.7492 0.63656-9.2584 0-4.6679-0.10456-8.5922-0.23253-8.7199-0.12794-0.12824-1.3674 0.0619-2.7546 0.4225-1.3871 0.36053-4.1491 0.79156-6.1379 0.95775-4.1892 0.35015-4.0879 0-1.7688 6.0917 1.6183 4.2499 3.751 7.5936 6.1626 9.6625 2.0822 1.7861 3.1375 2.0036 4.0947 0.84421zm29.583-8.9672c2.3196-1.1858 6.7426-5.579 9.4625-9.3981 2.3605-3.3149 5.1498-8.8917 4.8693-9.735-0.0965-0.29047-1.9415-1.0371-4.0996-1.6592-2.1582-0.62191-5.6714-2.0754-7.8071-3.2298l-3.8833-2.0989-1.2884 1.237c-1.273 1.2221-1.2918 1.3243-1.5678 8.516-0.15361 4.0037-0.37236 8.991-0.48596 11.083-0.17939 3.3028-0.0833 3.961 0.73003 4.9971 1.1236 1.4315 1.746 1.4754 4.0707 0.28688zm-70.001-3.3935c-1.2982-2.803-3.7083-10.465-3.7083-11.79 0-0.51469-1.1271-0.64532-5.5664-0.64532-3.0615 0-5.5664 0.13011-5.5664 0.28926 0 0.92539 15.013 15.051 15.997 15.051 0.10446 0-0.41557-1.3074-1.1556-2.9053zm88.796-5.6527c3.2524-2.1856 11.534-8.6811 11.306-8.8678-0.0747-0.0612-2.4877-0.23048-5.362-0.37568l-5.2258-0.26388-2.3516 3.4922c-2.5397 3.7716-4.2016 7.1049-3.837 7.6959 0.53743 0.87146 2.6281 0.22906 5.4703-1.6809z",fill:"#13ad73","stroke-width":".99999"}),e("path",{transform:"matrix(.59152 0 0 .59277 140.53 626.83)",d:"m193.13 174.15c-0.25843 0.0513-1.2403 0.54538-4.7008 0.5216-3.6326-0.025 4.9518 9.2271 6.9489 34.403-8.0479-0.46778-8.4257-3.2854-24.866-5.9993-15.917-2.6275-25.338 5.8833-29.636 18.883-3.4366 10.393-3.5997 23.655-1.1922 36.388 2.3533 12.447 7.1628 24.388 13.773 32.65 6.9238 8.6544 14.083 10.964 20.534 10.824 8.4768-0.18419 15.73-4.5989 19.618-4.4067 6.7401 0.33311 13.32 8.8942 28.225 2.1514 6.0595-2.7411 12.057-7.8194 17.112-14.81 4.9947-6.9067 9.0697-15.68 11.375-25.912 2.0218-8.972 2.6837-19.065 1.4126-30.004-1.6191-13.934-7.5347-20.424-14.929-23.011-8.3037-2.9053-18.473-0.88964-26.518 1.0321-4.7195 1.1273-8.7079 2.2227-11.161 2.2727-0.15558 3e-3 -0.28348-4.1e-4 -0.43441 2e-3 -0.0105-0.0693-0.019-0.13852-0.0325-0.20769-2.2542-11.578-4.7139-29.288-5.4127-33.639-0.15235-0.94859 0.0845-1.1766-0.11653-1.1368z","fill-rule":"evenodd",filter:"url(#filter1988)",opacity:".31422"}),e("path",{d:"m254.76 730.07c-0.15287 0.0304-0.73365 0.32329-2.7806 0.30919-2.1488-0.0147 2.9291 5.4696 4.1104 20.393-4.7605-0.27728-4.9839-1.9475-14.709-3.5562-9.4153-1.5575-14.988 3.4874-17.53 11.193-2.0328 6.1606-2.1293 14.022-0.70522 21.57 1.392 7.3779 4.2369 14.456 8.1467 19.354 4.0955 5.1301 8.3304 6.499 12.146 6.4159 5.0142-0.10919 9.3044-2.726 11.604-2.6121 3.9869 0.19746 7.8789 5.2722 16.696 1.2753 3.5843-1.6248 7.1318-4.6351 10.122-8.7791 2.9544-4.0941 5.3649-9.2946 6.7288-15.36 1.1959-5.3183 1.5874-11.301 0.8356-17.786-0.95776-8.2596-4.4569-12.107-8.8309-13.64-4.9118-1.7222-10.927-0.52734-15.686 0.61181-2.7917 0.66825-5.1509 1.3175-6.6017 1.3472-0.0919 2e-3 -0.16767-2.4e-4 -0.25695 1e-3 -6e-3 -0.0412-0.0117-0.0821-0.0192-0.12317-1.3334-6.8632-2.7884-17.361-3.2017-19.94-0.0901-0.5623 0.0501-0.69749-0.0689-0.67385z",fill:"#f2ac05","fill-rule":"evenodd"})])],-1),A=e("svg",{id:"sun",version:"1.1",viewBox:"0 0 512 128",xmlns:"http://www.w3.org/2000/svg","xmlns:xlink":"http://www.w3.org/1999/xlink"},[e("path",{d:"m512 128a128 128 0 0 1-90.51-37.49 128 128 0 0 1-37.49-90.51h128z",fill:"#f1e8db","stroke-width":"6"})],-1);function q(n,s){return g(),h(k,null,[j,A],64)}var S=L(I,[["render",q]]);const H=e("input",{type:"password",name:"pass"},null,-1),K=f("."),W=["value"],J={class:"image"},Q={setup(n){const{t:s}=z(),u=/^[^@]+@[^@]+$/,r=p(""),t=p({msg:"",level:""});function a(l){if(l.preventDefault(),t.value={},r.value.search(u)==-1)return t.value={msg:s("login.errmail"),level:"err"},!1;fetch("/auth/local",{method:"POST",headers:{"Content-Type":"application/x-www-form-urlencoded"},body:new URLSearchParams(new FormData(l.target.closest("form")))}).then(c=>{c.ok?window.location="/app":t.value={msg:s("login.err"+c.status),level:"err"}})}return(l,c)=>{const o=B("Message");return g(),h(k,null,[e("div",null,[e("form",null,[e("h1",null,i(l.$t("login.title")),1),m(o,G(R(t.value)),null,16),e("label",null,i(l.$t("form.email")),1),w(e("input",{type:"email",name:"email","onUpdate:modelValue":c[0]||(c[0]=d=>r.value=d)},null,512),[[y,r.value]]),e("label",null,[f(i(l.$t("form.pass"))+" ",1),m(v(x),{to:"/reset"},{default:b(()=>[f(i(l.$t("form.reset")),1)]),_:1})]),H,e("footer",null,[e("span",null,[f(i(l.$t("login.signup"))+" ",1),m(v(x),{to:"/signup"},{default:b(()=>[f(i(l.$t("register.action")),1)]),_:1}),K])]),e("input",{type:"submit",value:l.$t("login.action"),onClick:a},null,8,W)])]),e("div",J,[m(S)])],64)}}},X={},Y={version:"1.1",viewBox:"0 0 512 1024",xmlns:"http://www.w3.org/2000/svg","xmlns:xlink":"http://www.w3.org/1999/xlink"},Z=M(`<defs id="defs2"><filter id="filter30864" x="-.071351" y="-.078107" width="1.1427" height="1.1562" color-interpolation-filters="sRGB"><feGaussianBlur id="feGaussianBlur30866" stdDeviation="4.1668536"></feGaussianBlur></filter><linearGradient id="linearGradient3290" x1="-512" x2="158.74" y1="-256" y2="-256" gradientTransform="translate(512)" gradientUnits="userSpaceOnUse"><stop id="stop3284" stop-color="#c9d6df" offset="0"></stop><stop id="stop3286" stop-color="#dfd6cc" offset="1"></stop></linearGradient></defs><g id="morning"><rect transform="rotate(90)" y="-512" width="1024" height="512" fill="url(#linearGradient3290)"></rect><ellipse cx="261.39" cy="765.21" rx="128.33" ry="127.45" fill="#eee2d3" stroke-width="3"></ellipse><rect transform="rotate(90)" x="768" y="-512" width="256" height="512" fill="#bcc2b1"></rect><path d="m243.13 929.49 0.63824-0.35909 0.58171-0.12013 0.53023-0.29457s0.39449-0.23367 0.59105-0.23367 0.69664-0.21416 1.1113-0.47592c0.41465-0.26175 1.1161-0.57345 1.5588-0.69267 0.44272-0.11921 0.85872-0.30375 0.92444-0.41009 0.23279-0.37666 2.4006-1.7046 2.9612-1.8139 0.31292-0.061 1.2694 0.0209 2.1254 0.18208 2.2163 0.41728 2.5052 0.39411 3.4842-0.27948 1.1652-0.80166 2.2051-0.99847 2.7362-0.51784 0.30806 0.27879 0.43158 0.30307 0.55184 0.10848 0.37884-0.61296 0.99041-12.716 0.98882-19.568-5.9e-4 -2.1602-0.0836-5.057-0.18454-6.4375-0.10099-1.3804-0.27511-5.4013-0.38692-8.9353-0.30641-9.6846-0.58482-15.346-0.90987-18.501-0.31205-3.0288-0.47239-3.4837-1.8274-5.1842-0.33458-0.41988-0.60831-0.85387-0.60831-0.96441 0-0.48847-3.3957-4.7058-5.7779-7.176-1.5893-1.648-5.6569-5.0649-8.2762-6.9524-1.8802-1.3548-1.8828-1.3558-6.8604-2.5842-1.0801-0.26655-3.4219-0.79809-5.2041-1.1812-5.524-1.1875-11.527-3.1671-18.032-5.9467-7.3708-3.1495-9.0493-3.6442-15.802-4.6576-5.3287-0.79976-5.7642-0.88895-8.7401-1.7899-2.7549-0.83402-4.621-1.0742-8.3446-1.0742-2.9128 0-3.7058-0.23367-1.569-0.46233 1.1932-0.12769 7.4908-0.61052 10.658-0.81713 1.0261-0.0669 3.567-0.29629 5.6465-0.50968 2.0795-0.2134 3.8911-0.37023 4.0258-0.34854 0.13468 0.0217 0.40684-0.031 0.60481-0.11697 0.55144-0.23964 10.264 0.0333 13.583 0.38182 9.4618 0.9933 18.646 2.8728 22.997 4.7065 4.3216 1.8211 7.5546 4.2865 8.6861 6.624 0.79301 1.6381 0.47695 3.3819-0.8865 4.8911-0.40251 0.44553-0.63595 0.85801-0.53558 0.94631 0.0979 0.0861 0.86805 0.66839 1.7115 1.294 2.966 2.1998 7.7518 6.8871 12.401 12.146l1.3581 1.5361 0.30693-0.58791 0.30693-0.5879 0.46968 0.38229c0.25832 0.21025 0.50642 0.34553 0.55134 0.30061 0.0449-0.0449 0.25881-0.95293 0.47532-2.0178 0.44901-2.2084 2.1523-7.5815 3.4152-10.773 1.0054-2.5411 1.1802-3.5738 0.82795-4.8924-0.25529-0.95565-0.27182-0.97072-1.7454-1.5905-4.1834-1.7596-8.9754-6.4546-11.118-10.893-0.62722-1.2992-0.72269-1.6966-0.79517-3.3101-0.19229-4.2805 1.1546-8.4171 4.2944-13.189 1.8486-2.8098 3.4222-4.8103 5.3756-6.8341 2.6824-2.779 5.0157-4.3526 7.7738-5.2425 1.5379-0.49623 2.0762-0.78941 3.6755-2.0018 0.77939-0.59084 1.57-1.0448 1.757-1.0088 0.23632 0.0455 0.40294 0.3802 0.54673 1.0982 0.2574 1.2854 1.3121 2.44 3.0988 3.3922 1.3609 0.72536 2.8836 2.0542 3.5039 3.0579 0.6 0.97083 1.3611 3.0174 1.3611 3.66 0 0.30992 0.0663 0.52265 0.14728 0.47272 0.081-0.0499 0.15693 0.36463 0.16871 0.92122 0.0189 0.89357 0.33915 2.1592 0.85237 3.3686 1.0397 2.4501 1.445 4.4081 1.4786 7.1436l0.0323 2.6269-0.74856 1.4637c-0.41171 0.80505-1.2396 2.0973-1.8397 2.8718-1.5681 2.0235-9.2567 10.52-10.303 11.386-0.48604 0.40206-1.5465 0.9843-2.3566 1.2938-1.2895 0.49275-1.554 0.68188-2.1242 1.519-1.271 1.8659-3.739 7.037-4.2523 8.9096-1.3616 4.967-1.6444 6.7087-1.7967 11.063-0.11486 3.285 0.19499 19.838 0.48767 26.053l0.12485 2.6511 0.47525-1.1783c0.7681-1.9043 3.3006-6.8985 4.6693-9.208 1.4338-2.4193 1.6818-3.3718 1.5058-5.7832-0.13945-1.9108 0.39228-3.6798 2.1454-7.1378 3.3974-6.7012 6.1363-9.3515 13.43-12.996 3.0068-1.5022 7.1403-3.1542 10.939-4.3718 5.4053-1.7326 9.1712-2.4583 14.163-2.729 2.022-0.10965 2.5996-0.22156 4.0684-0.78831 1.9106-0.73723 2.9626-0.98489 3.2994-0.77675 0.14399 0.089-0.031 0.32174-0.48422 0.64425-1.2802 0.91088-1.5652 1.6173-1.6801 4.1649-0.0939 2.0811-0.15728 2.374-0.80724 3.7312-0.85769 1.791-2.4388 3.9541-5.4552 7.4629-1.2537 1.4584-2.7213 3.1784-3.2614 3.8222-3.326 3.9653-7.634 6.9849-12.513 8.7713-1.9985 0.73165-5.3754 1.6251-8.945 2.3665-2.6696 0.5545-3.2789 0.60797-7.9004 0.69339-4.295 0.0794-5.0608 0.13889-5.381 0.41819-0.20531 0.17911-1.0717 1.3788-1.9254 2.6659-1.3066 1.9701-2.027 3.3112-3.7008 6.8896-0.26483 0.56616-0.73776 1.9282-1.2956 3.7312-0.39596 1.2799-0.49854 2.0806-0.55351 4.3204-0.21275 8.6691-0.5192 18.336-0.68166 21.504-0.26032 5.0749-0.28927 4.8219 0.54722 4.7826 0.37974-0.0179 0.91734 0.0583 1.1946 0.1691 0.85408 0.34143 2.0847 0.50643 4.2845 0.57447 0.83073 0.0257 1.2274 0.1173 1.2274 0.28344 0 0.13501-0.10538 0.2806-0.23417 0.32353-0.12879 0.0429 0.15841 0.0677 0.63823 0.0551s1.1375 0.11033 1.
4616 0.27325c0.32402 0.16292 0.659 0.27686 0.74438 0.25319 0.0854-0.0236 0.32487 0.0478 0.53221 0.15872 0.20732 0.11096 0.88863 0.23916 1.514 0.28489 0.86687 0.0634 1.1242 0.15198 1.0833 0.37288-0.0392 0.21116 0.066 0.25843 0.38769 0.17429 0.24281-0.0635 0.33327-0.1515 0.20106-0.19557-0.51332-0.17111-0.20267-0.38661 0.55729-0.38661 0.73833 0 1.1264 0.75874 1.1264 0.75874l0.57695 0.75874 0.24825 0.22421z" filter="url(#filter30864)" opacity=".20938"></path><path d="m243.13 929.49 0.63824-0.35909 0.58171-0.12013 0.53023-0.29457s0.39449-0.23367 0.59105-0.23367 0.69664-0.21416 1.1113-0.47592c0.41465-0.26175 1.1161-0.57345 1.5588-0.69267 0.44272-0.11921 0.85872-0.30375 0.92444-0.41009 0.23279-0.37666 2.4006-1.7046 2.9612-1.8139 0.31292-0.061 1.2694 0.0209 2.1254 0.18208 2.2163 0.41728 2.5052 0.39411 3.4842-0.27948 1.1652-0.80166 2.2051-0.99847 2.7362-0.51784 0.30806 0.27879 0.43158 0.30307 0.55184 0.10848 0.37884-0.61296 0.99041-12.716 0.98882-19.568-5.9e-4 -2.1602-0.0836-5.057-0.18454-6.4375-0.10099-1.3804-0.27511-5.4013-0.38692-8.9353-0.30641-9.6846-0.58482-15.346-0.90987-18.501-0.31205-3.0288-0.47239-3.4837-1.8274-5.1842-0.33458-0.41988-0.60831-0.85387-0.60831-0.96441 0-0.48847-3.3957-4.7058-5.7779-7.176-1.5893-1.648-5.6569-5.0649-8.2762-6.9524-1.8802-1.3548-1.8828-1.3558-6.8604-2.5842-1.0801-0.26655-3.4219-0.79809-5.2041-1.1812-5.524-1.1875-11.527-3.1671-18.032-5.9467-7.3708-3.1495-9.0493-3.6442-15.802-4.6576-5.3287-0.79976-5.7642-0.88895-8.7401-1.7899-2.7549-0.83402-4.621-1.0742-8.3446-1.0742-2.9128 0-3.7058-0.23367-1.569-0.46233 1.1932-0.12769 7.4908-0.61052 10.658-0.81713 1.0261-0.0669 3.567-0.29629 5.6465-0.50968 2.0795-0.2134 3.8911-0.37023 4.0258-0.34854 0.13468 0.0217 0.40684-0.031 0.60481-0.11697 0.55144-0.23964 10.264 0.0333 13.583 0.38182 9.4618 0.9933 18.646 2.8728 22.997 4.7065 4.3216 1.8211 7.5546 4.2865 8.6861 6.624 0.79301 1.6381 0.47695 3.3819-0.8865 4.8911-0.40251 0.44553-0.63595 0.85801-0.53558 0.94631 0.0979 0.0861 0.86805 0.66839 1.7115 1.294 2.966 2.1998 7.7518 6.8871 12.401 12.146l1.3581 1.5361 0.30693-0.58791 0.30693-0.5879 0.46968 0.38229c0.25832 0.21025 0.50642 0.34553 0.55134 0.30061 0.0449-0.0449 0.25881-0.95293 0.47532-2.0178 0.44901-2.2084 2.1523-7.5815 3.4152-10.773 1.0054-2.5411 1.1802-3.5738 0.82795-4.8924-0.25529-0.95565-0.27182-0.97072-1.7454-1.5905-4.1834-1.7596-8.9754-6.4546-11.118-10.893-0.62722-1.2992-0.72269-1.6966-0.79517-3.3101-0.19229-4.2805 1.1546-8.4171 4.2944-13.189 1.8486-2.8098 3.4222-4.8103 5.3756-6.8341 2.6824-2.779 5.0157-4.3526 7.7738-5.2425 1.5379-0.49623 2.0762-0.78941 3.6755-2.0018 0.77939-0.59084 1.57-1.0448 1.757-1.0088 0.23632 0.0455 0.40294 0.3802 0.54673 1.0982 0.2574 1.2854 1.3121 2.44 3.0988 3.3922 1.3609 0.72536 2.8836 2.0542 3.5039 3.0579 0.6 0.97083 1.3611 3.0174 1.3611 3.66 0 0.30992 0.0663 0.52265 0.14728 0.47272 0.081-0.0499 0.15693 0.36463 0.16871 0.92122 0.0189 0.89357 0.33915 2.1592 0.85237 3.3686 1.0397 2.4501 1.445 4.4081 1.4786 7.1436l0.0323 2.6269-0.74856 1.4637c-0.41171 0.80505-1.2396 2.0973-1.8397 2.8718-1.5681 2.0235-9.2567 10.52-10.303 11.386-0.48604 0.40206-1.5465 0.9843-2.3566 1.2938-1.2895 0.49275-1.554 0.68188-2.1242 1.519-1.271 1.8659-3.739 7.037-4.2523 8.9096-1.3616 4.967-1.6444 6.7087-1.7967 11.063-0.11486 3.285 0.19499 19.838 0.48767 26.053l0.12485 2.6511 0.47525-1.1783c0.7681-1.9043 3.3006-6.8985 4.6693-9.208 1.4338-2.4193 1.6818-3.3718 1.5058-5.7832-0.13945-1.9108 0.39228-3.6798 2.1454-7.1378 3.3974-6.7012 6.1363-9.3515 13.43-12.996 3.0068-1.5022 7.1403-3.1542 10.939-4.3718 5.4053-1.7326 9.1712-2.4583 14.163-2.729 2.022-0.10965 2.5996-0.22156 4.0684-0.78831 1.9106-0.73723 2.9626-0.98489 3.2994-0.77675 0.14399 0.089-0.031 0.32174-0.48422 0.64425-1.2802 0.91088-1.5652 1.6173-1.6801 4.1649-0.0939 2.0811-0.15728 2.374-0.80724 3.7312-0.85769 1.791-2.4388 3.9541-5.4552 7.4629-1.2537 1.4584-2.7213 3.1784-3.2614 3.8222-3.326 3.9653-7.634 6.9849-12.513 8.7713-1.9985 0.73165-5.3754 1.6251-8.945 2.3665-2.6696 0.5545-3.2789 0.60797-7.9004 0.69339-4.295 0.0794-5.0608 0.13889-5.381 0.41819-0.20531 0.17911-1.0717 1.3788-1.9254 2.6659-1.3066 1.9701-2.027 3.3112-3.7008 6.8896-0.26483 0.56616-0.73776 1.9282-1.2956 3.7312-0.39596 1.2799-0.49854 2.0806-0.55351 4.3204-0.21275 8.6691-0.5192 18.336-0.68166 21.504-0.26032 5.0749-0.28927 4.8219 0.54722 4.7826 0.37974-0.0179 0.91734 0.0583 1.1946 0.1691 0.85408 0.34143 2.0847 0.50643 4.2845 0.57447 0.83073 0.0257 1.2274 0.1173 1.2274 0.28344 0 0.13501-0.10538 0.2806-0.23417 0.32353-0.12879 0.0429 0.15841 0.0677 0.63823 0.0551s1.1375 0.11033 1.
4616 0.27325c0.32402 0.16292 0.659 0.27686 0.74438 0.25319 0.0854-0.0236 0.32487 0.0478 0.53221 0.15872 0.20732 0.11096 0.88863 0.23916 1.514 0.28489 0.86687 0.0634 1.1242 0.15198 1.0833 0.37288-0.0392 0.21116 0.066 0.25843 0.38769 0.17429 0.24281-0.0635 0.33327-0.1515 0.20106-0.19557-0.51332-0.17111-0.20267-0.38661 0.55729-0.38661 0.73833 0 1.1264 0.75874 1.1264 0.75874l0.57695 0.75874 0.24825 0.22421z" fill="#12a66f"></path></g>`,2),e1=[Z];function t1(n,s){return g(),h("svg",Y,e1)}var l1=L(X,[["render",t1]]);const s1=f("."),r1=["value"],a1={class:"image register-image"},n1={setup(n){const{t:s}=z(),u=/^[^@]+@[^@]+$/,r=p(""),t=p(""),a=p(""),l=p({msg:"",level:""});function c(o){if(o.preventDefault(),l.value={},r.value.search(u)==-1)return l.value={msg:s("login.errmail"),level:"err"},!1;if(t.value==""||a.value=="")return l.value={msg:s("register.errpassempty"),level:"err"},!1;if(t.value!=a.value){l.value={msg:s("register.errpassmatch"),level:"err"};return}let d=r.value;fetch("/api/v1/user",{method:"POST",headers:{"Content-Type":"application/x-www-form-urlencoded"},body:new URLSearchParams(new FormData(o.target.closest("form")))}).then($=>{$.ok?l.value={msg:s("register.success",{addr:d})}:l.value={msg:s("register.err"+$.status),level:"err"}})}return(o,d)=>{const $=B("Message");return g(),h(k,null,[e("div",null,[e("form",null,[e("h1",null,i(o.$t("register.title")),1),m($,G(R(l.value)),null,16),e("label",null,i(o.$t("form.email")),1),w(e("input",{type:"email",name:"email","onUpdate:modelValue":d[0]||(d[0]=_=>r.value=_)},null,512),[[y,r.value]]),e("label",null,i(o.$t("form.pass")),1),w(e("input",{type:"password",name:"pass","onUpdate:modelValue":d[1]||(d[1]=_=>t.value=_)},null,512),[[y,t.value]]),e("label",null,i(o.$t("form.confirm")),1),w(e("input",{type:"password",name:"pass","onUpdate:modelValue":d[2]||(d[2]=_=>a.value=_)},null,512),[[y,a.value]]),e("footer",null,[e("span",null,[f(i(o.$t("register.signin"))+" ",1),m(v(x),{to:"/"},{default:b(()=>[f(i(o.$t("login.action")),1)]),_:1}),s1])]),e("input",{type:"submit",value:o.$t("register.action"),onClick:c},null,8,r1)])]),e("div",a1,[m(l1)])],64)}}};const o1=["value"],i1={class:"image"},u1={setup(n){const{t:s}=z(),u=/^[^@]+@[^@]+$/,r=p(""),t=p({msg:"",level:""});function a(l){if(l.preventDefault(),t.value={},r.value.search(u)==-1)return t.value={msg:s("login.errmail"),level:"err"},!1;let c=r.value;fetch("/auth/reset",{method:"POST",headers:{"Content-Type":"application/x-www-form-urlencoded"},body:new URLSearchParams(new FormData(l.target.closest("form")))}).then(o=>{o.ok?t.value={msg:s("reset.success",{addr:c})}:t.value={msg:s("reset.err"+o.status),level:"err"}})}return(l,c)=>{const o=B("Message");return g(),h(k,null,[e("div",null,[e("form",null,[e("h1",null,i(l.$t("reset.title")),1),m(o,G(R(t.value)),null,16),e("p",null,i(l.$t("reset.hint")),1),e("label",null,i(l.$t("form.email")),1),w(e("input",{type:"email",name:"email","onUpdate:modelValue":c[0]||(c[0]=d=>r.value=d)},null,512),[[y,r.value]]),e("input",{type:"submit",value:l.$t("reset.action"),onClick:a},null,8,o1)]),m(v(x),{to:"/",class:"back"},{default:b(()=>[f("\u1438 "+i(l.$t("form.back")),1)]),_:1})]),e("div",i1,[m(S)])],64)}}},c1=O({history:P(),routes:[{path:"/",name:"login",component:Q},{path:"/signup",name:"signup",component:n1},{path:"/reset",name:"reset",component:u1}]});const d1={props:{msg:String,level:Number},setup(n){const s=n,u=T(()=>s.msg);return(r,t)=>(g(),h("div",{class:V(["message",[{hidden:!v(u)},n.level]])},i(n.msg),3))}};fetch("/login/l10n.json").then(n=>n.json()).then(n=>{let s=navigator.language;navigator.languages!=null&&(s=navigator.languages[0]);const u=C({locale:s.split("-")[0],fallbackLocale:"en",messages:n}),r=N(E);r.component("Message",d1),r.use(c1),r.use(u),r.mount("#app")});
