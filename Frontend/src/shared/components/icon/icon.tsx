import { Icons } from "@/shared/assets/icons";
import styles from "./icon.module.css";

type Props = {
  IconComponent?: React.ComponentType<{ className?: string }>;
};

export const Icon = ({ IconComponent = Icons.default }: Props) => {
  return <IconComponent className={styles.icon} />;
};
