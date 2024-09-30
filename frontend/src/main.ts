import { createApp } from "vue";
import { createPinia } from "pinia";

import App from "./App.vue";
import router from "./router";
import i18n from "./i18n";
import piniaPluginPersistedstate from "pinia-plugin-persistedstate";
import "./style.scss";

const app = createApp(App);
const pinia = createPinia();
app.use(pinia);
pinia.use(piniaPluginPersistedstate);
app.use(router);
app.use(i18n);

app.mount("#app");
