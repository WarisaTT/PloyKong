import Swal from "sweetalert2";
import "sweetalert2/dist/sweetalert2.min.css";

const Toast = Swal.mixin({
  toast: true,
  position: "top-end",
  showConfirmButton: false,
  timer: 4000,
});

export const showError = (msg: string) => {
  try {
    return Swal.fire({ icon: "error", title: "ข้อผิดพลาด", text: msg });
  } catch (e) {
    // eslint-disable-next-line no-console
    console.error("showError failed", e);
  }
};

export const showSuccess = (msg: string) => {
  try {
    return Swal.fire({ icon: "success", title: "สำเร็จ", text: msg });
  } catch (e) {
    // eslint-disable-next-line no-console
    console.error("showSuccess failed", e);
  }
};

export const toastError = (msg: string) => {
  try {
    return Toast.fire({ icon: "error", title: msg });
  } catch (e) {
    // eslint-disable-next-line no-console
    console.error("toastError failed", e);
  }
};

export const toastSuccess = (msg: string) => {
  try {
    return Toast.fire({ icon: "success", title: msg });
  } catch (e) {
    // eslint-disable-next-line no-console
    console.error("toastSuccess failed", e);
  }
};

export default Swal;
