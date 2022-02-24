<script setup>
import RegisterImage from "./images/ImageRegister.vue";
import { RouterLink } from "vue-router";
</script>

<template>
	<div>
		<form>
			<h1>Sign Up</h1>
			<label >E-Mail</label>
			<input type="email" nam="email" />
			<label >Password</label>
			<input type="password" name="pass" />
			<label >Confirm Password</label>
			<input type="password" name="pass" />
			<footer>
				<span>Already have an account? <RouterLink to="/">Sign in</RouterLink>.</span>
			</footer>
			<input type="submit" value="Sign up" />
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