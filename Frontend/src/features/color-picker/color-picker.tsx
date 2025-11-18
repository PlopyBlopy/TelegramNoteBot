import styles from "./color-picker.module.css";
import type { ColorInfo } from "@/shared/api";

type Props = {
  options: ColorInfo[];
  value: number;
  onColorSelectId: (colorSelectId: number) => void;
  placeholder: string;
};

export const ColorPicker = ({ options, value, onColorSelectId, placeholder }: Props) => {
  const currentSelected = options.find((color) => color.id === value);

  return (
    <div className={styles.container}>
      {currentSelected && (
        <div className={styles.selectedColor}>
          <div className={styles.selectedColorPreview} style={{ backgroundColor: currentSelected.variable }} />
          <span>
            {placeholder}: {currentSelected.name}
          </span>
        </div>
      )}
      <div className={styles.colorGrid}>
        {options.map((color) => (
          <div
            key={color.variable}
            className={`${styles.colorItem} ${value === color.id ? styles.selected : ""}`}
            style={{ backgroundColor: color.variable }}
            onClick={() => onColorSelectId(color.id)}
          >
            <div className={styles.tooltip}>{color.name}</div>
          </div>
        ))}
      </div>
    </div>
  );
};
