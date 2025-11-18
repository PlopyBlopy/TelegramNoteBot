import { useState, useEffect, useRef } from "react";
import styles from "./search-bar.module.css";
import { useDebounce } from "@/shared/hook/debounce";

interface SearchBarProps {
  value?: string;
  onSearch: (value: string) => void;
  delay?: number; // Задержка для debounce
  placeholder?: string;
}

export const SearchBar = ({ value = "", delay = 500, onSearch, placeholder = "Поиск..." }: SearchBarProps) => {
  const [inputValue, setInputValue] = useState<string>(value);
  const debouncedSearchValue = useDebounce(inputValue, delay);
  const onSearchRef = useRef(onSearch);

  useEffect(() => {
    if (debouncedSearchValue) {
      onSearchRef.current = onSearch;
    }
  }, [onSearch]);

  useEffect(() => {
    if (debouncedSearchValue) {
      onSearchRef.current(debouncedSearchValue);
    }
  }, [debouncedSearchValue]);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setInputValue(e.target.value);
  };

  return (
    <div className={styles.container}>
      <div className={styles.searchBar}>
        <div className={styles.searchIcon}>
          <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2">
            <circle cx="11" cy="11" r="8" />
            <path d="m21 21-4.3-4.3" />
          </svg>
        </div>

        <input
          type="text"
          value={inputValue}
          onChange={handleChange}
          placeholder={placeholder}
          disabled={false}
          autoFocus={false}
          className={styles.input}
        />
      </div>
    </div>
  );
};
