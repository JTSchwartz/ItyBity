import Vue from 'vue';
import Vuetify from 'vuetify';
import 'vuetify/dist/vuetify.min.css';

Vue.use(Vuetify);

export default new Vuetify({
  theme: {
      options: {
        customProperties: true,
      },
    themes: {
      light: {
        primary: '#131313',
        secondary: '#003136',
        accent: '#006663',
        error: '#7d0000',
        info: '#2196F3',
        success: '#005804',
        warning: '#FFC107'
      },
    },
  },
  icons: {
    iconfont: 'fa',
    defaultAssets: false,
    values: {
      checkboxOn: "fa-regular fa-square-check",
      checkboxOff: "fa-regular fa-square",
      checkboxIndeterminate: "fa-regular fa-square-minus",
    }
  },
});
