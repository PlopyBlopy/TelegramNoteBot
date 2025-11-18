import { HeaderLayout } from "@/widgets/layouts/header";
import { MainLayout } from "@/widgets/layouts/main";
import { FooterLayout } from "@/widgets/layouts/footer";
import styles from "./page-wrapper.module.css";

export const PageWrapper = () => {
  return (
    <>
      <div className={styles.base}>
        <HeaderLayout />
        <MainLayout />
        <FooterLayout />
      </div>
    </>
  );
};
