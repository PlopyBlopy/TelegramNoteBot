import { MainPage } from "@/pages/main-page/main-page";
import { ROUTES } from "@/shared/routes";
import { PageWrapper } from "@/widgets/layouts/page-wrapper";
import { createBrowserRouter } from "react-router-dom";

export const Router = createBrowserRouter([
  {
    path: ROUTES.Main.base,
    element: <PageWrapper />,
    children: [
      {
        index: true,
        element: <MainPage />,
      },
    ],
  },
]);
