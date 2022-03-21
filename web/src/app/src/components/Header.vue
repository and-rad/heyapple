<script setup>
import { RouterLink } from "vue-router";
import { inject } from "vue";
import { useI18n } from "vue-i18n";
import HeaderImage from "./images/ImageHeader.vue";
import ProfileImage from "./images/ImageProfile.vue";
import BackArrow from "./images/ImageRightArrow.vue";

const { t } = useI18n();
const csrf = inject("csrfToken");

function confirm(evt) {
	evt.preventDefault();
	fetch("/auth/local", {
		method: "DELETE",
		headers: { "X-CSRF-Token": csrf },
	}).then((response) => {
		if (response.ok) {
			window.location = "/";
		} else {
			window.dispatchEvent(
				new CustomEvent("error", {
					detail: { msg: t("signout.err" + response.status) },
				})
			);
		}
	});
}

function toggleMenu(evt) {
	evt.stopPropagation();
	document.querySelector("header nav").classList.toggle("open");
}

document.addEventListener("click", function () {
	document.querySelector("header nav").classList.remove("open");
});
</script>

<template>
	<header>
		<HeaderImage id="logo" />
		<div id="app-name"><span>Hey</span><span>Apple<sup>beta</sup></span></div>
		<nav>
			<button @click="toggleMenu">
				<BackArrow />
			</button>
			<ul id="nav-main">
				<li>
					<RouterLink to="/">{{ t("nav.diary") }}</RouterLink>
				</li>
				<li>
					<RouterLink to="/food">{{ t("nav.food") }}</RouterLink>
				</li>
				<li>
					<RouterLink to="/recipes">{{ t("nav.recipes") }}</RouterLink>
				</li>
				<li>
					<RouterLink to="/shopping">{{ t("nav.shopping") }}</RouterLink>
				</li>
			</ul>
			<ul id="nav-user">
				<li>
					<RouterLink to="/profile">{{ t("nav.profile") }}</RouterLink>
				</li>
				<li>
					<RouterLink to="/settings">{{ t("nav.settings") }}</RouterLink>
				</li>
				<li>
					<a href="https://docs.heyapple.org" target="_blank">{{ t("nav.help") }}</a>
				</li>
				<li>
					<a href="#" @click="confirm">{{ t("nav.signout") }}</a>
				</li>
			</ul>
		</nav>
		<button @click="toggleMenu">
			<ProfileImage />
		</button>
	</header>
</template>

<style>
header {
	position: fixed;
	width: 100%;
	padding: 0.5em;
	border-bottom: var(--border-light);
	background-color: #fff;
	z-index: 100;
	display: flex;
	align-items: center;
}

header #logo {
	height: 3em;
	width: 3em;
}

header #app-name {
	margin-left: 0.5em;
	font-size: 2em;
	text-transform: uppercase;
	flex-grow: 1;
	color: var(--color-primary);
}

header #app-name span:last-child {
	font-weight: 300;
}

header #app-name sup {
	text-transform: lowercase;
	font-size: 0.5em;
	margin-left: 0.25em;
}

header nav {
	position: fixed;
	top: 0;
	right: -340px;
	bottom: 0;
	background-color: #ffffff;
	box-shadow: var(--shadow-menu);
	z-index: 110;
	padding-top: 4em;
	transition: right 0.25s ease-in-out;
}

header nav.open {
	right: 0;
}

header nav > button {
	display: none;
	position: absolute;
	top: 0.5em;
	right: 0.5em;
	padding: 0.75em;
}

header nav > button svg {
	width: 1.5em;
	height: 1.5em;
}

header nav.open > button {
	display: block;
}

header nav ul {
	list-style: none;
	white-space: nowrap;
	padding: 0;
	width: 320px;
}

header #nav-user li:last-child,
header #nav-user {
	border-top: var(--border-light);
}

header nav a {
	display: block;
	padding: 0.5em 1em;
	color: var(--color-text);
	border-left: 3px solid transparent;
	transition: color 0.2s;
}

header nav a:hover,
header nav a:active,
header nav a:focus {
	color: var(--color-primary);
}

header nav a.router-link-active {
	border-color: var(--color-secondary);
}

header button {
	width: 3em;
	height: 3em;
	padding: 0.5em;
	background: none;
}

header button svg {
	width: 2em;
	height: 2em;
	fill: var(--color-primary);
	transition: 0.2s;
}

header button:hover {
	box-shadow: none;
}

header button:hover svg {
	fill: var(--color-primary-dark);
}

@media only screen and (min-width: 800px) {
	header nav {
		padding: 0;
		position: static;
		box-shadow: none;
	}

	header nav > button {
		display: none !important;
	}

	header #nav-main {
		font-size: 1.2em;
		display: flex;
		align-items: center;
		width: auto;
	}

	header #nav-main a {
		color: var(--color-primary);
		padding: 0.25em 1em;
		border: none;
	}

	header #nav-main a:hover,
	header #nav-main a:active,
	header #nav-main a:focus {
		color: var(--color-primary-dark);
	}

	header #nav-main a.router-link-active {
		color: var(--color-secondary);
	}

	header #nav-user {
		position: fixed;
		top: 4.5em;
		right: -340px;
		background-color: #ffffff;
		box-shadow: var(--shadow-menu);
		z-index: 90;
		margin: 0;
		padding: 0.5em 0;
		opacity: 0;
		transition: opacity 0.25s ease-in-out;
	}

	header nav.open #nav-user {
		right: 0.5em;
		opacity: 1;
	}
}
</style>
