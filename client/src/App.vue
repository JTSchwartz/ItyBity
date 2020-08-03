<template>
	<v-app>
		<Logo id="logo"/>
		<Intro/>
		<v-container fluid>
			<v-row align="start" justify="space-around">
				<Cluster
						title="Create"
						desc="Create your short link"
						sm-order="1"
						lg-order="1"
				>
					<v-text-field
							label="Destination*"
							placeholder=" "
							:rules="destRules"
							v-model="create.dest"
					/>
					<v-text-field
							label="Link"
							prefix="itybity.xyz/"
							placeholder=" "
							:rules="slugRules"
							v-model="create.slug"
					/>
					<v-btn @click="createLink">
						Create
					</v-btn>
					<v-text-field
							label="Generated Secret"
							placeholder=" "
							readonly
							v-model="create.secret"
					/>
				</Cluster>
				<Cluster
						title="Delete"
						desc="Free up your short link for someone else to use"
						sm-order="3"
						lg-order="2"
				>
					<v-text-field
							label="Link*"
							prefix="itybity.xyz/"
							placeholder=" "
							:rules="slugRules"
							v-model="remove.slug"
					/>
					<v-text-field
							label="Secret*"
							placeholder=" "
							v-model="remove.secret"
					/>
					<v-btn @click="removeLink">
						Delete
					</v-btn>
				</Cluster>
				<Cluster
						title="Change"
						desc="Change your short link's destination"
						sm-order="2"
						lg-order="3"
				>
					<v-text-field
							label="Link*"
							prefix="itybity.xyz/"
							placeholder=" "
							:rules="slugRules"
							v-model="change.slug"
					/>
					<v-text-field
							label="Secret*"
							placeholder=" "
							v-model="change.secret"
					/>
					<v-text-field
							label="Destination*"
							placeholder=" "
							:rules="destRules"
							v-model="change.dest"
					/>
					<v-btn @click="changeLink">
						Change
					</v-btn>
				</Cluster>
			</v-row>
		</v-container>
		<Footer/>
		<v-snackbar
				v-model="infoBar"
				:bottom="true"
				:color="infoColor"
				:timeout="6000"
		>
			{{ infoMsg }}
			
			<template v-slot:action="{ attrs }">
				<v-btn
						icon
						color="black"
						v-bind="attrs"
						@click="infoBar = false"
				>
					<v-icon>
						fad fa-times-circle
					</v-icon>
				</v-btn>
			</template>
		</v-snackbar>
	</v-app>
</template>

