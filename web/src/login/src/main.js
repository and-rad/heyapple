import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import Message from "./components/Message.vue";

const app = createApp(App);

app.use(router);
app.component("Message", Message);
app.mount("#app");
