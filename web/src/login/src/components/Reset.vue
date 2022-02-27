<script setup>
import LoginImage from "./images/ImageLogin.vue";
import { RouterLink } from "vue-router";
import { ref } from "vue";
import { useI18n } from "vue-i18n";

const { t } = useI18n();
const reEmail = /^[^@]+@[^@]+$/;

const email = ref("");
const msg = ref({ msg: "", level: "" });

function confirm(evt) {
	evt.preventDefault();
	msg.value = {};

	if (email.value.search(reEmail) == -1) {
		msg.value = { msg: t("login.errmail"), level: "err" };
		return false;
	}

	let addr = email.value;
	fetch("/auth/reset", {
		method: "POST",
		headers: { "Content-Type": "application/x-www-form-urlencoded" },
		body: new URLSearchParams(new FormData(evt.target.closest("form"))),
	}).then((response) => {
		if (response.ok) {
			msg.value = {
				msg: t("reset.success", { addr: addr }),
			};
		} else {
			msg.value = { msg: t("reset.err" + response.status), level: "err" };
		}
	});
}
</script>

<template>
	<div>
		<form>
			<h1>{{ $t("reset.title") }}</h1>
			<Message v-bind="msg" />
			<p>
				{{ $t("reset.hint") }}
			</p>
			<label>{{ $t("form.email") }}</label>
			<input type="email" name="email" v-model="email" />
			<input type="submit" :value="$t('reset.action')" @click="confirm" />
		</form>
		<RouterLink to="/" class="back">&#5176; {{ $t("form.back") }}</RouterLink>
	</div>
	<div class="image">
		<LoginImage />
	</div>
</template>

<style>
a.back {
	position: absolute;
	display: block;
	padding: 1em;
	top: 0;
	left: 0;
}
</style>