<script>
	import Intro   from "./components/Intro.vue"
	import Logo    from "./components/Logo";
	import Cluster from "./components/Cluster";
	import Footer  from "./components/Footer";
	
	export default {
		name:       "App",
		components: {
			Footer,
			Cluster,
			Logo,
			Intro
		},
		data:       () => ({
			create:    {
				dest:   "",
				secret: "",
				slug:   ""
			},
			change:    {
				dest:   "",
				secret: "",
				slug:   ""
			},
			destRegex: new RegExp(
				"^" +
				// protocol identifier (optional)
				// short syntax // still required
				"(?:(?:(?:https?|ftp):)?\\/\\/)" +
				// user:pass BasicAuth (optional)
				"(?:\\S+(?::\\S*)?@)?" +
				"(?:" +
				// IP address exclusion
				// private & local networks
				"(?!(?:10|127)(?:\\.\\d{1,3}){3})" +
				"(?!(?:169\\.254|192\\.168)(?:\\.\\d{1,3}){2})" +
				"(?!172\\.(?:1[6-9]|2\\d|3[0-1])(?:\\.\\d{1,3}){2})" +
				// IP address dotted notation octets
				// excludes loopback network 0.0.0.0
				// excludes reserved space >= 224.0.0.0
				// excludes network & broadcast addresses
				// (first & last IP address of each class)
				"(?:[1-9]\\d?|1\\d\\d|2[01]\\d|22[0-3])" +
				"(?:\\.(?:1?\\d{1,2}|2[0-4]\\d|25[0-5])){2}" +
				"(?:\\.(?:[1-9]\\d?|1\\d\\d|2[0-4]\\d|25[0-4]))" +
				"|" +
				// host & domain names, may end with dot
				// can be replaced by a shortest alternative
				// (?![-_])(?:[-\\w\\u00a1-\\uffff]{0,63}[^-_]\\.)+
				"(?:" +
				"(?:" +
				"[a-z0-9\\u00a1-\\uffff]" +
				"[a-z0-9\\u00a1-\\uffff_-]{0,62}" +
				")?" +
				"[a-z0-9\\u00a1-\\uffff]\\." +
				")+" +
				// TLD identifier name, may end with dot
				"(?:[a-z\\u00a1-\\uffff]{2,}\\.?)" +
				")" +
				// port number (optional)
				"(?::\\d{2,5})?" +
				// resource path (optional)
				"(?:[/?#]\\S*)?" +
				"$", "i"
			),
			destRules: [
				dest => RegExp(
					"^" +
					// protocol identifier (optional)
					// short syntax // still required
					"(?:(?:(?:https?|ftp):)?\\/\\/)" +
					// user:pass BasicAuth (optional)
					"(?:\\S+(?::\\S*)?@)?" +
					"(?:" +
					// IP address exclusion
					// private & local networks
					"(?!(?:10|127)(?:\\.\\d{1,3}){3})" +
					"(?!(?:169\\.254|192\\.168)(?:\\.\\d{1,3}){2})" +
					"(?!172\\.(?:1[6-9]|2\\d|3[0-1])(?:\\.\\d{1,3}){2})" +
					// IP address dotted notation octets
					// excludes loopback network 0.0.0.0
					// excludes reserved space >= 224.0.0.0
					// excludes network & broadcast addresses
					// (first & last IP address of each class)
					"(?:[1-9]\\d?|1\\d\\d|2[01]\\d|22[0-3])" +
					"(?:\\.(?:1?\\d{1,2}|2[0-4]\\d|25[0-5])){2}" +
					"(?:\\.(?:[1-9]\\d?|1\\d\\d|2[0-4]\\d|25[0-4]))" +
					"|" +
					// host & domain names, may end with dot
					// can be replaced by a shortest alternative
					// (?![-_])(?:[-\\w\\u00a1-\\uffff]{0,63}[^-_]\\.)+
					"(?:" +
					"(?:" +
					"[a-z0-9\\u00a1-\\uffff]" +
					"[a-z0-9\\u00a1-\\uffff_-]{0,62}" +
					")?" +
					"[a-z0-9\\u00a1-\\uffff]\\." +
					")+" +
					// TLD identifier name, may end with dot
					"(?:[a-z\\u00a1-\\uffff]{2,}\\.?)" +
					")" +
					// port number (optional)
					"(?::\\d{2,5})?" +
					// resource path (optional)
					"(?:[/?#]\\S*)?" +
					"$", "i"
				).test(dest)
			],
			infoBar:   false,
			infoColor: "error",
			infoMsg:   "Error",
			remove:    {
				secret: "",
				slug:   ""
			},
			slugRegex: /^([a-z0-9][a-z0-9-]*)?[a-z0-9]+$/,
			slugRules: [
				slug => /^([a-z0-9][a-z0-9-]*)?[a-z0-9]+$/.test(slug)
			],
		}),
		computed:   {
			host: function () {
				return `${(process.env.NODE_ENV === "development") ? "http://127.0.0.1" : "https://itybity.xyz"}:7070`
			}
		},
		methods:    {
			apiRequest:            async function (request, data) {
				return await fetch(`${this.host}/${request}`, {
					"method":  "POST",
					"body":    JSON.stringify(data),
					"headers": {
						"Content-type": "application/json; charset=UTF-8"
					}
				})
			},
			changeLink:            function () {
				if (this.failsDestVerification(this.change.dest)) {
					this.setInfo("Destination is not valid", "error")
					return
				} else if (this.failsSlugVerification(this.change.slug)) {
					this.setInfo("Slug is not valid", "error")
					return
				}
				
				this.apiRequest("change", {
						"full":   this.change.dest,
						"secret": this.change.secret,
						"slug":   this.change.slug
					})
					.then(resp => resp.json())
					.then(data => {
						if (data.err) {
							this.setInfo(data.err, "error")
						} else {
							this.setInfo("Success", "success")
						}
					})
			},
			createLink:            function () {
				if (this.failsDestVerification(this.create.dest)) {
					this.setInfo("Destination is not valid", "error")
					return
				} else if (this.failsSlugVerification(this.create.slug)) {
					this.setInfo("Slug is not valid", "error")
					return
				}
				
				this.apiRequest("create", {
						"full": this.create.dest,
						"slug": this.create.slug
					})
					.then(resp => resp.json())
					.then(data => {
						if (data.secret) {
							this.create.secret = data.secret
							this.setInfo("Success", "success")
						} else {
							this.setInfo(data.err, "error")
						}
					})
			},
			removeLink:            function () {
				if (this.failsSlugVerification(this.remove.slug)) {
					this.setInfo("Slug is not valid", "error")
					return
				}
				
				this.apiRequest("remove", {
						"slug":   this.remove.slug,
						"secret": this.remove.secret
					})
					.then(resp => resp.json())
					.then(data => {
						if (data.err) {
							this.setInfo(data.err, "error")
						} else {
							this.setInfo("Success", "success")
						}
					})
			},
			failsDestVerification: function (testString) {
				return !this.destRegex.test(testString)
			},
			failsSlugVerification: function (testString) {
				return !this.slugRegex.test(testString)
			},
			setInfo:               function (msg, color) {
				this.infoMsg = msg
				this.infoColor = color
				this.infoBar = true
			}
		}
	}
</script>

<style>
	html::-webkit-scrollbar {
		display: none;
	}
	
	html {
		-ms-overflow-style: none;
		scrollbar-width: none;
	}
	
	body, #app {
		background: var(--v-primary-base);
	}
	
	#app {
		font-family: Avenir, Helvetica, Arial, sans-serif;
		-webkit-font-smoothing: antialiased;
		-moz-osx-font-smoothing: grayscale;
		text-align: center;
		color: var(--v-secondary-base);
	}
	
	#logo {
		padding-top: 60px;
		color: var(--v-accent-base);
		margin: 0 auto;
	}
</style>
