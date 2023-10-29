import React from "react";
import { SideBar } from "../SideBar";
import { NotificationBar } from "../NotificationBar";
import "./MainLayout.scss";
import { Outlet } from "react-router-dom";
import { useSelector } from "../../core";
import { selectConnection } from "../../core/slices/connection";
import { LoadingConnection } from "../LoadingConnection";

export const MainLayout: React.FC = () => {
  const isConnected = useSelector(selectConnection);
  return (
    <div className="mainlayout">
      <SideBar />
      {isConnected ? <Outlet /> : <LoadingConnection />}
      <NotificationBar />
    </div>
  );
};
