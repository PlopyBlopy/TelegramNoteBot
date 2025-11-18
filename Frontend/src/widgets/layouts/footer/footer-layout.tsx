import styles from "./footer-layout.module.css";

export const FooterLayout = () => {
  return (
    <div className={styles.items}>
      <p className={styles.inDevText}>In development. </p>
      <p className={styles.descriptionText}>Сайт заметок, в качестве практики Backend, Frontend, DevOps.</p>
    </div>
  );
};
