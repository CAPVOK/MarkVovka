import React from "react"
import "./MainPage.scss"
import { Photo } from "../../componets"
import { useSelector } from "../../core"
import { selectStationInfo } from "../../core/slices/station"

export const MainPage: React.FC  = () => {

  const info = useSelector(selectStationInfo);

  return (
    <div className="mainPage">
      <h1>Россия-1</h1>
      
      <Photo/>
    {/*   <div className="color"></div>
      <div className="color"></div>
      <div className="color"></div>
      <div className="color"></div>
      <div className="color"></div>
      <div className="color"></div>
      <div className="color"></div>
      <div className="color"></div>
      <div className="color"></div>
      <div className="color"></div>
      <div className="color"></div> */}
    </div>
  )
  
}
