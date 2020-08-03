<template>
	<v-col cols="10" md="3" :order="parseInt(smOrder)" :order-md="parseInt(lgOrder)">
		<v-card color="primary lighten-1" dark>
			<v-card-title><strong>{{ title }}</strong></v-card-title>
			<v-card-subtitle><strong>{{ desc }}</strong></v-card-subtitle>
			<v-card-text>
				<v-form>
					<slot>

					</slot>
				</v-form>
			</v-card-text>
		</v-card>
	</v-col>
</template>

<script>
	export default {
		name: "Cluster",
		props: {
			title: String,
			desc: String,
			smOrder: String,
			lgOrder: String
		}
	}
</script>

<style scoped>
	* {
		color: var(--v-secondary-lighten4);
	}
</style>