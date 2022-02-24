<script setup>
import LoginImage from "./images/ImageLogin.vue";
import { RouterLink } from "vue-router";
</script>

<template>
	<div>
		<form>
			<h1>Sign In</h1>
			<label for="">E-Mail</label>
			<input type="email" name="" />
			<label for="">Password</label>
			<input type="password" name="" />
			<footer>
				<span>Don't have an account yet? <RouterLink to="/signup">Sign up</RouterLink>.</span>
				<RouterLink to="/reset">Reset password</RouterLink>
			</footer>
			<input type="submit" value="Sign up" />
		</form>
	</div>
	<div class="image">
		<LoginImage />
	</div>
</template>

<style>
svg#sun {
	bottom: unset !important;
	top: 0;
	right: 0;
}
</style>
