import { useState } from "react";
import styles from "./note-create.module.css";
import type { NoteForm } from "./note-create.model";
import { ButtonIcon } from "@/shared/components/button-icon";
import { Icons } from "@/shared/assets/icons";
import { PrimaryButtonSubmit } from "@/shared/components/primary-button-submit";
import { ColorPicker } from "@/features/color-picker";
import { NoteColorData, TagData, ThemeData, type CreateNote } from "@/shared/api";
import { DropdownTheme } from "@/features/dropdown-theme";
import { TagsSelect } from "../tag-select";

type Props = {
  onClose: () => void;
  onSubmit: (note: CreateNote) => void;
};

export const NoteCreate = ({ onClose, onSubmit }: Props) => {
  // TODO: MobX или Singleton для themes, tags

  const [form, setForm] = useState<NoteForm>({
    title: "",
    description: "",
    themeId: 0,
    tagsId: [],
    noteColorId: 0,
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const { name, value } = e.target;
    setForm((prev) => ({ ...prev, [name]: value }));
  };

  const handleSubmitForm = () => {
    const note: CreateNote = {
      title: form.title,
      description: form.description,
      themeId: form.themeId,
      tagsId: form.tagsId,
      noteColorId: form.noteColorId,
    };
    onSubmit(note);
  };

  const tags = TagData;
  const themes = ThemeData;
  const noteColorData = NoteColorData;

  return (
    <div>
      <div className={styles.header}>
        Создать новую заметку
        <ButtonIcon onClick={onClose} IconComponent={Icons.elements.close} color="var(--color-red)" />
      </div>
      <form className={styles.middle} onSubmit={handleSubmitForm}>
        <div>
          <label>Название</label>
          <input
            type="text"
            id="title"
            name="title"
            required={true}
            value={form.title}
            onChange={handleChange}
            placeholder="Название"
            className={styles.formInput}
          />
        </div>

        <div>
          <label>Описание</label>
          <textarea
            id="description"
            name="description"
            required={true}
            value={form.description}
            onChange={handleChange}
            placeholder="Описание..."
            className={styles.formInput}
            rows={4}
          />
        </div>
        <div>
          <label>Темы</label>
          <DropdownTheme
            options={themes}
            value={themes[form.themeId]}
            onChange={(themeId) => setForm((prev) => ({ ...prev, themeId: themeId }))}
            placeholder="Тема"
          />
        </div>
        <div>
          <label>Теги</label>
          <TagsSelect options={tags} value={form.tagsId} onChange={(tags) => setForm((prev) => ({ ...prev, tagsId: tags }))} />
        </div>
        <div>
          <ColorPicker
            options={noteColorData}
            value={form.noteColorId}
            onColorSelectId={(color) => setForm((prev) => ({ ...prev, noteColorId: color }))}
            placeholder="Цвет темы"
          />
        </div>

        <div className={styles.footer}>
          <PrimaryButtonSubmit text="Создать" />
        </div>
      </form>
    </div>
  );
};
