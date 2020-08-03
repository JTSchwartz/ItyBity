<template>
	<div class="hello">
		<br/>
		<h1 class="text-h2">Welcome to ItyBity.xyz</h1>
		<p class="text-h6">
			ItyBity is a free link shortner, allowing you to choose your own address, as well as update and delete them.
			<br/>
			When creating a shortened link, ItyBity will generate a secret for you, save this in order to make changes later, secrets cannot be retrieved again.
		</p>
	</div>
</template>

<script>
	export default {
		name: 'HelloWorld'
	}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
	h1 {
		color: var(--v-secondary-lighten1);
		margin: 20px 0 0;
	}

	p {
		color: var(--v-secondary-lighten4);
		margin: 15px 20px;
	}
</style>
