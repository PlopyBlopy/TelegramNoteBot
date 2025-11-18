import styles from "./main-layout.module.css";
import { Outlet } from "react-router-dom";
export const MainLayout = () => {
  return (
    <div className={styles.container}>
      <div className={styles.items}>
        <Outlet />
      </div>
    </div>
  );
};
