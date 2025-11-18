import styles from "./primary-button-submit.module.css";

type Props = {
  text?: string;
};

export const PrimaryButtonSubmit = ({ text = "PBUTTON" }: Props) => {
  return (
    <button type="submit" className={styles.button}>
      {text}
    </button>
  );
};
