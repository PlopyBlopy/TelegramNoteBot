import { ThemeToggle } from "@/widgets/theme-toggle";
import styles from "./header-layout.module.css";
import { Icon } from "@/shared/components/icon";
import { Icons } from "@/shared/assets/icons";

export const HeaderLayout = () => {
  return (
    <div className={styles.container}>
      <div className={styles.items}>
        <Icon IconComponent={Icons.elements.note} />
        <ThemeToggle />
      </div>
    </div>
  );
};
