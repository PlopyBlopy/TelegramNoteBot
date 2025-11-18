import { Icons } from "@/shared/assets/icons";
import styles from "./primary-button-icon.module.css";

export interface Style {
  color?: string;
}

interface Props extends Style {
  onClick: () => void;
  text?: string;
  IconComponent?: React.ComponentType<{ className?: string }>;
}

export const PrimaryButtonIcon = ({ onClick, text, IconComponent = Icons.default, color }: Props) => {
  const style = {
    color: color,
  };
  return (
    <button style={style} className={styles.button} onClick={onClick}>
      <IconComponent className={styles.icon} />
      {text}
    </button>
  );
};
