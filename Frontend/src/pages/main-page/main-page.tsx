import { NoteList } from "@/widgets/note-list/note-list";
import styles from "./main-page.module.css";
import { ThemeRow } from "@/widgets/theme-row";
import { useEffect, useState } from "react";
import { TagsSelect } from "@/widgets/tag-select";
import { SearchBar } from "@/widgets/search-bar";
import { getNotes, getTags, getThemes, type NoteMetadata, type NotesFilter, type TagInfo, type Theme } from "@/shared/api";

//TODO: сохранить палитру цветов и их имен в массиве, что бы не делать каждый раз запрос в api

export const MainPage = () => {
  const [themes, setThemes] = useState<Theme[]>([]);
  const [tags, setTags] = useState<TagInfo[]>([]);
  const [notes, setNotes] = useState<NoteMetadata[]>([]);
  // const [search, setSearch] = useState<string>("");

  const [isLoading, setIsLoading] = useState({
    notes: true,
    tags: true,
    themes: true,
  });

  const [filter, setFilter] = useState<NotesFilter>({
    limit: 20,
    search: "",
    theme: 0,
    tags: [],
  });

  useEffect(() => {
    const loadTags = async () => {
      const data = await getTags();
      setTags(data);
      setIsLoading((prev) => ({ ...prev, tags: false }));
    };

    loadTags();
  }, []);

  useEffect(() => {
    const loadThemes = async () => {
      const data = await getThemes();
      setThemes(data);
      setIsLoading((prev) => ({ ...prev, themes: false }));
    };

    loadThemes();
  }, []);

  useEffect(() => {
    const loadNotes = async () => {
      setIsLoading((prev) => ({ ...prev, notes: true }));
      const data = await getNotes(filter);
      setNotes(data);
      setIsLoading((prev) => ({ ...prev, notes: false }));
    };

    loadNotes();
  }, [filter]);

  const handleTagsChange = (selectedTags: number[]) => {
    setFilter((prev) => ({ ...prev, tags: selectedTags }));
  };

  const handleThemeChange = (selectedTheme: number) => {
    setFilter((prev) => ({ ...prev, theme: selectedTheme }));
  };

  const handleSearchChange = (search: string) => {
    setFilter((prev) => ({ ...prev, search: search }));
  };

  // const handleFilter = (e: React.ChangeEvent<HTMLInputElement>) => {
  //   const { name, value } = e.target;
  //   setFilter((prev) => ({ ...prev, [name]: value }));
  // };
  return (
    <div className={styles.container}>
      {isLoading.themes ? <>Loading themes...</> : <ThemeRow options={themes} value={filter.theme} onChange={handleThemeChange} />}
      <SearchBar value={filter.search} onSearch={handleSearchChange} placeholder="Поиск заметок..." delay={300} />
      {isLoading.tags ? (
        <>Loading tags...</>
      ) : (
        <TagsSelect options={tags} value={filter.tags} onChange={handleTagsChange} placeholder="Теги еще не созданы" />
      )}
      {isLoading.notes ? <>Loading notes...</> : <NoteList notesMetadata={notes} />}
    </div>
  );
};
