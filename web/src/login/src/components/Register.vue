<script setup>
import RegisterImage from "./images/ImageRegister.vue";
import { RouterLink } from "vue-router";
import { ref } from "vue";

const reEmail = /^[^@]+@[^@]+$/;

const email = ref("");
const pass1 = ref("");
const pass2 = ref("");

const msg = ref({ msg: "", level: "" });

function confirm(evt) {
	evt.preventDefault();
	msg.value = {};

	if (email.value.search(reEmail) == -1) {
		msg.value = { msg: "Not a valid e-mail address", level: "err" };
		return false;
	}
	if (pass1.value == "" || pass2.value == "") {
		msg.value = { msg: "Fill out both password fields", level: "err" };
		return false;
	}
	if (pass1.value != pass2.value) {
		msg.value = { msg: "The passwords don't match", level: "err" };
		return;
	}

	let addr = email.value;
	fetch("/api/v1/user", {
		method: "POST",
		headers: { "Content-Type": "application/x-www-form-urlencoded" },
		body: new URLSearchParams(new FormData(evt.target.closest("form"))),
	}).then((response) => {
		if (response.ok) {
			msg.value = {
				msg: `Registration successful! We've sent an email to ${addr}. Open it up to activate your account.`,
			};
		} else {
			msg.value = { msg: "Error: " + response.status, level: "err" };
		}
	});
}
</script>

<template>
	<div>
		<form>
			<h1>Sign Up</h1>
			<Message v-bind="msg" />
			<label>E-Mail</label>
			<input type="email" name="email" v-model="email" />
			<label>Password</label>
			<input type="password" name="pass" v-model="pass1" />
			<label>Confirm Password</label>
			<input type="password" name="pass" v-model="pass2" />
			<footer>
				<span>Already have an account? <RouterLink to="/">Sign in</RouterLink>.</span>
			</footer>
			<input type="submit" value="Sign up" @click="confirm" />
		</form>
	</div>
	<div class="image register-image">
		<RegisterImage />
	</div>
</template>

<style>
#app > .image.register-image {
	background-color: #c9d6df;
}
</style>
