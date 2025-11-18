import { NoteColorData, Notes, TagColorData, TagData, ThemeData } from "./api.config";
import type { ColorInfo, CreateNote, NoteMetadata, NotesFilter, TagInfo, Theme, UpdateNote } from "./api.model";

const apiurl = "http://localhost:4153";

export const postNote = async (note: CreateNote) => {
  const response = await fetch(`${apiurl}/note`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(note),
  });
  return await response.json();
};

// export const postTag = async (tag: TagInfo) => {};

// export const postCompleteNote = async (id: number) => {};

export const getNotes = async (filter: NotesFilter): Promise<NoteMetadata[]> => {
  // const params = new URLSearchParams({
  //   limit: `${filter.limit}`,
  //   search: `${filter.search}`,
  //   theme: `${filter.theme}`,
  //   tags: `${filter.tags}`,
  // });

  // const response = await fetch(`${apiurl}/note?${params}`);
  // const data = await response.json();
  // return data;
  return await Notes(filter.theme, filter.tags);
};

export const getTags = async (): Promise<TagInfo[]> => {
  // const response = await fetch(`${apiurl}/note/tags`);
  // const data = await response.json();
  // return data;

  return await TagData;
};

export const getThemes = async (): Promise<Theme[]> => {
  // const response = await fetch(`${apiurl}/note/themes`);
  // const data = await response.json();
  // return data;

  return await ThemeData;
};

export const getTagsColors = (): ColorInfo[] => {
  // const response = await fetch(`${apiurl}/note/colors`);
  // const data = await response.json();
  // return data;

  return TagColorData;
};

export const getNoteColors = (): ColorInfo[] => {
  // const response = await fetch(`${apiurl}/note/tags/colors`);
  // const data = await response.json();
  // return data;

  return NoteColorData;
};

// export const patchNote = async (note: UpdateNote) => {};

// export const deleteNote = async (id: number) => {};
