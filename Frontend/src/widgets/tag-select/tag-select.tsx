import type { TagInfo } from "@/shared/api";
import styles from "./tag-select.module.css";
import { Tag } from "@/features/tag";

interface TagsSelectProps {
  options: TagInfo[];
  value: number[];
  onChange: (value: number[]) => void;
  placeholder?: string;
}

export const TagsSelect = ({ options, value = [], onChange, placeholder = "Теги еще не созданы" }: TagsSelectProps) => {
  const handleToggle = (option: number) => {
    if (value.includes(option)) {
      // Если тег уже выбран - удаляем
      onChange(value.filter((item) => item !== option));
    } else {
      // Если тег не выбран - добавляем
      onChange([...value, option]);
    }
  };

  return (
    <div className={styles.container}>
      <div className={styles.tagsRow}>
        {options
          .filter((option) => value.includes(option.colorId))
          .map((option, index) => (
            <div key={`selected-${index}`} className={styles.selectedTag} onClick={() => handleToggle(option.colorId)} title="Кликните чтобы удалить">
              <Tag tag={option} />
            </div>
          ))}
        {/* Затем доступные теги */}
        {options
          .filter((option) => !value.includes(option.colorId))
          .map((option, index) => (
            <div
              key={`available-${index}`}
              className={styles.availableTag}
              onClick={() => handleToggle(option.colorId)}
              title="Кликните чтобы добавить"
            >
              <Tag tag={option} />
            </div>
          ))}

        {/* Плейсхолдер, когда ничего не выбрано и нет тегов */}
        {value.length === 0 && options.length === 0 && <span className={styles.placeholder}>{placeholder}</span>}
      </div>
    </div>
  );
};
