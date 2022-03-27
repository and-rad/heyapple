import { o as openBlock, c as createElementBlock, a as createBaseVNode, u as useI18n, i as inject, b as createVNode, w as withCtx, d as createTextVNode, t as toDisplayString, e as unref, R as RouterLink, f as onMounted, g as createBlock, h as createCommentVNode, n as normalizeClass, r as ref, F as Fragment, j as renderList, k as RouterView, l as computed, m as renderSlot, p as watch, q as normalizeStyle, s as withDirectives, v as vModelText, x as vModelSelect, D as DateTime, y as createSlots, z as onBeforeMount, A as onUnmounted, B as createRouter, C as createWebHashHistory, S as Settings, E as createApp, G as createI18n } from "./vendor.js";
const p = function polyfill() {
  const relList = document.createElement("link").relList;
  if (relList && relList.supports && relList.supports("modulepreload")) {
    return;
  }
  for (const link of document.querySelectorAll('link[rel="modulepreload"]')) {
    processPreload(link);
  }
  new MutationObserver((mutations) => {
    for (const mutation of mutations) {
      if (mutation.type !== "childList") {
        continue;
      }
      for (const node of mutation.addedNodes) {
        if (node.tagName === "LINK" && node.rel === "modulepreload")
          processPreload(node);
      }
    }
  }).observe(document, { childList: true, subtree: true });
  function getFetchOpts(script) {
    const fetchOpts = {};
    if (script.integrity)
      fetchOpts.integrity = script.integrity;
    if (script.referrerpolicy)
      fetchOpts.referrerPolicy = script.referrerpolicy;
    if (script.crossorigin === "use-credentials")
      fetchOpts.credentials = "include";
    else if (script.crossorigin === "anonymous")
      fetchOpts.credentials = "omit";
    else
      fetchOpts.credentials = "same-origin";
    return fetchOpts;
  }
  function processPreload(link) {
    if (link.ep)
      return;
    link.ep = true;
    const fetchOpts = getFetchOpts(link);
    fetch(link.href, fetchOpts);
  }
};
p();
var _export_sfc = (sfc, props) => {
  const target = sfc.__vccOpts || sfc;
  for (const [key, val] of props) {
    target[key] = val;
  }
  return target;
};
const _sfc_main$x = {};
const _hoisted_1$v = {
  width: "512",
  height: "512",
  version: "1.1",
  viewBox: "0 0 512 512",
  xmlns: "http://www.w3.org/2000/svg"
};
const _hoisted_2$r = /* @__PURE__ */ createBaseVNode("path", {
  d: "m114.6 507.3c13.82-0.8729 90.06-3.083 106.2-3.079h13.49l2.639-5.293c5.191-10.41 10.35-29.57 13.13-48.74 1.811-12.5 1.488-16.68-1.826-23.56-2.528-5.249-11.04-13.06-20.54-18.86-28.38-17.32-51.21-34.77-76.05-58.12-6.673-6.272-9.377-8.209-16.08-11.53-12.37-6.124-20.36-12.88-29.19-24.7-1.45-1.941-1.795-2.038-4.695-1.334-1.722 0.4182-7.18 0.8166-12.13 0.8854-19.6 0.2722-38.18-7.715-52.49-22.56-15.97-16.57-23.43-35.49-23.41-59.37 0.01786-16.86 3.547-29.88 11.93-44 9.617-16.19 27.39-30.13 43.89-34.41 2.517-0.653 4.748-1.36 4.958-1.571 0.2103-0.2107-0.05169-3.078-0.5816-6.371-1.41-8.759-0.556-21.84 1.997-30.61 7.023-24.1 24.96-41.95 48.52-48.31 11.42-3.082 29.45-2.167 40.42 2.049 0.919 0.3533 1.637-0.7202 2.994-4.474 6.21-17.18 18.79-33.75 33.11-43.62 35.65-24.56 80.99-19.9 111.3 11.44l5.539 5.719 5.025-1.229c2.764-0.6762 9.075-1.428 14.02-1.672 16.17-0.7944 31.64 3.865 45.31 13.64l5.1 3.647 12.55 0.1171c9.984 0.09312 13.77 0.4491 18.49 1.739 29.1 7.945 50.37 32.58 56 64.86 1.354 7.764 1.096 21.6-0.5502 29.5l-0.916 4.394 5.367 7.086c20.1 26.54 25.52 61.05 14.68 93.52-7.862 23.55-24.99 43.53-46.38 54.11-10.16 5.023-23.1 8.587-31.28 8.613-3.79 0.0118-4.555 0.333-8.217 3.446-6.958 5.916-20.59 13.71-28.3 16.17-1.722 0.551-4.011 1.599-5.087 2.328-1.076 0.7297-8.119 4.802-15.65 9.05-18.57 10.47-37.53 22.69-48.72 31.39-13.27 10.32-16.83 22.63-14.72 50.94 2.648 35.64 8.059 59.41 16.33 71.75l2.89 4.313 29.4 0.4963c34.14 0.5764 92.53 2.291 93.01 2.731 0.1793 0.166-74.3 0.2503-165.5 0.1877-91.21-0.0631-161.4-0.392-156-0.732zm130.9-101.8c-0.1092-1.007-3.147-5.947-6.75-10.98-7.889-11.01-19.71-28.99-22.37-34.01-1.594-3.014-2.402-3.727-4.695-4.139-7.259-1.304-17.91-5.324-27.58-10.41-1.809-0.9509 5.669 20.15 9.032 25.49 4.859 7.714 27.69 24.94 45.08 34 5.177 2.7 7.575 2.715 7.285 0.0442zm6.168-19.46c0.8235-0.9976 1.074-4.638 1.074-15.62 0-7.875-0.1764-14.5-0.3923-14.71-0.2158-0.2163-2.307 0.1044-4.647 0.7128-2.34 0.6082-7 1.335-10.35 1.616-7.067 0.5907-6.896 0-2.984 10.28 2.73 7.17 6.328 12.81 10.4 16.3 3.513 3.013 5.293 3.38 6.908 1.424zm49.91-15.13c3.913-2 11.37-9.412 15.96-15.85 3.982-5.592 8.688-15 8.214-16.42-0.1628-0.49-3.275-1.75-6.916-2.799-3.641-1.049-9.568-3.501-13.17-5.449l-6.551-3.541-2.173 2.087c-2.148 2.062-2.179 2.234-2.645 14.37-0.2592 6.754-0.6282 15.17-0.8198 18.7-0.3026 5.572-0.1405 6.682 1.232 8.43 1.896 2.415 2.946 2.489 6.867 0.484zm-118.1-5.725c-2.19-4.729-6.256-17.66-6.256-19.89 0-0.8683-1.901-1.089-9.391-1.089-5.165 0-9.391 0.2195-9.391 0.488 0 1.561 25.33 25.39 26.99 25.39 0.1762 0-0.7011-2.206-1.949-4.901zm149.8-9.536c5.487-3.687 19.46-14.65 19.07-14.96-0.126-0.1032-4.197-0.3888-9.046-0.6338l-8.816-0.4452-3.967 5.891c-4.284 6.363-7.088 11.99-6.473 12.98 0.9066 1.47 4.434 0.3864 9.228-2.836z",
  fill: "#13ad73"
}, null, -1);
const _hoisted_3$o = /* @__PURE__ */ createBaseVNode("path", {
  d: "m253.9 167.7c-0.2579 0.0513-1.238 0.5454-4.691 0.5216-3.625-0.0248 4.941 9.227 6.934 34.4-8.031-0.4678-8.408-3.285-24.81-5.999-15.88-2.627-25.28 5.883-29.57 18.88-3.429 10.39-3.592 23.66-1.19 36.39 2.348 12.45 7.148 24.39 13.74 32.65 6.909 8.654 14.05 10.96 20.49 10.82 8.459-0.1842 15.7-4.599 19.58-4.407 6.726 0.3331 13.29 8.894 28.17 2.151 6.047-2.741 12.03-7.819 17.08-14.81 4.984-6.907 9.051-15.68 11.35-25.91 2.018-8.972 2.678-19.07 1.41-30-1.616-13.93-7.519-20.42-14.9-23.01-8.286-2.905-18.43-0.8896-26.46 1.032-4.71 1.127-8.69 2.223-11.14 2.273-0.155 3e-3 -0.2829-4.1e-4 -0.4335 2e-3 -0.01-0.0695-0.0197-0.1385-0.0323-0.2078-2.249-11.58-4.704-29.29-5.401-33.64-0.152-0.9486 0.0845-1.177-0.1162-1.137z",
  fill: "#f2ac05",
  "fill-rule": "evenodd"
}, null, -1);
const _hoisted_4$f = [
  _hoisted_2$r,
  _hoisted_3$o
];
function _sfc_render$a(_ctx, _cache) {
  return openBlock(), createElementBlock("svg", _hoisted_1$v, _hoisted_4$f);
}
var HeaderImage = /* @__PURE__ */ _export_sfc(_sfc_main$x, [["render", _sfc_render$a]]);
const _sfc_main$w = {};
const _hoisted_1$u = {
  width: "24",
  height: "24",
  version: "1.1",
  viewBox: "0 0 24 24",
  xmlns: "http://www.w3.org/2000/svg"
};
const _hoisted_2$q = /* @__PURE__ */ createBaseVNode("path", { d: "m12 0c-6.624 0-12 5.376-12 12s5.376 12 12 12 12-5.376 12-12-5.376-12-12-12zm0 3.6c1.992 0 3.6 1.608 3.6 3.6s-1.608 3.6-3.6 3.6-3.6-1.608-3.6-3.6 1.608-3.6 3.6-3.6zm0 17.04c-3 0-5.652-1.536-7.2-3.864 0.036-2.388 4.8-3.696 7.2-3.696 2.388 0 7.164 1.308 7.2 3.696-1.548 2.328-4.2 3.864-7.2 3.864z" }, null, -1);
const _hoisted_3$n = [
  _hoisted_2$q
];
function _sfc_render$9(_ctx, _cache) {
  return openBlock(), createElementBlock("svg", _hoisted_1$u, _hoisted_3$n);
}
var ProfileImage = /* @__PURE__ */ _export_sfc(_sfc_main$w, [["render", _sfc_render$9]]);
const _sfc_main$v = {};
const _hoisted_1$t = {
  width: "24",
  height: "24",
  version: "1.1",
  viewBox: "0 0 24 24",
  xmlns: "http://www.w3.org/2000/svg"
};
const _hoisted_2$p = /* @__PURE__ */ createBaseVNode("path", { d: "m12 0-2.115 2.115 8.37 8.385h-18.255v3h18.255l-8.37 8.385 2.115 2.115 12-12z" }, null, -1);
const _hoisted_3$m = [
  _hoisted_2$p
];
function _sfc_render$8(_ctx, _cache) {
  return openBlock(), createElementBlock("svg", _hoisted_1$t, _hoisted_3$m);
}
var ArrowImage = /* @__PURE__ */ _export_sfc(_sfc_main$v, [["render", _sfc_render$8]]);
var Header_vue_vue_type_style_index_0_lang = "";
const _hoisted_1$s = /* @__PURE__ */ createBaseVNode("div", { id: "app-name" }, [
  /* @__PURE__ */ createBaseVNode("span", null, "Hey"),
  /* @__PURE__ */ createBaseVNode("span", null, [
    /* @__PURE__ */ createTextVNode("Apple"),
    /* @__PURE__ */ createBaseVNode("sup", null, "beta")
  ])
], -1);
const _hoisted_2$o = { id: "nav-main" };
const _hoisted_3$l = { id: "nav-user" };
const _hoisted_4$e = {
  href: "https://docs.heyapple.org",
  target: "_blank"
};
const _sfc_main$u = {
  setup(__props) {
    const { t } = useI18n();
    const csrf = inject("csrfToken");
    function confirm(evt) {
      evt.preventDefault();
      fetch("/auth/local", {
        method: "DELETE",
        headers: { "X-CSRF-Token": csrf }
      }).then((response) => {
        if (response.ok) {
          window.location = "/";
        } else {
          window.dispatchEvent(new CustomEvent("error", {
            detail: { msg: t("signout.err" + response.status) }
          }));
        }
      });
    }
    function toggleMenu(evt) {
      evt.stopPropagation();
      document.querySelector("header nav").classList.toggle("open");
    }
    document.addEventListener("click", function() {
      document.querySelector("header nav").classList.remove("open");
    });
    return (_ctx, _cache) => {
      return openBlock(), createElementBlock("header", null, [
        createVNode(HeaderImage, { id: "logo" }),
        _hoisted_1$s,
        createBaseVNode("nav", null, [
          createBaseVNode("button", { onClick: toggleMenu }, [
            createVNode(ArrowImage)
          ]),
          createBaseVNode("ul", _hoisted_2$o, [
            createBaseVNode("li", null, [
              createVNode(unref(RouterLink), { to: "/" }, {
                default: withCtx(() => [
                  createTextVNode(toDisplayString(unref(t)("nav.diary")), 1)
                ]),
                _: 1
              })
            ]),
            createBaseVNode("li", null, [
              createVNode(unref(RouterLink), { to: "/food" }, {
                default: withCtx(() => [
                  createTextVNode(toDisplayString(unref(t)("nav.food")), 1)
                ]),
                _: 1
              })
            ]),
            createBaseVNode("li", null, [
              createVNode(unref(RouterLink), { to: "/recipes" }, {
                default: withCtx(() => [
                  createTextVNode(toDisplayString(unref(t)("nav.recipes")), 1)
                ]),
                _: 1
              })
            ]),
            createBaseVNode("li", null, [
              createVNode(unref(RouterLink), { to: "/shopping" }, {
                default: withCtx(() => [
                  createTextVNode(toDisplayString(unref(t)("nav.shopping")), 1)
                ]),
                _: 1
              })
            ])
          ]),
          createBaseVNode("ul", _hoisted_3$l, [
            createBaseVNode("li", null, [
              createVNode(unref(RouterLink), { to: "/profile" }, {
                default: withCtx(() => [
                  createTextVNode(toDisplayString(unref(t)("nav.profile")), 1)
                ]),
                _: 1
              })
            ]),
            createBaseVNode("li", null, [
              createVNode(unref(RouterLink), { to: "/settings" }, {
                default: withCtx(() => [
                  createTextVNode(toDisplayString(unref(t)("nav.settings")), 1)
                ]),
                _: 1
              })
            ]),
            createBaseVNode("li", null, [
              createBaseVNode("a", _hoisted_4$e, toDisplayString(unref(t)("nav.help")), 1)
            ]),
            createBaseVNode("li", null, [
              createBaseVNode("a", {
                href: "#",
                onClick: confirm
              }, toDisplayString(unref(t)("nav.signout")), 1)
            ])
          ])
        ]),
        createBaseVNode("button", { onClick: toggleMenu }, [
          createVNode(ProfileImage)
        ])
      ]);
    };
  }
};
const _sfc_main$t = {};
const _hoisted_1$r = {
  width: "24",
  height: "24",
  version: "1.1",
  viewBox: "0 0 24 24",
  xmlns: "http://www.w3.org/2000/svg"
};
const _hoisted_2$n = /* @__PURE__ */ createBaseVNode("path", { d: "m7.6364 17.318-5.7273-5.7273-1.9091 1.9091 7.6364 7.6364 16.364-16.364-1.9091-1.9091z" }, null, -1);
const _hoisted_3$k = [
  _hoisted_2$n
];
function _sfc_render$7(_ctx, _cache) {
  return openBlock(), createElementBlock("svg", _hoisted_1$r, _hoisted_3$k);
}
var SaveImage = /* @__PURE__ */ _export_sfc(_sfc_main$t, [["render", _sfc_render$7]]);
const _sfc_main$s = {};
const _hoisted_1$q = {
  width: "24",
  height: "24",
  version: "1.1",
  viewBox: "0 0 24 24",
  xmlns: "http://www.w3.org/2000/svg"
};
const _hoisted_2$m = /* @__PURE__ */ createBaseVNode("path", { d: "m12 1.6364-12 20.727h24zm0 4.353 8.2159 14.192h-16.43zm-1.0909 4.3743v5.4545h2.1818v-5.4545zm0 6.5455v2.1818h2.1818v-2.1818z" }, null, -1);
const _hoisted_3$j = [
  _hoisted_2$m
];
function _sfc_render$6(_ctx, _cache) {
  return openBlock(), createElementBlock("svg", _hoisted_1$q, _hoisted_3$j);
}
var WarnImage = /* @__PURE__ */ _export_sfc(_sfc_main$s, [["render", _sfc_render$6]]);
var Message_vue_vue_type_style_index_0_lang = "";
const _sfc_main$r = {
  props: ["msg"],
  emits: ["timeout"],
  setup(__props, { emit }) {
    const prop = __props;
    onMounted(() => {
      setTimeout(function() {
        emit("timeout", prop.msg.id);
      }, prop.msg.time);
    });
    return (_ctx, _cache) => {
      return openBlock(), createElementBlock("div", {
        class: normalizeClass(["message", [__props.msg.type, __props.msg.id]])
      }, [
        __props.msg.type == "message" ? (openBlock(), createBlock(SaveImage, { key: 0 })) : createCommentVNode("", true),
        __props.msg.type != "message" ? (openBlock(), createBlock(WarnImage, { key: 1 })) : createCommentVNode("", true),
        createBaseVNode("p", null, toDisplayString(__props.msg.msg), 1)
      ], 2);
    };
  }
};
var MessageList_vue_vue_type_style_index_0_lang = "";
const _hoisted_1$p = { id: "messages" };
const _sfc_main$q = {
  setup(__props) {
    const messages = ref([]);
    let nextId = 0;
    function onMessage(evt) {
      messages.value.push({
        id: nextId++,
        type: evt.type,
        msg: evt.detail.msg,
        time: evt.detail.timeout
      });
    }
    function onTimeout(id) {
      messages.value = messages.value.filter((m) => m.id != id);
    }
    onMounted(() => {
      window.addEventListener("message", onMessage);
      window.addEventListener("warning", onMessage);
      window.addEventListener("error", onMessage);
    });
    return (_ctx, _cache) => {
      return openBlock(), createElementBlock("div", _hoisted_1$p, [
        (openBlock(true), createElementBlock(Fragment, null, renderList(messages.value, (msg) => {
          return openBlock(), createBlock(_sfc_main$r, {
            key: msg.id,
            msg,
            onTimeout
          }, null, 8, ["msg"]);
        }), 128))
      ]);
    };
  }
};
var App_vue_vue_type_style_index_0_lang = "";
const _sfc_main$p = {
  setup(__props) {
    onMounted(() => {
      document.querySelector("body > .spinner-container").remove();
    });
    return (_ctx, _cache) => {
      return openBlock(), createElementBlock(Fragment, null, [
        createVNode(_sfc_main$u),
        createVNode(unref(RouterView)),
        createVNode(_sfc_main$q)
      ], 64);
    };
  }
};
const _sfc_main$o = {};
const _hoisted_1$o = {
  version: "1.1",
  viewBox: "0 0 24 24",
  xmlns: "http://www.w3.org/2000/svg"
};
const _hoisted_2$l = /* @__PURE__ */ createBaseVNode("path", {
  id: "path-top",
  d: "m0 3v2h24v-2z"
}, null, -1);
const _hoisted_3$i = /* @__PURE__ */ createBaseVNode("path", {
  id: "path-mid",
  d: "m0 13h24v-2h-24z"
}, null, -1);
const _hoisted_4$d = /* @__PURE__ */ createBaseVNode("path", {
  id: "path-bottom",
  d: "m0 21h24v-2h-24z"
}, null, -1);
const _hoisted_5$a = [
  _hoisted_2$l,
  _hoisted_3$i,
  _hoisted_4$d
];
function _sfc_render$5(_ctx, _cache) {
  return openBlock(), createElementBlock("svg", _hoisted_1$o, _hoisted_5$a);
}
var MenuImage = /* @__PURE__ */ _export_sfc(_sfc_main$o, [["render", _sfc_render$5]]);
const _sfc_main$n = {};
const _hoisted_1$n = {
  width: "512",
  height: "512",
  version: "1.1",
  viewBox: "0 0 512 512",
  xmlns: "http://www.w3.org/2000/svg"
};
const _hoisted_2$k = /* @__PURE__ */ createBaseVNode("path", { d: "m251.2 4.007c-17.24-0.2277-34.69 4.924-50.29 15.67-14.33 9.866-26.9 26.43-33.11 43.62-1.357 3.754-2.073 4.828-2.992 4.475-10.97-4.217-29-5.131-40.42-2.049-23.57 6.359-41.5 24.22-48.52 48.31-2.553 8.761-3.407 21.84-1.997 30.6 0.5299 3.293 0.7902 6.162 0.5799 6.372-0.2104 0.2108-2.439 0.9186-4.956 1.572-16.5 4.283-34.28 18.21-43.89 34.41-8.382 14.11-11.91 27.14-11.93 44-0.02401 23.88 7.433 42.79 23.41 59.36 14.31 14.85 32.89 22.84 52.49 22.56 4.95-0.0689 10.41-0.4681 12.13-0.8863 2.9-0.7045 3.245-0.6061 4.695 1.334 8.833 11.82 16.82 18.58 29.19 24.7 6.707 3.32 9.412 5.257 16.09 11.53 24.85 23.35 47.68 40.8 76.05 58.12 9.503 5.8 18.01 13.61 20.54 18.86 3.313 6.878 3.637 11.06 1.825 23.56-2.778 19.18-7.936 38.33-13.13 48.74l-2.639 5.295h-13.49c-16.17-3e-3 -92.41 2.205-106.2 3.077l0.01-3e-3c-5.381 0.3399 64.84 0.6684 156 0.7315 91.2 0.0625 165.7-0.0216 165.5-0.1878-0.4753-0.4399-58.86-2.155-93-2.731l-29.4-0.4942-2.89-4.313c-8.274-12.35-13.68-36.12-16.33-71.76-2.105-28.31 1.454-40.62 14.72-50.94 11.18-8.699 30.14-20.92 48.72-31.39 7.533-4.248 14.57-8.318 15.65-9.048 1.076-0.7296 3.366-1.779 5.087-2.33 7.71-2.467 21.34-10.26 28.3-16.17 3.662-3.113 4.427-3.435 8.218-3.446 8.183-0.0267 21.12-3.59 31.28-8.613 21.39-10.58 38.52-30.56 46.38-54.11 10.84-32.47 5.42-66.98-14.68-93.52l-5.364-7.084 0.916-4.395c1.646-7.896 1.904-21.74 0.5502-29.5-5.632-32.29-26.9-56.92-56-64.86-4.726-1.29-8.51-1.647-18.49-1.74l-12.55-0.1186-5.101-3.647c-13.67-9.776-29.15-14.43-45.31-13.64-4.95 0.2432-11.26 0.9944-14.02 1.671l-5.028 1.229-5.539-5.72c-17.07-17.63-38.89-26.81-61.05-27.1zm2.728 163.7c0.2008-0.0398-0.0367 0.1881 0.1154 1.137 0.6972 4.35 3.151 22.06 5.4 33.64 0.0118 0.0694 0.0228 0.138 0.0328 0.2076 0.1506-2e-3 0.2799 9e-5 0.4349-3e-3 2.448-0.0501 6.427-1.143 11.14-2.27 8.028-1.922 18.18-3.94 26.46-1.035 7.379 2.587 13.28 9.078 14.9 23.01 1.268 10.94 0.6106 21.03-1.407 30-2.301 10.23-6.37 19.01-11.35 25.91-5.045 6.991-11.03 12.07-17.07 14.81-14.87 6.743-21.44-1.818-28.16-2.152-3.88-0.1922-11.12 4.224-19.58 4.409-6.437 0.1402-13.58-2.169-20.49-10.82-6.596-8.262-11.39-20.2-13.74-32.65-2.402-12.73-2.24-26 1.189-36.39 4.29-13 13.69-21.51 29.58-18.88 16.41 2.714 16.78 5.532 24.81 6-1.993-25.18-10.56-34.43-6.936-34.4 3.453 0.0238 4.434-0.4726 4.692-0.5239zm45.16 159.1 6.55 3.542c3.603 1.947 9.529 4.401 13.17 5.45 3.641 1.049 6.753 2.307 6.916 2.797 0.4733 1.423-4.232 10.83-8.214 16.42-4.588 6.442-12.05 13.85-15.96 15.86-3.919 2.004-4.971 1.93-6.867-0.4843-1.372-1.748-1.532-2.86-1.229-8.432 0.1916-3.529 0.558-11.94 0.8171-18.7 0.4657-12.13 0.4983-12.3 2.646-14.37zm35.38 12.77 8.817 0.4448c4.849 0.245 8.919 0.5294 9.045 0.6326 0.3835 0.315-13.58 11.27-19.07 14.96-4.795 3.222-8.322 4.307-9.229 2.837-0.6151-0.9971 2.187-6.619 6.471-12.98zm-166.7 4.596c7.489 0 9.391 0.219 9.391 1.087 0 2.234 4.067 15.16 6.257 19.89 1.248 2.694 2.123 4.9 1.947 4.9-1.66 0-26.99-23.83-26.99-25.39 0-0.2685 4.226-0.4876 9.391-0.4876zm16.13 1.727c0.0388-0.0118 0.0851-6e-3 0.1417 0.0231 9.674 5.087 20.32 9.107 27.58 10.41 2.293 0.4121 3.101 1.124 4.695 4.138 2.657 5.026 14.48 23 22.37 34.02 3.603 5.03 6.639 9.968 6.748 10.98 0.2897 2.671-2.108 2.658-7.285-0.0429-17.39-9.068-40.22-26.29-45.07-34.01-3.258-5.172-10.38-25.14-9.173-25.52zm67.9 9.71c0.263-9e-3 0.4337 9e-3 0.4876 0.0625 0.2159 0.2154 0.3921 6.837 0.3921 14.71 0 10.98-0.2506 14.62-1.074 15.62-1.615 1.956-3.397 1.59-6.909-1.423-4.068-3.49-7.665-9.13-10.4-16.3-3.912-10.28-4.082-9.686 2.985-10.28 3.355-0.2804 8.013-1.006 10.35-1.615 1.755-0.4562 3.373-0.7502 4.162-0.7776z" }, null, -1);
const _hoisted_3$h = [
  _hoisted_2$k
];
function _sfc_render$4(_ctx, _cache) {
  return openBlock(), createElementBlock("svg", _hoisted_1$n, _hoisted_3$h);
}
var DetailsImage = /* @__PURE__ */ _export_sfc(_sfc_main$n, [["render", _sfc_render$4]]);
const _sfc_main$m = {};
const _hoisted_1$m = {
  width: "24",
  height: "24",
  version: "1.1",
  viewBox: "0 0 24 24",
  xmlns: "http://www.w3.org/2000/svg"
};
const _hoisted_2$j = /* @__PURE__ */ createBaseVNode("path", { d: "m12 6c1.65 0 3-1.35 3-3s-1.35-3-3-3-3 1.35-3 3 1.35 3 3 3zm0 3c-1.65 0-3 1.35-3 3s1.35 3 3 3 3-1.35 3-3-1.35-3-3-3zm0 9c-1.65 0-3 1.35-3 3s1.35 3 3 3 3-1.35 3-3-1.35-3-3-3z" }, null, -1);
const _hoisted_3$g = [
  _hoisted_2$j
];
function _sfc_render$3(_ctx, _cache) {
  return openBlock(), createElementBlock("svg", _hoisted_1$m, _hoisted_3$g);
}
var MoreImage = /* @__PURE__ */ _export_sfc(_sfc_main$m, [["render", _sfc_render$3]]);
var Main_vue_vue_type_style_index_0_lang = "";
const _hoisted_1$l = { id: "filter" };
const _hoisted_2$i = /* @__PURE__ */ createTextVNode(" This is the main search & filter area");
const _hoisted_3$f = { id: "main" };
const _hoisted_4$c = { class: "controls" };
const _hoisted_5$9 = /* @__PURE__ */ createBaseVNode("span", { class: "spacer" }, null, -1);
const _hoisted_6$8 = { class: "content" };
const _hoisted_7$7 = /* @__PURE__ */ createTextVNode("This is the main area");
const _hoisted_8$7 = { id: "details" };
const _hoisted_9$7 = { class: "controls" };
const _hoisted_10$7 = /* @__PURE__ */ createBaseVNode("span", { class: "spacer" }, null, -1);
const _hoisted_11$5 = { class: "placeholder" };
const _sfc_main$l = {
  emits: ["detailVisibility"],
  setup(__props, { expose, emit }) {
    const { t } = useI18n();
    const prefs2 = inject("prefs");
    const filter = ref(false);
    const details = ref(false);
    const mainClass = computed(() => ({
      "open-filter": filter.value,
      "open-details": details.value,
      "neutral-charts": prefs2.value.ui.neutralCharts
    }));
    function toggleFilter() {
      if (!filter.value) {
        filter.value = true;
        details.value = false;
        emit("detailVisibility");
      } else {
        filter.value = false;
      }
    }
    function toggleDetails() {
      emit("detailVisibility");
      if (!details.value) {
        details.value = true;
        filter.value = false;
      } else {
        details.value = false;
      }
    }
    function showDetails() {
      emit("detailVisibility");
      if (!details.value) {
        details.value = true;
        filter.value = false;
      }
    }
    expose({ showDetails });
    return (_ctx, _cache) => {
      return openBlock(), createElementBlock("main", {
        class: normalizeClass(unref(mainClass))
      }, [
        createBaseVNode("div", _hoisted_1$l, [
          renderSlot(_ctx.$slots, "filter", {}, () => [
            _hoisted_2$i
          ])
        ]),
        createBaseVNode("div", _hoisted_3$f, [
          createBaseVNode("div", _hoisted_4$c, [
            createBaseVNode("button", {
              onClick: toggleFilter,
              class: "open-filter icon"
            }, [
              createVNode(MenuImage)
            ]),
            _hoisted_5$9,
            createBaseVNode("button", {
              onClick: toggleDetails,
              class: "open-details icon"
            }, [
              createVNode(MoreImage)
            ])
          ]),
          createBaseVNode("div", _hoisted_6$8, [
            renderSlot(_ctx.$slots, "main", {}, () => [
              _hoisted_7$7
            ])
          ])
        ]),
        createBaseVNode("div", _hoisted_8$7, [
          createBaseVNode("div", _hoisted_9$7, [
            renderSlot(_ctx.$slots, "head-details"),
            _hoisted_10$7,
            createBaseVNode("button", {
              onClick: toggleDetails,
              class: "open-details icon"
            }, [
              createVNode(ArrowImage)
            ])
          ]),
          renderSlot(_ctx.$slots, "details", {}, () => [
            createBaseVNode("div", _hoisted_11$5, [
              createVNode(DetailsImage),
              createBaseVNode("p", null, toDisplayString(unref(t)("details.noitem")), 1)
            ])
          ])
        ])
      ], 2);
    };
  }
};
const _hoisted_1$k = ["placeholder"];
const _hoisted_2$h = /* @__PURE__ */ createTextVNode("Additional filters");
const _sfc_main$k = {
  props: ["data", "placeholder"],
  emits: ["result"],
  setup(__props, { emit }) {
    const prop = __props;
    var timeoutHandle = void 0;
    function confirm(evt) {
      evt.preventDefault();
      clearTimeout(timeoutHandle);
      timeoutHandle = setTimeout(function() {
        fetchData(evt.target.closest("form"));
      }, 500);
    }
    function fetchData(form) {
      let formData = new FormData(form);
      let filtered = prop.data.filter((d) => {
        let size = d.size || 1;
        for (let k of formData.keys()) {
          if (k == "name") {
            let name = formData.get(k).toLowerCase();
            if (!d[k].toLowerCase().includes(name)) {
              return false;
            }
            continue;
          }
          let [first, last] = formData.getAll(k).map((v) => parseFloat(v));
          if (!isNaN(first) && !isNaN(last)) {
            if (d[k] / size < first || last < d[k] / size) {
              return false;
            }
          }
        }
        return true;
      });
      emit("result", filtered);
    }
    return (_ctx, _cache) => {
      return openBlock(), createElementBlock("form", null, [
        createBaseVNode("input", {
          type: "text",
          name: "name",
          autocomplete: "off",
          placeholder: __props.placeholder,
          onInput: confirm
        }, null, 40, _hoisted_1$k),
        renderSlot(_ctx.$slots, "default", { confirm }, () => [
          _hoisted_2$h
        ])
      ]);
    };
  }
};
var Slider_vue_vue_type_style_index_0_lang = "";
const _hoisted_1$j = { class: "slider" };
const _hoisted_2$g = ["name", "value"];
const _hoisted_3$e = /* @__PURE__ */ createTextVNode(" \xA0\u2013\xA0 ");
const _hoisted_4$b = ["name", "value"];
const _hoisted_5$8 = { class: "bar" };
const _hoisted_6$7 = { class: "slide" };
const _hoisted_7$6 = /* @__PURE__ */ createBaseVNode("div", { class: "interact-area" }, null, -1);
const _hoisted_8$6 = [
  _hoisted_7$6
];
const _hoisted_9$6 = /* @__PURE__ */ createBaseVNode("div", { class: "interact-area" }, null, -1);
const _hoisted_10$6 = [
  _hoisted_9$6
];
const _sfc_main$j = {
  props: ["label", "name", "unit", "min", "max", "frac"],
  emits: ["input"],
  setup(__props, { emit }) {
    const prop = __props;
    const { t } = useI18n();
    const minVal = ref(parseFloat(prop.min).toFixed(prop.frac));
    const maxVal = ref(parseFloat(prop.max).toFixed(prop.frac));
    const minPercent = ref(0);
    const maxPercent = ref(100);
    watch(() => prop.min, (val, old) => setMin(val));
    watch(() => prop.max, (val, old) => setMax(val));
    function setMin(value) {
      let min = parseFloat(prop.min) || 0;
      let val = parseFloat(value) || min;
      minVal.value = val;
      maxVal.value = Math.max(minVal.value, maxVal.value);
      let max = parseFloat(prop.max) || 0;
      minPercent.value = (val - min) * 100 / (max - min);
      maxPercent.value = Math.max(minPercent.value, maxPercent.value);
    }
    function setMax(value) {
      let max = parseFloat(prop.max) || 0;
      let val = parseFloat(value) || max;
      maxVal.value = val;
      minVal.value = Math.min(minVal.value, maxVal.value);
      let min = parseFloat(prop.min) || 0;
      maxPercent.value = (val - min) * 100 / (max - min);
      minPercent.value = Math.min(minPercent.value, maxPercent.value);
    }
    function onMin(evt) {
      evt.target.blur();
      setMin(evt.target.value);
      emit("input", evt);
    }
    function onMax(evt) {
      evt.target.blur();
      setMax(evt.target.value);
      emit("input", evt);
    }
    function onSlide(evt) {
      let bar = evt.target.closest(".slide");
      let rect = bar.getBoundingClientRect();
      let pos = evt.pageX !== void 0 ? evt.pageX : evt.changedTouches[0].pageX;
      pos = Math.min(Math.max(pos - rect.left, 0), rect.width);
      let percent = pos * 100 / rect.width;
      let min = parseFloat(prop.min) || 0;
      let max = parseFloat(prop.max) || 0;
      let val = percent / 100 * (max - min) + min;
      if (evt.target.closest("button").classList.contains("min")) {
        minVal.value = val.toFixed(prop.frac);
        maxVal.value = Math.max(minVal.value, maxVal.value);
        minPercent.value = percent;
        maxPercent.value = Math.max(minPercent.value, maxPercent.value);
      } else {
        maxVal.value = val.toFixed(prop.frac);
        minVal.value = Math.min(minVal.value, maxVal.value);
        maxPercent.value = percent;
        minPercent.value = Math.min(minPercent.value, maxPercent.value);
      }
      emit("input", evt);
    }
    function onPress(evt) {
      let handle = evt.target.closest("button");
      handle.addEventListener("mousemove", onSlide);
      handle.addEventListener("touchmove", onSlide);
      handle.addEventListener("mouseup", onRelease);
      handle.addEventListener("touchend", onRelease);
      handle.addEventListener("mouseleave", onRelease);
      handle.addEventListener("touchcancel", onRelease);
      handle.classList.add("active");
    }
    function onRelease(evt) {
      let handle = evt.target.closest("button");
      handle.removeEventListener("mousemove", onSlide);
      handle.removeEventListener("touchmove", onSlide);
      handle.removeEventListener("mouseup", onRelease);
      handle.removeEventListener("touchend", onRelease);
      handle.removeEventListener("mouseleave", onRelease);
      handle.removeEventListener("touchcancel", onRelease);
      handle.classList.remove("active");
    }
    return (_ctx, _cache) => {
      return openBlock(), createElementBlock("div", _hoisted_1$j, [
        createBaseVNode("label", null, [
          createBaseVNode("span", null, toDisplayString(__props.label) + " (" + toDisplayString(unref(t)("unit." + __props.unit)) + ")", 1),
          createBaseVNode("input", {
            type: "text",
            name: __props.name,
            value: minVal.value,
            onChange: onMin
          }, null, 40, _hoisted_2$g),
          _hoisted_3$e,
          createBaseVNode("input", {
            type: "text",
            name: __props.name,
            value: maxVal.value,
            onChange: onMax
          }, null, 40, _hoisted_4$b)
        ]),
        createBaseVNode("div", _hoisted_5$8, [
          createBaseVNode("div", _hoisted_6$7, [
            createBaseVNode("div", {
              class: "overlay",
              style: normalizeStyle({ left: minPercent.value + "%", right: 100 - maxPercent.value + "%" })
            }, null, 4),
            createBaseVNode("button", {
              type: "button",
              class: "handle min",
              style: normalizeStyle({ left: minPercent.value + "%" }),
              onMousedown: onPress,
              onTouchstart: onPress
            }, _hoisted_8$6, 36),
            createBaseVNode("button", {
              type: "button",
              class: "handle max",
              style: normalizeStyle({ left: maxPercent.value + "%" }),
              onMousedown: onPress,
              onTouchstart: onPress
            }, _hoisted_10$6, 36)
          ])
        ])
      ]);
    };
  }
};
var ClickableInput_vue_vue_type_style_index_0_lang = "";
const _hoisted_1$i = { class: "clickable-input" };
const _hoisted_2$f = ["placeholder"];
const _hoisted_3$d = ["value"];
const _sfc_main$i = {
  props: ["label", "placeholder"],
  emits: ["confirm"],
  setup(__props, { emit }) {
    const value = ref("");
    function confirm(evt) {
      evt.preventDefault();
      emit("confirm", value.value);
      value.value = "";
    }
    return (_ctx, _cache) => {
      return openBlock(), createElementBlock("form", _hoisted_1$i, [
        withDirectives(createBaseVNode("input", {
          type: "text",
          autocomplete: "off",
          "onUpdate:modelValue": _cache[0] || (_cache[0] = ($event) => value.value = $event),
          placeholder: __props.placeholder
        }, null, 8, _hoisted_2$f), [
          [vModelText, value.value]
        ]),
        createBaseVNode("input", {
          type: "submit",
          onClick: confirm,
          value: __props.label
        }, null, 8, _hoisted_3$d)
      ]);
    };
  }
};
var ClickableSelect_vue_vue_type_style_index_0_lang = "";
const _hoisted_1$h = { class: "clickable-select" };
const _hoisted_2$e = {
  value: "0",
  selected: ""
};
const _hoisted_3$c = ["value"];
const _hoisted_4$a = ["disabled"];
const _sfc_main$h = {
  props: ["label", "placeholder", "items", "disabled"],
  emits: ["confirm"],
  setup(__props, { emit }) {
    const prop = __props;
    const { locale: locale2 } = useI18n();
    const value = ref(0);
    const collator = new Intl.Collator(locale2.value, { numeric: true });
    const sortedItems = computed(() => {
      return [...prop.items].sort((a, b) => collator.compare(a.name, b.name));
    });
    function confirm(evt) {
      evt.preventDefault();
      if (value.value) {
        emit("confirm", value.value);
      }
    }
    return (_ctx, _cache) => {
      return openBlock(), createElementBlock("form", _hoisted_1$h, [
        withDirectives(createBaseVNode("select", {
          "onUpdate:modelValue": _cache[0] || (_cache[0] = ($event) => value.value = $event)
        }, [
          createBaseVNode("option", _hoisted_2$e, toDisplayString(__props.placeholder), 1),
          (openBlock(true), createElementBlock(Fragment, null, renderList(unref(sortedItems), (item) => {
            return openBlock(), createElementBlock("option", {
              key: item.id,
              value: item.id
            }, toDisplayString(item.name), 9, _hoisted_3$c);
          }), 128))
        ], 512), [
          [vModelSelect, value.value]
        ]),
        createBaseVNode("button", {
          type: "submit",
          class: "async",
          onClick: confirm,
          disabled: __props.disabled
        }, toDisplayString(__props.label), 9, _hoisted_4$a)
      ]);
    };
  }
};
var ClickableDate_vue_vue_type_style_index_0_lang = "";
const _hoisted_1$g = { class: "clickable-date" };
const _hoisted_2$d = ["value"];
const _hoisted_3$b = ["value"];
const _hoisted_4$9 = ["disabled"];
const _sfc_main$g = {
  props: ["label", "date", "time", "disabled"],
  emits: ["confirm"],
  setup(__props, { emit }) {
    function confirm(evt) {
      evt.preventDefault();
      let data = new FormData(evt.target.closest("form"));
      emit("confirm", data.get("date"), data.get("time"));
    }
    return (_ctx, _cache) => {
      return openBlock(), createElementBlock("form", _hoisted_1$g, [
        createBaseVNode("input", {
          type: "date",
          name: "date",
          value: __props.date
        }, null, 8, _hoisted_2$d),
        createBaseVNode("input", {
          type: "time",
          name: "time",
          value: __props.time
        }, null, 8, _hoisted_3$b),
        createBaseVNode("button", {
          type: "submit",
          class: "async",
          onClick: confirm,
          disabled: __props.disabled
        }, toDisplayString(__props.label), 9, _hoisted_4$9)
      ]);
    };
  }
};
var ImageSortArrow_vue_vue_type_style_index_0_lang = "";
const _sfc_main$f = {};
const _hoisted_1$f = { class: "icon sort-arrow" };
function _sfc_render$2(_ctx, _cache) {
  return openBlock(), createElementBlock("span", _hoisted_1$f);
}
var Arrow = /* @__PURE__ */ _export_sfc(_sfc_main$f, [["render", _sfc_render$2]]);
const _hoisted_1$e = ["onClick"];
const _hoisted_2$c = { class: "name" };
const _hoisted_3$a = { class: "num" };
const _hoisted_4$8 = { class: "unit" };
const _hoisted_5$7 = { class: "m num" };
const _hoisted_6$6 = { class: "unit" };
const _hoisted_7$5 = { class: "m num" };
const _hoisted_8$5 = { class: "unit" };
const _hoisted_9$5 = { class: "m num" };
const _hoisted_10$5 = { class: "unit" };
const _sfc_main$e = {
  props: ["items"],
  emits: "selected",
  setup(__props, { emit }) {
    const prop = __props;
    const { t, locale: locale2 } = useI18n();
    const sortBy = ref("name");
    const sortDir = ref("asc");
    const collator = new Intl.Collator(locale2.value, { numeric: true });
    const sortedItems = computed(() => {
      if (sortDir.value == "asc") {
        return [...prop.items].sort((a, b) => {
          return collator.compare(a[sortBy.value], b[sortBy.value]);
        });
      } else {
        return [...prop.items].sort((a, b) => {
          return -collator.compare(a[sortBy.value], b[sortBy.value]);
        });
      }
    });
    function perServing(val, size, frac = 1) {
      if (size) {
        return +parseFloat(val / size).toFixed(frac);
      }
      return val;
    }
    function setActive(evt) {
      let cat = evt.target.dataset.sort;
      if (sortBy.value == cat) {
        sortDir.value = sortDir.value == "asc" ? "desc" : "asc";
      } else {
        sortBy.value = cat;
      }
    }
    return (_ctx, _cache) => {
      return openBlock(), createElementBlock("table", null, [
        createBaseVNode("thead", null, [
          createBaseVNode("tr", {
            class: normalizeClass(sortDir.value)
          }, [
            createBaseVNode("th", {
              class: normalizeClass(["name sort", { active: sortBy.value == "name" }]),
              onClick: setActive,
              "data-sort": "name"
            }, [
              createTextVNode(toDisplayString(unref(t)("food.name")) + " ", 1),
              createVNode(Arrow)
            ], 2),
            createBaseVNode("th", {
              class: normalizeClass(["num sort", { active: sortBy.value == "kcal" }]),
              onClick: setActive,
              "data-sort": "kcal"
            }, [
              createVNode(Arrow),
              createTextVNode(" " + toDisplayString(unref(t)("food.energy")), 1)
            ], 2),
            createBaseVNode("th", {
              class: normalizeClass(["m num sort", { active: sortBy.value == "fat" }]),
              onClick: setActive,
              "data-sort": "fat"
            }, [
              createVNode(Arrow),
              createTextVNode(" " + toDisplayString(unref(t)("food.fat")), 1)
            ], 2),
            createBaseVNode("th", {
              class: normalizeClass(["m num sort", { active: sortBy.value == "carb" }]),
              onClick: setActive,
              "data-sort": "carb"
            }, [
              createVNode(Arrow),
              createTextVNode(" " + toDisplayString(unref(t)("food.carbs2")), 1)
            ], 2),
            createBaseVNode("th", {
              class: normalizeClass(["m num sort", { active: sortBy.value == "prot" }]),
              onClick: setActive,
              "data-sort": "prot"
            }, [
              createVNode(Arrow),
              createTextVNode(" " + toDisplayString(unref(t)("food.protein")), 1)
            ], 2)
          ], 2)
        ]),
        createBaseVNode("tbody", null, [
          (openBlock(true), createElementBlock(Fragment, null, renderList(unref(sortedItems), (item) => {
            return openBlock(), createElementBlock("tr", {
              key: item.id,
              onClick: ($event) => _ctx.$emit("selected", item.id)
            }, [
              createBaseVNode("td", _hoisted_2$c, toDisplayString(item.name), 1),
              createBaseVNode("td", _hoisted_3$a, [
                createTextVNode(toDisplayString(perServing(item.kcal, item.size)) + " ", 1),
                createBaseVNode("span", _hoisted_4$8, toDisplayString(unref(t)("unit.cal")), 1)
              ]),
              createBaseVNode("td", _hoisted_5$7, [
                createTextVNode(toDisplayString(perServing(item.fat, item.size)) + " ", 1),
                createBaseVNode("span", _hoisted_6$6, toDisplayString(unref(t)("unit.g")), 1)
              ]),
              createBaseVNode("td", _hoisted_7$5, [
                createTextVNode(toDisplayString(perServing(item.carb, item.size)) + " ", 1),
                createBaseVNode("span", _hoisted_8$5, toDisplayString(unref(t)("unit.g")), 1)
              ]),
              createBaseVNode("td", _hoisted_9$5, [
                createTextVNode(toDisplayString(perServing(item.prot, item.size)) + " ", 1),
                createBaseVNode("span", _hoisted_10$5, toDisplayString(unref(t)("unit.g")), 1)
              ])
            ], 8, _hoisted_1$e);
          }), 128))
        ])
      ]);
    };
  }
};
const _sfc_main$d = {};
const _hoisted_1$d = {
  width: "24",
  height: "24",
  version: "1.1",
  viewBox: "0 0 24 24",
  xmlns: "http://www.w3.org/2000/svg"
};
const _hoisted_2$b = /* @__PURE__ */ createBaseVNode("path", { d: "m-7.5e-8 19.001v4.9993h4.9993l14.745-14.745-4.9993-4.9993zm23.61-13.611c0.51993-0.51993 0.51993-1.3598 0-1.8797l-3.1196-3.1196c-0.51993-0.51993-1.3598-0.51993-1.8797 0l-2.4397 2.4397 4.9993 4.9993z" }, null, -1);
const _hoisted_3$9 = [
  _hoisted_2$b
];
function _sfc_render$1(_ctx, _cache) {
  return openBlock(), createElementBlock("svg", _hoisted_1$d, _hoisted_3$9);
}
var EditImage = /* @__PURE__ */ _export_sfc(_sfc_main$d, [["render", _sfc_render$1]]);
var FoodView_vue_vue_type_style_index_0_lang = "";
const _hoisted_1$c = {
  key: 0,
  class: "new-item"
};
const _hoisted_2$a = /* @__PURE__ */ createBaseVNode("hr", null, null, -1);
const _hoisted_3$8 = /* @__PURE__ */ createBaseVNode("section", { class: "subtitle" }, "Some food category", -1);
const _hoisted_4$7 = { class: "tags" };
const _hoisted_5$6 = /* @__PURE__ */ createBaseVNode("span", { class: "tag" }, "Tag 1", -1);
const _hoisted_6$5 = /* @__PURE__ */ createBaseVNode("span", { class: "tag" }, "Tag 2", -1);
const _hoisted_7$4 = /* @__PURE__ */ createBaseVNode("span", { class: "tag" }, "Tag 3", -1);
const _hoisted_8$4 = ["disabled"];
const _hoisted_9$4 = /* @__PURE__ */ createBaseVNode("hr", null, null, -1);
const _hoisted_10$4 = { class: "tracking no-edit-mode" };
const _hoisted_11$4 = { class: "tracking-amount" };
const _hoisted_12$4 = { class: "unit" };
const _hoisted_13$3 = /* @__PURE__ */ createBaseVNode("label", null, "Add to recipe", -1);
const _hoisted_14$3 = /* @__PURE__ */ createBaseVNode("label", null, "Add to diary", -1);
const _hoisted_15$2 = /* @__PURE__ */ createBaseVNode("hr", null, null, -1);
const _hoisted_16$1 = { class: "nutrient-block" };
const _hoisted_17$1 = ["disabled"];
const _hoisted_18$1 = ["value"];
const _hoisted_19$1 = { class: "unit" };
const _hoisted_20$1 = ["value"];
const _hoisted_21$1 = { class: "unit" };
const _hoisted_22$1 = ["value"];
const _hoisted_23$1 = { class: "unit" };
const _hoisted_24$1 = ["value"];
const _hoisted_25$1 = { class: "unit" };
const _hoisted_26$1 = ["value"];
const _hoisted_27$1 = { class: "unit" };
const _hoisted_28$1 = ["disabled"];
const _hoisted_29$1 = ["value"];
const _hoisted_30$1 = { class: "unit" };
const _hoisted_31$1 = ["value"];
const _hoisted_32$1 = { class: "unit" };
const _hoisted_33$1 = ["value"];
const _hoisted_34$1 = { class: "unit" };
const _hoisted_35$1 = ["value"];
const _hoisted_36$1 = { class: "unit" };
const _hoisted_37$1 = ["value"];
const _hoisted_38$1 = { class: "unit" };
const _sfc_main$c = {
  setup(__props) {
    const { t } = useI18n();
    const log2 = inject("log");
    const csrf = inject("csrfToken");
    const perms2 = inject("perms");
    const foods = inject("food");
    const recipes2 = inject("recipes");
    const diary2 = inject("diary");
    const filtered = ref([]);
    const current = ref(null);
    const editMode = ref(false);
    const disableSave = ref(false);
    const disableToRecipe = ref(false);
    const disableToDiary = ref(false);
    const amount = ref(100);
    const today = ref(DateTime.now().toISODate());
    const now = ref(DateTime.now().toLocaleString(DateTime.TIME_24_SIMPLE));
    const main = ref(null);
    const form = ref(null);
    function newFood(name) {
      fetch("/api/v1/food", {
        method: "POST",
        headers: { "X-CSRF-Token": csrf }
      }).then((response) => {
        if (!response.ok) {
          throw t("createfood.err" + response.status);
        }
        return response.json();
      }).then((data) => {
        data.name = name;
        foods.value.push(data);
        filtered.value.push(data);
        log2.msg(t("createfood.ok"));
        showDetails(data.id);
      }).catch((err) => log2.err(err));
    }
    function saveFood() {
      disableSave.value = true;
      let id = current.value.id;
      fetch("/api/v1/food/" + id, {
        method: "PUT",
        headers: {
          "Content-Type": "application/x-www-form-urlencoded",
          "X-CSRF-Token": csrf
        },
        body: new URLSearchParams(new FormData(form.value))
      }).then((response) => {
        if (!response.ok) {
          throw t("savefood.err" + response.status);
        }
        editMode.value = false;
        return fetch("/api/v1/food/" + id);
      }).then((response) => response.json()).then((data) => {
        data.name = t(data.id.toString());
        foods.value = foods.value.map((f) => data.id == f.id ? data : f);
        filtered.value = filtered.value.map((f) => data.id == f.id ? data : f);
        current.value = current.value.id == data.id ? data : current.value;
        log2.msg(t("savefood.ok"));
      }).catch((err) => log2.err(err)).finally(() => {
        setTimeout(function() {
          disableSave.value = false;
        }, 500);
      });
    }
    function addToRecipe(id) {
      disableToRecipe.value = true;
      fetch(`/api/v1/recipe/${id}/ingredient/${current.value.id}`, {
        method: "POST",
        headers: {
          "Content-Type": "application/x-www-form-urlencoded",
          "X-CSRF-Token": csrf
        },
        body: new URLSearchParams({ amount: amount.value })
      }).then((response) => {
        if (!response.ok) {
          throw t("addfood.err" + response.status);
        }
        return fetch("/api/v1/recipe/" + id);
      }).then((response) => response.json()).then((data) => {
        recipes2.value = recipes2.value.map((r) => data.id == r.id ? data : r);
        amount.value = 100;
        log2.msg(t("addfood.ok"));
      }).catch((err) => log2.err(err)).finally(() => {
        setTimeout(function() {
          disableToRecipe.value = false;
        }, 500);
      });
    }
    function addToDiary(date, time) {
      disableToDiary.value = true;
      fetch(`/api/v1/diary/${date}`, {
        method: "POST",
        headers: {
          "Content-Type": "application/x-www-form-urlencoded",
          "X-CSRF-Token": csrf
        },
        body: new URLSearchParams({
          id: current.value.id,
          amount: amount.value,
          time,
          recipe: ""
        })
      }).then((response) => {
        if (!response.ok) {
          throw t("savediary.err" + response.status);
        }
        return fetch("/api/v1/diary/" + date.replaceAll("-", "/"));
      }).then((response) => response.json()).then((data) => {
        diary2.value[date] = data[0];
        amount.value = 100;
        log2.msg(t("savediary.ok"));
      }).catch((err) => log2.err(err)).finally(() => {
        setTimeout(function() {
          disableToDiary.value = false;
        }, 500);
      });
    }
    function updateList(items) {
      filtered.value = items;
      if (current.value) {
        if (filtered.value.filter((f) => f.id == current.value.id).length == 0) {
          current.value = null;
        }
      }
    }
    function showDetails(id) {
      current.value = filtered.value.filter((f) => f.id == id)[0];
      today.value = DateTime.now().toISODate();
      now.value = DateTime.now().toLocaleString(DateTime.TIME_24_SIMPLE);
      amount.value = 100;
      main.value.showDetails();
    }
    function onEditMode() {
      editMode.value ? saveFood() : editMode.value = true;
    }
    function onInput(evt) {
      evt.target.blur();
      if (isNaN(parseFloat(evt.target.value))) {
        evt.target.value = current.value[evt.target.name];
      }
    }
    return (_ctx, _cache) => {
      return openBlock(), createBlock(_sfc_main$l, {
        ref_key: "main",
        ref: main,
        onDetailVisibility: _cache[1] || (_cache[1] = ($event) => editMode.value = false),
        class: normalizeClass({ "edit-mode": editMode.value })
      }, createSlots({
        filter: withCtx(() => [
          unref(perms2).canCreateFood ? (openBlock(), createElementBlock("section", _hoisted_1$c, [
            createBaseVNode("h2", null, toDisplayString(unref(t)("aria.headnew")), 1),
            createVNode(_sfc_main$i, {
              label: unref(t)("btn.new"),
              placeholder: unref(t)("food.hintnew"),
              onConfirm: newFood
            }, null, 8, ["label", "placeholder"])
          ])) : createCommentVNode("", true),
          _hoisted_2$a,
          createBaseVNode("section", null, [
            createBaseVNode("h2", null, toDisplayString(unref(t)("aria.headsearch")), 1),
            createVNode(_sfc_main$k, {
              data: unref(foods),
              placeholder: unref(t)("food.hintsearch"),
              onResult: updateList
            }, {
              default: withCtx((slotProps) => [
                createBaseVNode("fieldset", null, [
                  createBaseVNode("legend", null, toDisplayString(unref(t)("aria.headmacro1")), 1),
                  createVNode(_sfc_main$j, {
                    label: unref(t)("food.energy"),
                    onInput: slotProps.confirm,
                    name: "kcal",
                    unit: "cal",
                    min: "0",
                    max: "900",
                    frac: "0"
                  }, null, 8, ["label", "onInput"]),
                  createVNode(_sfc_main$j, {
                    label: unref(t)("food.fat"),
                    onInput: slotProps.confirm,
                    name: "fat",
                    unit: "g",
                    min: "0",
                    max: "100",
                    frac: "0"
                  }, null, 8, ["label", "onInput"]),
                  createVNode(_sfc_main$j, {
                    label: unref(t)("food.carbs"),
                    onInput: slotProps.confirm,
                    name: "carb",
                    unit: "g",
                    min: "0",
                    max: "100",
                    frac: "0"
                  }, null, 8, ["label", "onInput"]),
                  createVNode(_sfc_main$j, {
                    label: unref(t)("food.protein"),
                    onInput: slotProps.confirm,
                    name: "prot",
                    unit: "g",
                    min: "0",
                    max: "89",
                    frac: "0"
                  }, null, 8, ["label", "onInput"]),
                  createVNode(_sfc_main$j, {
                    label: unref(t)("food.fiber"),
                    onInput: slotProps.confirm,
                    name: "fib",
                    unit: "g",
                    min: "0",
                    max: "71",
                    frac: "0"
                  }, null, 8, ["label", "onInput"])
                ]),
                createBaseVNode("fieldset", null, [
                  createBaseVNode("legend", null, toDisplayString(unref(t)("aria.headmacro2")), 1),
                  createVNode(_sfc_main$j, {
                    label: unref(t)("food.fatsat"),
                    onInput: slotProps.confirm,
                    name: "fatsat",
                    unit: "g",
                    min: "0",
                    max: "83",
                    frac: "0"
                  }, null, 8, ["label", "onInput"]),
                  createVNode(_sfc_main$j, {
                    label: unref(t)("food.fato3"),
                    onInput: slotProps.confirm,
                    name: "fato3",
                    unit: "g",
                    min: "0",
                    max: "54",
                    frac: "0"
                  }, null, 8, ["label", "onInput"]),
                  createVNode(_sfc_main$j, {
                    label: unref(t)("food.fato6"),
                    onInput: slotProps.confirm,
                    name: "fato6",
                    unit: "g",
                    min: "0",
                    max: "70",
                    frac: "0"
                  }, null, 8, ["label", "onInput"]),
                  createVNode(_sfc_main$j, {
                    label: unref(t)("food.sugar"),
                    onInput: slotProps.confirm,
                    name: "sug",
                    unit: "g",
                    min: "0",
                    max: "100",
                    frac: "0"
                  }, null, 8, ["label", "onInput"]),
                  createVNode(_sfc_main$j, {
                    label: unref(t)("food.salt"),
                    onInput: slotProps.confirm,
                    name: "salt",
                    unit: "g",
                    min: "0",
                    max: "100",
                    frac: "0"
                  }, null, 8, ["label", "onInput"])
                ])
              ]),
              _: 1
            }, 8, ["data", "placeholder"])
          ])
        ]),
        main: withCtx(() => [
          createVNode(_sfc_main$e, {
            items: filtered.value,
            onSelected: showDetails
          }, null, 8, ["items"])
        ]),
        _: 2
      }, [
        current.value ? {
          name: "head-details",
          fn: withCtx(() => [
            createBaseVNode("h2", null, toDisplayString(current.value.name), 1)
          ])
        } : void 0,
        current.value ? {
          name: "details",
          fn: withCtx(() => [
            _hoisted_3$8,
            createBaseVNode("section", _hoisted_4$7, [
              _hoisted_5$6,
              _hoisted_6$5,
              _hoisted_7$4,
              unref(perms2).canCreateFood || unref(perms2).canEditFood ? (openBlock(), createElementBlock("button", {
                key: 0,
                class: "icon async",
                disabled: disableSave.value,
                onClick: onEditMode
              }, [
                !editMode.value ? (openBlock(), createBlock(EditImage, { key: 0 })) : createCommentVNode("", true),
                editMode.value ? (openBlock(), createBlock(SaveImage, { key: 1 })) : createCommentVNode("", true)
              ], 8, _hoisted_8$4)) : createCommentVNode("", true)
            ]),
            _hoisted_9$4,
            createBaseVNode("section", _hoisted_10$4, [
              createBaseVNode("h2", null, toDisplayString(unref(t)("aria.headtrack")), 1),
              createBaseVNode("fieldset", _hoisted_11$4, [
                createBaseVNode("div", null, [
                  createBaseVNode("label", null, toDisplayString(unref(t)("food.amount")), 1),
                  withDirectives(createBaseVNode("input", {
                    type: "number",
                    "onUpdate:modelValue": _cache[0] || (_cache[0] = ($event) => amount.value = $event),
                    name: "amount"
                  }, null, 512), [
                    [vModelText, amount.value]
                  ]),
                  createBaseVNode("span", _hoisted_12$4, toDisplayString(unref(t)("unit.g")), 1)
                ])
              ]),
              createBaseVNode("fieldset", null, [
                _hoisted_13$3,
                createVNode(_sfc_main$h, {
                  label: unref(t)("btn.add"),
                  placeholder: unref(t)("food.hintrec"),
                  items: unref(recipes2),
                  disabled: disableToRecipe.value,
                  onConfirm: addToRecipe
                }, null, 8, ["label", "placeholder", "items", "disabled"])
              ]),
              createBaseVNode("fieldset", null, [
                _hoisted_14$3,
                createVNode(_sfc_main$g, {
                  label: unref(t)("btn.add"),
                  time: now.value,
                  date: today.value,
                  disabled: disableToDiary.value,
                  onConfirm: addToDiary
                }, null, 8, ["label", "time", "date", "disabled"])
              ])
            ]),
            _hoisted_15$2,
            createBaseVNode("section", null, [
              createBaseVNode("h2", null, toDisplayString(unref(t)("aria.headnutrients")), 1),
              createBaseVNode("form", {
                ref_key: "form",
                ref: form
              }, [
                createBaseVNode("div", _hoisted_16$1, [
                  createBaseVNode("fieldset", {
                    disabled: !editMode.value,
                    class: "col50"
                  }, [
                    createBaseVNode("div", null, [
                      createBaseVNode("label", null, toDisplayString(unref(t)("food.energy")), 1),
                      createBaseVNode("input", {
                        type: "text",
                        value: current.value.kcal,
                        name: "kcal",
                        onChange: onInput
                      }, null, 40, _hoisted_18$1),
                      createBaseVNode("span", _hoisted_19$1, toDisplayString(unref(t)("unit.cal")), 1)
                    ]),
                    createBaseVNode("div", null, [
                      createBaseVNode("label", null, toDisplayString(unref(t)("food.fat")), 1),
                      createBaseVNode("input", {
                        type: "text",
                        value: current.value.fat,
                        name: "fat",
                        onChange: onInput
                      }, null, 40, _hoisted_20$1),
                      createBaseVNode("span", _hoisted_21$1, toDisplayString(unref(t)("unit.g")), 1)
                    ]),
                    createBaseVNode("div", null, [
                      createBaseVNode("label", null, toDisplayString(unref(t)("food.carbs2")), 1),
                      createBaseVNode("input", {
                        type: "text",
                        value: current.value.carb,
                        name: "carb",
                        onChange: onInput
                      }, null, 40, _hoisted_22$1),
                      createBaseVNode("span", _hoisted_23$1, toDisplayString(unref(t)("unit.g")), 1)
                    ]),
                    createBaseVNode("div", null, [
                      createBaseVNode("label", null, toDisplayString(unref(t)("food.protein")), 1),
                      createBaseVNode("input", {
                        type: "text",
                        value: current.value.prot,
                        name: "prot",
                        onChange: onInput
                      }, null, 40, _hoisted_24$1),
                      createBaseVNode("span", _hoisted_25$1, toDisplayString(unref(t)("unit.g")), 1)
                    ]),
                    createBaseVNode("div", null, [
                      createBaseVNode("label", null, toDisplayString(unref(t)("food.fiber")), 1),
                      createBaseVNode("input", {
                        type: "text",
                        value: current.value.fib,
                        name: "fib",
                        onChange: onInput
                      }, null, 40, _hoisted_26$1),
                      createBaseVNode("span", _hoisted_27$1, toDisplayString(unref(t)("unit.g")), 1)
                    ])
                  ], 8, _hoisted_17$1),
                  createBaseVNode("fieldset", {
                    disabled: !editMode.value,
                    class: "col50"
                  }, [
                    createBaseVNode("div", null, [
                      createBaseVNode("label", null, toDisplayString(unref(t)("food.fatsat")), 1),
                      createBaseVNode("input", {
                        type: "text",
                        value: current.value.fatsat,
                        name: "fatsat",
                        onChange: onInput
                      }, null, 40, _hoisted_29$1),
                      createBaseVNode("span", _hoisted_30$1, toDisplayString(unref(t)("unit.g")), 1)
                    ]),
                    createBaseVNode("div", null, [
                      createBaseVNode("label", null, toDisplayString(unref(t)("food.fato3")), 1),
                      createBaseVNode("input", {
                        type: "text",
                        value: current.value.fato3,
                        name: "fato3",
                        onChange: onInput
                      }, null, 40, _hoisted_31$1),
                      createBaseVNode("span", _hoisted_32$1, toDisplayString(unref(t)("unit.g")), 1)
                    ]),
                    createBaseVNode("div", null, [
                      createBaseVNode("label", null, toDisplayString(unref(t)("food.fato6")), 1),
                      createBaseVNode("input", {
                        type: "text",
                        value: current.value.fato6,
                        name: "fato6",
                        onChange: onInput
                      }, null, 40, _hoisted_33$1),
                      createBaseVNode("span", _hoisted_34$1, toDisplayString(unref(t)("unit.g")), 1)
                    ]),
                    createBaseVNode("div", null, [
                      createBaseVNode("label", null, toDisplayString(unref(t)("food.sugar")), 1),
                      createBaseVNode("input", {
                        type: "text",
                        value: current.value.sug,
                        name: "sug",
                        onChange: onInput
                      }, null, 40, _hoisted_35$1),
                      createBaseVNode("span", _hoisted_36$1, toDisplayString(unref(t)("unit.g")), 1)
                    ]),
                    createBaseVNode("div", null, [
                      createBaseVNode("label", null, toDisplayString(unref(t)("food.salt")), 1),
                      createBaseVNode("input", {
                        type: "text",
                        value: current.value.salt,
                        name: "salt",
                        onChange: onInput
                      }, null, 40, _hoisted_37$1),
                      createBaseVNode("span", _hoisted_38$1, toDisplayString(unref(t)("unit.g")), 1)
                    ])
                  ], 8, _hoisted_28$1)
                ])
              ], 512)
            ])
          ])
        } : void 0
      ]), 1032, ["class"]);
    };
  }
};
var IngredientList_vue_vue_type_style_index_0_lang = "";
const _hoisted_1$b = ["disabled"];
const _hoisted_2$9 = ["value"];
const _hoisted_3$7 = { class: "unit" };
const _hoisted_4$6 = ["value"];
const _sfc_main$b = {
  props: ["items", "disabled"],
  setup(__props, { expose }) {
    const prop = __props;
    const { t, locale: locale2 } = useI18n();
    const sortBy = ref("name");
    const form = ref(null);
    const collator = new Intl.Collator(locale2.value, { numeric: true });
    const sortedItems = computed(() => {
      let items = prop.items.map((i) => ({ id: i.id, amount: i.amount, name: t(i.id.toString()) }));
      return items.sort((a, b) => collator.compare(a[sortBy.value], b[sortBy.value]));
    });
    function onInput(evt) {
      evt.target.blur();
      let val = parseFloat(evt.target.value);
      if (isNaN(val) || val < 0) {
        evt.target.value = 0;
      }
    }
    function getDiff() {
      let data = new FormData(form.value);
      let ids = data.getAll("id");
      let amounts = data.getAll("amount");
      let result = [];
      prop.items.forEach((item) => {
        let idx = parseInt(ids.indexOf(item.id.toString()));
        let amount = parseFloat(amounts[idx]);
        if (amount != item.amount) {
          result.push({ id: item.id, amount });
        }
      });
      return result;
    }
    expose({ getDiff });
    return (_ctx, _cache) => {
      return openBlock(), createElementBlock("form", {
        class: "ingredients",
        ref_key: "form",
        ref: form
      }, [
        createBaseVNode("fieldset", { disabled: __props.disabled }, [
          (openBlock(true), createElementBlock(Fragment, null, renderList(unref(sortedItems), (item) => {
            return openBlock(), createElementBlock("div", {
              key: item.id
            }, [
              createBaseVNode("label", null, toDisplayString(item.name), 1),
              createBaseVNode("input", {
                type: "number",
                name: "amount",
                value: item.amount,
                onChange: onInput
              }, null, 40, _hoisted_2$9),
              createBaseVNode("span", _hoisted_3$7, toDisplayString(unref(t)("unit.g")), 1),
              createBaseVNode("input", {
                type: "hidden",
                name: "id",
                value: item.id
              }, null, 8, _hoisted_4$6)
            ]);
          }), 128))
        ], 8, _hoisted_1$b)
      ], 512);
    };
  }
};
const _sfc_main$a = {};
const _hoisted_1$a = {
  width: "24",
  height: "24",
  version: "1.1",
  viewBox: "0 0 24 24",
  xmlns: "http://www.w3.org/2000/svg"
};
const _hoisted_2$8 = /* @__PURE__ */ createBaseVNode("path", { d: "m0 18.316h2.5263v0.63158h-1.2632v1.2632h1.2632v0.63158h-2.5263v1.2632h3.7895v-5.0526h-3.7895zm1.2632-11.368h1.2632v-5.0526h-2.5263v1.2632h1.2632zm-1.2632 3.7895h2.2737l-2.2737 2.6526v1.1368h3.7895v-1.2632h-2.2737l2.2737-2.6526v-1.1368h-3.7895zm6.3158-7.5789v2.5263h17.684v-2.5263zm0 17.684h17.684v-2.5263h-17.684zm0-7.5789h17.684v-2.5263h-17.684z" }, null, -1);
const _hoisted_3$6 = [
  _hoisted_2$8
];
function _sfc_render(_ctx, _cache) {
  return openBlock(), createElementBlock("svg", _hoisted_1$a, _hoisted_3$6);
}
var ListImage = /* @__PURE__ */ _export_sfc(_sfc_main$a, [["render", _sfc_render]]);
var RecipeView_vue_vue_type_style_index_0_lang = "";
const _hoisted_1$9 = { class: "new-item" };
const _hoisted_2$7 = /* @__PURE__ */ createBaseVNode("hr", null, null, -1);
const _hoisted_3$5 = ["disabled"];
const _hoisted_4$5 = ["value"];
const _hoisted_5$5 = ["innerHTML"];
const _hoisted_6$4 = { class: "tags" };
const _hoisted_7$3 = /* @__PURE__ */ createBaseVNode("span", { class: "tag no-edit-mode" }, "Tag 1", -1);
const _hoisted_8$3 = /* @__PURE__ */ createBaseVNode("span", { class: "tag no-edit-mode" }, "Tag 2", -1);
const _hoisted_9$3 = /* @__PURE__ */ createBaseVNode("span", { class: "tag no-edit-mode" }, "Tag 3", -1);
const _hoisted_10$3 = ["disabled"];
const _hoisted_11$3 = /* @__PURE__ */ createBaseVNode("hr", null, null, -1);
const _hoisted_12$3 = {
  key: 0,
  class: "tracking no-edit-mode"
};
const _hoisted_13$2 = { class: "tracking-amount" };
const _hoisted_14$2 = { class: "unit" };
const _hoisted_15$1 = /* @__PURE__ */ createBaseVNode("label", null, "Add to diary", -1);
const _hoisted_16 = /* @__PURE__ */ createBaseVNode("hr", null, null, -1);
const _hoisted_17 = ["innerHTML"];
const _hoisted_18 = /* @__PURE__ */ createBaseVNode("hr", null, null, -1);
const _hoisted_19 = { class: "no-edit-mode" };
const _hoisted_20 = { class: "nutrient-block" };
const _hoisted_21 = { class: "col50" };
const _hoisted_22 = { class: "unit" };
const _hoisted_23 = { class: "unit" };
const _hoisted_24 = { class: "col50" };
const _hoisted_25 = { class: "unit" };
const _hoisted_26 = { class: "unit" };
const _hoisted_27 = /* @__PURE__ */ createBaseVNode("hr", null, null, -1);
const _hoisted_28 = { class: "prep" };
const _hoisted_29 = ["disabled"];
const _hoisted_30 = { class: "prep-size" };
const _hoisted_31 = ["value"];
const _hoisted_32 = ["disabled"];
const _hoisted_33 = ["value"];
const _hoisted_34 = { class: "unit" };
const _hoisted_35 = ["value"];
const _hoisted_36 = { class: "unit" };
const _hoisted_37 = ["value"];
const _hoisted_38 = { class: "unit" };
const _hoisted_39 = ["value"];
const _hoisted_40 = { class: "unit" };
const _hoisted_41 = { class: "placeholder" };
const _sfc_main$9 = {
  setup(__props) {
    const { t } = useI18n();
    const log2 = inject("log");
    const csrf = inject("csrfToken");
    inject("perms");
    const recipes2 = inject("recipes");
    const diary2 = inject("diary");
    const filtered = ref([]);
    const current = ref(null);
    const editMode = ref(false);
    const isSaving = ref(false);
    const disableToDiary = ref(false);
    const amount = ref(1);
    const ownerInfo = ref("&nbsp;");
    const today = ref(DateTime.now().toISODate());
    const now = ref(DateTime.now().toLocaleString(DateTime.TIME_24_SIMPLE));
    const main = ref(null);
    const form = ref(null);
    const ingredients = ref(null);
    function perServing(val, frac = 2) {
      return +parseFloat(val / (current.value.size || 1)).toFixed(frac);
    }
    function minSearchAttr(attr) {
      return Math.min.apply(Math, recipes2.value.map((r) => Math.floor(r[attr] / r.size)));
    }
    function maxSearchAttr(attr) {
      return Math.max.apply(Math, recipes2.value.map((r) => Math.ceil(r[attr] / r.size)));
    }
    function newRecipe(name) {
      fetch("/api/v1/recipe", {
        method: "POST",
        headers: {
          "Content-Type": "application/x-www-form-urlencoded",
          "X-CSRF-Token": csrf
        },
        body: new URLSearchParams({ name })
      }).then((response) => {
        if (!response.ok) {
          throw t("createrec.err" + response.status);
        }
        return response.json();
      }).then((data) => {
        data.isowner = true;
        recipes2.value.push(data);
        filtered.value.push(data);
        log2.msg(t("createrec.ok"));
        showDetails(data.id);
      }).catch((err) => log2.err(err));
    }
    function saveRecipe() {
      isSaving.value = true;
      let id = current.value.id;
      let owner = current.value.owner;
      let isowner = current.value.isowner;
      fetch("/api/v1/recipe/" + id, {
        method: "PUT",
        headers: {
          "Content-Type": "application/x-www-form-urlencoded",
          "X-CSRF-Token": csrf
        },
        body: new URLSearchParams(new FormData(form.value))
      }).then((response) => {
        if (!response.ok) {
          throw t("saverec.err" + response.status);
        }
        return saveIngredients(id);
      }).then(() => fetch("/api/v1/recipe/" + id)).then((response) => response.json()).then((data) => {
        data.owner = owner;
        data.isowner = isowner;
        recipes2.value = recipes2.value.map((r) => data.id == r.id ? data : r);
        filtered.value = filtered.value.map((r) => data.id == r.id ? data : r);
        current.value = current.value.id == data.id ? data : current.value;
        log2.msg(t("saverec.ok"));
      }).catch((err) => log2.err(err)).finally(() => {
        setTimeout(function() {
          isSaving.value = false;
        }, 500);
      });
    }
    function saveIngredients(id) {
      let items = ingredients.value.getDiff();
      if (items.length == 0) {
        return new Promise((resolve) => {
          editMode.value = false;
          resolve();
        });
      }
      return new Promise((resolve, reject) => {
        let count = 0;
        let error = void 0;
        items.forEach((item) => {
          fetch(`/api/v1/recipe/${id}/ingredient/${item.id}`, {
            method: "PUT",
            headers: {
              "Content-Type": "application/x-www-form-urlencoded",
              "X-CSRF-Token": csrf
            },
            body: new URLSearchParams({ amount: item.amount })
          }).then((response) => {
            ++count;
            if (!response.ok) {
              throw t("saverec.err" + response.status);
            }
          }).catch((err) => error = err).finally(() => {
            if (count == items.length) {
              if (error) {
                reject(error);
              } else {
                editMode.value = false;
                resolve();
              }
            }
          });
        });
      });
    }
    function addToDiary(date, time) {
      disableToDiary.value = true;
      let params = new URLSearchParams();
      current.value.items.forEach((i) => {
        params.append("id", i.id);
        params.append("amount", perServing(i.amount) * amount.value);
        params.append("time", time);
        params.append("recipe", current.value.name);
      });
      fetch(`/api/v1/diary/${date}`, {
        method: "POST",
        body: params,
        headers: {
          "Content-Type": "application/x-www-form-urlencoded",
          "X-CSRF-Token": csrf
        }
      }).then((response) => {
        if (!response.ok) {
          throw t("savediary.err" + response.status);
        }
        return fetch("/api/v1/diary/" + date.replaceAll("-", "/"));
      }).then((response) => response.json()).then((data) => {
        diary2.value[date] = data[0];
        amount.value = 1;
        log2.msg(t("savediary.ok"));
      }).catch((err) => log2.err(err)).finally(() => {
        setTimeout(function() {
          disableToDiary.value = false;
        }, 500);
      });
    }
    function showDetails(id) {
      current.value = filtered.value.filter((r) => r.id == id)[0];
      today.value = DateTime.now().toISODate();
      now.value = DateTime.now().toLocaleString(DateTime.TIME_24_SIMPLE);
      amount.value = 1;
      main.value.showDetails();
      if ("isowner" in current.value) {
        updateOwnerInfo();
      } else {
        fetchOwnerInfo();
      }
    }
    function fetchOwnerInfo() {
      fetch(`/api/v1/recipe/${current.value.id}/owner`).then((response) => {
        if (!response.ok) {
          throw response;
        }
        return response.json();
      }).then((data) => {
        current.value.isowner = data.isowner;
        current.value.owner = data.owner;
        updateOwnerInfo();
      }).catch(() => {
        log2.err(t("recowner.err"));
        ownerInfo.value = "&nbsp;";
      });
    }
    function updateOwnerInfo() {
      if (current.value.isowner) {
        ownerInfo.value = t("recipe.isowner");
      } else if (current.value.owner) {
        ownerInfo.value = t("recipe.owner", { name: current.value.owner });
      } else {
        ownerInfo.value = t("recipe.ispublic");
      }
    }
    function updateList(items) {
      filtered.value = items;
      if (current.value) {
        if (filtered.value.filter((r) => r.id == current.value.id).length == 0) {
          current.value = null;
        }
      }
    }
    function onEditMode() {
      editMode.value ? saveRecipe() : editMode.value = true;
    }
    function onInput(evt) {
      evt.target.blur();
      if (isNaN(parseFloat(evt.target.value))) {
        evt.target.value = current.value[evt.target.name];
      }
    }
    onMounted(() => filtered.value = [...recipes2.value]);
    return (_ctx, _cache) => {
      return openBlock(), createBlock(_sfc_main$l, {
        ref_key: "main",
        ref: main,
        onDetailVisibility: _cache[1] || (_cache[1] = ($event) => editMode.value = false),
        class: normalizeClass({ "edit-mode": editMode.value })
      }, createSlots({
        filter: withCtx(() => [
          createBaseVNode("section", _hoisted_1$9, [
            createBaseVNode("h2", null, toDisplayString(unref(t)("aria.headnewrec")), 1),
            createVNode(_sfc_main$i, {
              label: unref(t)("btn.new"),
              placeholder: unref(t)("recipe.hintnew"),
              onConfirm: newRecipe
            }, null, 8, ["label", "placeholder"])
          ]),
          _hoisted_2$7,
          createBaseVNode("section", null, [
            createBaseVNode("h2", null, toDisplayString(unref(t)("aria.headsearch")), 1),
            createVNode(_sfc_main$k, {
              data: unref(recipes2),
              placeholder: unref(t)("recipe.hintsearch"),
              onResult: updateList
            }, {
              default: withCtx((slotProps) => [
                createBaseVNode("fieldset", null, [
                  createVNode(_sfc_main$j, {
                    label: unref(t)("food.energy"),
                    min: minSearchAttr("kcal"),
                    max: maxSearchAttr("kcal"),
                    onInput: slotProps.confirm,
                    name: "kcal",
                    unit: "cal",
                    frac: "0"
                  }, null, 8, ["label", "min", "max", "onInput"]),
                  createVNode(_sfc_main$j, {
                    label: unref(t)("food.fat"),
                    min: minSearchAttr("fat"),
                    max: maxSearchAttr("fat"),
                    onInput: slotProps.confirm,
                    name: "fat",
                    unit: "g",
                    frac: "0"
                  }, null, 8, ["label", "min", "max", "onInput"]),
                  createVNode(_sfc_main$j, {
                    label: unref(t)("food.carbs"),
                    min: minSearchAttr("carb"),
                    max: maxSearchAttr("carb"),
                    onInput: slotProps.confirm,
                    name: "carb",
                    unit: "g",
                    frac: "0"
                  }, null, 8, ["label", "min", "max", "onInput"]),
                  createVNode(_sfc_main$j, {
                    label: unref(t)("food.protein"),
                    min: minSearchAttr("prot"),
                    max: maxSearchAttr("prot"),
                    onInput: slotProps.confirm,
                    name: "prot",
                    unit: "g",
                    frac: "0"
                  }, null, 8, ["label", "min", "max", "onInput"])
                ])
              ]),
              _: 1
            }, 8, ["data", "placeholder"])
          ])
        ]),
        main: withCtx(() => [
          createVNode(_sfc_main$e, {
            items: filtered.value,
            onSelected: showDetails
          }, null, 8, ["items"])
        ]),
        _: 2
      }, [
        current.value ? {
          name: "head-details",
          fn: withCtx(() => [
            createBaseVNode("form", {
              ref_key: "form",
              ref: form,
              autocomplete: "off",
              id: "form-recipe"
            }, [
              createBaseVNode("fieldset", {
                disabled: !editMode.value
              }, [
                createBaseVNode("input", {
                  type: "text",
                  name: "name",
                  value: current.value.name
                }, null, 8, _hoisted_4$5)
              ], 8, _hoisted_3$5)
            ], 512)
          ])
        } : void 0,
        current.value ? {
          name: "details",
          fn: withCtx(() => [
            createBaseVNode("section", {
              class: "subtitle no-edit-mode",
              innerHTML: ownerInfo.value
            }, null, 8, _hoisted_5$5),
            createBaseVNode("section", _hoisted_6$4, [
              _hoisted_7$3,
              _hoisted_8$3,
              _hoisted_9$3,
              current.value.isowner ? (openBlock(), createElementBlock("button", {
                key: 0,
                class: "icon async",
                disabled: isSaving.value,
                onClick: onEditMode
              }, [
                !editMode.value ? (openBlock(), createBlock(EditImage, { key: 0 })) : createCommentVNode("", true),
                editMode.value ? (openBlock(), createBlock(SaveImage, { key: 1 })) : createCommentVNode("", true)
              ], 8, _hoisted_10$3)) : createCommentVNode("", true)
            ]),
            _hoisted_11$3,
            current.value.items.length ? (openBlock(), createElementBlock("section", _hoisted_12$3, [
              createBaseVNode("h2", null, toDisplayString(unref(t)("aria.headtrack")), 1),
              createBaseVNode("fieldset", _hoisted_13$2, [
                createBaseVNode("div", null, [
                  createBaseVNode("label", null, toDisplayString(unref(t)("food.amount")), 1),
                  withDirectives(createBaseVNode("input", {
                    type: "number",
                    "onUpdate:modelValue": _cache[0] || (_cache[0] = ($event) => amount.value = $event),
                    name: "amount"
                  }, null, 512), [
                    [vModelText, amount.value]
                  ]),
                  createBaseVNode("span", _hoisted_14$2, toDisplayString(unref(t)("recipe.size", amount.value)), 1)
                ])
              ]),
              createBaseVNode("fieldset", null, [
                _hoisted_15$1,
                createVNode(_sfc_main$g, {
                  label: unref(t)("btn.add"),
                  time: now.value,
                  date: today.value,
                  disabled: disableToDiary.value,
                  onConfirm: addToDiary
                }, null, 8, ["label", "time", "date", "disabled"])
              ])
            ])) : createCommentVNode("", true),
            _hoisted_16,
            createBaseVNode("section", null, [
              createBaseVNode("h2", null, toDisplayString(unref(t)("aria.headingred")), 1),
              !current.value.items.length ? (openBlock(), createElementBlock("p", {
                key: 0,
                class: "msg-noitems",
                innerHTML: unref(t)("recipe.noitems")
              }, null, 8, _hoisted_17)) : createCommentVNode("", true),
              createVNode(_sfc_main$b, {
                ref_key: "ingredients",
                ref: ingredients,
                items: current.value.items,
                disabled: !editMode.value
              }, null, 8, ["items", "disabled"])
            ]),
            _hoisted_18,
            createBaseVNode("section", _hoisted_19, [
              createBaseVNode("h2", null, toDisplayString(unref(t)("aria.headnutrients")), 1),
              createBaseVNode("div", _hoisted_20, [
                createBaseVNode("div", _hoisted_21, [
                  createBaseVNode("div", null, [
                    createBaseVNode("label", null, toDisplayString(unref(t)("food.energy")), 1),
                    createBaseVNode("span", null, toDisplayString(perServing(current.value.kcal, 1)), 1),
                    createBaseVNode("span", _hoisted_22, toDisplayString(unref(t)("unit.cal")), 1)
                  ]),
                  createBaseVNode("div", null, [
                    createBaseVNode("label", null, toDisplayString(unref(t)("food.fat")), 1),
                    createBaseVNode("span", null, toDisplayString(perServing(current.value.fat, 1)), 1),
                    createBaseVNode("span", _hoisted_23, toDisplayString(unref(t)("unit.g")), 1)
                  ])
                ]),
                createBaseVNode("div", _hoisted_24, [
                  createBaseVNode("div", null, [
                    createBaseVNode("label", null, toDisplayString(unref(t)("food.carbs2")), 1),
                    createBaseVNode("span", null, toDisplayString(perServing(current.value.carb, 1)), 1),
                    createBaseVNode("span", _hoisted_25, toDisplayString(unref(t)("unit.g")), 1)
                  ]),
                  createBaseVNode("div", null, [
                    createBaseVNode("label", null, toDisplayString(unref(t)("food.protein")), 1),
                    createBaseVNode("span", null, toDisplayString(perServing(current.value.prot, 1)), 1),
                    createBaseVNode("span", _hoisted_26, toDisplayString(unref(t)("unit.g")), 1)
                  ])
                ])
              ])
            ]),
            _hoisted_27,
            createBaseVNode("section", _hoisted_28, [
              createBaseVNode("h2", null, toDisplayString(unref(t)("aria.headprep")), 1),
              createBaseVNode("div", null, [
                createBaseVNode("fieldset", {
                  disabled: !editMode.value,
                  class: "col50"
                }, [
                  createBaseVNode("div", _hoisted_30, [
                    createBaseVNode("label", null, toDisplayString(unref(t)("recipe.size", 2)), 1),
                    createBaseVNode("input", {
                      type: "text",
                      name: "size",
                      form: "form-recipe",
                      value: current.value.size,
                      onChange: onInput
                    }, null, 40, _hoisted_31),
                    createBaseVNode("label", null, toDisplayString(unref(t)("recipe.size", current.value.size)), 1)
                  ])
                ], 8, _hoisted_29),
                createBaseVNode("fieldset", {
                  disabled: !editMode.value,
                  class: "col50"
                }, [
                  createBaseVNode("div", null, [
                    createBaseVNode("label", null, toDisplayString(unref(t)("recipe.time")), 1),
                    createBaseVNode("input", {
                      type: "text",
                      disabled: "",
                      value: current.value.preptime + current.value.cooktime + current.value.misctime
                    }, null, 8, _hoisted_33),
                    createBaseVNode("span", _hoisted_34, toDisplayString(unref(t)("unit.min")), 1)
                  ]),
                  createBaseVNode("div", null, [
                    createBaseVNode("label", null, toDisplayString(unref(t)("recipe.preptime")), 1),
                    createBaseVNode("input", {
                      type: "text",
                      name: "preptime",
                      form: "form-recipe",
                      value: current.value.preptime,
                      onChange: onInput
                    }, null, 40, _hoisted_35),
                    createBaseVNode("span", _hoisted_36, toDisplayString(unref(t)("unit.min")), 1)
                  ]),
                  createBaseVNode("div", null, [
                    createBaseVNode("label", null, toDisplayString(unref(t)("recipe.cooktime")), 1),
                    createBaseVNode("input", {
                      type: "text",
                      name: "cooktime",
                      form: "form-recipe",
                      value: current.value.cooktime,
                      onChange: onInput
                    }, null, 40, _hoisted_37),
                    createBaseVNode("span", _hoisted_38, toDisplayString(unref(t)("unit.min")), 1)
                  ]),
                  createBaseVNode("div", null, [
                    createBaseVNode("label", null, toDisplayString(unref(t)("recipe.misctime")), 1),
                    createBaseVNode("input", {
                      type: "text",
                      name: "misctime",
                      form: "form-recipe",
                      value: current.value.misctime,
                      onChange: onInput
                    }, null, 40, _hoisted_39),
                    createBaseVNode("span", _hoisted_40, toDisplayString(unref(t)("unit.min")), 1)
                  ])
                ], 8, _hoisted_32)
              ]),
              createBaseVNode("div", _hoisted_41, [
                createVNode(ListImage),
                createBaseVNode("p", null, toDisplayString(unref(t)("todo.instructions")), 1)
              ])
            ])
          ])
        } : void 0
      ]), 1032, ["class"]);
    };
  }
};
var Calendar_vue_vue_type_style_index_0_lang = "";
const _hoisted_1$8 = ["disabled"];
const _hoisted_2$6 = { value: "1" };
const _hoisted_3$4 = { value: "2" };
const _hoisted_4$4 = { value: "3" };
const _hoisted_5$4 = { value: "4" };
const _hoisted_6$3 = { value: "5" };
const _hoisted_7$2 = { value: "6" };
const _hoisted_8$2 = { value: "7" };
const _hoisted_9$2 = { value: "8" };
const _hoisted_10$2 = { value: "9" };
const _hoisted_11$2 = { value: "10" };
const _hoisted_12$2 = { value: "11" };
const _hoisted_13$1 = { value: "12" };
const _hoisted_14$1 = ["value"];
const _hoisted_15 = ["disabled"];
const _sfc_main$8 = {
  props: ["mode", "items", "storage"],
  emits: ["selection"],
  setup(__props, { expose, emit }) {
    const prop = __props;
    const { t } = useI18n();
    const month = ref(DateTime.now().month);
    const year = ref(DateTime.now().year);
    const selection = ref([]);
    const calendar = ref(null);
    const years = computed(() => {
      let dates = prop.items.map((i) => parseInt(i.split("-")[0])).filter((i, idx, self) => self.indexOf(i) === idx);
      if (dates.length == 0) {
        dates = [DateTime.now().year];
      }
      return [Math.min(...dates) - 1, ...dates, Math.max(...dates) + 1];
    });
    const hasPrev = computed(() => {
      return year.value != years.value[0] || month.value > 1;
    });
    const hasNext = computed(() => {
      return year.value != years.value[years.value.length - 1] || month.value < 12;
    });
    watch(() => prop.items, () => {
      calendar.value.querySelectorAll("td>button").forEach((btn) => {
        if (prop.items.indexOf(btn.dataset.date) != -1) {
          btn.parentNode.classList.add("has-entries");
        } else {
          btn.parentNode.classList.remove("has-entries");
        }
      });
    });
    function onCalendarChanged() {
      let today = DateTime.now().toISODate();
      let date = DateTime.local(year.value, month.value);
      date = date.minus({ days: date.weekday - 1 });
      calendar.value.querySelectorAll("td").forEach((cell) => {
        let iso = date.toISODate();
        cell.firstElementChild.textContent = date.day;
        cell.firstElementChild.dataset.date = iso;
        if (date.month != month.value) {
          cell.classList.add("outside");
        } else {
          cell.classList.remove("outside");
        }
        if (iso == today) {
          cell.classList.add("today");
        } else {
          cell.classList.remove("today");
        }
        if (prop.items.indexOf(iso) != -1) {
          cell.classList.add("has-entries");
        } else {
          cell.classList.remove("has-entries");
        }
        if (selection.value.indexOf(iso) != -1) {
          cell.classList.add("active");
        } else {
          cell.classList.remove("active");
        }
        date = date.plus({ days: 1 });
      });
    }
    function onPrev() {
      if (--month.value < 1) {
        month.value = 12;
        --year.value;
      }
      onCalendarChanged();
    }
    function onNext() {
      if (++month.value > 12) {
        month.value = 1;
        ++year.value;
      }
      onCalendarChanged();
    }
    function onDay(evt) {
      let iso = evt.target.dataset.date;
      if (prop.mode == "toggle") {
        let idx = selection.value.indexOf(iso);
        if (idx == -1) {
          selection.value.push(iso);
        } else {
          selection.value.splice(idx, 1);
        }
      } else {
        selection.value = [iso];
      }
      calendar.value.querySelectorAll("td>button").forEach((btn) => {
        if (selection.value.indexOf(btn.dataset.date) != -1) {
          btn.parentNode.classList.add("active");
        } else {
          btn.parentNode.classList.remove("active");
        }
      });
      saveDatesLocal();
      emit("selection", selection.value);
    }
    function saveDatesLocal() {
      if (prop.storage) {
        let str = JSON.stringify(selection.value);
        window.localStorage.setItem(prop.storage, str);
      }
    }
    function loadDatesLocal() {
      if (prop.storage) {
        let str = window.localStorage.getItem(prop.storage);
        selection.value = JSON.parse(str) || [];
      }
      if (!selection.value.length && prop.mode != "toggle") {
        selection.value = [DateTime.now().toISODate()];
      }
    }
    expose({ onDay, selection });
    onBeforeMount(() => {
      loadDatesLocal();
    });
    onMounted(() => {
      onCalendarChanged();
    });
    return (_ctx, _cache) => {
      return openBlock(), createElementBlock("div", {
        class: normalizeClass(["calendar", [__props.mode]])
      }, [
        createBaseVNode("div", null, [
          createBaseVNode("button", {
            class: "prev icon",
            onClick: onPrev,
            disabled: !unref(hasPrev)
          }, [
            createVNode(ArrowImage)
          ], 8, _hoisted_1$8),
          createBaseVNode("div", null, [
            withDirectives(createBaseVNode("select", {
              class: "month",
              "onUpdate:modelValue": _cache[0] || (_cache[0] = ($event) => month.value = $event),
              onChange: onCalendarChanged
            }, [
              createBaseVNode("option", _hoisted_2$6, toDisplayString(unref(t)("month.1")), 1),
              createBaseVNode("option", _hoisted_3$4, toDisplayString(unref(t)("month.2")), 1),
              createBaseVNode("option", _hoisted_4$4, toDisplayString(unref(t)("month.3")), 1),
              createBaseVNode("option", _hoisted_5$4, toDisplayString(unref(t)("month.4")), 1),
              createBaseVNode("option", _hoisted_6$3, toDisplayString(unref(t)("month.5")), 1),
              createBaseVNode("option", _hoisted_7$2, toDisplayString(unref(t)("month.6")), 1),
              createBaseVNode("option", _hoisted_8$2, toDisplayString(unref(t)("month.7")), 1),
              createBaseVNode("option", _hoisted_9$2, toDisplayString(unref(t)("month.8")), 1),
              createBaseVNode("option", _hoisted_10$2, toDisplayString(unref(t)("month.9")), 1),
              createBaseVNode("option", _hoisted_11$2, toDisplayString(unref(t)("month.10")), 1),
              createBaseVNode("option", _hoisted_12$2, toDisplayString(unref(t)("month.11")), 1),
              createBaseVNode("option", _hoisted_13$1, toDisplayString(unref(t)("month.12")), 1)
            ], 544), [
              [
                vModelSelect,
                month.value,
                void 0,
                { number: true }
              ]
            ]),
            withDirectives(createBaseVNode("select", {
              class: "year",
              "onUpdate:modelValue": _cache[1] || (_cache[1] = ($event) => year.value = $event),
              onChange: onCalendarChanged
            }, [
              (openBlock(true), createElementBlock(Fragment, null, renderList(unref(years), (y) => {
                return openBlock(), createElementBlock("option", { value: y }, toDisplayString(y), 9, _hoisted_14$1);
              }), 256))
            ], 544), [
              [
                vModelSelect,
                year.value,
                void 0,
                { number: true }
              ]
            ])
          ]),
          createBaseVNode("button", {
            class: "next icon",
            onClick: onNext,
            disabled: !unref(hasNext)
          }, [
            createVNode(ArrowImage)
          ], 8, _hoisted_15)
        ]),
        createBaseVNode("table", {
          ref_key: "calendar",
          ref: calendar
        }, [
          createBaseVNode("thead", null, [
            createBaseVNode("tr", null, [
              createBaseVNode("th", null, toDisplayString(unref(t)("day.cal1")), 1),
              createBaseVNode("th", null, toDisplayString(unref(t)("day.cal2")), 1),
              createBaseVNode("th", null, toDisplayString(unref(t)("day.cal3")), 1),
              createBaseVNode("th", null, toDisplayString(unref(t)("day.cal4")), 1),
              createBaseVNode("th", null, toDisplayString(unref(t)("day.cal5")), 1),
              createBaseVNode("th", null, toDisplayString(unref(t)("day.cal6")), 1),
              createBaseVNode("th", null, toDisplayString(unref(t)("day.cal7")), 1)
            ])
          ]),
          createBaseVNode("tbody", null, [
            createBaseVNode("tr", null, [
              createBaseVNode("td", null, [
                createBaseVNode("button", { onClick: onDay })
              ]),
              createBaseVNode("td", null, [
                createBaseVNode("button", { onClick: onDay })
              ]),
              createBaseVNode("td", null, [
                createBaseVNode("button", { onClick: onDay })
              ]),
              createBaseVNode("td", null, [
                createBaseVNode("button", { onClick: onDay })
              ]),
              createBaseVNode("td", null, [
                createBaseVNode("button", { onClick: onDay })
              ]),
              createBaseVNode("td", null, [
                createBaseVNode("button", { onClick: onDay })
              ]),
              createBaseVNode("td", null, [
                createBaseVNode("button", { onClick: onDay })
              ])
            ]),
            createBaseVNode("tr", null, [
              createBaseVNode("td", null, [
                createBaseVNode("button", { onClick: onDay })
              ]),
              createBaseVNode("td", null, [
                createBaseVNode("button", { onClick: onDay })
              ]),
              createBaseVNode("td", null, [
                createBaseVNode("button", { onClick: onDay })
              ]),
              createBaseVNode("td", null, [
                createBaseVNode("button", { onClick: onDay })
              ]),
              createBaseVNode("td", null, [
                createBaseVNode("button", { onClick: onDay })
              ]),
              createBaseVNode("td", null, [
                createBaseVNode("button", { onClick: onDay })
              ]),
              createBaseVNode("td", null, [
                createBaseVNode("button", { onClick: onDay })
              ])
            ]),
            createBaseVNode("tr", null, [
              createBaseVNode("td", null, [
                createBaseVNode("button", { onClick: onDay })
              ]),
              createBaseVNode("td", null, [
                createBaseVNode("button", { onClick: onDay })
              ]),
              createBaseVNode("td", null, [
                createBaseVNode("button", { onClick: onDay })
              ]),
              createBaseVNode("td", null, [
                createBaseVNode("button", { onClick: onDay })
              ]),
              createBaseVNode("td", null, [
                createBaseVNode("button", { onClick: onDay })
              ]),
              createBaseVNode("td", null, [
                createBaseVNode("button", { onClick: onDay })
              ]),
              createBaseVNode("td", null, [
                createBaseVNode("button", { onClick: onDay })
              ])
            ]),
            createBaseVNode("tr", null, [
              createBaseVNode("td", null, [
                createBaseVNode("button", { onClick: onDay })
              ]),
              createBaseVNode("td", null, [
                createBaseVNode("button", { onClick: onDay })
              ]),
              createBaseVNode("td", null, [
                createBaseVNode("button", { onClick: onDay })
              ]),
              createBaseVNode("td", null, [
                createBaseVNode("button", { onClick: onDay })
              ]),
              createBaseVNode("td", null, [
                createBaseVNode("button", { onClick: onDay })
              ]),
              createBaseVNode("td", null, [
                createBaseVNode("button", { onClick: onDay })
              ]),
              createBaseVNode("td", null, [
                createBaseVNode("button", { onClick: onDay })
              ])
            ]),
            createBaseVNode("tr", null, [
              createBaseVNode("td", null, [
                createBaseVNode("button", { onClick: onDay })
              ]),
              createBaseVNode("td", null, [
                createBaseVNode("button", { onClick: onDay })
              ]),
              createBaseVNode("td", null, [
                createBaseVNode("button", { onClick: onDay })
              ]),
              createBaseVNode("td", null, [
                createBaseVNode("button", { onClick: onDay })
              ]),
              createBaseVNode("td", null, [
                createBaseVNode("button", { onClick: onDay })
              ]),
              createBaseVNode("td", null, [
                createBaseVNode("button", { onClick: onDay })
              ]),
              createBaseVNode("td", null, [
                createBaseVNode("button", { onClick: onDay })
              ])
            ]),
            createBaseVNode("tr", null, [
              createBaseVNode("td", null, [
                createBaseVNode("button", { onClick: onDay })
              ]),
              createBaseVNode("td", null, [
                createBaseVNode("button", { onClick: onDay })
              ]),
              createBaseVNode("td", null, [
                createBaseVNode("button", { onClick: onDay })
              ]),
              createBaseVNode("td", null, [
                createBaseVNode("button", { onClick: onDay })
              ]),
              createBaseVNode("td", null, [
                createBaseVNode("button", { onClick: onDay })
              ]),
              createBaseVNode("td", null, [
                createBaseVNode("button", { onClick: onDay })
              ]),
              createBaseVNode("td", null, [
                createBaseVNode("button", { onClick: onDay })
              ])
            ])
          ])
        ], 512)
      ], 2);
    };
  }
};
var DiaryEntryList_vue_vue_type_style_index_0_lang = "";
const _hoisted_1$7 = { class: "unit" };
const _hoisted_2$5 = ["disabled"];
const _hoisted_3$3 = ["value"];
const _hoisted_4$3 = { class: "unit" };
const _hoisted_5$3 = ["value"];
const _hoisted_6$2 = ["value"];
const _hoisted_7$1 = ["value"];
const _hoisted_8$1 = ["value", "disabled"];
const _hoisted_9$1 = { class: "unit" };
const _hoisted_10$1 = ["value"];
const _hoisted_11$1 = ["value"];
const _hoisted_12$1 = ["value"];
const _hoisted_13 = { key: 0 };
const _hoisted_14 = { key: 1 };
const _sfc_main$7 = {
  props: ["entries", "nutrient", "disabled"],
  setup(__props, { expose }) {
    const prop = __props;
    const { t } = useI18n();
    const foods = inject("food");
    const prefs2 = inject("prefs");
    const groupMode = ref("hour");
    const nutrientMode = ref("relative");
    const form = ref(null);
    const groupedEntries = computed(() => {
      switch (groupMode.value) {
        case "hour":
          return groupedByHour();
        case "day":
          return groupedByDay();
        default:
          return [];
      }
    });
    const nutrientUnit = computed(() => {
      if (nutrientMode.value == "relative") {
        return "%";
      }
      switch (prop.nutrient) {
        case "kcal":
          return " " + t("unit.cal");
        case "fat":
        case "carb":
        case "prot":
          return " " + t("unit.g");
        default:
          return " " + t("unit.mg");
      }
    });
    function groupedByHour() {
      let result = [];
      if (prop.entries) {
        let groups = {};
        prop.entries.forEach((entry) => {
          let hour = parseInt(entry.time);
          let start = DateTime.fromObject({ hours: hour - hour % 2 });
          let name = start.toFormat("t") + " - " + start.plus({ hours: 2 }).toFormat("t");
          if (!(name in groups)) {
            groups[name] = { name, entries: [] };
          }
          let next = {
            id: entry.food.id,
            name: t(entry.food.id.toString()),
            amount: entry.food.amount,
            nutrient: getNutrient(entry.food),
            recipe: entry.recipe,
            time: entry.time,
            isrec: false
          };
          if (entry.recipe) {
            let exists = false;
            groups[name].entries.every((e) => {
              if (e.isrec && e.name == entry.recipe) {
                e.entries.push(next);
                e.amount += next.amount;
                e.nutrient += next.nutrient;
                exists = true;
                return false;
              }
              return true;
            });
            if (!exists) {
              groups[name].entries.push({
                name: entry.recipe,
                isrec: true,
                amount: next.amount,
                nutrient: next.nutrient,
                entries: [next]
              });
            }
          } else {
            groups[name].entries.push(next);
          }
        });
        result = Object.values(groups);
      }
      return result;
    }
    function groupedByDay() {
      return [
        { name: "Morning", entries: [{ type: "recipe" }, { type: "single" }] },
        { name: "Noon", entries: [{ type: "recipe" }, { type: "single" }] }
      ];
    }
    function getNutrient(food2) {
      let data = foods.value.filter((f) => f.id == food2.id)[0];
      let amount = data[prop.nutrient] * food2.amount * 0.01;
      if (nutrientMode.value == "metric") {
        return Math.round(amount);
      }
      let rdi = prefs2.value.rdi[prop.nutrient];
      return Math.round(amount * 100 / rdi);
    }
    function toggleNutrientMode() {
      if (nutrientMode.value == "relative") {
        nutrientMode.value = "metric";
      } else if (nutrientMode.value == "metric") {
        nutrientMode.value = "relative";
      }
    }
    function onRecipeDetails(evt) {
      let parent = evt.target.closest("div");
      parent.classList.toggle("open");
    }
    function onKeydown(evt) {
      if (evt.key == "Enter") {
        evt.preventDefault();
        onInput(evt);
      }
    }
    function onInput(evt) {
      evt.target.blur();
      let val = parseFloat(evt.target.value);
      if (isNaN(val) || val < 0) {
        evt.target.value = 0;
      }
    }
    function getDiff() {
      let data = new FormData(form.value);
      let ids = data.getAll("id");
      let amounts = data.getAll("amount");
      let recipes2 = data.getAll("recipe");
      let times = data.getAll("time");
      let result = [];
      prop.entries.forEach((entry) => {
        for (let i = 0; i < ids.length; ++i) {
          if (ids[i] != entry.food.id) {
            continue;
          }
          if (recipes2[i] != entry.recipe) {
            continue;
          }
          if (times[i] != entry.time) {
            continue;
          }
          if (amounts[i] != entry.food.amount) {
            result.push({ id: ids[i], amount: amounts[i], recipe: recipes2[i], time: times[i] });
          }
          break;
        }
      });
      return result;
    }
    expose({ getDiff });
    return (_ctx, _cache) => {
      return openBlock(), createElementBlock("form", {
        class: "diary-entry-list",
        ref_key: "form",
        ref: form
      }, [
        (openBlock(true), createElementBlock(Fragment, null, renderList(unref(groupedEntries), (group) => {
          return openBlock(), createElementBlock("fieldset", {
            key: group.name
          }, [
            createBaseVNode("legend", null, toDisplayString(group.name), 1),
            (openBlock(true), createElementBlock(Fragment, null, renderList(group.entries, (entry) => {
              return openBlock(), createElementBlock("div", {
                class: normalizeClass({ recipe: entry.isrec })
              }, [
                entry.isrec ? (openBlock(), createElementBlock(Fragment, { key: 0 }, [
                  createBaseVNode("label", null, toDisplayString(entry.name), 1),
                  createBaseVNode("button", {
                    class: "icon",
                    type: "button",
                    onClick: onRecipeDetails
                  }, [
                    createVNode(MoreImage)
                  ]),
                  createBaseVNode("span", null, toDisplayString(entry.amount), 1),
                  createBaseVNode("span", _hoisted_1$7, toDisplayString(unref(t)("unit.g")), 1),
                  createBaseVNode("span", {
                    class: normalizeClass(["nutrient", [__props.nutrient, nutrientMode.value]])
                  }, toDisplayString(entry.nutrient) + toDisplayString(unref(nutrientUnit)), 3),
                  createBaseVNode("fieldset", {
                    disabled: __props.disabled,
                    style: normalizeStyle({ "--max-height": entry.entries.length * 41 + "px" })
                  }, [
                    (openBlock(true), createElementBlock(Fragment, null, renderList(entry.entries, (food2) => {
                      return openBlock(), createElementBlock("div", {
                        key: food2.id
                      }, [
                        createBaseVNode("label", null, toDisplayString(food2.name), 1),
                        createBaseVNode("input", {
                          type: "number",
                          name: "amount",
                          value: food2.amount,
                          onKeydown,
                          onChange: onInput
                        }, null, 40, _hoisted_3$3),
                        createBaseVNode("span", _hoisted_4$3, toDisplayString(unref(t)("unit.g")), 1),
                        createBaseVNode("span", {
                          class: normalizeClass(["nutrient", [__props.nutrient, nutrientMode.value]])
                        }, toDisplayString(food2.nutrient) + toDisplayString(unref(nutrientUnit)), 3),
                        createBaseVNode("input", {
                          type: "hidden",
                          name: "id",
                          value: food2.id
                        }, null, 8, _hoisted_5$3),
                        createBaseVNode("input", {
                          type: "hidden",
                          name: "recipe",
                          value: food2.recipe
                        }, null, 8, _hoisted_6$2),
                        createBaseVNode("input", {
                          type: "hidden",
                          name: "time",
                          value: food2.time
                        }, null, 8, _hoisted_7$1)
                      ]);
                    }), 128))
                  ], 12, _hoisted_2$5)
                ], 64)) : (openBlock(), createElementBlock(Fragment, { key: 1 }, [
                  createBaseVNode("label", null, toDisplayString(entry.name), 1),
                  createBaseVNode("input", {
                    type: "number",
                    name: "amount",
                    value: entry.amount,
                    disabled: __props.disabled,
                    onKeydown,
                    onChange: onInput
                  }, null, 40, _hoisted_8$1),
                  createBaseVNode("span", _hoisted_9$1, toDisplayString(unref(t)("unit.g")), 1),
                  createBaseVNode("span", {
                    class: normalizeClass(["nutrient", [__props.nutrient, nutrientMode.value]])
                  }, toDisplayString(entry.nutrient) + toDisplayString(unref(nutrientUnit)), 3),
                  createBaseVNode("input", {
                    type: "hidden",
                    name: "id",
                    value: entry.id
                  }, null, 8, _hoisted_10$1),
                  createBaseVNode("input", {
                    type: "hidden",
                    name: "recipe",
                    value: entry.recipe
                  }, null, 8, _hoisted_11$1),
                  createBaseVNode("input", {
                    type: "hidden",
                    name: "time",
                    value: entry.time
                  }, null, 8, _hoisted_12$1)
                ], 64))
              ], 2);
            }), 256))
          ]);
        }), 128)),
        createBaseVNode("button", {
          type: "button",
          class: "nutrient-mode-switch icon",
          onClick: toggleNutrientMode
        }, [
          nutrientMode.value == "relative" ? (openBlock(), createElementBlock("span", _hoisted_13, "%")) : createCommentVNode("", true),
          nutrientMode.value == "metric" ? (openBlock(), createElementBlock("span", _hoisted_14, "g")) : createCommentVNode("", true)
        ])
      ], 512);
    };
  }
};
var PieChart_vue_vue_type_style_index_0_lang = "";
const _hoisted_1$6 = { class: "pie-chart" };
const _hoisted_2$4 = { viewBox: "0 0 256 256" };
const _hoisted_3$2 = ["transform", "stroke-dasharray"];
const _hoisted_4$2 = ["stroke-dasharray", "transform"];
const _hoisted_5$2 = ["stroke-dasharray", "transform"];
const _hoisted_6$1 = /* @__PURE__ */ createBaseVNode("hr", null, null, -1);
const _sfc_main$6 = {
  props: ["start", "range", "value", "max", "label", "unit", "frac"],
  setup(__props) {
    const prop = __props;
    const radius = 112;
    const formattedValue = computed(() => +(parseFloat(prop.value) || 0).toFixed(prop.frac || 0));
    const formattedMax = computed(() => +(parseFloat(prop.max) || 0).toFixed(prop.frac || 0));
    const transform = computed(() => {
      let rot = -90 + (parseFloat(prop.start) || 0);
      return `rotate(${rot} 128 128)`;
    });
    const fullArc = computed(() => {
      let ratio = (parseFloat(prop.range) || 0) / 360;
      return Math.PI * 2 * radius * ratio + " 710";
    });
    const valueArc = computed(() => {
      let normalized = (parseFloat(prop.value) || 0) / (parseFloat(prop.max) || 1);
      let range = parseFloat(prop.range) || 0;
      let ratio = Math.min(normalized, 1) * range / 360;
      return Math.PI * 2 * radius * ratio + " 710";
    });
    const overArc = computed(() => {
      let normalized = (parseFloat(prop.value) || 0) / (parseFloat(prop.max) || 1);
      let range = parseFloat(prop.range) || 0;
      let ratio = Math.min(Math.max(0, normalized - 1), 1) * range / 360;
      return Math.PI * 2 * radius * ratio + " 710";
    });
    return (_ctx, _cache) => {
      return openBlock(), createElementBlock("figure", _hoisted_1$6, [
        (openBlock(), createElementBlock("svg", _hoisted_2$4, [
          createBaseVNode("circle", {
            class: "base",
            r: "112",
            cx: "128",
            cy: "128",
            transform: unref(transform),
            "stroke-dasharray": unref(fullArc)
          }, null, 8, _hoisted_3$2),
          createBaseVNode("circle", {
            class: "good",
            r: "112",
            cx: "128",
            cy: "128",
            "stroke-dasharray": unref(valueArc),
            transform: unref(transform)
          }, null, 8, _hoisted_4$2),
          createBaseVNode("circle", {
            class: "bad",
            r: "112",
            cx: "128",
            cy: "128",
            "stroke-dasharray": unref(overArc),
            transform: unref(transform)
          }, null, 8, _hoisted_5$2)
        ])),
        createBaseVNode("div", null, [
          renderSlot(_ctx.$slots, "details", {}, () => [
            createBaseVNode("span", null, toDisplayString(unref(formattedValue)) + " " + toDisplayString(__props.unit), 1),
            _hoisted_6$1,
            createBaseVNode("span", null, toDisplayString(unref(formattedMax)) + " " + toDisplayString(__props.unit), 1)
          ])
        ]),
        createBaseVNode("figcaption", null, toDisplayString(__props.label), 1)
      ]);
    };
  }
};
var DiaryView_vue_vue_type_style_index_0_lang = "";
const _hoisted_1$5 = /* @__PURE__ */ createBaseVNode("hr", null, null, -1);
const _hoisted_2$3 = /* @__PURE__ */ createBaseVNode("section", null, null, -1);
const _hoisted_3$1 = { id: "charts-week" };
const _hoisted_4$1 = ["data-date"];
const _hoisted_5$1 = { id: "charts-macro" };
const _hoisted_6 = { class: "no-edit-mode" };
const _hoisted_7 = { class: "subtitle no-edit-mode" };
const _hoisted_8 = { class: "tabs" };
const _hoisted_9 = ["disabled"];
const _hoisted_10 = /* @__PURE__ */ createBaseVNode("hr", null, null, -1);
const _hoisted_11 = { key: 0 };
const _hoisted_12 = ["innerHTML"];
const _sfc_main$5 = {
  setup(__props) {
    const { t } = useI18n();
    const log2 = inject("log");
    const csrf = inject("csrfToken");
    const diary2 = inject("diary");
    const prefs2 = inject("prefs");
    const current = ref(null);
    const currentDate = ref(DateTime.now());
    const currentNutrient = ref("kcal");
    const disableSave = ref(false);
    const editMode = ref(false);
    const calendar = ref(null);
    const entries = ref(null);
    const main = ref(null);
    const daysWithEntries = computed(() => Object.keys(diary2.value));
    const currentWeek = computed(() => {
      let result = [];
      let isoNow = DateTime.now().toISODate();
      let isoCurrent = currentDate.value.toISODate();
      let date = currentDate.value.minus({ days: currentDate.value.weekday - 1 });
      for (let i = 0; i < 7; ++i) {
        let iso = date.toISODate();
        result.push({
          weekday: t("day.cal" + date.weekday),
          calday: date.day,
          date: iso,
          active: iso == isoCurrent,
          today: iso == isoNow,
          value: diary2.value[iso] ? diary2.value[iso].kcal : 0
        });
        date = date.plus({ days: 1 });
      }
      return result;
    });
    let hasTabDrag = false;
    function onTabSlide(evt) {
      evt.stopPropagation();
      evt.preventDefault();
      moveTabBar(evt.target.closest("ul"), evt.movementX);
      hasTabDrag = true;
    }
    function onTabWheel(evt) {
      evt.stopPropagation();
      evt.preventDefault();
      let delta = Math.max(-16, Math.min(-evt.deltaY, 16));
      moveTabBar(evt.target.closest("ul"), delta);
    }
    function moveTabBar(elem, delta) {
      let offset = parseInt(elem.style.getPropertyValue("--offset")) || 0;
      let final = Math.max(elem.clientWidth - elem.scrollWidth, Math.min(offset + delta, 0));
      elem.style.setProperty("--offset", final + "px");
    }
    function onTabsPress(evt) {
      let handle = evt.target.closest("ul");
      handle.addEventListener("pointermove", onTabSlide);
      handle.addEventListener("mouseup", onTabsRelease);
      handle.addEventListener("touchend", onTabsRelease);
      handle.addEventListener("mouseleave", onTabsRelease);
      handle.addEventListener("touchcancel", onTabsRelease);
      hasTabDrag = false;
    }
    function onTabsRelease(evt) {
      let handle = evt.target.closest("ul");
      handle.removeEventListener("pointermove", onTabSlide);
      handle.removeEventListener("mouseup", onTabsRelease);
      handle.removeEventListener("touchend", onTabsRelease);
      handle.removeEventListener("mouseleave", onTabsRelease);
      handle.removeEventListener("touchcancel", onTabsRelease);
    }
    function onTabClick(evt) {
      if (!hasTabDrag) {
        currentNutrient.value = evt.target.dataset.name;
      }
    }
    function onEditMode() {
      if (current.value && current.value.entries) {
        editMode.value ? saveEntries() : editMode.value = true;
      }
    }
    function saveEntries() {
      let items = entries.value.getDiff();
      if (items.length == 0) {
        editMode.value = false;
        return;
      }
      disableSave.value = true;
      let date = current.value.date;
      let params = new URLSearchParams();
      items.forEach((i) => {
        params.append("id", i.id);
        params.append("amount", i.amount);
        params.append("time", i.time);
        params.append("recipe", i.recipe);
      });
      fetch(`/api/v1/diary/${date}`, {
        method: "PUT",
        body: params,
        headers: {
          "Content-Type": "application/x-www-form-urlencoded",
          "X-CSRF-Token": csrf
        }
      }).then((response) => {
        if (!response.ok) {
          throw t("savediary.err" + response.status);
        }
        editMode.value = false;
        return fetch("/api/v1/diary/" + date.replaceAll("-", "/"));
      }).then((response) => response.json()).then((data) => {
        if (data[0] && data[0].kcal) {
          diary2.value[date] = data[0];
        } else {
          delete diary2.value[date];
        }
        onDateSelected([date]);
        log2.msg(t("savediary.ok"));
      }).catch((err) => log2.err(err)).finally(() => {
        setTimeout(function() {
          disableSave.value = false;
        }, 500);
      });
    }
    function onDateSelected(dates) {
      currentDate.value = DateTime.fromISO(dates[0]);
      let date = currentDate.value.toISODate();
      if (!diary2.value[date]) {
        current.value = { date };
        return;
      }
      if (diary2.value[date].entries) {
        current.value = diary2.value[date];
        return;
      }
      fetch(`/api/v1/diary/${date.replaceAll("-", "/")}/entries`).then((response) => response.json()).then((data) => {
        data.forEach((d) => d.time = d.date.split("T")[1].slice(0, 5));
        data.sort((a, b) => {
          if (a.time < b.time)
            return -1;
          if (b.time < a.time)
            return 1;
          return 0;
        });
        diary2.value[date].entries = data;
      }).catch((err) => log2.err(err)).finally(() => current.value = diary2.value[date]);
    }
    onMounted(() => {
      onDateSelected(calendar.value.selection);
    });
    return (_ctx, _cache) => {
      return openBlock(), createBlock(_sfc_main$l, {
        ref_key: "main",
        ref: main,
        class: normalizeClass(["diary", { "edit-mode": editMode.value }]),
        onDetailVisibility: _cache[1] || (_cache[1] = ($event) => editMode.value = false)
      }, {
        filter: withCtx(() => [
          createBaseVNode("section", null, [
            createBaseVNode("h2", null, toDisplayString(unref(t)("aria.headcal")), 1),
            createVNode(_sfc_main$8, {
              ref_key: "calendar",
              ref: calendar,
              storage: "caldiary",
              items: unref(daysWithEntries),
              onSelection: onDateSelected
            }, null, 8, ["items"])
          ]),
          _hoisted_1$5,
          _hoisted_2$3
        ]),
        main: withCtx(() => [
          createBaseVNode("section", _hoisted_3$1, [
            (openBlock(true), createElementBlock(Fragment, null, renderList(unref(currentWeek), (day) => {
              return openBlock(), createElementBlock("button", {
                "data-date": day.date,
                class: normalizeClass({ today: day.today, active: day.active }),
                onClick: _cache[0] || (_cache[0] = (...args) => calendar.value.onDay && calendar.value.onDay(...args))
              }, [
                createVNode(_sfc_main$6, {
                  range: "360",
                  value: day.value,
                  max: unref(prefs2).rdi.kcal
                }, {
                  details: withCtx(() => [
                    createBaseVNode("span", null, toDisplayString(day.weekday), 1),
                    createBaseVNode("span", null, toDisplayString(day.calday), 1)
                  ]),
                  _: 2
                }, 1032, ["value", "max"])
              ], 10, _hoisted_4$1);
            }), 256))
          ]),
          createBaseVNode("section", _hoisted_5$1, [
            createVNode(_sfc_main$6, {
              class: "kcal",
              start: "225",
              range: "270",
              frac: "0",
              label: unref(t)("food.energy"),
              unit: unref(t)("unit.cal"),
              value: current.value ? current.value.kcal : 0,
              max: unref(prefs2).rdi.kcal
            }, null, 8, ["label", "unit", "value", "max"]),
            createVNode(_sfc_main$6, {
              class: "fat",
              start: "225",
              range: "270",
              frac: "0",
              label: unref(t)("food.fat"),
              unit: unref(t)("unit.g"),
              value: current.value ? current.value.fat : 0,
              max: unref(prefs2).rdi.fat
            }, null, 8, ["label", "unit", "value", "max"]),
            createVNode(_sfc_main$6, {
              class: "carb",
              start: "225",
              range: "270",
              frac: "0",
              label: unref(t)("food.carbs2"),
              unit: unref(t)("unit.g"),
              value: current.value ? current.value.carb : 0,
              max: unref(prefs2).rdi.carb
            }, null, 8, ["label", "unit", "value", "max"]),
            createVNode(_sfc_main$6, {
              class: "prot",
              start: "225",
              range: "270",
              frac: "0",
              label: unref(t)("food.protein"),
              unit: unref(t)("unit.g"),
              value: current.value ? current.value.prot : 0,
              max: unref(prefs2).rdi.prot
            }, null, 8, ["label", "unit", "value", "max"])
          ])
        ]),
        "head-details": withCtx(() => [
          createBaseVNode("h2", _hoisted_6, toDisplayString(currentDate.value.weekdayLong), 1)
        ]),
        details: withCtx(() => [
          createBaseVNode("section", _hoisted_7, toDisplayString(currentDate.value.toLocaleString(unref(DateTime).DATE_FULL)), 1),
          createBaseVNode("section", _hoisted_8, [
            createBaseVNode("div", null, [
              createBaseVNode("ul", {
                onMousedown: onTabsPress,
                onTouchstart: onTabsPress,
                onWheel: onTabWheel
              }, [
                createBaseVNode("li", {
                  class: normalizeClass({ active: currentNutrient.value == "kcal" })
                }, [
                  createBaseVNode("button", {
                    "data-name": "kcal",
                    onClick: onTabClick
                  }, toDisplayString(unref(t)("food.energy")), 1)
                ], 2),
                createBaseVNode("li", {
                  class: normalizeClass({ active: currentNutrient.value == "fat" })
                }, [
                  createBaseVNode("button", {
                    "data-name": "fat",
                    onClick: onTabClick
                  }, toDisplayString(unref(t)("food.fat")), 1)
                ], 2),
                createBaseVNode("li", {
                  class: normalizeClass({ active: currentNutrient.value == "carb" })
                }, [
                  createBaseVNode("button", {
                    "data-name": "carb",
                    onClick: onTabClick
                  }, toDisplayString(unref(t)("food.carbs2")), 1)
                ], 2),
                createBaseVNode("li", {
                  class: normalizeClass({ active: currentNutrient.value == "prot" })
                }, [
                  createBaseVNode("button", {
                    "data-name": "prot",
                    onClick: onTabClick
                  }, toDisplayString(unref(t)("food.protein")), 1)
                ], 2)
              ], 32)
            ]),
            createBaseVNode("button", {
              class: "icon async",
              disabled: disableSave.value,
              onClick: onEditMode
            }, [
              !editMode.value ? (openBlock(), createBlock(EditImage, { key: 0 })) : createCommentVNode("", true),
              editMode.value ? (openBlock(), createBlock(SaveImage, { key: 1 })) : createCommentVNode("", true)
            ], 8, _hoisted_9)
          ]),
          _hoisted_10,
          current.value ? (openBlock(), createElementBlock("section", _hoisted_11, [
            createBaseVNode("h2", null, toDisplayString(unref(t)("aria.headlog")), 1),
            !current.value.entries ? (openBlock(), createElementBlock("p", {
              key: 0,
              class: "msg-noitems",
              innerHTML: unref(t)("diary.noitems")
            }, null, 8, _hoisted_12)) : createCommentVNode("", true),
            current.value.entries ? (openBlock(), createBlock(_sfc_main$7, {
              key: 1,
              ref_key: "entries",
              ref: entries,
              entries: current.value.entries,
              nutrient: currentNutrient.value,
              disabled: !editMode.value
            }, null, 8, ["entries", "nutrient", "disabled"])) : createCommentVNode("", true)
          ])) : createCommentVNode("", true)
        ]),
        _: 1
      }, 8, ["class"]);
    };
  }
};
var Checkbox_vue_vue_type_style_index_0_lang = "";
const _hoisted_1$4 = ["checked"];
const _hoisted_2$2 = /* @__PURE__ */ createBaseVNode("div", null, null, -1);
const _sfc_main$4 = {
  props: ["checked"],
  emits: ["change", "click"],
  setup(__props, { emit }) {
    function onChange(evt) {
      emit("change", evt);
    }
    function onClick(evt) {
      emit("click", evt);
    }
    return (_ctx, _cache) => {
      return openBlock(), createElementBlock("label", null, [
        createBaseVNode("input", {
          type: "checkbox",
          checked: __props.checked,
          onChange,
          onClick
        }, null, 40, _hoisted_1$4),
        _hoisted_2$2
      ]);
    };
  }
};
const _hoisted_1$3 = { class: "select" };
const _hoisted_2$1 = { class: "select" };
const _hoisted_3 = ["innerHTML"];
const _hoisted_4 = ["onClick"];
const _hoisted_5 = { class: "s" };
const _sfc_main$3 = {
  props: ["items", "offline"],
  emits: "selected",
  setup(__props, { emit }) {
    const prop = __props;
    const { t, locale: locale2 } = useI18n();
    inject("log");
    const csrf = inject("csrfToken");
    const sortBy = ref("amount");
    const sortDir = ref("desc");
    const collator = new Intl.Collator(locale2.value, { numeric: true });
    const sortedItems = computed(() => {
      if (sortDir.value == "asc") {
        return [...prop.items].sort((a, b) => {
          return collator.compare(a[sortBy.value], b[sortBy.value]);
        });
      } else {
        return [...prop.items].sort((a, b) => {
          return -collator.compare(a[sortBy.value], b[sortBy.value]);
        });
      }
    });
    const allChecked = computed(() => {
      if (prop.items && prop.items.length) {
        return prop.items.length == prop.items.filter((i) => i.done).length;
      }
      return false;
    });
    watch(() => prop.offline, resync);
    function formattedAmount(item) {
      if (item.amount > 999) {
        var amount = +parseFloat(item.amount * 1e-3).toFixed(1);
        var unit = t("unit.kg");
      } else {
        var amount = item.amount;
        var unit = t("unit.g");
      }
      return `${amount} <span class="unit">${unit}</span>`;
    }
    function setActive(evt) {
      let cat = evt.target.dataset.sort;
      if (sortBy.value == cat) {
        sortDir.value = sortDir.value == "asc" ? "desc" : "asc";
      } else {
        sortBy.value = cat;
      }
    }
    function resync() {
      if (!prop.offline) {
        let params = new URLSearchParams();
        prop.items.forEach((i) => {
          if (i.local) {
            params.append("id", i.id);
            params.append("done", i.done);
          }
        });
        if (params.has("id")) {
          fetch("/api/v1/list/diary/done", {
            method: "PUT",
            body: params,
            headers: {
              "Content-Type": "application/x-www-form-urlencoded",
              "X-CSRF-Token": csrf
            }
          }).then((response) => {
            if (response.ok) {
              prop.items.forEach((i) => i.local = false);
            }
          }).catch(() => {
          });
        }
      }
    }
    function onCheckedAll(evt) {
      let val = evt.target.checked;
      let params = new URLSearchParams();
      prop.items.forEach((i) => {
        i.done = val;
        i.local = prop.offline;
        params.append("id", i.id);
        params.append("done", i.done);
      });
      if (!prop.offline) {
        fetch("/api/v1/list/diary/done", {
          method: "PUT",
          body: params,
          headers: {
            "Content-Type": "application/x-www-form-urlencoded",
            "X-CSRF-Token": csrf
          }
        }).catch(() => {
        });
      }
    }
    function onChecked(evt) {
      let id = evt.target.closest("label").dataset.id;
      let val = evt.target.checked;
      let item = prop.items.filter((i) => i.id == id)[0];
      item.done = val;
      item.local = prop.offline;
      if (!prop.offline) {
        fetch("/api/v1/list/diary/done", {
          method: "PUT",
          body: new URLSearchParams({ id, done: val }),
          headers: {
            "Content-Type": "application/x-www-form-urlencoded",
            "X-CSRF-Token": csrf
          }
        }).catch(() => {
        });
      }
    }
    onMounted(resync);
    return (_ctx, _cache) => {
      return openBlock(), createElementBlock("table", null, [
        createBaseVNode("thead", null, [
          createBaseVNode("tr", {
            class: normalizeClass(sortDir.value)
          }, [
            createBaseVNode("th", _hoisted_1$3, [
              createVNode(_sfc_main$4, {
                checked: unref(allChecked),
                onClick: onCheckedAll
              }, null, 8, ["checked"])
            ]),
            createBaseVNode("th", {
              class: normalizeClass(["num sort", { active: sortBy.value == "amount" }]),
              onClick: setActive,
              "data-sort": "amount"
            }, [
              createVNode(Arrow),
              createTextVNode(" " + toDisplayString(unref(t)("food.amount")), 1)
            ], 2),
            createBaseVNode("th", {
              class: normalizeClass(["name sort", { active: sortBy.value == "name" }]),
              onClick: setActive,
              "data-sort": "name"
            }, [
              createTextVNode(toDisplayString(unref(t)("food.name")) + " ", 1),
              createVNode(Arrow)
            ], 2),
            createBaseVNode("th", {
              class: normalizeClass(["s sort", { active: sortBy.value == "aisle" }]),
              onClick: setActive,
              "data-sort": "aisle"
            }, [
              createVNode(Arrow),
              createTextVNode(" " + toDisplayString(unref(t)("food.aisle")), 1)
            ], 2)
          ], 2)
        ]),
        createBaseVNode("tbody", null, [
          (openBlock(true), createElementBlock(Fragment, null, renderList(unref(sortedItems), (item) => {
            return openBlock(), createElementBlock("tr", {
              key: item.id,
              class: normalizeClass({ done: item.done })
            }, [
              createBaseVNode("td", _hoisted_2$1, [
                createVNode(_sfc_main$4, {
                  "data-id": item.id,
                  checked: item.done,
                  onChange: onChecked
                }, null, 8, ["data-id", "checked"])
              ]),
              createBaseVNode("td", {
                class: "num",
                innerHTML: formattedAmount(item)
              }, null, 8, _hoisted_3),
              createBaseVNode("td", {
                class: "name",
                onClick: ($event) => _ctx.$emit("selected", item.id)
              }, toDisplayString(item.name), 9, _hoisted_4),
              createBaseVNode("td", _hoisted_5, toDisplayString(unref(t)("aisle." + item.aisle)), 1)
            ], 2);
          }), 128))
        ])
      ]);
    };
  }
};
var ShoppingView_vue_vue_type_style_index_0_lang = "";
const _hoisted_1$2 = /* @__PURE__ */ createBaseVNode("hr", null, null, -1);
const _hoisted_2 = /* @__PURE__ */ createBaseVNode("section", null, null, -1);
const _sfc_main$2 = {
  setup(__props) {
    const { t } = useI18n();
    const log2 = inject("log");
    const diary2 = inject("diary");
    const filtered = ref([]);
    const offline = ref(false);
    const main = ref(null);
    const calendar = ref(null);
    const daysWithEntries = computed(() => Object.keys(diary2.value));
    let pingIntervalHandle = void 0;
    watch(filtered, (val) => saveItemsLocal(), { deep: true });
    function onDateSelected(dates) {
      if (dates.length == 0) {
        filtered.value = [];
        return;
      }
      if (offline.value) {
        return;
      }
      let params = new URLSearchParams(dates.map((d) => ["date", d]));
      fetch("/api/v1/list/diary?" + params).then((response) => {
        if (!response.ok) {
          throw t("getlist.err" + response.status);
        }
        return response.json();
      }).then((data) => {
        data.forEach((d) => {
          d.name = t(d.id.toString());
          let old = filtered.value.filter((i) => i.id == d.id)[0];
          if (old && old.local) {
            d.local = true;
            d.done = old.done;
          }
        });
        filtered.value = data;
      }).catch((err) => log2.err(err));
    }
    function onOffline() {
      if (!offline.value) {
        offline.value = true;
        log2.warn(t("conn.off"));
      }
    }
    function onOnline() {
      if (offline.value) {
        offline.value = false;
        log2.msg(t("conn.on"));
      }
    }
    function showDetails(id) {
      main.value.showDetails();
    }
    function saveItemsLocal() {
      let str = JSON.stringify(filtered.value);
      window.localStorage.setItem("listdiary", str);
    }
    function loadItemsLocal() {
      let str = window.localStorage.getItem("listdiary") || "[]";
      filtered.value = JSON.parse(str);
    }
    function ping() {
      return fetch("/ping", { method: "HEAD" }).then(onOnline, onOffline);
    }
    onBeforeMount(() => {
      loadItemsLocal();
    });
    onMounted(() => {
      pingIntervalHandle = setInterval(ping, 15e3);
      ping().then(() => onDateSelected(calendar.value.selection));
    });
    onUnmounted(() => {
      clearInterval(pingIntervalHandle);
    });
    return (_ctx, _cache) => {
      return openBlock(), createBlock(_sfc_main$l, {
        ref_key: "main",
        ref: main
      }, {
        filter: withCtx(() => [
          createBaseVNode("section", {
            class: normalizeClass({ offline: offline.value })
          }, [
            createBaseVNode("h2", null, toDisplayString(unref(t)("aria.headcal")), 1),
            createVNode(_sfc_main$8, {
              ref_key: "calendar",
              ref: calendar,
              mode: "toggle",
              storage: "calshop",
              items: unref(daysWithEntries),
              onSelection: onDateSelected
            }, null, 8, ["items"])
          ], 2),
          _hoisted_1$2,
          _hoisted_2
        ]),
        main: withCtx(() => [
          createVNode(_sfc_main$3, {
            items: filtered.value,
            offline: offline.value,
            onSelected: showDetails
          }, null, 8, ["items", "offline"])
        ]),
        _: 1
      }, 512);
    };
  }
};
const _hoisted_1$1 = /* @__PURE__ */ createTextVNode(" Profile ");
const _sfc_main$1 = {
  setup(__props) {
    return (_ctx, _cache) => {
      return openBlock(), createBlock(_sfc_main$l, null, {
        main: withCtx(() => [
          _hoisted_1$1
        ]),
        _: 1
      });
    };
  }
};
const _hoisted_1 = /* @__PURE__ */ createTextVNode(" Settings ");
const _sfc_main = {
  setup(__props) {
    return (_ctx, _cache) => {
      return openBlock(), createBlock(_sfc_main$l, null, {
        main: withCtx(() => [
          _hoisted_1
        ]),
        _: 1
      });
    };
  }
};
const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: "/",
      name: "diary",
      component: _sfc_main$5
    },
    {
      path: "/recipes",
      name: "recipes",
      component: _sfc_main$9
    },
    {
      path: "/food",
      name: "food",
      component: _sfc_main$c
    },
    {
      path: "/shopping",
      name: "shopping",
      component: _sfc_main$2
    },
    {
      path: "/profile",
      name: "profile",
      component: _sfc_main$1
    },
    {
      path: "/settings",
      name: "settings",
      component: _sfc_main
    }
  ]
});
let lang = document.documentElement.lang || navigator.language;
if (!lang && navigator.languages != void 0) {
  lang = navigator.languages[0];
}
Settings.defaultLocale = lang;
const csrfMeta = document.querySelector("meta[name='_csrf']");
const csrfToken = csrfMeta ? csrfMeta.content : "";
const perms = function() {
  const _current = document.documentElement.dataset.perm || 1;
  const _createFood = 65536;
  const _editFood = 131072;
  function _check(perm) {
    return (_current & perm) == perm;
  }
  return {
    canCreateFood: _check(_createFood),
    canEditFood: _check(_editFood)
  };
}();
const log = function() {
  function _getMessage(obj) {
    if (typeof obj === "string") {
      return obj;
    }
    if ("message" in obj) {
      return obj.message;
    }
    return locale.global.t("err.err");
  }
  return {
    msg: function(obj) {
      let payload = { msg: _getMessage(obj), timeout: 3e3 };
      window.dispatchEvent(new CustomEvent("message", { detail: payload }));
    },
    warn: function(obj) {
      let payload = { msg: _getMessage(obj), timeout: 4e3 };
      window.dispatchEvent(new CustomEvent("warning", { detail: payload }));
    },
    err: function(obj) {
      let payload = { msg: _getMessage(obj), timeout: 5e3 };
      window.dispatchEvent(new CustomEvent("error", { detail: payload }));
    }
  };
}();
const app = createApp(_sfc_main$p);
let locale = void 0;
let food = void 0;
let recipes = void 0;
let diary = void 0;
let prefs = void 0;
function initLocale(messages) {
  locale = createI18n({
    legacy: false,
    locale: lang.split("-")[0],
    fallbackLocale: "en",
    messages
  });
  if (food && recipes && diary && prefs) {
    mountApp();
  }
}
function initFoods(data) {
  food = data;
  if (locale && recipes && diary && prefs) {
    mountApp();
  }
}
function initRecipes(data) {
  recipes = data;
  if (locale && food && diary && prefs) {
    mountApp();
  }
}
function initDiary(data) {
  diary = {};
  data.forEach((d) => diary[d.date] = d);
  if (locale && food && recipes && prefs) {
    mountApp();
  }
}
function initPrefs(data) {
  prefs = data;
  if (locale && food && recipes && diary) {
    mountApp();
  }
}
function mountApp() {
  food.forEach((f) => {
    f.name = locale.global.t(f.id.toString());
  });
  app.provide("csrfToken", csrfToken);
  app.provide("perms", perms);
  app.provide("food", ref(food));
  app.provide("recipes", ref(recipes));
  app.provide("diary", ref(diary));
  app.provide("prefs", ref(prefs));
  app.provide("log", log);
  app.use(router);
  app.use(locale);
  app.mount("#app");
}
fetch("/app/l10n.json").then((response) => response.json()).then(initLocale);
fetch("/api/v1/foods").then((response) => response.json()).then(initFoods);
fetch("/api/v1/recipes").then((response) => response.json()).then(initRecipes);
fetch("/api/v1/diary").then((response) => response.json()).then(initDiary);
fetch("/api/v1/prefs").then((response) => response.json()).then(initPrefs).catch(() => {
  initPrefs({
    rdi: {
      kcal: 2e3,
      fat: 60,
      carb: 270,
      prot: 80
    },
    ui: {
      neutralCharts: false
    }
  });
});
