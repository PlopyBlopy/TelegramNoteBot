import { useNavigate } from "react-router-dom";
import { ROUTES } from "./routes.config";

export const useAppNavigate = () => {
  const navigate = useNavigate();

  return {
    goToMain: () => navigate(ROUTES.Main.path),
    customNavigate: navigate,
  };
};
