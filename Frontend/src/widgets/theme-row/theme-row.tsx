import styles from "./theme-row.module.css";
import type { Theme } from "@/shared/api";

type Props = {
  options: Theme[];
  value: number;
  onChange: (value: number) => void;
};

export const ThemeRow = ({ options, value, onChange }: Props) => {
  return (
    <div className={styles.container}>
      <div className={styles.themeContainer}>
        {options.map((t, i) => (
          <button key={`theme-${i}`} className={value === t.id ? styles.selectButton : styles.button} onClick={() => onChange(t.id)}>
            {t.title}
          </button>
        ))}
      </div>
      <div className={styles.footer}></div>
    </div>
  );
};
