import styles from "./marked-word-button.module.css";

export interface Style {
  color?: string;
}

interface Props extends Style {
  onClick: () => void;
  text: string;
}

export const MarkedWordButton = ({ onClick, text, color }: Props) => {
  const style = {
    color: color,
  };
  return (
    <button style={style} className={styles.button} onClick={onClick}>
      {text}
    </button>
  );
};
