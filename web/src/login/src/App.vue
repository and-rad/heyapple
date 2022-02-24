<script setup>
import { RouterView } from "vue-router";
</script>

<template>
	<RouterView />
</template>

<style>
@import "@/assets/base.css";

#app {
	width: 960px;
	height: 640px;
	margin: 0 auto;
	transform: translateY(-50%) translateX(-50%);
	background-color: #fff;
	border-radius: 4px;
	box-shadow: 0 0 24px #ccc;
	overflow-y: hidden;
	display: flex;
	justify-content: space-between;
	position: absolute;
	top: 50%;
	left: 50%;
}

#app > *:first-child {
	flex-basis: 60%;
	padding: 3em;
}

#app > *:nth-child(2) {
	flex-basis: 40%;
}

#app > .image svg {
	position: absolute;
	bottom: 0;
}

h1 {
	margin-bottom: 1em;
}

label {
	margin-top: 1em;
	display: block;
}

input[type="text"],
input[type="email"],
input[type="password"] {
	display: block;
	line-height: 150%;
	font-size: 16px;
	padding: 0.5em;
	border: 1px solid #aaa;
	border-radius: 4px;
	width: 100%;
}

input[type="text"]:focus,
input[type="text"]:active,
input[type="email"]:focus,
input[type="email"]:active,
input[type="password"]:focus,
input[type="password"]:active {
	outline: none;
	border-color: var(--color-primary);
}

input[type="submit"] {
	display: block;
	line-height: 150%;
	font-size: 16px;
	padding: 0.5em;
	border: none;
	border-radius: 4px;
	width: 100%;
	background-color: var(--color-primary);
	margin-top: 2em;
	color: #fff;
	transition: box-shadow 0.2s;
}

input[type="submit"]:hover {
	box-shadow: inset 0 0 100px rgba(0,0,0,0.1);
}

form footer {
	margin-top: 0.5em;
	display: flex;
	justify-content: space-between;
}
</style>
