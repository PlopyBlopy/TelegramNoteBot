import { useState } from "react";
import styles from "./Dropdown.module.css";

interface DropdownProps {
  options: string[];
  value: string | null;
  onChange: (value: string | null) => void;
  placeholder?: string;
}

export const Dropdown = ({ options, value, onChange, placeholder = "Выберите опцию" }: DropdownProps) => {
  const [isOpen, setIsOpen] = useState<boolean>(false);

  const handleSelect = (option: string) => {
    onChange(option);
    setIsOpen(false);
  };

  const handleClear = (e: React.MouseEvent) => {
    e.stopPropagation();
    onChange(null);
    setIsOpen(false);
  };

  return (
    <div className={styles.container}>
      <div className={styles.header} onClick={() => setIsOpen(!isOpen)}>
        <span className={value ? styles.selected : styles.placeholder}>{value || placeholder}</span>
        <div className={styles.controls}>
          {value && (
            <span className={styles.clear} onClick={handleClear}>
              ×
            </span>
          )}
          <span className={styles.arrow}>{isOpen ? "▲" : "▼"}</span>
        </div>
      </div>

      {isOpen && (
        <div className={styles.list}>
          {options.map((option, index) => (
            <div key={index} className={`${styles.item} ${value === option ? styles.itemSelected : ""}`} onClick={() => handleSelect(option)}>
              {option}
            </div>
          ))}
        </div>
      )}
    </div>
  );
};
