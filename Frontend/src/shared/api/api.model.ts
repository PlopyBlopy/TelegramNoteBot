export type NoteMetadata = {
  note: Note;
  completed: boolean;
  theme: number;
  tagsId: number[];
  noteColorId: number;
  createdAt: string;
};

export type Note = {
  id: number;
  title: string;
  description: string;
};

export type CreateNote = {
  title: string;
  description: string;
  themeId: number;
  tagsId: number[];
  noteColorId: number;
};

export type UpdateNote = {
  id: number;
  title: string;
  description: string;
  theme: string;
  tags: string[];
  noteColorId: number;
};

export type Theme = {
  id: number;
  title: string;
};

export type TagInfo = {
  id: number;
  title: string;
  colorId: number;
};

export type ColorInfo = {
  id: number;
  variable: string;
  name: string;
};

export type NotesFilter = {
  limit: number;
  search: string;
  theme: number;
  tags: number[];
};
