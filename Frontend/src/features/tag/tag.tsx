import { TagColorData, type TagInfo } from "@/shared/api";
import { MarkedWord } from "@/shared/components/marked-word";

type Props = {
  tag: TagInfo | undefined;
};

export const Tag = ({ tag }: Props) => {
  if (!tag) {
    console.warn("Tag component received undefined tag");
    return null; // или <div>Неизвестный тег</div>
  }
  const tagColor = TagColorData[tag.colorId];

  return <MarkedWord text={tag.title} color={"var(--tag-text-color)"} backgroundColor={tagColor.variable} />;
};
