<script setup>
import LoginImage from "./images/ImageLogin.vue";
import { RouterLink } from "vue-router";
</script>

<template>
	<div>
		<form>
			<h1>Password Reset</h1>
			<p>
				Enter your account's e-mail address. Instructions on how to reset your password will be sent to you.
			</p>
			<label>E-Mail</label>
			<input type="email" name="" />
			<input type="submit" value="Send password reset e-mail" />
		</form>
		<RouterLink to="/" class="back">&#5176; Back</RouterLink>
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