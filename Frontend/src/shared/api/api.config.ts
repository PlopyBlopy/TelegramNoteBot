import type { ColorInfo, NoteMetadata, TagInfo, Theme } from "./api.model";

/* Temp */

export const TagColorData: ColorInfo[] = [
  { id: 0, variable: "var(--tag-color-1)", name: "Голубой" },
  { id: 1, variable: "var(--tag-color-2)", name: "Фиолетовый" },
  { id: 2, variable: "var(--tag-color-3)", name: "Зеленый" },
  { id: 3, variable: "var(--tag-color-4)", name: "Оранжевый" },
  { id: 4, variable: "var(--tag-color-5)", name: "Розовый" },
  { id: 5, variable: "var(--tag-color-6)", name: "Бирюзовый" },
  { id: 6, variable: "var(--tag-color-7)", name: "Желтый" },
  { id: 7, variable: "var(--tag-color-8)", name: "Бежевый" },
  { id: 8, variable: "var(--tag-color-9)", name: "Синий" },
  { id: 9, variable: "var(--tag-color-10)", name: "Лаймовый" },
  { id: 10, variable: "var(--tag-color-11)", name: "Кремовый" },
  { id: 11, variable: "var(--tag-color-12)", name: "Лимонный" },
];

export const NoteColorData: ColorInfo[] = [
  { id: 0, variable: "var(--note-color-1)", name: "Розовый" },
  { id: 1, variable: "var(--note-color-2)", name: "Персиковый" },
  { id: 2, variable: "var(--note-color-3)", name: "Светло-желтый" },
  { id: 3, variable: "var(--note-color-4)", name: "Мятный" },
  { id: 4, variable: "var(--note-color-5)", name: "Голубой" },
  { id: 5, variable: "var(--note-color-6)", name: "Лавандовый" },
  { id: 6, variable: "var(--note-color-7)", name: "Розово-персиковый" },
  { id: 7, variable: "var(--note-color-8)", name: "Аквамариновый" },
  { id: 8, variable: "var(--note-color-12)", name: "Сиреневый" },
  { id: 9, variable: "var(--note-color-9)", name: "Песочный" },
  { id: 10, variable: "var(--note-color-10)", name: "Светло-зеленый" },
  { id: 11, variable: "var(--note-color-11)", name: "Светло-синий" },
];

export const TagData: TagInfo[] = [
  { id: 0, title: "first", colorId: 0 },
  { id: 1, title: "second", colorId: 1 },
  { id: 2, title: "third", colorId: 2 },
  { id: 3, title: "fourth", colorId: 3 },
  { id: 4, title: "fifth", colorId: 4 },
  { id: 5, title: "sixth", colorId: 5 },
  { id: 6, title: "seventh", colorId: 6 },
  { id: 7, title: "eigth", colorId: 7 },
  { id: 8, title: "nineth", colorId: 8 },
  { id: 9, title: "tenth", colorId: 9 },
  { id: 10, title: "eleventh", colorId: 10 },
  { id: 11, title: "twelfth", colorId: 11 },
];

export const ThemeData: Theme[] = [
  { id: 0, title: "Все" },
  { id: 1, title: "First" },
  { id: 2, title: "Second" },
  { id: 3, title: "Third" },
  { id: 4, title: "Fourth" },
  { id: 5, title: "Fifth" },
  { id: 6, title: "Sixth" },
  { id: 7, title: "Seventh" },
  { id: 8, title: "Eigth" },
  { id: 9, title: "Nineth" },
  { id: 10, title: "Tenth" },
  { id: 11, title: "Eleventh" },
  { id: 12, title: "Twelfth" },
];

export const Notes = (theme: number, tags: number[]): NoteMetadata[] => {
  return Array.from({ length: 20 }, (_, i) => {
    const date = new Date();
    date.setDate(date.getDate() + i);
    const options: Intl.DateTimeFormatOptions = {
      year: "numeric",
      month: "short",
      day: "numeric",
      hour: "numeric",
      minute: "numeric",
      timeZone: "Europe/Moscow",
    };

    return {
      note: {
        id: i,
        title: `Title ${i}-${Math.floor(Math.random() * (999 - 100 + 1)) + 100}`,
        description:
          "Описание \nЭто с новой строки **Description** for some note, and this note have *some length* for test max length on note Prewiew or any.",
      },
      completed: false,
      theme: theme == ThemeData[0].id ? ThemeData[Math.floor(Math.random() * NoteColorData.length)].id : theme,
      tagsId: tags,
      noteColorId: NoteColorData[Math.floor(Math.random() * NoteColorData.length)].id,
      createdAt: date.toLocaleDateString("ru-RU", options),
    };
  });
};
