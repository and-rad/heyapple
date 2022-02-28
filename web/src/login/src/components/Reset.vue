<script setup>
import { RouterLink } from "vue-router";
import { ref, inject } from "vue";
import { useI18n } from "vue-i18n";

const { t } = useI18n();
const reEmail = /^[^@]+@[^@]+$/;
const csrf = inject("csrfToken");

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
		headers: { 
			"Content-Type": "application/x-www-form-urlencoded",
			"X-CSRF-Token": csrf,
		},
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
	<RouterLink to="/" class="back">&#5176; {{ $t("form.back") }}</RouterLink>
	<header>
		<HeaderImage id="logo" />
		<div id="app-name">
			<span>{{ $t("reset.title1") }}</span><span>{{ $t("reset.title2") }}</span>
		</div>
	</header>
	<form>
		<Message v-bind="msg" />
		<p>{{ $t("reset.hint") }}</p>
		<label>{{ $t("form.email") }}</label>
		<input type="email" name="email" v-model="email" />
		<input type="submit" :value="$t('reset.action')" @click="confirm" />
	</form>
</template>

<style>
a.back {
	display: block;
	margin-bottom: 2em;
	transition: color 0.2s;
}

a.back:hover {
	color: var(--color-primary-dark);
}
</style>
