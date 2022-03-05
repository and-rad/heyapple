<script setup>
import PasswordField from "./Password.vue"
import { RouterLink } from "vue-router";
import { ref, inject } from "vue";
import { useI18n } from "vue-i18n";

const { t } = useI18n();
const reEmail = /^[^@]+@[^@]+$/;
const csrf = inject("csrfToken");

const email = ref("");
const pass1 = ref("");
const pass2 = ref("");

const msg = ref({ msg: "", level: "" });

function confirm(evt) {
	evt.preventDefault();

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
	evt.target.disabled = true;
	fetch("/api/v1/user", {
		method: "POST",
		headers: {
			"Content-Type": "application/x-www-form-urlencoded",
			"X-CSRF-Token": csrf,
		},
		body: new URLSearchParams(new FormData(evt.target.closest("form"))),
	}).then((response) => {
		if (response.ok) {
			msg.value = { msg: t("register.success", { addr: addr })};
			clearForm(evt);
		} else {
			msg.value = { msg: t("register.err" + response.status), level: "err" };
			evt.target.disabled = false;
		}
	});
}

function clearForm(evt) {
	evt.target.disabled = false;
	email.value = "";
	pass1.value = "";
	pass2.value = "";
}
</script>

<template>
	<header>
		<HeaderImage id="logo" />
		<div id="app-name">
			<span>{{ $t("register.title1") }}</span><span>{{ $t("register.title2") }}</span>
		</div>
	</header>
	<form>
		<p v-html="$t('register.hint')"></p>
		<Message v-bind="msg" />
		<label>{{ $t("form.email") }}</label>
		<input type="email" name="email" v-model="email" />
		<label>{{ $t("form.pass") }}</label>
		<PasswordField name="pass" withBar=true v-model="pass1" />
		<label>{{ $t("form.confirm") }}</label>
		<PasswordField name="pass" v-model="pass2" />
		<button type="submit" @click="confirm" class="async">{{ $t("register.action") }}</button>
		<p>{{ $t("register.signin") }} <RouterLink to="/">{{ $t("login.action") }}</RouterLink>.</p>
	</form>
</template>

<style>
</style>
