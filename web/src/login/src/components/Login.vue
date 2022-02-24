<script setup>
import LoginImage from "./images/ImageLogin.vue";
import { RouterLink } from "vue-router";
</script>

<template>
	<div>
		<form>
			<h1>Sign In</h1>
			<label >E-Mail</label>
			<input type="email" name="email" />
			<label >Password <RouterLink to="/reset">Reset password</RouterLink></label>
			<input type="password" name="pass" />
			<footer>
				<span>Don't have an account yet? <RouterLink to="/signup">Sign up</RouterLink>.</span>
			</footer>
			<input type="submit" value="Sign up" />
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
