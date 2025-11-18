import { TagData, ThemeData, type NoteMetadata } from "@/shared/api";
import styles from "./note-card.module.css";
import { ButtonIcon } from "@/shared/components/button-icon";
import { Icons } from "@/shared/assets/icons";
import { MarkedWord } from "@/shared/components/marked-word";
import { Checkbox } from "@/shared/components/checkbox/checkbox";
import { Tag } from "../tag";

interface Style {
  backgroundColor?: string;
}

interface PropNoteContainer extends Style {
  noteMetadata: NoteMetadata;
  // onComplete: () => void;
}

export const NoteCard = ({ noteMetadata, backgroundColor }: PropNoteContainer) => {
  if (!noteMetadata) {
    console.warn("noteMetadata component received undefined noteMetadata");
    return null; // или <div>Неизвестный тег</div>
  }

  const handleEdit = () => {};
  const handleDelete = () => {};

  const style: React.CSSProperties = {
    backgroundColor: backgroundColor,
  };

  const themeTitle = ThemeData[noteMetadata.theme].title;

  return (
    <div style={style} className={styles.container}>
      <div className={styles.header}>
        <div className={styles.headerLeft}>
          <Checkbox />
          <div className={styles.title}>{noteMetadata.note.title}</div>
        </div>
        <MarkedWord text={themeTitle} color="var(--text-color-primary)" backgroundColor="var(--color-light)" />
      </div>
      <div className={styles.middle}>
        <div>{noteMetadata.note.description}</div>
        <div className={styles.tags}>
          {noteMetadata.tagsId.map((nm, index) => (
            <Tag key={nm} tag={TagData[index]} />
          ))}
        </div>
      </div>
      <div className={styles.footer}>
        <div>{noteMetadata.createdAt}</div>
        <div className={styles.footerLeft}>
          <ButtonIcon onClick={handleEdit} IconComponent={Icons.elements.edit} />
          <ButtonIcon onClick={handleDelete} IconComponent={Icons.elements.delete} color="red" />
        </div>
      </div>
    </div>
  );
};
