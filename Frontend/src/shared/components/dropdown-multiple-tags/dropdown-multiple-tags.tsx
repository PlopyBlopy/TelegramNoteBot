import { useState, useRef, useEffect } from "react";
import styles from "./dropdown-multiple-tags.module.css";
import { Tag } from "@/features/tag";
import type { TagInfo } from "@/shared/api";

type Props = {
  options: TagInfo[];
  value: number[];
  onChange: (value: number[]) => void;
  placeholder?: string;
};

export const MultipleTagsDropdown = ({ options, value = [], onChange, placeholder = "Выберите опции" }: Props) => {
  const [isOpen, setIsOpen] = useState<boolean>(false);
  const dropdownRef = useRef<HTMLDivElement>(null);

  // Закрытие dropdown при клике вне компонента
  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (dropdownRef.current && !dropdownRef.current.contains(event.target as Node)) {
        setIsOpen(false);
      }
    };

    document.addEventListener("mousedown", handleClickOutside);
    return () => {
      document.removeEventListener("mousedown", handleClickOutside);
    };
  }, []);

  const handleSelect = (option: number) => {
    if (!value.includes(option)) {
      onChange([...value, option]);
    }
  };

  const handleRemove = (optionToRemove: number, e: React.MouseEvent) => {
    e.stopPropagation();
    onChange(value.filter((option) => option !== optionToRemove));
  };

  const handleClear = (e: React.MouseEvent) => {
    e.stopPropagation();
    onChange([]);
  };

  const handleCloseDropdown = () => {
    setIsOpen(false);
  };

  const availableOptions = options.filter((option) => !value.includes(option.id));

  return (
    <div className={styles.container} ref={dropdownRef}>
      <div className={styles.header} onClick={() => setIsOpen(!isOpen)}>
        <div className={styles.tagsContainer}>
          {value.length > 0 ? (
            options
              .filter((option) => value.includes(option.id))
              .map((option, index) => (
                <div key={index} className={styles.tagWrapper}>
                  <Tag tag={option} />
                  <span className={styles.removeTag} onClick={(e) => handleRemove(option.id, e)}>
                    ×
                  </span>
                </div>
              ))
          ) : (
            <span className={styles.placeholder}>{placeholder}</span>
          )}
        </div>

        <div className={styles.controls}>
          {value.length > 0 && (
            <span className={styles.clear} onClick={handleClear}>
              ×
            </span>
          )}
          <span className={styles.arrow}>{isOpen ? "▲" : "▼"}</span>
        </div>
      </div>

      {isOpen && (
        <div className={styles.list}>
          <div className={styles.tagsGrid}>
            {availableOptions.map((option, index) => (
              <div key={index} className={styles.option} onClick={() => handleSelect(option.id)}>
                <Tag tag={option} />
              </div>
            ))}
          </div>

          <div className={styles.footer}>
            <button className={styles.closeButton} onClick={handleCloseDropdown}>
              Готово
            </button>
          </div>
        </div>
      )}
    </div>
  );
};
