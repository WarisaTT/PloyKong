import { createApp } from "vue";
import { createPinia } from "pinia";
import App from "./App.vue";
import router from "./router";
import "./assets/main.css";
import { intersectDirective } from "./directives/intersect";
import { showError } from "@/utils/alert";

const app = createApp(App);

app.directive("intersect", intersectDirective);

app.use(createPinia());
app.use(router);

// Global Vue error handler — show alert for uncaught component errors
app.config.errorHandler = (err: unknown, vm, info) => {
  try {
    let message = "Unknown error";
    if (err instanceof Error) {
      message = err.message;
    } else {
      message = String(err);
    }
    showError("เกิดข้อผิดพลาด โปรดลองอีกครั้ง");
    console.error("Error details:", message);
  } catch (e) {
    console.error("Failed to show global error alert", e);
  }
  console.error(err, info);
};

// Global window error and unhandled rejection handlers
window.addEventListener("error", (ev) => {
  try {
    const msg =
      (ev && (ev as any).message) ||
      (ev && (ev as any).error && (ev as any).error.message) ||
      "Unknown error";
    showError("เกิดข้อผิดพลาด โปรดลองรีเฟรชหน้าหรือติดต่อผู้ดูแล");
    // eslint-disable-next-line no-console
    console.error("Window error details:", msg);
  } catch (e) {
    // ignore
  }
});

window.addEventListener("unhandledrejection", (ev) => {
  try {
    const reason = (ev && (ev as any).reason) || "Unhandled rejection";
    const msg = reason && reason.message ? reason.message : String(reason);
    showError("เกิดข้อผิดพลาดในการประมวลผลคำขอ โปรดลองอีกครั้ง");
    // eslint-disable-next-line no-console
    console.error("Unhandled rejection details:", msg);
  } catch (e) {
    // ignore
  }
});

app.mount("#app");
