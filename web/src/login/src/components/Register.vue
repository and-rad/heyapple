<script setup>
import RegisterImage from "./images/ImageRegister.vue";
import { RouterLink } from "vue-router";
</script>

<template>
	<div>
		<form>
			<h1>Sign Up</h1>
			<label for="">E-Mail</label>
			<input type="email" name="" id="" />
			<label for="">Password</label>
			<input type="password" name="" />
			<label for="">Confirm Password</label>
			<input type="password" name="" />
			<footer>
				<span>Already have an account? <RouterLink to="/">Sign in</RouterLink>.</span>
			</footer>
			<input type="submit" value="Sign up" />
		</form>
	</div>
	<div class="image">
		<RegisterImage />
	</div>
</template>
