import { useState } from "react";
import { ILogInInfo, useLogInMutation } from "../../core/api/authApi";
import "./LogIn.scss";
import { ChangeEvent } from "../../App.typig";

function LogInForm() {
  const [logIn] = useLogInMutation();

  const [formData, setFormData] = useState<ILogInInfo>({
    userName: "",
    password: "",
  });
  const [message, setMessage] = useState("");

  function handleChange(event: ChangeEvent) {
    const { id, value } = event.target;
    setFormData((prevState) => ({ ...prevState, [id]: value }));
  }

  function clickSignUp() {
    if (formData.password && formData.userName) {
      logIn(formData);
    } else [setMessage("Введите все данные")];
  }

  return (
    <div className="login">
      <h3>Войти</h3>
      <label htmlFor="userName">Username</label>
      <input
        type="text"
        placeholder="capvok"
        id="userName"
        value={formData.userName}
        onChange={handleChange}
      />

      <label htmlFor="password">Password</label>
      <input
        type="password"
        placeholder="123123"
        id="password"
        value={formData.password}
        onChange={handleChange}
      />

      <button onClick={clickSignUp}>Log In</button>
      <p>{message}</p>
    </div>
  );
}

export { LogInForm };
