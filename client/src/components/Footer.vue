<template>
	<v-footer fixed tile padless style="background: transparent">
		<v-dialog
				v-model="dialog"
				width="500"
		>
			<template v-slot:activator="{ on, attrs }">
				<v-btn
						icon
						v-bind="attrs"
						v-on="on"
				>
					<v-icon color="white">
						fad fa-info-circle
					</v-icon>
				</v-btn>
			</template>
			
			<v-card>
				<v-card-title
						class="headline grey lighten-2"
						primary-title
				>
					ItyBity's Policy
				</v-card-title>
				
				<v-card-text>
					<br/>
					ItyBity reserves all rights to remove or edit any URLs that it deems to be malicious or used in an attempt to take advantage of the service. ItyBity is a free service but reserves the right to change this policy at any time.
					<br/><br/>
					Any questions should be directed to <a href="mailto:support@itybity.xyz">Support@ItyBity.xyz</a>
				</v-card-text>
				
				<v-divider></v-divider>
				
				<v-card-actions>
					<v-spacer></v-spacer>
					<v-btn
							color="primary"
							text
							@click="dialog = false"
					>
						Close
					</v-btn>
				</v-card-actions>
			</v-card>
		</v-dialog>
	</v-footer>
</template>

<script>
	export default {
		name: "Footer",
		data: () => ({
			dialog: false
		})
	}
</script>

<style scoped>

</style>
