<script setup>
import LoginImage from "./images/ImageLogin.vue";
import { RouterLink } from "vue-router";
import { ref } from "vue";

const reEmail = /^[^@]+@[^@]+$/;

const email = ref("");
const pass = ref("");

const msg = ref({ msg: "", level: "" });

function confirm(evt) {
	evt.preventDefault();
	msg.value = {};

	if (email.value.search(reEmail) == -1) {
		msg.value = { msg: "Not a valid e-mail address", level: "err" };
		return false;
	}

	fetch("/auth/local", {
		method: "POST",
		headers: { "Content-Type": "application/x-www-form-urlencoded" },
		body: new URLSearchParams(new FormData(evt.target.closest("form"))),
	}).then((response) => {
		if (response.ok) {
			window.location = "/app";
		} else {
			msg.value = { msg: "Error: " + response.status, level: "err" };
		}
	});
}
</script>

<template>
	<div>
		<form>
			<h1>Sign In</h1>
			<Message v-bind="msg" />
			<label>E-Mail</label>
			<input type="email" name="email" v-model="email" />
			<label>Password <!--<RouterLink to="/reset">Reset password</RouterLink>--></label>
			<input type="password" name="pass" v-model="pass" />
			<footer>
				<span
					>Don't have an account yet? <RouterLink to="/signup">Sign up</RouterLink>.</span
				>
			</footer>
			<input type="submit" value="Sign in" @click="confirm" />
		</form>
	</div>
	<div class="image">
		<LoginImage />
	</div>
</template>

<style>
#app > .image {
	background-color: #cad0db;
}

svg#sun {
	bottom: unset !important;
	top: 0;
	right: 0;
}
</style>
