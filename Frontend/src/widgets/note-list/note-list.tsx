import styles from "./note-list.module.css";
import { NoteCard } from "@/features/note-card";
import { NoteColorData, type NoteMetadata } from "@/shared/api";

type Props = {
  // onClick: () => {};
  notesMetadata: NoteMetadata[];
};

export const NoteList = ({ notesMetadata }: Props) => {
  return (
    <div className={styles.container}>
      {notesMetadata.map((nm) => (
        <NoteCard key={nm.note.id} noteMetadata={nm} backgroundColor={NoteColorData[nm.noteColorId].variable} />
      ))}
    </div>
  );
};
