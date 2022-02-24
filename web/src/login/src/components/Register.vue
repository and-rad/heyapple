<script setup>
import RegisterImage from "./images/ImageRegister.vue";
import { RouterLink } from "vue-router";
import { ref } from "vue";
import { useI18n } from "vue-i18n";

const { t } = useI18n();
const reEmail = /^[^@]+@[^@]+$/;

const email = ref("");
const pass1 = ref("");
const pass2 = ref("");

const msg = ref({ msg: "", level: "" });

function confirm(evt) {
	evt.preventDefault();
	msg.value = {};

	if (email.value.search(reEmail) == -1) {
		msg.value = { msg: t("login.errmail"), level: "err" };
		return false;
	}
	if (pass1.value == "" || pass2.value == "") {
		msg.value = { msg: t("register.errpassempty"), level: "err" };
		return false;
	}
	if (pass1.value != pass2.value) {
		msg.value = { msg: t("register.errpassmatch"), level: "err" };
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
				msg: t("register.success",{addr: addr}),
			};
		} else {
			msg.value = { msg: t("register.err"+response.status), level: "err" };
		}
	});
}
</script>

<template>
	<div>
		<form>
			<h1>{{ $t("register.title") }}</h1>
			<Message v-bind="msg" />
			<label>{{ $t("form.email") }}</label>
			<input type="email" name="email" v-model="email" />
			<label>{{ $t("form.pass") }}</label>
			<input type="password" name="pass" v-model="pass1" />
			<label>{{ $t("form.confirm") }}</label>
			<input type="password" name="pass" v-model="pass2" />
			<footer>
				<span>{{ $t("register.signin") }} <RouterLink to="/">{{ $t("login.action") }}</RouterLink>.</span>
			</footer>
			<input type="submit" :value="$t('register.action')" @click="confirm" />
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
