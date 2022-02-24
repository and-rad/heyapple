<script setup>
import LoginImage from "./images/ImageLogin.vue";
import { RouterLink } from "vue-router";
</script>

<template>
	<div>
		<form>
			<h1>{{ $t("reset.title") }}</h1>
			<p>
				{{ $t("reset.hint") }}
			</p>
			<label>{{ $t("form.email") }}</label>
			<input type="email" name="email" />
			<input type="submit" :value="$t('reset.action')" />
		</form>
		<RouterLink to="/" class="back">&#5176; {{ $t("form.back") }}</RouterLink>
	</div>
	<div class="image">
		<LoginImage />
	</div>
</template>

<style>
a.back {
	position:absolute;
	display: block;
	padding: 1em;
	top: 0;
	left: 0;
}
</style>