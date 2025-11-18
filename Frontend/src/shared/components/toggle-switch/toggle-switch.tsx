import React from "react";
import styles from "./toggle-switch.module.css";

type Props = {
  isDark: boolean;
  onChange: (isDark: boolean) => void;
  IconComponent?: React.ComponentType<{ className?: string }>;
};

export const ToggleSwitch = ({ isDark, onChange, IconComponent }: Props) => {
  const handleToggle = () => {
    onChange(!isDark);
  };

  return (
    <button
      className={`${styles.toggle} ${isDark ? styles.dark : styles.light}`}
      onClick={handleToggle}
      aria-label={`Switch to ${isDark ? "light" : "dark"} theme`}
      type="button"
    >
      <div className={styles.track}>
        <div className={styles.thumb}>{IconComponent && <IconComponent className={styles.icon} />}</div>
      </div>
    </button>
  );
};
