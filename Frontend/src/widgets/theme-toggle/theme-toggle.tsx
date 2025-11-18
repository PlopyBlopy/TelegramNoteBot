import { Icons } from "@/shared/assets/icons";
import { ToggleSwitch } from "@/shared/components/toggle-switch";
import { useTheme } from "@/shared/hook/theme";

export const ThemeToggle = () => {
  const { theme, toggleTheme } = useTheme();

  const handleChange = (checked: boolean) => {
    checked ? theme === "light" : theme === "dark";

    toggleTheme();
  };

  return (
    <ToggleSwitch isDark={theme === "dark"} onChange={handleChange} IconComponent={theme === "dark" ? Icons.elements.moon : Icons.elements.sun} />
  );
};
