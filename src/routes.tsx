import { RouteObject } from "react-router-dom";
import { MainLayout } from "./componets";
import { AuthPage, ControlPage, ErrorPage, MainPage, } from "./pages";
import { SatelliteIcon, ControlIcon } from "./componets/icons";

export interface IAppRoute {
  label?: string;
  path: string;
  index?: boolean;
  icon?: JSX.Element;
  element: JSX.Element;
  children?: IAppRoute[];
}

export const routes: IAppRoute[] = [
  {
    path: "/",
    element: <MainLayout />,
    children: [
      {
        label: "Станция",
        path: "/",
        index: true,
        element: <MainPage />,
        icon: < SatelliteIcon/>,
      },
      {
        label: "Управленние",
        path: "/control",
        element: <ControlPage />,
        icon: <ControlIcon />,
      },
    ],
  },
  {
    path: "/auth",
    element: <AuthPage />,
  },
  {
    path: "*",
    element: <ErrorPage />,
  },
];

export const realRoutes: RouteObject[] = convertRoutes(routes);

function convertRoutes(routes: IAppRoute[]): RouteObject[] {
  return routes.map((route) => {
    const convertedRoute: RouteObject = {
      path: route.path,
      element: route.element,
    };
    if (route.children) {
      convertedRoute.children = convertRoutes(route.children);
    }
    return convertedRoute;
  });
}
